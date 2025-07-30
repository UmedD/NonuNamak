package config

import (
    "github.com/joho/godotenv"
    "log"
    "os"
)

func LoadEnv() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Ошибка загрузки .env файла")
    }
}

func GetEnv(key string) string {
    return os.Getenv(key)
}
