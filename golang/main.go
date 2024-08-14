package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
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

const (
	host = "database"
	port = 5432
)

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
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")

	fmt.Println("POSTGRES_USER:", user)
	fmt.Println("POSTGRES_PASSWORD:", password)
	fmt.Println("POSTGRES_DB:", dbname)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("DB successfully connected!")

	// Проверка существования таблицы
	var tableExists bool
	err = db.QueryRow("SELECT EXISTS (SELECT FROM information_schema.tables WHERE table_name = 'users')").Scan(&tableExists)
	if err != nil {
		log.Fatal(err)
	}

	if !tableExists {
		log.Fatal("Table 'users' does not exist")
	}

	rows, err := db.Query("SELECT id, name FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, Name: %s\n", id, name)
	}

	// Обработка ошибок после завершения итерации
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/api", handleRequest)
	http.HandleFunc("/api/user", handleUserExists)

	fmt.Println("Starting server on :1234")
	log.Fatal(http.ListenAndServe(":1234", nil))
}
