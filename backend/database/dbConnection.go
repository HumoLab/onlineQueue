package database

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/perman/backend/models"
)

// ConnVar - connection variable to DB
var ConnVar *sql.DB

// ConnectToDB - connection function to DB
func ConnectToDB() error {
	var (
		dbConfigFilePath = "./configs/databaseConfig.json"
		dbProps          models.DBprops
		err              error
	)

	err = json.Unmarshal(readFromFile(dbConfigFilePath), &dbProps)
	//fmt.Printf("%+v", dbProps)
	if err != nil {
		return errors.New("DBConnection: couldn't parse JSON")
	}
	// teproraty sql connection
	// after we'll use sqlx or gorm
	dbInfo := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", dbProps.User, dbProps.Password, dbProps.Host, dbProps.Port, dbProps.Name, dbProps.SslMode)

	ConnVar, err = sql.Open(dbProps.Driver, dbInfo)
	if err != nil {
		return errors.New("DBConnection: couldn't connect to DB")
	}

	err = ConnVar.Ping()
	if err != nil {
		return errors.New("DBConnection: there no connection to DB")
	}

	return nil
}

// Disconnect from DB
func Disconnect() error {
	err := ConnVar.Close()
	if err != nil {
		return errors.New("DBConnection: couldn't disconnect from DB")
	}

	return nil
}

// how to print out JSON?
// Here: dbProps - JSON object
//-> fmt.Printf("%+v", dbProps)
//result: {Driver:"postgres", User:"postgres", Name:"perman", Password:"postgres", SslMode:"disable"}
//-> fmt.Printf("%#v", dbProps)
//result: models.DBprops{Driver:"postgres", User:"postgres", Name:"perman", Password:"postgres", SslMode:"disable"}
//
//!!!!!!!!!!!!!!!! SSLMODE - often adds to the URL like a query parameter !!!!!!!!!!!!!!!!
