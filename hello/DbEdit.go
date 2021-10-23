package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "12345"
	dbname   = "first_db"
)

func registrationDb(Id int, ParentId int) {
	KeyWord := RandStringBytes()
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	defer db.Close()

	insertStmp := fmt.Sprintf(`insert into "Users" ("Id", "ParentId", "KeyWords") values (%d, %d, '{%s}')`, Id, ParentId, KeyWord)
	_, e := db.Exec(insertStmp)
	CheckError(e)
	fmt.Println("ALL IS OK!")
}
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
