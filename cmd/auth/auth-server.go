package auth

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/p2034/HACK-2022-DREAM-DAIRY/internal/authrequest"
	"github.com/p2034/HACK-2022-DREAM-DAIRY/internal/database"
)

func Server(ip string) {
	dbhost := os.Getenv("DATABASE_HOST")
	dbport, _ := strconv.Atoi(os.Getenv("DATABASE_PORT"))
	dbuser := os.Getenv("DATABASE_USER")
	dbpassword := os.Getenv("DATABASE_PASSWORD")
	dbname := os.Getenv("DATABASE_NAME")

	database.Connection_string = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s",
		dbhost, dbport, dbuser, dbpassword, dbname)
	fmt.Println("DB on: " + database.Connection_string)

	fmt.Println("Database connected")
	http.HandleFunc("/register", authrequest.RegisterRequest)
	http.HandleFunc("/login", authrequest.LoginRequest)
	http.HandleFunc("/logout", authrequest.LogoutRequest)
	http.HandleFunc("/checktoken", authrequest.CheckTokenRequest)

	host := ip + ":10000"
	fmt.Println("Server is listening on host: " + host)
	http.ListenAndServe(host, nil)
}
