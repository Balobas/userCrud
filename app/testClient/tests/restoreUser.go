package tests

import (
	"context"
	"log"
	"userCrud/app/netInterfaces/interfaces"
)

func(t *UserCrudServerTester) testRestoreUser() {
	log.Println("Restore user test")

	testUser := &interfaces.User{
		Uid:             "",
		LastName:        "Bo",
		Name:            "Joh",
		Patronymic:      "Johovich",
		CreatedDateTime: nil,
		UpdatedDateTime: nil,
		IsArchived:      true,
	}
	resp, err := t.Client.CreateUser(context.Background(), testUser)
	if err != nil {
		log.Fatalf("Error when calling CreateUser: %s", err)
	}
	t.StoredUids[resp.Uid] = true

	testUser.Uid = resp.Uid

	// Good case
	_, err = t.Client.RestoreUser(context.Background(), &interfaces.UserUID{Uid:testUser.Uid})
	if err != nil {
		log.Fatalf("Error when calling RestoreUser(uid: %s): %s", testUser.Uid, err)
	}
	log.Printf("Successfully restored user with uid %s\n", testUser.Uid)

	// Bad case
	// user not archived
	_, err = t.Client.RestoreUser(context.Background(), &interfaces.UserUID{Uid:testUser.Uid})
	if err == nil {
		log.Fatalf("Successfully restored user with uid %s", testUser.Uid)
	} else {
		log.Printf("Error as expected when calling RestoreUser(uid: %s): %s\n", testUser.Uid, err)
	}

	// user not found
	_, err = t.Client.RestoreUser(context.Background(), &interfaces.UserUID{Uid:"kaonsdlkassndlkandklad"})
	if err == nil {
		log.Fatalf("Successfully restored user with uid %s", testUser.Uid)
	} else {
		log.Printf("Error as expected when calling RestoreUser(uid: %s): %s\n", testUser.Uid, err)
	}

	log.Println("Restore user test passed\n\n")
}
