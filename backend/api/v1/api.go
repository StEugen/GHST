package api

import (
    "net/http"
    "os"
    "path/filepath"
    "github.com/steugen/ghst/backend/models"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    
)

func UploadImage(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        username := c.PostForm("username")
        tag := c.PostForm("tag")
        digest := c.PostForm("digest")
        file, err := c.FormFile("file")
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "File upload error"})
            return
        }

        var user models.User
        if err := db.Where("username = ?", username).First(&user).Error; err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
            return
        }

        filePath := filepath.Join("uploads", file.Filename)
        if err := c.SaveUploadedFile(file, filePath); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to save file"})
            return
        }

        fileInfo, err := os.Stat(filePath)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to retrieve file info"})
            return
        }

        image := models.Image{
            Name:   file.Filename,
            Tag:    tag,
            Digest: digest,
            Size:   fileInfo.Size(),
            UserID: user.ID,
        }

        if err := db.Create(&image).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to save image record"})
            return
        }

        c.JSON(http.StatusOK, gin.H{"message": "Image uploaded successfully"})
    }
}

