package main

import (
    "fmt"
    "NonuNamak/pkg/config"
    "NonuNamak/pkg/database"
    "NonuNamak/internal/model"
    "NonuNamak/internal/service"
)

func main() {
    fmt.Println("NonuNamak — backend запускается...")

    config.LoadEnv()
    database.Connect()

    // TODO: Запуск роутера и серверной логики
    err := database.DB.AutoMigrate(&model.User{})
    if err != nil {
        fmt.Printf("Ошибка миграции базы данных: %v\n", err)
        return
    }

    fmt.Println("Миграция базы данных выполнена успешно")

    user, err := service.CreateUser("John Doe", "jon@example.com", "securepassword")
    if err != nil {
       panic("Ошибка создании пользователя" + err.Error())
    }

    fmt.Printf("✅ Новый пользователь: ID=%d, Email=%s\n", user.ID, user.Email)
}
