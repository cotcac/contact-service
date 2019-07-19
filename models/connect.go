package models

import (
	"database/sql"
	"encoding/json"
	"os"
	"time"
)

// Config is...
type Config struct {
	Database struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Db       string `json:"db"`
	} `json:"database"`
}

// LoadConfiguration is ...
func LoadConfiguration(filename string) (Config, error) {
	var config Config
	configFile, err := os.Open(filename)
	defer configFile.Close()
	if err != nil {
		return config, err
	}
	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)
	return config, err

}

// DBConn will contact to mysql
func DBConn() (db *sql.DB) {
	config, err := LoadConfiguration("config.json")
	if err != nil {
		panic(err.Error())
	}
	dbDriver := "mysql"
	dbUser := config.Database.Username
	dbPass := config.Database.Password
	dbName := config.Database.Db
	db, err = sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	db.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	db.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	db.SetConnMaxLifetime(time.Hour)
	return db
}
