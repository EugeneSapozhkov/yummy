package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"yummyGo/utils"
)

var db *sql.DB

func main() {
	err := godotenv.Load()

	dbConfig := fmt.Sprintf("%s@/%s",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_NAME"))

	db, err := sql.Open("mysql", dbConfig)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	router := utils.RouterInit(db)
	log.Fatal(http.ListenAndServe(":8080", router))
}
