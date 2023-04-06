package repository

import (
	"context"
	"deepflower/internal/model"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
)

type ProjectStorage struct {
	Db *sqlx.DB
}

func NewProjectStorage(dbpool *sqlx.DB) ProjectStorage {
	return ProjectStorage{Db: dbpool}
}

func (s *ProjectStorage) CreateProject(ctx context.Context, name, info, graph, dreamName, creater string) (model.Project, error) {
	var m model.Project
	tx := s.Db.MustBegin()

	q1 := `
	INSERT INTO project (name, info ,published, status, creater, energy, graph) 
	VALUES ($1,$2,$3,$4,$5,$6,$7) 
	returning *;
	`
	err := tx.GetContext(ctx, &m, q1, name, info, false, "created", creater, 0, graph)
	if err != nil {
		tx.Rollback()
		return model.Project{}, err
	}
	q2 := `INSERT INTO dream_project(dreamid, projectid) 
	VALUES ($1, (SELECT id FROM dream WHERE name=$2));
	`
	result, err := tx.ExecContext(ctx, q2, m.ID, dreamName)
	if err != nil {
		tx.Rollback()
		return model.Project{}, err
	}
	n, err := result.RowsAffected()
	if err != nil {
		tx.Rollback()
		return model.Project{}, err
	}
	if n == 0 {
		tx.Rollback()
		return model.Project{}, fmt.Errorf("error: in dream_project no new row")
	}

	err = tx.Commit()
	if err != nil {
		return model.Project{}, err
	}
	return m, nil
}

// поиск по имени проекта (searchTerm), сортировка, простая offset пагинация
func (s *ProjectStorage) SearchProjects(ctx context.Context, userId string,
	limit uint64, offset uint64, onlyMyProjects bool, order string, searchTerm string,
	sort string) ([]model.Project, int, error) {
	var projects []model.Project
	var args []interface{}
	var query string
	var queryCnt string
	var count int

	filter := fmt.Sprintf(` ORDER BY %s %s LIMIT %d OFFSET %d;`, order, sort, limit, offset)
	switch {
	case searchTerm != "" && onlyMyProjects:
		query = `SELECT * FROM project WHERE LOWER(name) LIKE CONCAT('%%',$1::text,'%%') AND creater=$2`
		queryCnt = `SELECT count(id) FROM project WHERE LOWER(name) LIKE CONCAT('%%',$1::text,'%%') AND creater=$2`
		args = append(args, searchTerm, userId)
	case searchTerm != "" && !onlyMyProjects:
		query = `SELECT * FROM project WHERE LOWER(name) LIKE CONCAT('%%',$1::text,'%%')`
		queryCnt = `SELECT count(id) FROM project WHERE LOWER(name) LIKE CONCAT('%%',$1::text,'%%')`
		args = append(args, searchTerm)
	case searchTerm == "" && onlyMyProjects:
		query = `SELECT * FROM project WHERE creater=$1`
		queryCnt = `SELECT count(id) FROM project WHERE creater=$1`
		args = append(args, userId)
	case searchTerm == "" && !onlyMyProjects:
		query = `SELECT * FROM project`
		queryCnt = `SELECT count(id) FROM project`
	}
	q := query + filter

	if err := s.Db.SelectContext(ctx, &projects, q, args...); err != nil {
		return []model.Project{}, 0, err
	}
	s.Db.GetContext(ctx, &count, queryCnt, args...)
	return projects, count, nil
}

func (s *ProjectStorage) GetProjectById(ctx context.Context, projectId string) (model.Project, error) {
	var project model.Project
	q := `SELECT * FROM project WHERE id=$1;`
	if err := s.Db.GetContext(ctx, &project, q, projectId); err != nil {
		return model.Project{}, err
	}
	return project, nil
}

func (s *ProjectStorage) DeleteUserProject(ctx context.Context, projectId string) error {
	q := `DELETE FROM project WHERE id=$1;`
	result, err := s.Db.ExecContext(ctx, q, projectId)
	if err != nil {
		return err
	}
	count, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if count != 1 {
		return fmt.Errorf("count != 1")
	}
	return nil
}

// dangerous method. strictly check the input data to patch
func (s *ProjectStorage) UpdateUserProject(ctx context.Context, projectId string, patchProject map[string]interface{}) (model.Project, error) {
	var project model.Project
	sqlSet := `UPDATE project SET`
	for key := range patchProject {
		sqlSet += fmt.Sprintf(` %s=:%s,`, strings.ToLower(key), key)
	}
	sqlSet = strings.TrimSuffix(sqlSet, ",")
	sqlSet += fmt.Sprintf(` WHERE id='%s' returning *;`, projectId)
	rows, err := s.Db.NamedQueryContext(ctx, sqlSet, patchProject)
	if err != nil {
		return model.Project{}, err
	}
	for rows.Next() {
		err := rows.StructScan(&project)
		if err != nil {
			return model.Project{}, err
		}
	}
	return project, nil
}

func (s *ProjectStorage) EnergyTxUserToProject(ctx context.Context, userId, projectId string, energy uint64) error {
	tx := s.Db.MustBegin()
	query1 := `UPDATE users SET energy=energy-$1 WHERE id=$2;`
	query2 := `UPDATE project SET energy=energy+$1 WHERE id=$2;`

	_, err := tx.ExecContext(ctx, query1, energy, userId)
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = tx.ExecContext(ctx, query2, energy, projectId)
	if err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

// CODE PROTOTYPE !!!!!!!!!!!!!!
// / TODO :POSTGRES-> JSONB OR DECOMPOSE NODES DATA OR OTHER DB ?????
func (s *ProjectStorage) EnergyTxUserToTask(ctx context.Context, userId, projectId, nodeId string, energy uint64) error {
	tx := s.Db.MustBegin()
	query1 := `UPDATE users SET energy=energy-$1 WHERE id=$2;`
	query2 := `SELECT graph FROM project WHERE id=$1;`

	_, err := tx.ExecContext(ctx, query1, energy, userId)
	if err != nil {
		tx.Rollback()
		return err
	}
	var graphStr string
	err = tx.GetContext(ctx, &graphStr, query2, projectId)
	if err != nil {
		tx.Rollback()
		return err
	}
	var graph model.Graph
	err = json.Unmarshal([]byte(graphStr), &graph)
	if err != nil {
		tx.Rollback()
		return err
	}
	for n, c := range graph.Cells {
		if c.Shape == "slow-model" || c.Shape == "fast-model" {
			if c.Id == nodeId {
				graph.Cells[n].Data.Energy = c.Data.Energy + energy
			}
		}
	}
	updatedGraphByte, err := json.Marshal(graph)
	if err != nil {
		tx.Rollback()
		return err
	}

	var query3 = `UPDATE project SET graph=$1 WHERE id=$2;`
	_, err = tx.ExecContext(ctx, query3, string(updatedGraphByte), projectId)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	return nil

}

// CODE PROTOTYPE !!!!!!!!!!!!!!
// / TODO :POSTGRES-> JSONB OR DECOMPOSE NODES DATA OR OTHER DB ?????
func (s *ProjectStorage) UpdateTaskStatus(ctx context.Context, projectId, nodeId, newStatus string) error {
	tx := s.Db.MustBegin()
	query2 := `SELECT graph FROM project WHERE id=$1;`
	var graphStr string

	///// check
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	err := tx.GetContext(ctx, &graphStr, query2, projectId)
	if err != nil {
		tx.Rollback()
		return err
	}
	var graph model.Graph
	err = json.Unmarshal([]byte(graphStr), &graph)
	if err != nil {
		tx.Rollback()
		return err
	}
	for n, c := range graph.Cells {
		if c.Shape == "slow-model" || c.Shape == "fast-model" {
			if c.Id == nodeId {
				graph.Cells[n].Data.Status = newStatus
			}
		}
	}
	updatedGraphByte, err := json.Marshal(graph)
	if err != nil {
		tx.Rollback()
		return err
	}
	var query3 = `UPDATE project SET graph=$1 WHERE id=$2;`
	_, err = tx.ExecContext(ctx, query3, string(updatedGraphByte), projectId)
	if err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}
