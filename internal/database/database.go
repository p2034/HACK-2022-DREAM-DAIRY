package database

import (
	"database/sql"
	"log"
	"math/rand"
	"strings"
)

var Connection_string string

func OpenDB() *sql.DB {
	db, err := sql.Open("postgres", Connection_string)
	if err != nil {
		log.Println(err.Error())
	}
	return db
}

func Random(length uint, chars []rune) string {
	var b strings.Builder

	var i uint
	for i = 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	return b.String()
}
