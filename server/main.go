package main

import (
    "database/sql"
    "fmt"
    "log"
    "net/http"
    "os"

    _ "github.com/lib/pq"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!")
}

func main() {
 
    psqlInfo := "host=db port=5432 user=user1 dbname=mydatabase password=password123 sslmode=disable"

    // Open the connection
    db, err := sql.Open("postgres", psqlInfo)
    if err != nil {
        log.Fatalf("Error opening database: %q", err)
    }
    defer db.Close()

    // Check the connection
    err = db.Ping()
    if err != nil {
        log.Fatalf("Error connecting to the database: %q", err)
    }

    fmt.Println("Successfully connected to the database")

    http.HandleFunc("/", helloWorld)
    fmt.Println("Server starting on port 8080...")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Printf("Error starting server: %s\n", err)
    }
}