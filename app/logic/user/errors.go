package user

import "errors"

var (
	ErrEmptyUidField = errors.New(" logic: empty uid field")
	ErrUserAlreadyIsArchived = errors.New("logic: user already is archived")
	ErrUserNotFound = errors.New("logic: user not found")
	ErrUserWithSameNumberExist = errors.New("logic: user with the same number already exists")
	ErrUserIsNotArchived = errors.New("logic: user is not archived")
)
