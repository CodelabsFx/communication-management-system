package api

import (
    "net/http"

    "communication-system/backend/internal/auth"
    ws "communication-system/backend/internal/websocket"

    "github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, jwtSecret string) {
    api := router.Group("/api")
    {
        api.POST("/login", loginHandler(jwtSecret))
        api.GET("/health", healthHandler)
        api.GET("/ws", ws.HandleWebSocket)
        api.GET("/user/me", auth.JWTMiddleware(jwtSecret), currentUserHandler)
    }
}

func healthHandler(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func loginHandler(jwtSecret string) gin.HandlerFunc {
    return func(c *gin.Context) {
        var credentials struct {
            Email    string `json:"email"`
            Password string `json:"password"`
        }
        if err := c.ShouldBindJSON(&credentials); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
            return
        }

        token, err := auth.GenerateToken(jwtSecret, "user-123", "company-123", []string{"user"})
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create token"})
            return
        }

        c.JSON(http.StatusOK, gin.H{"token": token})
    }
}

func currentUserHandler(c *gin.Context) {
    claims, _ := c.Get("claims")
    c.JSON(http.StatusOK, gin.H{"user": claims})
}
