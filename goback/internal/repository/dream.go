package repository

import (
	"context"
	"deepflower/internal/model"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
)

type DreamStorage struct {
	Db *sqlx.DB
}

func NewDreamStorage(dbpool *sqlx.DB) *DreamStorage {
	return &DreamStorage{Db: dbpool}
}

func (s *DreamStorage) CreateDream(ctx context.Context, name, info, location, creater string) (model.Dream, error) {
	var m model.Dream
	tx := s.Db.MustBegin()
	q1 := `
	INSERT INTO dream (name, info,published, status, creater, energy, location, countG) 
	VALUES ($1,$2,$3,$4,$5,$6,$7,$8) 
	returning *;
	`
	err := tx.GetContext(ctx, &m, q1, name, info, false, "created", creater, 0, location, 0)
	if err != nil {
		tx.Rollback()
		return model.Dream{}, err
	}
	q2 := `INSERT INTO dream_location(dreamid, locationid) 
	VALUES ($1, (SELECT id FROM location WHERE name=$2));
	`
	result, err := tx.ExecContext(ctx, q2, m.ID, location)
	if err != nil {
		tx.Rollback()
		return model.Dream{}, err
	}
	n, err := result.RowsAffected()
	if err != nil {
		tx.Rollback()
		return model.Dream{}, err
	}
	if n == 0 {
		tx.Rollback()
		return model.Dream{}, fmt.Errorf("error: in dream_location no new row")
	}

	err = tx.Commit()
	if err != nil {
		return model.Dream{}, err
	}
	return m, nil
}

// no used
func (s *DreamStorage) GetAllUserDreams(ctx context.Context, userId string) ([]model.Dream, error) {
	var dreams []model.Dream
	q := `SELECT * FROM dream WHERE creater=$1;`
	if err := s.Db.Select(&dreams, q, userId); err != nil {
		return []model.Dream{}, err
	}
	return dreams, nil
}

// поиск по имени мечты(searchTerm), сортировка, простая offset пагинация
func (s *DreamStorage) SearchDreams(ctx context.Context, userId string,
	limit uint64, offset uint64, onlyMyDreams bool, order string, searchTerm string,
	sort string) ([]model.Dream, int, error) {
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

	if err := s.Db.SelectContext(ctx, &dreams, q, args...); err != nil {
		return []model.Dream{}, 0, err
	}
	s.Db.GetContext(ctx, &count, queryCnt, args...)
	return dreams, count, nil
}

func (s *DreamStorage) GetDreamById(ctx context.Context, dreamId string) (model.Dream, error) {
	var dream model.Dream
	q := `SELECT * FROM dream WHERE id=$1;`
	if err := s.Db.GetContext(ctx, &dream, q, dreamId); err != nil {
		return model.Dream{}, err
	}
	return dream, nil
}

func (s *DreamStorage) DeleteUserDream(ctx context.Context, dreamId string) error {
	q := `DELETE FROM dream WHERE id=$1;`
	result, err := s.Db.ExecContext(ctx, q, dreamId)
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

// транзакция энергии от пользователя к мечте
func (s *DreamStorage) EnergyTxUserToDream(ctx context.Context, userId, dreamId string, energy uint64) error {
	tx := s.Db.MustBegin()
	query1 := `UPDATE users SET energy=energy-$1 WHERE id=$2;`
	query2 := `UPDATE dream SET energy=energy+$1 WHERE id=$2;`

	_, err := tx.ExecContext(ctx, query1, energy, userId)
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = tx.ExecContext(ctx, query2, energy, dreamId)
	if err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}
