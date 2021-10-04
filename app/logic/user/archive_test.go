package user

import (
	"testing"
	"userCrud/app/data"
)

func TestUserUseCases_ArchiveUser(t *testing.T) {
	// Good case.
	userRepo := UserRepoMock{
		GetByUidUserToReturn:  data.User{IsArchived: false},
		GetByUidErrorToReturn: nil,
		UpdateErrorToReturn:   nil,
	}

	useCases := NewUserUseCases(userRepo)

	if err := useCases.ArchiveUser("uid"); err != nil {
		t.Error(err)
		t.FailNow()
	} else {
		t.Log("user successfully archived")
	}

	// Bad case.

	// user already is archived
	userRepo.GetByUidUserToReturn = data.User{IsArchived: true}
	useCases = NewUserUseCases(userRepo)

	if err := useCases.ArchiveUser("uid"); err == nil {
		t.Error("expected error: ", ErrUserAlreadyIsArchived)
		t.FailNow()
	} else {
		t.Log(err)
	}

	// user not found
	userRepo.GetByUidErrorToReturn = ErrUserNotFound
	useCases = NewUserUseCases(userRepo)

	if err := useCases.ArchiveUser("uid"); err == nil {
		t.Error("expected error:", ErrUserNotFound)
		t.FailNow()
	} else {
		t.Log(err)
	}

}
