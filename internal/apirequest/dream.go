package apirequest

import (
	"encoding/json"
	"net/http"
)

type dream struct {
	Title       string `json:"title"`
	Description string `json:"description"`
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
	Userid  int    `json:"userid"`
	Token   string `json:"token"`
	Firstid int    `json:"firstid"`
	Lastid  int    `json:"lastid"`
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

	// set all in the database

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
		return
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
