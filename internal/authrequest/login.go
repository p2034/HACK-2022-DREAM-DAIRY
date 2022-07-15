package authrequest

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
	Error  []string `json:"error"`
	Token  string   `json:"token"`
	Userid int      `json:"userid"`
}

func LoginRequest(w http.ResponseWriter, r *http.Request) {
	var res login_res
	// check method
	if r.Method != http.MethodGet {
		w.WriteHeader(400)
		res.Error = append(res.Error, "Wrong method")
		return
	}
	// get params
	var body login_req
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		w.WriteHeader(400)
		res.Error = append(res.Error, "Can not get login body")
		return
	}

	// check password
	db := database.OpenDB()
	defer db.Close()
	var check bool
	check, res.Userid = database.Password_verify(db, body.Something, body.Password)
	if !check {
		w.WriteHeader(400)
		res.Error = append(res.Error, "Wrong password")
		return
	}

	// gen token
	res.Token = database.Token_gen(db, res.Userid)
	if res.Token == "" {
		w.WriteHeader(500)
		res.Error = append(res.Error, "Can not create token")
		return
	}

	// send
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	jsonResp, _ := json.Marshal(res)
	w.Write(jsonResp)
}
