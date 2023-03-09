package repository

import (
	"database/sql"
	"time"
)

type DreamStorage struct {
	Db *sql.DB
}

func NewDreamStorage(dbpool *sql.DB) DreamStorage {
	return DreamStorage{Db: dbpool}
}

func (s *DreamStorage) CreateDream(name, info, location, creater string) (int, error) {
	var id int

	q := `
	INSERT INTO dream (name, info, createdAt, publishAt,published, status, creater, energy, location, countG) 
	VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10) 
	returning id`
	err := s.Db.QueryRow(q, name, info, time.Now(), "", false, "created", creater, 0, location, 0).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil

}
