package apirequest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/p2034/HACK-2022-DREAM-DAIRY/internal/database"
)

type dream struct {
	Title        string `json:"title"`
	Descriptiont string `json:"description"`
}

type dream_post_req struct {
	Userid int    `json:"userid"`
	Token  string `json:"token"`
	Dream  dream  `json:"dream"`
}

type dream_post_res struct {
	Error []string `json:"error"`
}

type dream_get_req struct {
	Userid int    `json:"userid"`
	Token  string `json:"token"`
}

type dream_get_res struct {
	Error  []string `json:"error"`
	Dreams []dream  `json:""dreams`
}

func receiveDream(w http.ResponseWriter, r *http.Request) {
	var res dream_post_res
	// get params
	var body dream_post_req
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		w.WriteHeader(400)
		res.Error = append(res.Error, "Can not get body")
		return
	}

	// check token
	db := database.OpenDB()
	defer db.Close()
	if !database.Token_find(db, body.Token, body.Userid) {
		w.WriteHeader(400)
		res.Error = append(res.Error, "Wrong token")
		return
	}

	// set all in the database
	db.Exec(fmt.Sprintf("INSERT INTO dreams (userid, title, description) VALUE (%d, '%s', '%s');",
		body.Userid, body.Dream.Title, body.Dream.Descriptiont))

	// send
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	jsonResp, _ := json.Marshal(res)
	w.Write(jsonResp)
}

func sendDream(w http.ResponseWriter, r *http.Request) {
	var res dream_get_res
	// get params
	var body dream_get_req
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		w.WriteHeader(400)
		res.Error = append(res.Error, "Can not get body")
		return
	}

	// check token
	db := database.OpenDB()
	defer db.Close()
	if !database.Token_find(db, body.Token, body.Userid) {
		w.WriteHeader(400)
		res.Error = append(res.Error, "Wrong token")
		return
	}

	// select dat afrom db
	rows, err := db.Query("SELECT title, descrip FROM dreams WHERE userid = %d;", body.Userid)
	defer rows.Close()
	if err != nil {
		w.WriteHeader(400)
		res.Error = append(res.Error, "Can not query bd")
		return
	}
	for rows.Next() {
		var d dream
		rows.Scan(&d.Title, &d.Descriptiont)
		res.Dreams = append(res.Dreams, d)
	}

	// send
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	jsonResp, _ := json.Marshal(res)
	w.Write(jsonResp)
}

func DreamRequest(w http.ResponseWriter, r *http.Request) {
	// check method
	if r.Method == http.MethodGet {
		sendDream(w, r)
	} else if r.Method == http.MethodPost {
		receiveDream(w, r)
	} else {
		w.WriteHeader(400)
		return
	}
}
