package migrations

import "github.com/jmoiron/sqlx"

func Down(dbPool *sqlx.DB) error {
	q := `
		DROP TABLE IF EXISTS "users" CASCADE;
		
		DROP TABLE IF EXISTS "location" CASCADE;

		DROP TABLE IF EXISTS "dream" CASCADE;

		DROP TABLE IF EXISTS "dream_location" CASCADE;

		DROP TABLE IF EXISTS "project" CASCADE;

		DROP TABLE IF EXISTS "dream_project" CASCADE;

		DROP TABLE IF EXISTS "task_users" CASCADE;

		DROP TABLE IF EXISTS "task_process" CASCADE;
		`
	_, errDb := dbPool.Exec(q)
	if errDb != nil {
		return errDb
	}
	return nil

}

func Up(dbPool *sqlx.DB) error {
	q := `
		CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
	
		CREATE TABLE IF NOT EXISTS "users" (
		id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
		createdAt timestamp DEFAULT current_timestamp NOT NULL,
		updatedAt timestamp DEFAULT current_timestamp NOT NULL,
		username VARCHAR(64) UNIQUE NOT NULL,  
		energy bigint NOT NULL DEFAULT 0 CHECK (energy >= 0),
		hashedPassword VARCHAR(128) NOT NULL,
		active BOOLEAN NOT NULL,
		tgId integer UNIQUE NOT NULL,
		tgChatId integer NOT NULL,
		tgUserName VARCHAR(64) NOT NULL DEFAULT 'empty',
	 	tgFirstName VARCHAR(64) NOT NULL DEFAULT 'empty',
	    tgLastName VARCHAR(64) NOT NULL DEFAULT 'empty', 
	  	tgLanguageCode VARCHAR(12) NOT NULL DEFAULT 'empty');
		
		CREATE TABLE IF NOT EXISTS "location" (
		id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
		name VARCHAR(64) UNIQUE NOT NULL,
		info TEXT NOT NULL DEFAULT 'empty',
		createdat timestamp DEFAULT current_timestamp NOT NULL,
		updatedat timestamp DEFAULT current_timestamp NOT NULL,
		creater uuid REFERENCES users (id),
		geolocation point NOT NULL,
		radius bigint NOT NULL DEFAULT 0 CHECK (radius >= 0),
		height bigint NOT NULL DEFAULT 0,
		energy bigint NOT NULL DEFAULT 0 CHECK (energy >= 0),
		active BOOLEAN NOT NULL);

		CREATE TABLE IF NOT EXISTS "dream" (
			id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
			name VARCHAR(64) UNIQUE NOT NULL,
			info TEXT NOT NULL DEFAULT 'empty',
			createdAt timestamp DEFAULT current_timestamp NOT NULL,
			updatedAt timestamp DEFAULT current_timestamp NOT NULL,
			publishAt timestamp NOT NULL DEFAULT current_timestamp ,
			published BOOLEAN NOT NULL DEFAULT false,
			status VARCHAR(32) NOT NULL,
			creater uuid REFERENCES users (id),
			energy bigint NOT NULL DEFAULT 0 CHECK (energy >= 0),
			countG integer NOT NULL DEFAULT 0);

		CREATE TABLE IF NOT EXISTS "dream_location" (
		id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
		locationid uuid REFERENCES location (id) ON DELETE CASCADE,
		dreamid uuid REFERENCES dream (id) ON DELETE CASCADE);

		CREATE TABLE IF NOT EXISTS "project" (
			id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
			name VARCHAR(64) UNIQUE NOT NULL,
			info TEXT NOT NULL DEFAULT 'empty',
			createdAt timestamp DEFAULT current_timestamp NOT NULL,
			updatedAt timestamp DEFAULT current_timestamp NOT NULL,
			publishAt timestamp NOT NULL DEFAULT current_timestamp ,
			published BOOLEAN NOT NULL DEFAULT false,
			status VARCHAR(32) NOT NULL,
			creater uuid REFERENCES users (id),
			energy bigint NOT NULL DEFAULT 0 CHECK (energy >= 0),
			graph TEXT NOT NULL DEFAULT '{}');

		CREATE TABLE IF NOT EXISTS "dream_project" (
			id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
			projectid uuid REFERENCES project (id) ON DELETE CASCADE,
			dreamid uuid REFERENCES dream (id) ON DELETE CASCADE);

		CREATE TABLE IF NOT EXISTS "task_users" (
			id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
			projectid uuid REFERENCES project (id),
			nodeid uuid NOT NULL,
			userid uuid REFERENCES users (id),
			updatedAt timestamp DEFAULT current_timestamp NOT NULL,
			energy bigint NOT NULL DEFAULT 0 CHECK (energy >= 0),
			confirmed BOOLEAN NOT NULL DEFAULT false,
			UNIQUE (nodeid, userid)
			);

		CREATE TABLE IF NOT EXISTS "task_process" (
			id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
			projectid uuid REFERENCES project (id),
			nodeid uuid UNIQUE NOT NULL,
			createdAt timestamp DEFAULT current_timestamp NOT NULL,
			updatedAt timestamp DEFAULT current_timestamp NOT NULL,
			exec_userid uuid REFERENCES users (id),
			inspectors_total bigint NOT NULL DEFAULT 0 CHECK (inspectors_total >= 0),
			inspectors_confirmed bigint NOT NULL DEFAULT 0 CHECK (inspectors_confirmed >= 0),
			energy_total bigint NOT NULL DEFAULT 0 CHECK (inspectors_confirmed >= 0),
			leadtime bigint NOT NULL DEFAULT 0 CHECK (inspectors_confirmed >= 0),
			status VARCHAR(64) NOT NULL DEFAULT 'created',
			completed BOOLEAN NOT NULL DEFAULT false);
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
