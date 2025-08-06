package service

import (
	"NonuNamak/internal/model"
	"NonuNamak/internal/repository"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

// CreateUser - создание пользователя
func (s *UserService) CreateUser(name, email, password string) (*model.User, error) {
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
		Role:     "user",
	}

	if err := s.repo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

// GetUserByID - получение пользователя по ID
func (s *UserService) GetUserByID(id uint) (*model.User, error) {
	return s.repo.GetByID(id)
}

// GetAllUsers - список всех пользователей
func (s *UserService) GetAllUsers() ([]model.User, error) {
	return s.repo.GetAll()
}

// UpdateUser - обновление данных пользователя
func (s *UserService) UpdateUser(id uint, name, email, password string) (*model.User, error) {
	user, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if name != "" {
		user.Name = name
	}
	if email != "" {
		user.Email = email
	}
	if password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		user.Password = string(hashedPassword)
	}

	if err := s.repo.Update(user); err != nil {
		return nil, err
	}

	return user, nil
}

// DeleteUser - удаление пользователя
func (s *UserService) DeleteUser(id uint) error {
	return s.repo.Delete(id)
}

func (s *UserService) GetUserByEmail(email string) (*model.User, error) {
	return s.repo.GetByEmail(email)
}

func (s *UserService) PatchUser(id uint, updates map[string]interface{}) (*model.User, error) {
	user, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Если обновляется пароль — хэшируем
	if pass, ok := updates["password"]; ok {
		if strPass, ok := pass.(string); ok && strPass != "" {
			hashedPassword, err := bcrypt.GenerateFromPassword([]byte(strPass), bcrypt.DefaultCost)
			if err != nil {
				return nil, err
			}
			updates["password"] = string(hashedPassword)
		} else {
			delete(updates, "password")
		}
	}

	if err := s.repo.UpdatePartial(user.ID, updates); err != nil {
		return nil, err
	}

	// Обновляем модель для возврата
	for k, v := range updates {
		switch k {
		case "name":
			user.Name = v.(string)
		case "email":
			user.Email = v.(string)
		case "role":
			user.Role = v.(string)
		}
	}

	return user, nil
}
