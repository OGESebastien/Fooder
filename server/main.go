package main

import (
    "database/sql"
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    _ "github.com/lib/pq"
)

var db *sql.DB

func initDB() {
    var err error
    connStr := "postgres://user:password@db:5432/mydb?sslmode=disable"
    db, err = sql.Open("postgres", connStr)
    if err != nil {
        log.Fatalf("Error connecting to PostgreSQL: %s", err)
    }

    err = db.Ping()
    if err != nil {
        log.Fatalf("Error connecting to PostgreSQL: %s", err)
    }

    log.Println("Connected to PostgreSQL!")
}

func main() {
    initDB()

    http.HandleFunc("/", handler)
	
    log.Println("Server running on http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
