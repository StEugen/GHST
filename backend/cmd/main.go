package main

import (
    
    "os"
    "path/filepath"

    "github.com/gin-gonic/gin"
    "github.com/steugen/ghst/backend/api/v1"
    "github.com/steugen/ghst/backend/internal/models"
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

    r := gin.Default()

    r.POST("/api/v1/upload", api.UploadImage(db))

    dir, err := filepath.Abs("./frontend/.next")
    if err != nil {
        panic(err)
    }

    r.Static("/_next", dir)
    r.StaticFile("/", filepath.Join(dir, "index.html"))

    port := os.Getenv("PORT")
    if port == "" {
        port = "9000"
    }
    println("Server listening on port " + port)
    if err := r.Run(":" + port); err != nil {
        panic(err)
    }
}

