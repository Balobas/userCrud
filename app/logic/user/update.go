package user

import (
	"github.com/pkg/errors"
	"userCrud/app/data"
)

func (u *UseCases) UpdateUser(user data.User) error {
	if err := user.Validate(); err != nil {
		return errors.WithStack(err)
	}

	if len(user.UID) == 0 {
		return ErrEmptyUidField
	}

	storedUser, err := u.UserRepo.GetByUid(user.UID)
	if err != nil {
		return errors.WithStack(err)
	}

	user.UpdateFields(storedUser)

	if err := u.UserRepo.Update(user); err != nil {
		return errors.WithStack(err)
	}
	return nil
}
