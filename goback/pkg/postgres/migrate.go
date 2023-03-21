package postgres

import "github.com/jmoiron/sqlx"

func MigrateDb(dbPool *sqlx.DB) error {
	q := `
		CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
	
	
		CREATE TABLE IF NOT EXISTS "users" (
		id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
		createdAt timestamp DEFAULT current_timestamp NOT NULL,
		updatedAt timestamp DEFAULT current_timestamp NOT NULL,
		username VARCHAR(64) UNIQUE NOT NULL,
		hashedPassword VARCHAR(128) NOT NULL,
		active BOOLEAN NOT NULL,
		tgId integer UNIQUE NOT NULL,
		tgChatId integer NOT NULL,
		tgUserName VARCHAR(64) NOT NULL DEFAULT 'empty',
	 	tgFirstName VARCHAR(64) NOT NULL DEFAULT 'empty',
	    tgLastName VARCHAR(64) NOT NULL DEFAULT 'empty', 
	  	tgLanguageCode VARCHAR(12) NOT NULL DEFAULT 'empty');
		
		CREATE TABLE IF NOT EXISTS "dream" (
		id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
		name VARCHAR(64) UNIQUE NOT NULL DEFAULT 'empty',
		info TEXT NOT NULL DEFAULT 'empty',
		createdAt timestamp DEFAULT current_timestamp NOT NULL,
		updatedAt timestamp DEFAULT current_timestamp NOT NULL,
		publishAt timestamp NOT NULL DEFAULT current_timestamp ,
		published BOOLEAN NOT NULL,
		status VARCHAR(32) NOT NULL,
		creater uuid NOT NULL,
		energy bigint NOT NULL,
		location VARCHAR(128) NOT NULL DEFAULT 'empty',
		countG integer NOT NULL DEFAULT 0);
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