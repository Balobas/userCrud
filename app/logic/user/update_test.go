package user

import (
	"errors"
	"testing"
	"userCrud/app/data"
)

func TestUserUseCases_UpdateUser(t *testing.T) {
	userRepo := UserRepoMock{
		GetByUidUserToReturn:  data.User{},
		GetByUidErrorToReturn: nil,
		UpdateErrorToReturn:   nil,
	}

	// Good case.
	testUser := data.User{
		UID:             "uid1",
		LastName:        "lastName",
		Name:            "name",
		Patronymic:      "",
		CreatedDateTime: 0,
		UpdatedDateTime: 0,
		IsArchived:      false,
	}
	useCases := NewUserUseCases(userRepo)

	if err := useCases.UpdateUser(testUser); err != nil {
		t.Error(err)
		t.FailNow()
	} else {
		t.Log("user was successfully updated")
	}

	// Bad case.

	// invalid user fields
	testUser.Name = ""

	if err := useCases.UpdateUser(testUser); err == nil {
		t.Error("expected error invalid user fields")
		t.FailNow()
	} else {
		t.Log(err)
	}

	// user not found
	testUser.Name = "name"
	userRepo.GetByUidErrorToReturn = errors.New("user not found")
	useCases = NewUserUseCases(userRepo)

	if err := useCases.UpdateUser(testUser); err == nil {
		t.Error("expected error user not found")
		t.FailNow()
	} else {
		t.Log(err)
	}
}
