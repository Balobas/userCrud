package data

import (
	"errors"
	"time"
)

type User struct {
	UID             string `json:"uid" bson:"_id"`
	LastName        string `json:"lastName" bson:"lastName"`
	Name            string `json:"name" bson:"name"`
	Patronymic      string `json:"patronymic" bson:"patronymic"`
	PhoneNumber     string `json:"phoneNumber" bson:"phoneNumber"`
	CreatedDateTime int64  `json:"createdDateTime" bson:"createdDateTime"`
	UpdatedDateTime int64  `json:"updatedDateTime" bson:"updatedDateTime"`
	IsArchived      bool   `json:"isArchived" bson:"isArchived"`
}


/*
Имя и фамилия обязательные поля
 */
func (u *User) Validate() error {
	if len(u.Name) == 0 {
		return errors.New("empty name")
	}
	if len(u.LastName) == 0 {
		return errors.New("empty lastName")
	}
	return nil
}

// При обновлении полей посредством этого метода никак не изменяется createdDateTime
func (u *User) UpdateFields(user User) {
	if len(user.Name) != 0 {
		u.Name = user.Name
	}
	if len(user.LastName) != 0 {
		u.LastName = user.LastName
	}
	if len(user.Patronymic) != 0 {
		u.Patronymic = user.Patronymic
	}
	if len(user.PhoneNumber) != 0 {
		u.PhoneNumber = user.PhoneNumber
	}
	u.UpdatedDateTime = time.Now().Unix()
}

func (u *User) ChangeArchivedFlag() {
	u.IsArchived = !u.IsArchived
	u.UpdatedDateTime = time.Now().Unix()
}


/*
Интерфейс репозитория для работы с объектом пользователя
 */
type UserRepository interface {
	GetByUid(uid string) (User, error)                          // В случае если пользователь не найден, возвращает ошибку
	GetByPhoneNumber(phone string) (User, error)                // Так же как в getByUid
	GetByFIO(name, lastName, patronymic string) ([]User, error) // В случае если ни один пользователь не найден, возвращает пустой массив и nil в качестве ошибки
	Create(user User) error
	Update(user User) error
	Delete(uid string) error
}

var ErrorNoObjectsFound = errors.New("no objects found")
