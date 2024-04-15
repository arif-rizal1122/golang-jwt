package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)





func ParseJSON(r *http.Request, payload any) error {
	// jika request body tidak ada
	if r.Body == nil {
		return fmt.Errorf("missing request body")
	}
	// jika ada maka tangkap dan masukan ke payload
	return json.NewDecoder(r.Body).Decode(payload)
}



func WriteJson(w http.ResponseWriter, status int, v any ) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
    // w adl object yg digunakan untuk menulis respon http kembali ke client
	// v any adl data yg ingin dikirimkan sebagai format json
	return json.NewEncoder(w).Encode(v)
}

 
func WriteError(w http.ResponseWriter, status int, err error)  {
	WriteJson(w, status, map[string]string{"error": err.Error()})
}