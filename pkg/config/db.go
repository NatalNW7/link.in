package config

import (
	"github.com/NatalNW7/cliqtree/internal/cliqtree/schemas"
	"gorm.io/gorm"
)


func initDatabase(env string) (*gorm.DB, error){
	var db *gorm.DB
	var err error
	logger := GetLogger("database")

	if env == "local" {
		db, err = initSQLite()
		if err != nil {
			return nil, err
		}
		logger.Debug("using SQLite")
	} else {
		db, err = initPostgreSQL()
		if err != nil {
			return nil, err
		}
		logger.Debug("using PostgreSQL")
	}

	err = db.AutoMigrate(&schemas.Link{})
	if  err != nil {
		logger.Errorf("Error to automigrate schema: %v", err)
		return nil, err
	}

	return db, nil
}

type DatabaseInfo struct {
	Name string `json:"name"`
	Version any  `json:"version"`
	MaxConnections any `json:"max_connections"`
}

func DbInfo(env string) any {
	if env == "local" {
		return sqliteInfo()
	}
	
	return postgresInfo()
}