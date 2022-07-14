package cmd

import (
	"fmt"
	"net/http"
	"os"

	"github.com/p2034/HACK-2022-DREAM-DAIRY/internal/database"
	"github.com/p2034/HACK-2022-DREAM-DAIRY/internal/request"
)

func Auth_server() {
	dbhost := os.Getenv("DATABASE_HOST")
	dbport := os.Getenv("DATABASE_PORT")
	dbuser := os.Getenv("DATABASE_USER")
	dbpassword := os.Getenv("DATABASE_PASSWORD")
	dbname := os.Getenv("DATABASE_NAME")

	database.Connection_string = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbhost, dbport, dbuser, dbpassword, dbname)

	fmt.Println("Database connected")
	http.HandleFunc("/register", request.RequestRegister)

	host := ":10002"
	fmt.Println("Server is listening on host: " + host)
	http.ListenAndServe(host, nil)
}
