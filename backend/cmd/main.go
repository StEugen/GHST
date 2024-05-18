package main

import (
    "net/http"
    "os"
    "path/filepath"

    "github.com/steugen/ghst/backend/internal/models"
    "github.com/steugen/ghst/backend/api/v1"

    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

func main() {
    db, err := gorm.Open(sqlite.Open("db/test.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    err = db.AutoMigrate(&models.User{}, &models.Image{})
    if err != nil {
        panic("failed to migrate database")
    }

    dir, err := filepath.Abs("./frontend/.next")
    if err != nil {
        panic(err)
    }

    fileServer := http.FileServer(http.Dir(dir))
    http.Handle("/_next/", http.StripPrefix("/_next/", fileServer))

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, filepath.Join(dir, "index.html"))
    })

    port := os.Getenv("PORT")
    if port == "" {
        port = "9000"
    }
    println("Server listening on port " + port)
    err = http.ListenAndServe(":"+port, nil)
    if err != nil {
        panic(err)
    }
}


