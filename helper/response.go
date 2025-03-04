package helper

import (
	"encoding/json"
	"net/http"
)

func ResponseJson(w http.ResponseWriter, code int, payload interface{}){
	response, _ := json.Marshal(payload)
	w.Header().Add("conten-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}