package user

import (
	"testing"
	"userCrud/app/data"
)

func TestUserUseCases_CreateUser(t *testing.T) {
	userRepo := UserRepoMock{
		GetByFIOUsersToReturn: []data.User{},
		GetByFIOErrorToReturn: nil,
		CreateErrorToReturn:   nil,
	}

	// Good case.
	testUser := data.User{
		UID:             "",
		LastName:        "Balobanov",
		Name:            "Vladislav",
		Patronymic:      "Valerievich",
		CreatedDateTime: 0,
		UpdatedDateTime: 0,
		IsArchived:      false,
	}

	useCases := NewUserUseCases(userRepo)

	if uid, err := useCases.CreateUser(testUser); err != nil {
		t.Error(err)
		t.FailNow()
	} else {
		t.Logf("user successfuly created with uid %s", uid)
	}

	// пустое поле с отчеством
	testUser.Patronymic = ""
	if uid, err := useCases.CreateUser(testUser); err != nil {
		t.Error(err)
		t.FailNow()
	} else {
		t.Logf("user with empty patronymic successfully created with uid %s", uid)
	}

	// есть пользователь с такими же ФИО, но с другим номером телефона
	userRepo = UserRepoMock{
		GetByFIOUsersToReturn: []data.User{{
			UID:             "",
			LastName:        "",
			Name:            "",
			Patronymic:      "",
			CreatedDateTime: 0,
			UpdatedDateTime: 0,
			IsArchived:      false,
		}},
		GetByFIOErrorToReturn: nil,
		GetByPhoneNumberErrorToReturn: data.ErrorNoObjectsFound,
		CreateErrorToReturn:   nil,
	}

	useCases = NewUserUseCases(userRepo)
	testUser.PhoneNumber = "897192391"
	if uid, err := useCases.CreateUser(testUser); err != nil {
		t.Error(err)
		t.FailNow()
	} else {
		t.Logf("user with same FIO but different phone number successfully created with uid %s", uid)
	}

	// Bad case.

	// пустое поле с именем
	userRepo = UserRepoMock{
		GetByFIOUsersToReturn: []data.User{},
		GetByFIOErrorToReturn: nil,
		CreateErrorToReturn:   nil,
	}
	testUser.Name = ""

	if _, err := useCases.CreateUser(testUser); err == nil {
		t.Error("user with empty name was created. Error expected")
		t.FailNow()
	} else {
		t.Log(err)
	}

	// Пустое поле с фамилией
	testUser.Name = "name"
	testUser.LastName = ""

	if _, err := useCases.CreateUser(testUser); err == nil {
		t.Error("user with empty lastName was created. Error expected")
		t.FailNow()
	} else {
		t.Log(err)
	}

	// Не указан номер телефона и пользователь с такими же ФИО есть в базе
	userRepo = UserRepoMock{
		GetByFIOUsersToReturn: []data.User{{
			UID:             "",
			LastName:        "",
			Name:            "",
			Patronymic:      "",
			CreatedDateTime: 0,
			UpdatedDateTime: 0,
			IsArchived:      false,
		}},
		GetByFIOErrorToReturn: nil,
		CreateErrorToReturn:   nil,
	}

	useCases = NewUserUseCases(userRepo)
	testUser.LastName = "lastName"
	if _, err := useCases.CreateUser(testUser); err == nil {
		t.Error("user was created. Error expected: user with the same FIO already exist")
		t.FailNow()
	} else {
		t.Log(err)
	}
}
