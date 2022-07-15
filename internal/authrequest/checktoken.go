package authrequest

import (
	"encoding/json"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/p2034/HACK-2022-DREAM-DAIRY/internal/database"
)

type checktoken_req struct {
	Token  string `json:"token"`
	Userid int    `json:"userid"`
}

type checktoken_res struct {
	Error []string `json:"error"`
}

func CheckTokenRequest(w http.ResponseWriter, r *http.Request) {
	var res checktoken_res
	// check method
	if r.Method != http.MethodGet {
		w.WriteHeader(400)
		res.Error = append(res.Error, "Wrong method")
		return
	}
	// get params
	var body logout_req
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		w.WriteHeader(400)
		res.Error = append(res.Error, "Can not get login body")
		return
	}

	// delete token
	db := database.OpenDB()
	defer db.Close()
	if !database.Token_find(db, body.Token, body.Userid) {
		w.WriteHeader(400)
		res.Error = append(res.Error, "Token does not exist")
	}

	// send
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	jsonResp, _ := json.Marshal(res)
	w.Write(jsonResp)
}
