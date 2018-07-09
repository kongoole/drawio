package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", Index)
	log.Println("Listening on 8888...")
	http.ListenAndServe(":8888", nil)
}

// Index shows the index page
func Index(w http.ResponseWriter, r *http.Request) {
	fs := http.FileServer(http.Dir("src/main/webapp"))
	http.Handle("/static/", http.StripPrefix("/static", fs))
}
