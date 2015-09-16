package httpres

import (
	"encoding/json"
	"log"
	"net/http"
)

func Response(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Println(err)
	}
}

func ResponseFile(w http.ResponseWriter, r *http.Request, path string) {
	http.ServeFile(w, r, path)
}
