package user

import "userCrud/app/data"

func (u *UseCases) GetByFIO(name, lastName, patronymic string) ([]data.User, error) {
	return u.UserRepo.GetByFIO(name, lastName, patronymic)
}
