package service

type BaseUserServiceInterface interface {
	RegisterUser(login, password string) error
}
