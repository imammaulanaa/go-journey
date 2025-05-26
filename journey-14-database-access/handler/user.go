package handler

import (
	"encoding/json"
	"net/http"

	"database-access/config"
	"database-access/model"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query("SELECT id, name, age FROM users")
	if err != nil {
		http.Error(w, "Error fetching users", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var u model.User
		if err := rows.Scan(&u.ID, &u.Name, &u.Age); err != nil {
			http.Error(w, "Error scanning users", http.StatusInternalServerError)
			return
		}
		users = append(users, u)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var u model.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	stmt, err := config.DB.Prepare("INSERT INTO users(name, age) VALUES($1, $2)")
	if err != nil {
		http.Error(w, "Error preparing query", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(u.Name, u.Age)
	if err != nil {
		http.Error(w, "Error inserting user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
