package repository

import (
	"context"
	"deepflower/internal/model"
	"fmt"

	"deepflower/pkg/postgres"

	"github.com/jmoiron/sqlx"
)

type Tx interface {
	ExtractTx(ctx context.Context) (*sqlx.Tx, error)
}

type LocationStorage struct {
	Db *postgres.PG
}

func NewLocationStorage(db *postgres.PG) *LocationStorage {
	return &LocationStorage{Db: db}
}

func (s *LocationStorage) CreateLocation(ctx context.Context, creater string, name string, info string, geolocation string, radius uint64, height uint64) (model.Location, error) {
	tx := s.Db.ExtractTx(ctx)
	var m model.Location
	q := `
	INSERT INTO location (name, info, creater, energy, geolocation, radius, height, active) 
	VALUES ($1,$2,$3,$4,$5,$6,$7, $8) 
	returning *;
	`
	err := tx.GetContext(ctx, &m, q, name, info, creater, 0, geolocation, radius, height, true)
	if err != nil {
		return model.Location{}, err
	}
	return m, nil
}

// поиск по имени локации(searchTerm), сортировка, простая offset пагинация
func (s *LocationStorage) SearchLocations(ctx context.Context, userId string,
	limit uint64, offset uint64, onlyMyLocations bool,
	order string, searchTerm string,
	sort string) ([]model.Location, int, error) {
	tx := s.Db.ExtractTx(ctx)
	var locations []model.Location
	var args []interface{}
	var query string
	var queryCnt string
	var count int

	filter := fmt.Sprintf(` ORDER BY %s %s LIMIT %d OFFSET %d;`, order, sort, limit, offset)
	switch {
	case searchTerm != "" && onlyMyLocations:
		query = `SELECT * FROM location WHERE LOWER(name) LIKE CONCAT('%%',$1::text,'%%') AND creater=$2`
		queryCnt = `SELECT count(id) FROM location WHERE LOWER(name) LIKE CONCAT('%%',$1::text,'%%') AND creater=$2`
		args = append(args, searchTerm, userId)
	case searchTerm != "" && !onlyMyLocations:
		query = `SELECT * FROM location WHERE LOWER(name) LIKE CONCAT('%%',$1::text,'%%')`
		queryCnt = `SELECT count(id) FROM location WHERE LOWER(name) LIKE CONCAT('%%',$1::text,'%%')`
		args = append(args, searchTerm)
	case searchTerm == "" && onlyMyLocations:
		query = `SELECT * FROM location WHERE creater=$1`
		queryCnt = `SELECT count(id) FROM location WHERE creater=$1`
		args = append(args, userId)
	case searchTerm == "" && !onlyMyLocations:
		query = `SELECT * FROM location`
		queryCnt = `SELECT count(id) FROM location`
	}
	q := query + filter

	if err := tx.SelectContext(ctx, &locations, q, args...); err != nil {
		return []model.Location{}, 0, err
	}
	tx.GetContext(ctx, &count, queryCnt, args...)
	return locations, count, nil
}

func (s *LocationStorage) GetLocationDreams(ctx context.Context, locationId string) ([]model.Dream, error) {
	tx := s.Db.ExtractTx(ctx)
	var dreams []model.Dream
	q := `SELECT * FROM dream WHERE id IN (SELECT dreamid FROM dream_location WHERE locationid=$1);`
	if err := tx.SelectContext(ctx, &dreams, q, locationId); err != nil {
		return []model.Dream{}, err
	}
	return dreams, nil

}

func (s *LocationStorage) GetLocationById(ctx context.Context, locationId string) (model.Location, error) {
	tx := s.Db.ExtractTx(ctx)
	var location model.Location
	q := `SELECT * FROM location WHERE id=$1;`
	if err := tx.GetContext(ctx, &location, q, locationId); err != nil {
		return model.Location{}, err
	}
	return location, nil
}

func (s *LocationStorage) DeleteUserLocation(ctx context.Context, locationId string) error {
	tx := s.Db.ExtractTx(ctx)
	q := `DELETE FROM location WHERE id=$1;`
	result, err := tx.ExecContext(ctx, q, locationId)
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

/*

func (s *LocationStorage) UpdateUserLocation(ctx context.Context, locationId string, patchLocation map[string]interface{}) (model.Location, error) {
	var location model.Location
	sqlSet := `UPDATE location SET`
	for key := range patchLocation {
		sqlSet += fmt.Sprintf(` %s=:%s,`, strings.ToLower(key), key)
	}
	sqlSet = strings.TrimSuffix(sqlSet, ",")
	sqlSet += fmt.Sprintf(` WHERE id='%s' returning *;`, locationId)
	rows, err := s.Db.NamedQueryContext(ctx, sqlSet, patchLocation)
	if err != nil {
		return model.Location{}, err
	}
	for rows.Next() {
		err := rows.StructScan(&location)
		if err != nil {
			return model.Location{}, err
		}
	}
	return location, nil
}

*/

func (s *LocationStorage) AddEnergy(ctx context.Context, locationId string, energy uint64) error {
	tx := s.Db.ExtractTx(ctx)
	query2 := `UPDATE location SET energy=energy+$1 WHERE id=$2;`
	_, err := tx.ExecContext(ctx, query2, energy, locationId)
	if err != nil {
		return err
	}
	return nil
}

// // транзакция энергии от пользователя к локации
// func (s *LocationStorage) EnergyTxUserToLocation(ctx context.Context, userId, locationId string, energy uint64) error {
// 	tx := s.Db.MustBegin()
// 	query1 := `UPDATE users SET energy=energy-$1 WHERE id=$2;`
// 	query2 := `UPDATE location SET energy=energy+$1 WHERE id=$2;`

// 	_, err := tx.ExecContext(ctx, query1, energy, userId)
// 	if err != nil {
// 		tx.Rollback()
// 		return err
// 	}
// 	_, err = tx.ExecContext(ctx, query2, energy, locationId)
// 	if err != nil {
// 		tx.Rollback()
// 		return err
// 	}
// 	if err := tx.Commit(); err != nil {
// 		return err
// 	}
// 	return nil
// }
