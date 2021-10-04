package tests

import (
	"context"
	"log"
	"userCrud/app/netInterfaces/interfaces"
)

func(t *UserCrudServerTester) testGetUserByUid() {
	log.Println("Get user by uid test")
	response, err := t.Client.CreateUser(context.Background(), &interfaces.User{
		Uid:             "",
		LastName:        "Balo",
		Name:            "Vladi",
		Patronymic:      "Vale",
		CreatedDateTime: nil,
		UpdatedDateTime: nil,
		IsArchived:      false,
	})
	if err != nil {
		log.Fatalf("Error when calling CreateUser: %s", err)
	}
	t.StoredUids[response.Uid] = true
	log.Printf("Response from server: %s\n", response.Uid)

	// Good case
	resp, err := t.Client.GetUserByUID(context.Background(), &interfaces.UserUID{Uid: response.Uid})
	if err != nil {
		log.Fatalf("Error: %s", err)
	} else {
		log.Println(resp.User)
	}

	// Bad case
	resp, err = t.Client.GetUserByUID(context.Background(), &interfaces.UserUID{Uid: "sdkfmsdfmsklmflksmdflk"})
	if err == nil {
		log.Fatalf("expected error user not found")
	} else {
		log.Printf("error as expected: %s\n", err)
	}

	log.Println("Get user by uid test passed\n\n")
}
