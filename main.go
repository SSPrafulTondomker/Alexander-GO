package main

import (
"encoding/json"
    "fmt"
    "log"
    "net/http"
//     "math/rand"
//     "strconv"
    "github.com/gorilla/mux"
)

// Book struct represents
type Book struct {
    ID string `json:"id"`
    Isbn string `json:"isbn"`
    Title string `json:"title"`
    Author *Author `json:"author"`
}

type Author struct {
    Firstname string `json:"first_name"`
    Lastname string `json:"last_name"`
}

var books []Book

func getBooksHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w)
    json.NewEncoder(w).Encode(books)
}

func getBookIdHandler(w http.ResponseWriter, r *http.Request) {
}

func createBooksHandler(w http.ResponseWriter, r *http.Request) {
}

func updateBooksHandler(w http.ResponseWriter, r *http.Request) {
}

func deleteBooksHandler(w http.ResponseWriter, r *http.Request) {
}
func main() {

    //init router
    request := mux.NewRouter()

    books = append(books, Book{ID: "1", Isbn: "438227", Title: "Book One", Author: &Author{Firstname: "John", Lastname: "Doe"}})
    books = append(books, Book{ID: "2", Isbn: "454555", Title: "Book Two", Author: &Author{Firstname: "Steve", Lastname: "Smith"}})

    //route handler
    request.HandleFunc("/api/books", getBooksHandler).Methods("GET")
    request.HandleFunc("/api/books/{id}", getBooksHandler).Methods("GET")
    request.HandleFunc("/api/books", createBooksHandler).Methods("POST")
    request.HandleFunc("/api/books/{id}", getBooksHandler).Methods("PUT")
    request.HandleFunc("/api/books", getBooksHandler).Methods("Delete")

    err := http.ListenAndServe(":8000", request)

    if (err != nil) {
        log.Fatal(err, "Unable to listen on")
    }else {
        log.Println("server started on : 8000")
    }

    fmt.Println("Hello go");
}