package user

import (
	"github.com/pkg/errors"
)

func (u *UseCases) ArchiveUser(uid string) error {

	if len(uid) == 0 {
		return ErrEmptyUidField
	}

	user, err := u.UserRepo.GetByUid(uid)
	if err != nil {
		return errors.WithStack(err)
	}

	if user.IsArchived {
		return ErrUserAlreadyIsArchived
	}

	user.ChangeArchivedFlag()

	if err := u.UserRepo.Update(user); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
