package database

import (
	"clarifiveassignment/config"
	"clarifiveassignment/models"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	log "github.com/sirupsen/logrus"
)

type DBConnectionInterface interface {
	// DBConnect() *sql.DB
	InitDB() (*gorm.DB, error)
}

type Connection struct {
	DBConfig *config.DatabaseConfig
}

func DBConnection(DBConfig *config.DatabaseConfig) DBConnectionInterface {
	return Connection{DBConfig}
}

var db *gorm.DB

func (con Connection) InitDB() (*gorm.DB, error) {

	dbConfigs := con.DBConfig
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=True", dbConfigs.User, dbConfigs.Password, dbConfigs.Host, dbConfigs.DBName)
	db, err := gorm.Open("mysql", dataSourceName)

	if err != nil {
		log.Error("Failed to open db connection")
		return db, err
	}

	// Migration to create tables for family schema
	db.AutoMigrate(&models.Member{})

	return db, err
}
