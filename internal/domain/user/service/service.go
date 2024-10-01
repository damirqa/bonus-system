package service

import (
	"github.com.damirqa.gophermart/internal/domain/user/repository"
	"github.com.damirqa.gophermart/internal/infrastructure/logging"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repository repository.UserRepositoryInterface
}

func NewUserService(repository repository.UserRepositoryInterface) *UserService {
	return &UserService{repository: repository}
}

func (s *UserService) RegisterUser(login, password string) error {
	hashedPassword, err := hashPassword(password)

	err = s.repository.Insert(login, hashedPassword)
	if err != nil {
		return err
	}

	return nil
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		logging.GetLogger().Error("Failed to hash password: %v", zap.Error(err))
		return "", err
	}

	return string(hashedPassword), nil
}
