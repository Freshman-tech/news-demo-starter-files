package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func indexhandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello World</h1"))
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", indexhandler)
	http.ListenAndServe(":"+port, mux)
}
