package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("html/main.html", "html/header.html", "html/footer.html")
	if err != nil {
		panic(err)
	}
	t.ExecuteTemplate(w, "main", nil)
}

// ***************************************************************************************************************************************************************************

func StartFunc() {
	rtr := mux.NewRouter()

	http.Handle("/html/", http.StripPrefix("/html/", http.FileServer(http.Dir("/html/"))))

	rtr.HandleFunc("/", HomePage).Methods("GET")

	http.Handle("/", nil)
	http.ListenAndServe(":5500", nil)
}

func main() {
	StartFunc()
}
