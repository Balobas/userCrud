package user

import "userCrud/app/data"

type UseCases struct {
	UserRepo data.UserRepository
}

func NewUserUseCases(repository data.UserRepository) *UseCases {
	return &UseCases{repository}
}
