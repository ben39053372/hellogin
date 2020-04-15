package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// DB grom.DB
var DB *gorm.DB

// DBConfig database config
type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

// BuildDBConfig build database config
func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		Host:     "127.0.0.1",
		Port:     3306,
		User:     "root",
		DBName:   "testing",
		Password: "39053372",
	}
	return &dbConfig
}

// DbURL return db url
func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}
