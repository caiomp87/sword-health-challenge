package db

import (
	"database/sql"

	"github.com/caiomp87/sword-health-challenge/interfaces"
)

type databaseService struct {
	driver           string
	connectionString string
}

func NewDatabaseService(driver, connectionString string) interfaces.IDatabase {
	return &databaseService{
		driver,
		connectionString,
	}
}

func (d *databaseService) Connect() (*sql.DB, error) {
	db, err := sql.Open(d.driver, d.connectionString)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func (d *databaseService) Disconnect(db *sql.DB) error {
	return db.Close()
}
