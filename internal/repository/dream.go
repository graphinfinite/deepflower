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

	q := `
	INSERT INTO dream (name, info, createdAt, publishAt,published, status, creater, energy, location, countG) 
	VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10) 
	returning id, name, info, createdAt, publishAt,published, status, creater, energy, location, countG;
	`
	err := s.Db.QueryRow(q, name, info, time.Now(), time.Now(), false, "created", creater, 0, location, 0).Scan(&m)
	if err != nil {
		return model.Dream{}, err
	}
	return m, nil
}
