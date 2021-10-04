package tests

import (
	"context"
	"google.golang.org/grpc/grpclog"
	"log"
	"userCrud/app/netInterfaces/interfaces"
)

type UserCrudServerTester struct {
	Client interfaces.UserCrudClient
	StoredUids map[string]bool
}

func NewUserCrudServerTester(client interfaces.UserCrudClient) *UserCrudServerTester {
	return &UserCrudServerTester{Client:client, StoredUids: make(map[string]bool)}
}

func(t *UserCrudServerTester) Run() {
	defer t.ClearData()

	t.testCreateUser()
	t.testUpdateUser()
	t.testArchiveUser()
	t.testRestoreUser()
	t.testDeleteUser()
	t.testGetUserByUid()
	t.testGetUsersByFIO()
	t.testGetUserByPhoneNumber()

	log.Println("All tests passed")
}

func(t *UserCrudServerTester) ClearData() {

	for uid, _ := range t.StoredUids {
		_, err := t.Client.DeleteUser(context.Background(), &interfaces.UserUID{Uid: uid})
		if err != nil {
			grpclog.Errorf("Clear test data error: cant delete user with uid: %s", uid)
		}
	}
}
