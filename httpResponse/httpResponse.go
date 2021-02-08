package httpResponse

import (
	"encoding/json"
	"net/http"
)

func ResponseMessage(w http.ResponseWriter, code int, message string) {
	Response(w, code, map[string]string{"message": message})
}

func ResponseData(w http.ResponseWriter, code int, payload interface{}) {
	Response(w, code, map[string]interface{}{"data": payload})
}

func Response(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
