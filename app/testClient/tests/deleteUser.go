package tests

import (
	"context"
	"log"
	"userCrud/app/netInterfaces/interfaces"
)

func(t *UserCrudServerTester) testDeleteUser() {
	log.Println("Delete user test")

	response, err := t.Client.CreateUser(context.Background(), &interfaces.User{
		Uid:             "",
		LastName:        "Baloban",
		Name:            "Vladisl",
		Patronymic:      "Valer",
		CreatedDateTime: nil,
		UpdatedDateTime: nil,
		IsArchived:      false,
	})
	if err != nil {
		log.Fatalf("Error when calling CreateUser: %s", err)
	}
	log.Printf("Response from server: %s\n", response.Uid)

	// Good case
	_, err = t.Client.DeleteUser(context.Background(), &interfaces.UserUID{Uid:response.Uid})
	if err != nil {
		log.Fatalf("Error when calling DeleteUser: %s", err)
	}
	log.Printf("Successfully delete user with uid: %s\n", response.Uid)

	// Bad case
	// user not found
	_, err = t.Client.DeleteUser(context.Background(), &interfaces.UserUID{Uid:"skdnfksjndflkskndlfnsdlfnsl"})
	if err == nil {
		log.Fatal("expected error: user not found")
	} else {
		log.Printf("error as expected: %s\n", err)
	}

	log.Println("Delete user test passed\n\n")
}
