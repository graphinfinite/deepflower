package repository

import (
	"database/sql"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func MigrateDb(dbPool *sql.DB) error {
	q := `CREATE TABLE IF NOT EXISTS "users" (
		id serial PRIMARY KEY,
		createdAt timestamp DEFAULT current_timestamp NOT NULL,
		updatedAt timestamp DEFAULT current_timestamp NOT NULL,
		username VARCHAR(64) UNIQUE NOT NULL,
		hashedPassword VARCHAR(128) NOT NULL,
		active BOOLEAN NOT NULL,
		tgId integer UNIQUE NOT NULL,
		tgChatId integer NOT NULL,
		tgUserName VARCHAR(64),
	 	tgFirstName VARCHAR(64) NOT NULL,
	    tgLastName VARCHAR(64) NOT NULL, 
	  	tgLanguageCode VARCHAR(64) NOT NULL);
		
		CREATE TABLE IF NOT EXIST "dream" (
		name VARCHAR(64) UNIQUE NOT NULL,
		info TEXT,
		createdAt timestamp DEFAULT current_timestamp NOT NULL,
		publishAt timestamp DEFAULT current_timestamp NOT NULL,
		published BOOLEAN NOT NULL,
		status integer NOT NULL,
		creater integer NOT NULL,
		energy integer NOT NULL,
		location VARCHAR(64),
		countG integer NOT NULL);
		
		`
	_, errDb := dbPool.Exec(q)
	if errDb != nil {
		return errDb
	}
	return nil

	/*
		driver, err := postgres.WithInstance(db, &postgres.Config{})
		if err != nil {
			return err
		}
		m, err := migrate.NewWithDatabaseInstance(
			"file:///migrations",
			"postgres", driver)
		if err != nil {
			return err
		}
		m.Up()
		return nil
	*/
}

func NewPostgresPool(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
