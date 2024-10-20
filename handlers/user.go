package handlers

import (
	"database/sql"
	"encoding/json"
	"inventory/models"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Message struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func CreateUser(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")

	err := db.QueryRow("select username from users where username=?", username).Scan(&username)

	if err == nil {
		http.Error(w, "server error, unable to username exists", 500)
		return
	}

	hashed, _ := bcrypt.GenerateFromPassword([]byte(r.FormValue("password")), bcrypt.DefaultCost)
	password := string(hashed)

	_, insertErr := db.Exec("insert into users (username, password) values ($1,$2)", username, password)

	if insertErr != nil {
		http.Error(w, "Registration failed", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	msg := Message{
		Status:  "success",
		Message: "User Created",
	}
	json.NewEncoder(w).Encode(msg)
}

func Login(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	var hashedPassword string
	err := db.QueryRow("select id, username, password from users where username=?", user.Username).Scan(&user.ID, &user.Username, &hashedPassword)

	if err != nil {
		http.Error(w, "Unable to login", http.StatusBadRequest)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(user.Password)); err != nil {
		http.Error(w, "Unable to login", http.StatusBadRequest)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":   user.ID,
		"expirate": time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, _ := token.SignedString(jwtSecret)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})

}
