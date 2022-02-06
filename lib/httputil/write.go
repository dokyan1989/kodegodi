package httputil

import (
	"encoding/json"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, v interface{}, code int) {
	w.Header().Set(HeaderContentType, MIMETypeApplicationJSON)
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(v)
}
