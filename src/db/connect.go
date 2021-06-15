package db

import (
	"database/sql"
	"fmt"
)

const (
	host = "db"
	port = 5432
	user = "postgres"
	password = "12340"
	dbname = "marathon"
)


func Connect() *sql.DB {
	conn, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname))
	if err != nil{
		panic(err)
	}
	return conn
}