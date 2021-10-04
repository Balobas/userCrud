package tests

import (
	"context"
	"log"
	"userCrud/app/netInterfaces/interfaces"
)

func(t *UserCrudServerTester) testArchiveUser() {
	log.Println("Archive user test")

	testUser := &interfaces.User{
		Uid:             "",
		LastName:        "Bob",
		Name:            "John",
		Patronymic:      "Johnovich",
		PhoneNumber:     "999999999",
		CreatedDateTime: nil,
		UpdatedDateTime: nil,
		IsArchived:      false,
	}
	resp, err := t.Client.CreateUser(context.Background(), testUser)
	if err != nil {
		log.Fatalf("Error when calling CreateUser: %s", err)
	}
	t.StoredUids[resp.Uid] = true

	testUser.Uid = resp.Uid

	// Good case
	_, err = t.Client.ArchiveUser(context.Background(), &interfaces.UserUID{Uid:testUser.Uid})
	if err != nil {
		log.Fatalf("Error when calling ArchiveUser(uid: %s): %s", testUser.Uid, err)
	}
	log.Printf("Successfully archived user with uid %s\n", testUser.Uid)

	// Bad case
	// user already archived
	_, err = t.Client.ArchiveUser(context.Background(), &interfaces.UserUID{Uid:testUser.Uid})
	if err == nil {
		log.Fatalf("Successfully archived user with uid %s", testUser.Uid)
	} else {
		log.Printf("Error as expected when calling ArchiveUser(uid: %s): %s\n", testUser.Uid, err)
	}

	// user not found
	_, err = t.Client.ArchiveUser(context.Background(), &interfaces.UserUID{Uid:"kaonsdlkassndlkandklad"})
	if err == nil {
		log.Fatalf("Successfully archived user with uid %s", testUser.Uid)
	} else {
		log.Printf("Error as expected when calling ArchiveUser(uid: %s): %s\n", testUser.Uid, err)
	}

	log.Println("Archive user test passed\n\n")
}
