package main

import (
	"html/template"
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
	// Try to read template file 
	ts, err := template.ParseFiles("./ui/html/pages/home.html")
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

func transactionView(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(r.URL.Query().Get("id"))
    if err != nil || id < 1 {
        http.NotFound(w, r)
        return
    }

    ts, err := template.ParseFiles("./ui/html/pages/view.html")
    if err != nil {
        log.Print(err.Error())
        http.Error(w, "Internal Server Error", 500)
        return
    }

    // Create data to pass to template
    data := struct {
        ID int
    }{
        ID: id,
    }

    err = ts.Execute(w, data)
    if err != nil {
        log.Print(err.Error())
        http.Error(w, "Internal Server Error", 500)
    }
}

// Add a transactionCreate handler function
func transactionCreate(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        w.Header().Set("Allow", http.MethodPost)
        
        // Add template rendering for GET requests to show the form
        ts, err := template.ParseFiles("./ui/html/pages/create.html")
        if err != nil {
            log.Print(err.Error())
            http.Error(w, "Internal Server Error", 500)
            return
        }

        err = ts.Execute(w, nil)
        if err != nil {
            log.Print(err.Error())
            http.Error(w, "Internal Server Error", 500)
        }
        return
    }

    // Handle POST request here
    w.Write([]byte("Create a new transaction"))
}

func transactionList(w http.ResponseWriter, r *http.Request) {
    ts, err := template.ParseFiles("./ui/html/pages/list.html")
    if err != nil {
        log.Print(err.Error())
        http.Error(w, "Internal Server Error", 500)
        return
    }

    err = ts.Execute(w, nil)
    if err != nil {
        log.Print(err.Error())
        http.Error(w, "Internal Server Error", 500)
    }
}