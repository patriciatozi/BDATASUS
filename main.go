package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)


func ConnectDatabase() *sql.DB {
	db, err := sql.Open("mysql", "root:123456@/formularioEletronico")
	if err != nil {
		panic(err.Error())
	}
	return db
}

func main() {
	db := ConnectDatabase()
	defer db.Close()

	InitStatements(db)
	defer CleanStatements()

	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))

}
