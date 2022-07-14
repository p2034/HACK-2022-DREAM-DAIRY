package request

import (
	"encoding/json"
	"net/http"

	"github.com/p2034/AUTH-SERVER-dream-diary-hack-2022/internal/database"
)

type checktoken_req struct {
	Token  string `json:"token"`
	Userid int    `json:"userid"`
}

func CheckTokenRequest(w http.ResponseWriter, r *http.Request) {
	// check method
	if r.Method != http.MethodPost {
		w.WriteHeader(400)
		return
	}
	// get params
	var body logout_req
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	// delete token
	db := database.OpenDB()
	defer db.Close()
	if database.Token_find(db, body.Token, body.Userid) {
		w.WriteHeader(200)
	} else {
		w.WriteHeader(400)
	}
}
