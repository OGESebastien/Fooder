package main

import (
    "database/sql"
    "fmt"
    "log"
    "net/http"
    "os"

    _ "github.com/lib/pq"
    "github.com/joho/godotenv"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!")
}

func main() {
    // Load .env file
    err := godotenv.Load()

    if err != nil {
        log.Fatalf("Error loading .env file: %s", err)
    }

    // Get database credentials from environment variables
    dbUser := os.Getenv("POSTGRES_USER")
    dbPassword := os.Getenv("POSTGRES_PASSWORD")
    dbName := os.Getenv("POSTGRES_DB")

    // Connection string
    psqlInfo := fmt.Sprintf("host=db port=5432 user=%s dbname=%s password=%s sslmode=disable", dbUser, dbName, dbPassword)

    // Open the connection
    db, err := sql.Open("postgres", psqlInfo)
    if err != nil {
        log.Fatalf("Error opening database: %s", err)
    }
    defer db.Close()

    // Check the connection
    err = db.Ping()
    if err != nil {
        log.Fatalf("Error connecting to the database: %s", err)
    }

    fmt.Println("Successfully connected to the database")

    http.HandleFunc("/", helloWorld)
    fmt.Println("Server starting on port 8080...")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Printf("Error starting server: %s", err)
    }
}