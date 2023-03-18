package repository

import (
	"deepflower/internal/model"
	"time"

	"github.com/jmoiron/sqlx"
)

type DreamStorage struct {
	Db *sqlx.DB
}

func NewDreamStorage(dbpool *sqlx.DB) DreamStorage {
	return DreamStorage{Db: dbpool}
}

func (s *DreamStorage) CreateDream(name, info, location, creater string) (model.Dream, error) {
	var m model.Dream

	tx := s.Db.MustBegin()

	q := `
	INSERT INTO dream (name, info, createdAt, publishAt,published, status, creater, energy, location, countG) 
	VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10) 
	returning *;
	`
	err := tx.QueryRowx(q, name, info, time.Now(), time.Now(), false, "created", creater, 0, location, 0).StructScan(&m)
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

func (s *DreamStorage) GetAllUserDreams(userId string) ([]model.Dream, error) {
	var dreams []model.Dream
	q := `SELECT * FROM dream WHERE creater=$1;`

	if err := s.Db.Select(&dreams, q, userId); err != nil {
		return []model.Dream{}, err
	}
	return dreams, nil
}
