package core

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

type Database struct {
	Name     string
	Host     string
	User     string
	Password string
	Port     string
}

func NewDatabase() *Database {
	return &Database{
		Name:     "cooking_server",
		Host:     "localhost",
		User:     "root",
		Password: "p@ssWoRD",
		Port:     "3306",
	}
}

// ConnectDB to connect database
func (db *Database) Connect() (*gorm.DB, error) {
	newDB, err := gorm.Open("mysql",
		fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local&multiStatements=True&loc=UTC",
			db.User, db.Password, db.Host, db.Port, db.Name,
		))
	if err != nil {
		log.Println("DB error ===> ", err)
		return nil, err
	}
	newDB.LogMode(true)
	return newDB, nil
}

func (db *Database) PingDB() error {
	_, err := db.Connect()
	return err
}
