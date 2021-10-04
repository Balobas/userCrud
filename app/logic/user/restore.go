package user

import (
	"github.com/pkg/errors"
)

func (u *UseCases) RestoreUser(uid string) error {
	if len(uid) == 0 {
		return errors.New("empty uid")
	}

	user, err := u.UserRepo.GetByUid(uid)
	if err != nil {
		return errors.WithStack(err)
	}

	if !user.IsArchived {
		return ErrUserIsNotArchived
	}

	user.ChangeArchivedFlag()

	if err := u.UserRepo.Update(user); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
