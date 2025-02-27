package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {
    // Create a new Gin router
    router := gin.Default()

    // Middleware to handle CORS
    router.Use(func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(http.StatusOK)
            return
        }
        c.Next()
    })

    // Routes
    router.GET("/register-page", registerUser)
    router.POST("/api/register", registerUser)
    router.POST("/api/login", loginUser)
    router.POST("/api/forgot-password", forgotPassword)

    // Start the server on port 8080
    router.Run(":8080")
}

// Register User
func registerUser(c *gin.Context) {
    var user struct {
        FullName string `json:"fullName"`
        Email    string `json:"email"`
        Mobile   string `json:"mobile"`
        Password string `json:"possword"`
        Role     string `json:"role"`
    }

    // Bind JSON data from the request body
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    // For now, we'll just log the user data (replace this with database logic later)
    c.JSON(http.StatusOK, gin.H{
        "message": "User registered successfully",
        "user":    user,
    })
}

// Login User
func loginUser(c *gin.Context) {
    var credentials struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }

    // Bind JSON data from the request body
    if err := c.ShouldBindJSON(&credentials); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    // For now, we'll just check hardcoded credentials (replace this with database logic later)
    if credentials.Email == "test@example.com" && credentials.Password == "password123" {
        c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
    } else {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
    }
}

// Forgot Password
func forgotPassword(c *gin.Context) {
    var emailData struct {
        Email string `json:"email"`
    }

    // Bind JSON data from the request body
    if err := c.ShouldBindJSON(&emailData); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    // For now, we'll just log the email (replace this with email-sending logic later)
    c.JSON(http.StatusOK, gin.H{"message": "Password reset link sent to " + emailData.Email})
}