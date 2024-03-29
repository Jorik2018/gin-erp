package handlers

import (
	"net/http"
	"database/sql"
	"encoding/json"

	c "github.com/Jorik2018/go_crud/commons"
)

func GetHandler(db *sql.DB) (func(w http.ResponseWriter, r *http.Request)){
return func(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case http.MethodGet:
			getUsers(w, r, db)
		case http.MethodPost:
			createUser(w, r, db)
		case http.MethodPut:
			updateUser(w, r, db)
		case http.MethodDelete:
			deleteUser(w, r, db)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
}

func getUsers(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	rows, err := db.Query("SELECT uid, name, mail FROM dru_users")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	users := []c.User{}
	for rows.Next() {
		var u c.User
		err := rows.Scan(&u.ID, &u.Name, &u.Email)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		users = append(users, u)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func createUser(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var u c.User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := db.Exec("INSERT INTO users (name, email) VALUES (?, ?)", u.Name, u.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	u.ID = int(id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(u)
}

func updateUser(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var u c.User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = db.Exec("UPDATE users SET name = ?, email = ? WHERE id = ?", u.Name, u.Email, u.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(u)
}

func deleteUser(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var u c.User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = db.Exec("DELETE FROM users WHERE id = ?", u.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
