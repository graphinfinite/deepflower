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

func NewDreamStorage(dbpool *sqlx.DB) DreamStorage {
	return DreamStorage{Db: dbpool}
}

func (s *DreamStorage) CreateDream(ctx context.Context, name, info, location, creater string) (model.Dream, error) {
	var m model.Dream
	tx := s.Db.MustBegin()
	q := `
	INSERT INTO dream (name, info,published, status, creater, energy, location, countG) 
	VALUES ($1,$2,$3,$4,$5,$6,$7,$8) 
	returning *;
	`
	err := tx.GetContext(ctx, &m, q, name, info, false, "CREATE", creater, 0, location, 0)
	if err != nil {
		tx.Rollback()
		return model.Dream{}, err
	}
	err = tx.Commit()
	if err != nil {
		return model.Dream{}, err
	}
	return m, nil
}

func (s *DreamStorage) GetAllUserDreams(ctx context.Context, userId string) ([]model.Dream, error) {
	var dreams []model.Dream
	q := `SELECT * FROM dream WHERE creater=$1;`
	if err := s.Db.Select(&dreams, q, userId); err != nil {
		return []model.Dream{}, err
	}
	return dreams, nil
}

func (s *DreamStorage) GetDreamById(ctx context.Context, dreamId string) (model.Dream, error) {
	var dream model.Dream
	q := `SELECT * FROM "dream" WHERE id=$1;`
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

func (s *DreamStorage) EnergyTxUserToDream(ctx context.Context, userId, dreamId string, energy uint64) error {
	tx := s.Db.MustBegin()
	query1 := `UPDATE users SET energy=energy-$1 WHERE id='$2';`
	query2 := `UPDATE dream SET energy=energy+$1 WHERE id='$2';`

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
