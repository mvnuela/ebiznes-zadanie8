package main

import (
    "net/http"
    "time"

    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
    "github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte("your_secret_key")

type User struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

var users = make(map[string]string)

func main() {
    router := gin.Default()

    // Dodanie obsługi CORS
    router.Use(cors.New(cors.Config{
        AllowAllOrigins:  true,  // Umożliwienie wszystkich źródeł
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge:           12 * time.Hour,
    }))

    router.POST("/register", register)
    router.POST("/login", login)

    auth := router.Group("/auth")
    auth.Use(authMiddleware())
    auth.GET("/hello", hello)

    router.Run(":8080")
}

func register(c *gin.Context) {
    var user User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if _, exists := users[user.Username]; exists {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
        return
    }

    users[user.Username] = user.Password
    c.JSON(http.StatusOK, gin.H{"message": "Registration successful"})
}

func login(c *gin.Context) {
    var user User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if password, exists := users[user.Username]; !exists || password != user.Password {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "username": user.Username,
        "exp":      time.Now().Add(time.Hour * 72).Unix(),
    })

    tokenString, err := token.SignedString(jwtSecret)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func authMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenString := c.GetHeader("Authorization")

        if tokenString == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
            c.Abort()
            return
        }

        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            return jwtSecret, nil
        })

        if err != nil || !token.Valid {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }

        c.Next()
    }
}

func hello(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"message": "hello world"})
}
