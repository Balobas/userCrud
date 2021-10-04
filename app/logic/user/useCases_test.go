package user

import "userCrud/app/data"

type UserRepoMock struct {
	GetByUidUserToReturn  data.User
	GetByUidErrorToReturn error
	GetByFIOUsersToReturn []data.User
	GetByFIOErrorToReturn error
	GetByPhoneNumberUserToReturn data.User
	GetByPhoneNumberErrorToReturn error
	CreateErrorToReturn   error
	UpdateErrorToReturn   error
	DeleteErrorToReturn   error
}

func (u UserRepoMock) GetByUid(_ string) (data.User, error) {
	return u.GetByUidUserToReturn, u.GetByUidErrorToReturn
}

func (u UserRepoMock) GetByFIO(_, _, _ string) ([]data.User, error) {
	return u.GetByFIOUsersToReturn, u.GetByUidErrorToReturn
}

func (u UserRepoMock) Create(_ data.User) error {
	return u.CreateErrorToReturn
}

func (u UserRepoMock) Update(_ data.User) error {
	return u.UpdateErrorToReturn
}

func (u UserRepoMock) Delete(_ string) error {
	return u.DeleteErrorToReturn
}

func (u UserRepoMock) GetByPhoneNumber(_ string) (data.User, error) {
	return u.GetByPhoneNumberUserToReturn, u.GetByPhoneNumberErrorToReturn
}
