package main

import (
	"net/http"
	"os"
)

func indexhandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello World</h1"))
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexhandler)
	http.ListenAndServe(":"+port, mux)
}
