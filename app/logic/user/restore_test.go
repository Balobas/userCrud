package user

import (
	"github.com/pkg/errors"
	"testing"
	"userCrud/app/data"
)

func TestUserUseCases_RestoreUser(t *testing.T) {
	// Good case.
	userRepo := UserRepoMock{
		GetByUidUserToReturn:  data.User{IsArchived: true},
		GetByUidErrorToReturn: nil,
		UpdateErrorToReturn:   nil,
	}

	useCases := NewUserUseCases(userRepo)

	if err := useCases.RestoreUser("uid"); err != nil {
		t.Error(err)
		t.FailNow()
	} else {
		t.Log("user successfully restored")
	}

	// Bad case.

	// user already is archived
	userRepo.GetByUidUserToReturn = data.User{IsArchived: false}
	useCases = NewUserUseCases(userRepo)

	if err := useCases.RestoreUser("uid"); err == nil {
		t.Error("expected error: user already restored")
		t.FailNow()
	} else {
		t.Log(err)
	}

	// user not found
	userRepo.GetByUidErrorToReturn = errors.New("user not found")
	useCases = NewUserUseCases(userRepo)

	if err := useCases.RestoreUser("uid"); err == nil {
		t.Error("expected error: user not found")
		t.FailNow()
	} else {
		t.Log(err)
	}

}
