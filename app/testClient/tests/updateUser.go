package tests

import (
	"context"
	"log"
	"userCrud/app/netInterfaces/interfaces"
)

func(t *UserCrudServerTester) testUpdateUser() {
	log.Println("Update user test")

	resp, err := t.Client.CreateUser(context.Background(), &interfaces.User{
		Uid:             "",
		LastName:        "B",
		Name:            "V",
		Patronymic:      "V",
		CreatedDateTime: nil,
		UpdatedDateTime: nil,
		IsArchived:      false,
	})
	if err != nil {
		log.Fatalf("Error when calling UpdateUser: %s", err)
	}
	t.StoredUids[resp.Uid] = true

	existedUid := resp.Uid

	// Good case
	_, err = t.Client.UpdateUser(context.Background(), &interfaces.User{
		Uid:             existedUid,
		LastName:        "Bal",
		Name:            "Vlad",
		Patronymic:      "Val",
		CreatedDateTime: nil,
		UpdatedDateTime: nil,
		IsArchived:      false,
	})
	if err != nil {
		log.Fatalf("Error when calling UpdateUser: %s", err)
	}

	// Bad case
	// invalid user fields
	_, err = t.Client.UpdateUser(context.Background(), &interfaces.User{
		Uid:             "",
		LastName:        "last",
		Name:            "",
		Patronymic:      "",
		CreatedDateTime: nil,
		UpdatedDateTime: nil,
		IsArchived:      false,
	})
	if err == nil {
		log.Fatal("expected error: empty name")
	} else {
		log.Printf("user with invalid fields wasnt updated: %s\n", err)
	}

	// user not found
	_, err = t.Client.UpdateUser(context.Background(), &interfaces.User{
		Uid:             "asdasdassdasdassdasdawdasdasdasd",
		LastName:        "Balobanoffff",
		Name:            "Vladislav",
		Patronymic:      "Valerievich",
		CreatedDateTime: nil,
		UpdatedDateTime: nil,
		IsArchived:      false,
	})
	if err == nil {
		log.Fatal("error expected: user not found")
	} else {
		log.Printf("user wasnt updated: %s\n", err)
	}

	log.Println("Update user test passed\n\n")
}
