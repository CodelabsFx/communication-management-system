package main

import (
    "fmt"
    "os"

    "github.com/rs/zerolog/log"
    "github.com/spf13/viper"
)

func main() {
    if err := initConfig(); err != nil {
        fmt.Fprintf(os.Stderr, "configuration error: %v\n", err)
        os.Exit(1)
    }

    log.Info().Msg("starting background worker")
    // Placeholder for worker loop
    select {}
}

func initConfig() error {
    viper.AutomaticEnv()
    viper.SetDefault("DATABASE_URL", "postgres://postgres:postgres@localhost:5432/comm_system?sslmode=disable")
    viper.SetDefault("REDIS_ADDR", "localhost:6379")
    viper.SetDefault("S3_ENDPOINT", "http://localhost:9000")
    viper.SetDefault("S3_BUCKET", "comm-files")
    return nil
}
