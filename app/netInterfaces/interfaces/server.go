package interfaces

import (
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"userCrud/app/data"
	userLogic "userCrud/app/logic/user"
)

type UserServer struct {
	UseCases *userLogic.UseCases
}

func (u UserServer) CreateUser(ctx context.Context, user *User) (*CreateUserResponse, error) {
	log.Println("Called CreateUser method with params: ", *user)

	uid, err := u.UseCases.CreateUser(data.User{
		UID:             user.Uid,
		LastName:        user.LastName,
		Name:            user.Name,
		Patronymic:      user.Patronymic,
		PhoneNumber:     user.PhoneNumber,
		CreatedDateTime: 0,
		UpdatedDateTime: 0,
		IsArchived:      user.IsArchived,
	})

	log.Println("Returned : ", uid, err, "\n")
	return &CreateUserResponse{Uid: uid}, err
}

func (u UserServer) UpdateUser(ctx context.Context, user *User) (*Response, error) {
	log.Println("Called UpdateUser method with params: ", *user)

	err := u.UseCases.UpdateUser(data.User{
		UID:             user.Uid,
		LastName:        user.LastName,
		Name:            user.Name,
		Patronymic:      user.Patronymic,
		PhoneNumber:     user.PhoneNumber,
		CreatedDateTime: 0,
		UpdatedDateTime: 0,
		IsArchived:      user.IsArchived,
	})

	log.Println("Returned : ", err, "\n")
	return &Response{Message: ""}, err
}

func (u UserServer) ArchiveUser(ctx context.Context, user *UserUID) (*Response, error) {
	log.Println("Called ArchiveUser method with params: ", *user)

	err := u.UseCases.ArchiveUser(user.Uid)

	log.Println("Returned : ", err, "\n")
	return &Response{Message: ""}, err
}

func (u UserServer) RestoreUser(ctx context.Context, user *UserUID) (*Response, error) {
	log.Println("Called RestoreUser method with params: ", *user)

	err := u.UseCases.RestoreUser(user.Uid)

	log.Println("Returned : ", err, "\n")
	return &Response{Message: ""}, err
}

func (u UserServer) DeleteUser(ctx context.Context, user *UserUID) (*Response, error) {
	log.Println("Called DeleteUser method with params: ", *user)

	err := u.UseCases.DeleteUser(user.Uid)

	log.Println("Returned : ", err, "\n")
	return &Response{Message: ""}, err
}

func (u UserServer) GetUserByUID(ctx context.Context, user *UserUID) (*GetUserResponse, error) {
	log.Println("Called GetUserByUid method with params: ", *user)

	foundUser, err := u.UseCases.GetByUid(user.Uid)

	response := GetUserResponse{User: &User{
		Uid:             foundUser.UID,
		LastName:        foundUser.LastName,
		Name:            foundUser.Name,
		Patronymic:      foundUser.Patronymic,
		PhoneNumber:     foundUser.PhoneNumber,
		IsArchived:      foundUser.IsArchived,
	}}

	response.User.CreatedDateTime = &timestamppb.Timestamp{}
	response.User.UpdatedDateTime = &timestamppb.Timestamp{}

	response.User.CreatedDateTime.Seconds = foundUser.CreatedDateTime
	response.User.UpdatedDateTime.Seconds = foundUser.UpdatedDateTime

	log.Println("Returned : ", response, err, "\n")
	return &response, err
}

func (u UserServer) GetUsersByFIO(ctx context.Context, params *UserFIOParams) (*GetUsersResponse, error) {
	log.Println("Called GetUserByFIO method with params: ", *params)

	foundUsers, err := u.UseCases.GetByFIO(params.Name, params.LastName, params.Patronymic)

	responseUsers := []*User{}

	for i := 0; i < len(foundUsers); i++ {
		responseUsers = append(responseUsers, &User{
			Uid:             foundUsers[i].UID,
			LastName:        foundUsers[i].LastName,
			Name:            foundUsers[i].Name,
			Patronymic:      foundUsers[i].Patronymic,
			PhoneNumber:     foundUsers[i].PhoneNumber,
			CreatedDateTime: nil,
			UpdatedDateTime: nil,
			IsArchived:      foundUsers[i].IsArchived,
		})

		responseUsers[i].CreatedDateTime = &timestamppb.Timestamp{}
		responseUsers[i].UpdatedDateTime = &timestamppb.Timestamp{}

		responseUsers[i].CreatedDateTime.Seconds = foundUsers[i].CreatedDateTime
		responseUsers[i].UpdatedDateTime.Seconds = foundUsers[i].UpdatedDateTime
	}

	response := GetUsersResponse{Users: responseUsers}

	log.Println("Returned : ", response, err, "\n")
	return &response, err
}

func (u UserServer) GetUserByPhoneNumber(ctx context.Context, params *UserPhoneParams) (*GetUserResponse, error) {
	log.Println("Called GetUserByUid method with params: ", params.Phone)

	foundUser, err := u.UseCases.GetByPhoneNumber(params.Phone)

	response := GetUserResponse{User: &User{
		Uid:             foundUser.UID,
		LastName:        foundUser.LastName,
		Name:            foundUser.Name,
		PhoneNumber:     foundUser.PhoneNumber,
		Patronymic:      foundUser.Patronymic,
		IsArchived:      foundUser.IsArchived,
	}}

	response.User.CreatedDateTime = &timestamppb.Timestamp{}
	response.User.UpdatedDateTime = &timestamppb.Timestamp{}

	response.User.CreatedDateTime.Seconds = foundUser.CreatedDateTime
	response.User.UpdatedDateTime.Seconds = foundUser.UpdatedDateTime

	log.Println("Returned : ", response, err, "\n")
	return &response, err
}

func (u UserServer) mustEmbedUnimplementedUserCrudServer() {

}

