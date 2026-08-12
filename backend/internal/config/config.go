package config

import (
    "github.com/spf13/viper"
)

type Config struct {
    HTTPAddr   string
    DatabaseURL string
    RedisAddr   string
    S3Endpoint  string
    S3Bucket    string
    JWTSecret   string
}

func Load() Config {
    viper.AutomaticEnv()
    viper.SetDefault("HTTP_ADDR", ":8080")
    viper.SetDefault("DATABASE_URL", "postgres://postgres:postgres@localhost:5432/comm_system?sslmode=disable")
    viper.SetDefault("REDIS_ADDR", "localhost:6379")
    viper.SetDefault("S3_ENDPOINT", "http://localhost:9000")
    viper.SetDefault("S3_BUCKET", "comm-files")
    viper.SetDefault("JWT_SECRET", "change-this-secret")

    return Config{
        HTTPAddr:   viper.GetString("HTTP_ADDR"),
        DatabaseURL: viper.GetString("DATABASE_URL"),
        RedisAddr:   viper.GetString("REDIS_ADDR"),
        S3Endpoint:  viper.GetString("S3_ENDPOINT"),
        S3Bucket:    viper.GetString("S3_BUCKET"),
        JWTSecret:   viper.GetString("JWT_SECRET"),
    }
}
