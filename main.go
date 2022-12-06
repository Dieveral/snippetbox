package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	log.Println("Server started on http://127.0.0.1:4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Hello from SnippetBox"))
}

func showSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Here you will see note"))
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)

		http.Error(w, fmt.Sprintf("%s method is not allowed", r.Method), http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("Here you can create new note"))
}
