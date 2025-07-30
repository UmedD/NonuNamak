package service

import (
	"NonuNamak/internal/model"
	"NonuNamak/pkg/database"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(name, email, password string) (*model.User, error) {
	if name == "" || email == "" || password == "" {
		return nil, errors.New("имя, email и пароль не могут быть пустыми")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
		Role:    "user",
	}

	if err := database.DB.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
