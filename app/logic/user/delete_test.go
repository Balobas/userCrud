package user

import (
	"errors"
	"testing"
)

func TestUserUseCases_DeleteUser(t *testing.T) {
	userRepo := UserRepoMock{
		GetByUidErrorToReturn: nil,
		DeleteErrorToReturn:   nil,
	}

	// Good case
	useCase := NewUserUseCases(userRepo)
	if err := useCase.DeleteUser("uid"); err != nil {
		t.Error(err)
		t.FailNow()
	} else {
		t.Log("user was successfully deleted")
	}

	// Bad Case
	userRepo.GetByUidErrorToReturn = errors.New("user not found")
	useCase = NewUserUseCases(userRepo)
	if err := useCase.DeleteUser("uid"); err == nil {
		t.Error("Expected error user not found")
		t.FailNow()
	} else {
		t.Log(err)
	}
}
