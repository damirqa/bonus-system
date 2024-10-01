package repository

type UserRepositoryInterface interface {
	Insert(login, hashPassword string) error
}
