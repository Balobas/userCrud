package tests

import (
	"context"
	"log"
	"userCrud/app/netInterfaces/interfaces"
)

func(t *UserCrudServerTester) testGetUsersByFIO() {
	log.Println("Get users by FIO test")

	response, err := t.Client.CreateUser(context.Background(), &interfaces.User{
		Uid:             "",
		LastName:        "Baloba",
		Name:            "Vladi",
		Patronymic:      "Valet",
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
	resp, err := t.Client.GetUsersByFIO(context.Background(), &interfaces.UserFIOParams{
		LastName:   "Baloba",
		Name:       "Vladi",
		Patronymic: "Valet",
	})
	if err != nil {
		log.Fatalf("Error: %s", err)
	} else {
		log.Println(resp.Users)
	}

	// Bad case
	resp, err = t.Client.GetUsersByFIO(context.Background(), &interfaces.UserFIOParams{
		LastName:   "AAAAAAA",
		Name:       "AAAAAAA",
		Patronymic: "AAAAAAA",
	})
	if err != nil {
		log.Fatalf("error: %s", err)
	} else {

		if len(resp.Users) != 0 {
			log.Fatalf(" expected empty array, but given: %s", resp.Users)
		}
		log.Println("empty array as expected")
	}

	log.Println("Get users by FIO test passed\n\n")
}
