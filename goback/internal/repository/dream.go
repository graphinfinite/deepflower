package repository

import (
	"context"
	"deepflower/internal/model"
	"deepflower/pkg/postgres"
	"fmt"
)

type DreamStorage struct {
	Db *postgres.PG
}

func NewDreamStorage(db *postgres.PG) *DreamStorage {
	return &DreamStorage{Db: db}
}

// TODO
func (s *DreamStorage) CreateDream(ctx context.Context, name, info, location, creater string) (model.Dream, error) {
	tx := s.Db.ExtractTx(ctx)
	var m model.Dream
	q1 := `
	INSERT INTO dream (name, info,published, status, creater, energy, countG) 
	VALUES ($1,$2,$3,$4,$5,$6,$7,$8) 
	returning *;
	`
	err := tx.GetContext(ctx, &m, q1, name, info, false, "created", creater, 0, 0)
	if err != nil {
		return model.Dream{}, err
	}
	q2 := `INSERT INTO dream_location(dreamid, locationid) 
	VALUES ($1, (SELECT id FROM location WHERE name=$2));
	`
	result, err := tx.ExecContext(ctx, q2, m.ID, location)
	if err != nil {
		return model.Dream{}, err
	}
	n, err := result.RowsAffected()
	if err != nil {
		return model.Dream{}, err
	}
	if n == 0 {
		return model.Dream{}, fmt.Errorf("error: in dream_location no new row")
	}
	return m, nil
}

// no used
func (s *DreamStorage) GetAllUserDreams(ctx context.Context, userId string) ([]model.Dream, error) {
	tx := s.Db.ExtractTx(ctx)
	var dreams []model.Dream
	q := `SELECT * FROM dream WHERE creater=$1;`
	if err := tx.SelectContext(ctx, &dreams, q, userId); err != nil {
		return []model.Dream{}, err
	}
	return dreams, nil
}

// поиск по имени мечты(searchTerm), сортировка, простая offset пагинация
func (s *DreamStorage) SearchDreams(ctx context.Context, userId string,
	limit uint64, offset uint64, onlyMyDreams bool, order string, searchTerm string,
	sort string) ([]model.Dream, int, error) {
	tx := s.Db.ExtractTx(ctx)
	var dreams []model.Dream
	var args []interface{}
	var query string
	var queryCnt string
	var count int

	filter := fmt.Sprintf(` ORDER BY %s %s LIMIT %d OFFSET %d;`, order, sort, limit, offset)
	switch {
	case searchTerm != "" && onlyMyDreams:
		query = `SELECT * FROM dream WHERE LOWER(name) LIKE CONCAT('%%',$1::text,'%%') AND creater=$2`
		queryCnt = `SELECT count(id) FROM dream WHERE LOWER(name) LIKE CONCAT('%%',$1::text,'%%') AND creater=$2`
		args = append(args, searchTerm, userId)
	case searchTerm != "" && !onlyMyDreams:
		query = `SELECT * FROM dream WHERE LOWER(name) LIKE CONCAT('%%',$1::text,'%%')`
		queryCnt = `SELECT count(id) FROM dream WHERE LOWER(name) LIKE CONCAT('%%',$1::text,'%%')`
		args = append(args, searchTerm)
	case searchTerm == "" && onlyMyDreams:
		query = `SELECT * FROM dream WHERE creater=$1`
		queryCnt = `SELECT count(id) FROM dream WHERE creater=$1`
		args = append(args, userId)
	case searchTerm == "" && !onlyMyDreams:
		query = `SELECT * FROM dream`
		queryCnt = `SELECT count(id) FROM dream`
	}
	q := query + filter

	if err := tx.SelectContext(ctx, &dreams, q, args...); err != nil {
		return []model.Dream{}, 0, err
	}
	tx.GetContext(ctx, &count, queryCnt, args...)
	return dreams, count, nil
}

func (s *DreamStorage) GetDreamById(ctx context.Context, dreamId string) (model.Dream, error) {
	tx := s.Db.ExtractTx(ctx)
	var dream model.Dream
	q := `SELECT * FROM dream WHERE id=$1;`
	if err := tx.GetContext(ctx, &dream, q, dreamId); err != nil {
		return model.Dream{}, err
	}
	return dream, nil
}

func (s *DreamStorage) DeleteUserDream(ctx context.Context, dreamId string) error {
	tx := s.Db.ExtractTx(ctx)
	q := `DELETE FROM dream WHERE id=$1;`
	result, err := tx.ExecContext(ctx, q, dreamId)
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

// TODO result sql
func (s *DreamStorage) UpdateDreamPublished(ctx context.Context, dreamId string) error {
	tx := s.Db.ExtractTx(ctx)
	q := `UPDATE dream SET publish=true WHERE id=$1`
	if _, err := tx.ExecContext(ctx, q, dreamId); err != nil {
		return err
	}
	return nil
}

/*

// dangerous method. strictly check the input data to patch
func (s *DreamStorage) UpdateUserDream(ctx context.Context, dreamId string, patchDream map[string]interface{}) (model.Dream, error) {

	var dream model.Dream
	sqlSet := `UPDATE dream SET`
	for key := range patchDream {
		sqlSet += fmt.Sprintf(` %s=:%s,`, strings.ToLower(key), key)
	}
	sqlSet = strings.TrimSuffix(sqlSet, ",")
	sqlSet += fmt.Sprintf(` WHERE id='%s' returning *;`, dreamId)
	rows, err := s.Db.NamedQueryContext(ctx, sqlSet, patchDream)
	if err != nil {
		return model.Dream{}, err
	}
	for rows.Next() {
		err := rows.StructScan(&dream)
		if err != nil {
			return model.Dream{}, err
		}
	}
	return dream, nil
}
*/

func (s *DreamStorage) AddEnergy(ctx context.Context, dreamId string, energy uint64) error {
	tx := s.Db.ExtractTx(ctx)
	query2 := `UPDATE dream SET energy=energy+$1 WHERE id=$2;`
	_, err := tx.ExecContext(ctx, query2, energy, dreamId)
	if err != nil {
		return err
	}
	return nil
}
