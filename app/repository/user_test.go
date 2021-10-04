package repository

import (
	"fmt"
	"testing"
	"userCrud/app/data"
)

func TestUser_Create(t *testing.T) {
	userRepo, err := GetUserRepository()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if userRepo == nil {
		t.FailNow()
	}

	if err := userRepo.Create(data.User{
		UID:             "uid1",
		LastName:        "last",
		Name:            "name",
		Patronymic:      "patron",
		CreatedDateTime: 0,
		UpdatedDateTime: 0,
		IsArchived:      false,
	}); err != nil {
		t.Error(err)
		t.FailNow()
	}
}

func TestUser_Update(t *testing.T) {
	userRepo, err := GetUserRepository()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if userRepo == nil {
		t.FailNow()
	}

	if err := userRepo.Update(data.User{
		UID:             "uid1",
		LastName:        "loll",
		Name:            "name",
		Patronymic:      "patron",
		CreatedDateTime: 0,
		UpdatedDateTime: 0,
		IsArchived:      false,
	}); err != nil {
		t.Error(err)
		t.FailNow()
	}
}

func TestUser_GetByUid(t *testing.T) {
	userRepo, err := GetUserRepository()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if userRepo == nil {
		t.FailNow()
	}

	user, err := userRepo.GetByUid("uid1")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	fmt.Println(user)
}

func TestUser_GetByPhoneNumber(t *testing.T) {
	userRepo, err := GetUserRepository()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if userRepo == nil {
		t.FailNow()
	}

	user, err := userRepo.GetByPhoneNumber("89123123")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	fmt.Println(user)
}

func TestUser_GetByFIO(t *testing.T) {
	userRepo, err := GetUserRepository()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if userRepo == nil {
		t.FailNow()
	}

	users, err := userRepo.GetByFIO("name", "loll", "patron")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	fmt.Println(users)
}

func TestUser_Delete(t *testing.T) {
	userRepo, err := GetUserRepository()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if userRepo == nil {
		t.FailNow()
	}


	if err := userRepo.Delete("uid1"); err != nil {
		t.Error(err)
		t.FailNow()
	}

}
