package database

import (
	"go-url-shortener/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type SQLClient struct {
	Client *gorm.DB
}

// make sqlite connection and return that db instance
func ConnectSQLite() (*SQLClient, error) {
	db, err := gorm.Open(sqlite.Open("urls.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &SQLClient{
		Client: db,
	}, nil
}

// to run migrations
func RunMigrations(db *SQLClient) {
	db.Client.AutoMigrate(&model.URL{})
}
