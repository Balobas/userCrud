package user

import "userCrud/app/data"

func (u *UseCases) GetByUid(uid string) (data.User, error) {
	return u.UserRepo.GetByUid(uid)
}
