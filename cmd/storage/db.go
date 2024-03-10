package storage

import (
    "database/sql"
    "fmt"
    "log"
    "os"
    
    "github.com/joho/godotenv"
    _ "github.com/tursodatabase/libsql-client-go/libsql"
)

var db *sql.DB

func InitDB() { 
    err := godotenv.Load()
    if err != nil { 
        log.Fatal("error loading the .env file")
    }
    dbName := os.Getenv("DB_NAME")
    
    url := fmt.Sprintf("libsql://%s-wewantwebsites.turso.io?authToken=%s", dbName, os.Getenv("DB_AUTH_TOKEN"))

    db, err = sql.Open("libsql", url)
    if err != nil { 
        fmt.Fprintf(os.Stderr, "failed to open db %s: %s", url, err)
        log.Fatal("failed to connect to the db")
        os.Exit(1)
    }
    defer db.Close()
    fmt.Println("connected to the db@", dbName)
}
