package main

import (
	"fmt"
	"net/http"
	"strconv"
)

// This is like creating a recipe for what to do when someone comes to your stand.
// When a customer arrives, you give them a friendly "Hello from Snippetbox"
func home(w http.ResponseWriter, r *http.Request) {

    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }
    w.Write([]byte("Market Tax - Welcome"))
}

func transactionView(w http.ResponseWriter, r *http.Request){
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
	return
	}

	fmt.Fprintf(w, "Display a specific transaction with ID %d...", id)
}



// Add a transactionCreate handler function
func transactionCreate(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method Not Allowed"))
		return
	}

	w.Write([]byte("Create a new transaction"))
}

func transactionList(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("List all transactions..."))
}