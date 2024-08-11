package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type RequestData struct {
	Name string `json:"name"`
	Key  string `json:"id"`
}

type ResponseData struct {
	Message string `json:"message"`
}

type UserExistsResponse struct {
	Exists bool `json:"exists"`
}

var users = map[string]string{
	"Dmitry": "12345", // Заглушка с одним пользователем
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var reqData RequestData
	err := json.NewDecoder(r.Body).Decode(&reqData)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	response := ResponseData{
		Message: fmt.Sprintf("Hello, %s! Your pass is %s.", reqData.Name, reqData.Key),
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func handleUserExists(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "Missing name parameter", http.StatusBadRequest)
		return
	}

	exists := false
	if _, ok := users[name]; ok {
		exists = true
	}

	response := UserExistsResponse{
		Exists: exists,
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/api", handleRequest)
	http.HandleFunc("/api/user", handleUserExists)

	fmt.Println("Starting server on :1234")
	log.Fatal(http.ListenAndServe(":1234", nil))
}
