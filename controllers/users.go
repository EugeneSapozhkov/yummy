package controllers

import (
	"database/sql"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"net/http"
	"yummyGo/responders"
	"yummyGo/validators"
)

type User struct {
	ID    int `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
	Created string `json:"created,omitempty"`
}

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}


func GetUsers(db *sql.DB) http.HandlerFunc {
	fn := func(w http.ResponseWriter, req *http.Request) {

		rows, err := db.Query("SELECT * FROM users")
		checkErr(err)
		defer rows.Close()

		var users []User
		for rows.Next() {
			var r User
			err = rows.Scan(&r.ID, &r.Name, &r.Email, &r.Created)
			checkErr(err)
			users = append(users, r)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
	}

	return http.HandlerFunc(fn)
}

func GetUserById(db *sql.DB) http.HandlerFunc {
	fn := func(w http.ResponseWriter, req *http.Request) {
		var user User
		userId := mux.Vars(req)["id"]
		row := db.QueryRow("SELECT id, name, email, created FROM users WHERE id = ?", userId)

		switch err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Created); err {
		case sql.ErrNoRows:
			responders.Error(w, "User does not exist!", http.StatusNotAcceptable)
		case nil:
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(user)
		default:
			panic(err)
		}
	}

	return http.HandlerFunc(fn)
}


// Post new user
func PostUser(db *sql.DB) http.HandlerFunc {
	fn := func(w http.ResponseWriter, req *http.Request) {
		var user User
		err := json.NewDecoder(req.Body).Decode(&user)

		if err := validators.NameValid(user.Name); err != nil {
			panic(err)
		}

		if err := validators.EmailValid(user.Email); err != nil {
			panic(err)
		}

		name := user.Name
		email := user.Email

		var matchedEmail string
		row := db.QueryRow("SELECT email FROM users WHERE email = ?", email)
		err = row.Scan(&matchedEmail)
		checkErr(err)

		if matchedEmail != email {
			insert, err := db.Prepare("INSERT INTO users (name, email) VALUES (?,?)")
			checkErr(err)
			insert.Exec(name, email)

			w.Header().Set("Content-Type", "application/json")

			response := map[string]string{"message": "user was created", "status": "ok"}
			json.NewEncoder(w).Encode(response)
		} else {
			responders.Error(w, "email is already exist", http.StatusNotAcceptable)
		}
	}

	return http.HandlerFunc(fn)
}


// TODO
// Update user
func UpdateUser(db *sql.DB) http.HandlerFunc {
	fn := func(w http.ResponseWriter, req *http.Request) {
		var user User
		userId := mux.Vars(req)["id"]

		err := json.NewDecoder(req.Body).Decode(&user)
		checkErr(err)

		name := user.Name
		email := user.Email

		if err := validators.NameValid(user.Name); err != nil {
			panic(err)
		}

		if err := validators.EmailValid(user.Email); err != nil {
			panic(err)
		}

		_, err = db.Exec("UPDATE users SET name = ?, email = ? WHERE id = ?;", name, email, userId)
		checkErr(err)

		responders.Success(w,"user was updated", http.StatusOK)
	}
	return http.HandlerFunc(fn)
}