package main

import (
    "github.com/gin-gonic/gin"
    "github.com/jinzhu/gorm"
    _ "github.com/lib/pq"
    "log"
    "net/http"
    "os"
    "time"
)

// User represents a user in the database
type User struct {
    ID        uint      `gorm:"primary_key"`
    Username  string    `gorm:"unique;not null"`
    Email     string    `gorm:"unique;not null"`
    Password  string    `gorm:"not null"`
    CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
    UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

// Post represents a post in the database
type Post struct {
    ID        uint      `gorm:"primary_key"`
    UserID    uint      `gorm:"not null"`
    Title     string    `gorm:"not null"`
    Content   string    `gorm:"not null"`
    CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
    UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

var db *gorm.DB

func initDB() {
    var err error
    dsn := "host=localhost user=youruser password=yourpassword dbname=yourdb port=5432 sslmode=disable"
    db, err = gorm.Open("postgres", dsn)
    if err != nil {
        log.Fatal(err)
    }
    db.AutoMigrate(&User{}, &Post{})
}

func main() {
    initDB()

    r := gin.Default()

    r.GET("/users", func(c *gin.Context) {
        var users []User
        db.Find(&users)
        c.JSON(http.StatusOK, users)
    })

    r.POST("/users", func(c *gin.Context) {
        var user User
        if err := c.ShouldBindJSON(&user); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        db.Create(&user)
        c.JSON(http.StatusOK, user)
    })

    r.GET("/posts", func(c *gin.Context) {
        var posts []Post
        db.Find(&posts)
        c.JSON(http.StatusOK, posts)
    })

    r.POST("/posts", func(c *gin.Context) {
        var post Post
        if err := c.ShouldBindJSON(&post); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        db.Create(&post)
        c.JSON(http.StatusOK, post)
    })

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    r.Run(":" + port)
}
