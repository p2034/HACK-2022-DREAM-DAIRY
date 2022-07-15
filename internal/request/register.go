package request

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"

	_ "github.com/lib/pq"

	"github.com/p2034/HACK-2022-DREAM-DAIRY/internal/database"
)

type register_req struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func RegisterRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(400)
		return
	}
	var body register_req
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	check, _ := regexp.MatchString("^[w.-]+@([w-]+.)+[w-]{2,4}$", body.Email)
	if !check {
		w.WriteHeader(400)
		return
	}

	db := database.OpenDB()
	defer db.Close()

	var userid int
	cache := database.Password_cache_gen(body.Password)
	err = db.QueryRow(fmt.Sprintf("INSERT INTO users (username, email, password_salt, password_hash, password_iterations) VALUES "+
		"('%s', '%s', '%s', '%s', %d) RETURNING userid;",
		body.Username, body.Email, cache.Salt, cache.Hash, cache.Iterations)).Scan(&userid)
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(400)
		return
	}

	w.WriteHeader(200)
}
