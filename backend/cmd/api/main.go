package main

import (
    "context"
    "net/http"
    "time"

    "communication-system/backend/internal/api"
    "communication-system/backend/internal/config"

    "github.com/gin-gonic/gin"
    "github.com/rs/zerolog/log"
)

func main() {
    cfg := config.Load()

    router := gin.New()
    router.Use(gin.Recovery())
    router.Use(requestLogger())

    api.RegisterRoutes(router, cfg.JWTSecret)

    addr := cfg.HTTPAddr
    if addr == "" {
        addr = ":8080"
    }

    srv := &http.Server{
        Addr:    addr,
        Handler: router,
    }

    log.Info().Str("addr", addr).Msg("starting API server")
    if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
        log.Fatal().Err(err).Msg("server failed")
    }

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    _ = srv.Shutdown(ctx)
}

func requestLogger() gin.HandlerFunc {
    return func(c *gin.Context) {
        start := time.Now()
        c.Next()
        log.Info().Str("method", c.Request.Method).
            Str("path", c.Request.URL.Path).
            Int("status", c.Writer.Status()).
            Dur("duration", time.Since(start)).
            Msg("request completed")
    }
}

func healthHandler(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func loginHandler(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"token": "dev-token-placeholder"})
}

func websocketHandler(c *gin.Context) {
    c.JSON(http.StatusNotImplemented, gin.H{"message": "WebSocket endpoint not implemented yet"})
}
