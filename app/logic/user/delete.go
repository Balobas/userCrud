package user

import (
	"github.com/pkg/errors"
)

func (u *UseCases) DeleteUser(uid string) error {
	_, err := u.UserRepo.GetByUid(uid)
	if err != nil {
		return errors.WithStack(err)
	}

	if err := u.UserRepo.Delete(uid); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
