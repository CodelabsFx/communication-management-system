package auth

import (
    "net/http"
    "strings"

    "github.com/gin-gonic/gin"
)

func JWTMiddleware(secret string) gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing authorization header"})
            return
        }

        tokenParts := strings.SplitN(authHeader, " ", 2)
        if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization format"})
            return
        }

        claims, err := ParseToken(secret, tokenParts[1])
        if err != nil {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
            return
        }

        c.Set("claims", claims)
        c.Next()
    }
}
