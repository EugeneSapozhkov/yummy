package utils

import (
	"database/sql"
	"github.com/gorilla/mux"
	"yummyGo/controllers"
)

func RouterInit(db *sql.DB) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/users", controllers.GetUsers(db)).Methods("GET")
	r.HandleFunc("/user/{id}", controllers.GetUserById(db)).Methods("GET")
	r.HandleFunc("/user", controllers.PostUser(db)).Methods("POST")
	r.HandleFunc("/user/{id}", controllers.UpdateUser(db)).Methods("PATCH")
	return r
}

