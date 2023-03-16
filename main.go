package main

///we will come to html and css part later, the objective here is to learn go
import (
	"html/template"
	"log"
	"net/http"
	"net/url"
	"os"

	//
	"github.com/joho/godotenv"
)

var tpl = template.Must(template.ParseFiles("index.html"))

func indexhandler(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("<h1>Hello World</h1"))
	tpl.Execute(w, nil)
}
func searchhandler(w http.ResponseWriter, r *http.Request) {
	u, err := url.Parse(r.URL.String())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	params := u.Query()
	searchQuery := params.Get("q")
	page := params.Get("page")
	if page == "" {
		page = "1"
	}
	fmt.println("Search query is: ", searchQuery)
	fmt.println("Page is: ", page)
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
	fs := http.FileServer(http.Dir("assets"))
	mux := http.NewServeMux()
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
	mux.HandleFunc("/search", searchhandler)
	mux.HandleFunc("/", indexhandler)
	http.ListenAndServe(":"+port, mux)
}
