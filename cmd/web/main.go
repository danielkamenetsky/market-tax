package main

import (
	"log"
	"net/http"
)
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

