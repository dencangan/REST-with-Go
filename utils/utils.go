package utils

import (
	"fmt"
	"encoding/json"
	"net/http"
)

// Build message
func Message(status bool, message string) (map[string]interface{}) {
	x := map[string]interface{} {"status": status, "message": message}
	fmt.Println(x)
	return x
}

// Build json messages and return a json response
func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
