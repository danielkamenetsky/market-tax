package main

import (
	"fmt"
	"log"
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

func main() {
    // This is like creating your menu board.
    // You're making a new blank menu where you can write what drinks you offer
    mux := http.NewServeMux()

    // This is like writing on your menu:
    // "When someone comes to the main counter (/), use the home recipe"
    mux.HandleFunc("/", home)
	mux.HandleFunc("/transaction/", transactionView)
	mux.HandleFunc("/transaction/create", transactionCreate)
	mux.HandleFunc("/transaction/list", transactionList)

    // This is like opening your lemonade stand for business!
    // - ":4000" is like saying "I'm setting up my stand at spot number 4000"
    // - If something goes wrong (like someone took your spot), the program will tell you
    log.Println("Starting server on :4000")
    err := http.ListenAndServe(":4000", mux)
    if err!= nil { log.Fatal(err)}
}