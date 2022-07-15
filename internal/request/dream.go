package request

// import (
// 	"encoding/json"
// 	"net/http"
// )

// type dream_post_req struct {
// 	"user"
// }

// type dream_res struct {
// }

// /*
// {
//   "userid": 1234,
//   "token": "egy4983fbi3564342i78...",
//   "dream": {
//     "title": "Some Dream",
//     "description": "This dream is ...",
//     "photos": [
//       "BASE64 encrypted photo",
//       "BASE64 encrypted photo"
//     ]
//   }
// }
// */

// func DreamRequest(w http.ResponseWriter, r *http.Request) {
// 	var resp dream_res
// 	// check method
// 	if r.Method != http.MethodGet {
// 		w.WriteHeader(400)
// 		return
// 	}
// 	// get params
// 	var body dream_req
// 	err := json.NewDecoder(r.Body).Decode(&body)
// 	if err != nil {
// 		w.WriteHeader(400)
// 		return
// 	}

// 	// send
// 	w.WriteHeader(200)
// 	w.Header().Set("Content-Type", "application/json")
// 	jsonResp, _ := json.Marshal(resp)
// 	w.Write(jsonResp)
// }
