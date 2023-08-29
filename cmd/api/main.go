package main

import (
	"crypto/rand"
	"database/sql"
	_ "database/sql"
	"encoding/hex"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	db := connectDb()
	defer db.Close()

	mux := http.NewServeMux()
	mux.Handle("/health", &HealthHandler{db: db})
	mux.Handle("/api-keys", &ApiKeyHandler{db: db})

	fmt.Println("starting webserver")
	http.ListenAndServe(":8080", mux)
}

type HealthHandler struct {
	db *sql.DB
}

func (h HealthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	err := h.db.Ping()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("unable to connect to DB"))
		return
	}
	w.WriteHeader(200)
	return
}

type ApiKeyHandler struct {
	db *sql.DB
}

type ApiKeyModel struct {
	Name      string    `json:"name"`
	ApiKey    string    `json:"apiKey"`
	CreatedAt time.Time `json:"createdAt"`
}

type PostApiKeyBody struct {
	Name string `json:"name"`
}

func (h ApiKeyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	switch {
	case r.Method == http.MethodPost:
		var postApiKeyBody PostApiKeyBody
		if err := json.NewDecoder(r.Body).Decode(&postApiKeyBody); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			//ToDo - return error message in body.  Validation library?
			return
		}

		apiKeyModel := ApiKeyModel{
			Name:      postApiKeyBody.Name,
			ApiKey:    NewApiKey(),
			CreatedAt: time.Now(),
		}
		_, err := h.db.Exec(`INSERT INTO api_keys (name, api_key, created_at) VALUES ($1, $2, $3)`, apiKeyModel.Name, apiKeyModel.ApiKey, apiKeyModel.CreatedAt)
		if err != nil {
			w.WriteHeader(500)
			return
		}
		bytes, err := json.Marshal(apiKeyModel)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(bytes)
	}
}

func getEnvOrPanic(envName string) string {
	envValue, exists := os.LookupEnv(envName)
	if exists == false {
		panic(fmt.Errorf("environment variable '%s' does not exist", envName))
	}
	return envValue
}

func connectDb() *sql.DB {
	host := getEnvOrPanic("DB_HOST")
	portEnv := getEnvOrPanic("DB_PORT")
	port, err := strconv.Atoi(portEnv)
	if err != nil {
		panic(err)
	}
	user := getEnvOrPanic("DB_USER")
	password := getEnvOrPanic("DB_PASSWORD")
	dbName := getEnvOrPanic("DB_NAME")

	psqlConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)

	fmt.Println("Connecting to DB")
	db, err := sql.Open("postgres", psqlConn)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected to DB")

	return db
}

func NewApiKey() string {
	b := make([]byte, 40)
	rand.Read(b)
	return hex.EncodeToString(b)
}
