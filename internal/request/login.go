package request

import (
	"encoding/json"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/p2034/HACK-2022-DREAM-DAIRY/internal/database"
)

type login_req struct {
	Something string `json:"something"`
	Password  string `json:"password"`
}

type login_res struct {
	Token  string `json:"token"`
	Userid int    `json:"userid"`
}

func LoginRequest(w http.ResponseWriter, r *http.Request) {
	var resp login_res
	// check method
	if r.Method != http.MethodGet {
		w.WriteHeader(400)
		return
	}
	// get params
	var body login_req
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	// check password
	db := database.OpenDB()
	defer db.Close()
	var check bool
	check, resp.Userid = database.Password_verify(db, body.Something, body.Password)
	if !check {
		w.WriteHeader(400)
		return
	}

	// gen token
	resp.Token = database.Token_gen(db, resp.Userid)
	if resp.Token == "" {
		w.WriteHeader(500)
		return
	}

	// send
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	jsonResp, _ := json.Marshal(resp)
	w.Write(jsonResp)
}
