package user

import "userCrud/app/data"

func(u *UseCases) GetByPhoneNumber(phone string) (data.User, error) {
	return u.UserRepo.GetByPhoneNumber(phone)
}
