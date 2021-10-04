package tests

import (
	"context"
	"log"
	"userCrud/app/netInterfaces/interfaces"
)

func(t *UserCrudServerTester) testCreateUser() {
	log.Println("Create user test")

	// Good case
	response, err := t.Client.CreateUser(context.Background(), &interfaces.User{
		Uid:             "",
		LastName:        "Baloban",
		Name:            "Petr",
		Patronymic:      "Valerievich",
		PhoneNumber:     "99991213123",
		CreatedDateTime: nil,
		UpdatedDateTime: nil,
		IsArchived:      false,
	})
	if err != nil {
		log.Fatalf("Error when calling CreateUser: %s", err)
	}
	log.Printf("Response from server: %s\n", response.Uid)
	t.StoredUids[response.Uid] = true

	// empty patronymic
	response, err = t.Client.CreateUser(context.Background(), &interfaces.User{
		Uid:             "",
		LastName:        "last",
		Name:            "name",
		Patronymic:      "",
		PhoneNumber:     "911211312",
		CreatedDateTime: nil,
		UpdatedDateTime: nil,
		IsArchived:      false,
	})
	if err != nil {
		log.Fatalf("Error when calling CreateUser: %s", err)
	}
	t.StoredUids[response.Uid] = true
	log.Printf("Response from server: %s\n", response.Uid)

	// Bad case
	// empty name
	response, err = t.Client.CreateUser(context.Background(), &interfaces.User{
		Uid:             "",
		LastName:        "last",
		Name:            "",
		Patronymic:      "",
		PhoneNumber:     "123123123",
		CreatedDateTime: nil,
		UpdatedDateTime: nil,
		IsArchived:      false,
	})
	if err == nil {
		t.StoredUids[response.Uid] = true
		log.Fatal("expected error: empty name")
	} else {
		log.Printf("user with empty name wasnt created: %s\n", err)
	}

	// empty last name
	response, err = t.Client.CreateUser(context.Background(), &interfaces.User{
		Uid:             "",
		LastName:        "",
		Name:            "name",
		Patronymic:      "",
		CreatedDateTime: nil,
		UpdatedDateTime: nil,
		IsArchived:      false,
	})
	if err == nil {
		t.StoredUids[response.Uid] = true
		log.Fatal("expected error: empty last name")
	} else {
		log.Printf("user with empty last name wasnt created: %s\n", err)
	}

	// user with the same FIO already exist
	response, err = t.Client.CreateUser(context.Background(), &interfaces.User{
		Uid:             "",
		LastName:        "Baloban",
		Name:            "Petr",
		Patronymic:      "Valerievich",
		CreatedDateTime: nil,
		UpdatedDateTime: nil,
		IsArchived:      false,
	})
	if err == nil {
		t.StoredUids[response.Uid] = true
		log.Fatal("error expected: user with the same FIO already exist")
	} else {
		log.Printf("user with same FIO wasnt created: %s\n", err)
	}

	log.Println("Create user test passed\n\n")
}
