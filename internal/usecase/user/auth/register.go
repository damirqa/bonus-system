package auth

import (
	"github.com.damirqa.gophermart/internal/domain/user/service"
)

type UserRegisterUseCase struct {
	userService service.BaseUserServiceInterface
}

func NewUserRegisterUseCase(userService service.BaseUserServiceInterface) *UserRegisterUseCase {
	return &UserRegisterUseCase{userService: userService}
}

func (u *UserRegisterUseCase) Register(login, password string) error {
	err := u.userService.RegisterUser(login, password)
	if err != nil {
		return err
	}

	return nil
}
