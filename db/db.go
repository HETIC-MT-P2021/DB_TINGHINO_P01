package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // Go MySQL driver
	"github.com/spf13/viper"
)

var (
	DB *sql.DB
)

// GetDataSourceName returns environment variable for database connection.
func GetDataSourceName() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		viper.GetString("DB_USER"),
		viper.GetString("DB_PASSWORD"),
		viper.GetString("DB_HOST"),
		viper.GetString("DB_PORT"),
		viper.GetString("DB_NAME"),
	)
}

// Connect connect to the database.
func Connect() {
	var err error

	dsn := GetDataSourceName()

	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	if err := DB.Ping(); err != nil {
		panic(err)
	}
}
