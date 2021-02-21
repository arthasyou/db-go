package mysql

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql" // very important mysql need this
)

var cli *sql.DB

// Connect to mysql
func Connect(addr string, port uint16, user string, password string, dbname string) {
	url := user + ":" + password + "@/" + dbname
	cli, err := sql.Open("mysql", url)
	if err != nil {
		panic(err)
	}
	cli.SetConnMaxLifetime(time.Minute * 3)
	cli.SetMaxOpenConns(10)
	cli.SetMaxIdleConns(10)
}

// Close mysql connection
func Close() {
	cli.Close()
}

// Query database
func Query(query string, args ...interface{}) (*sql.Rows, error) {
	return cli.Query(query, args)
}

// Prepare database
func Prepare(query string) (*sql.Stmt, error) {
	return cli.Prepare(query)
}
