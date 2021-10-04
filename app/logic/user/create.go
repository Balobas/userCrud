package user

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"time"
	"userCrud/app/data"
)

func (u *UseCases) CreateUser(user data.User) (string, error) {
	if err := user.Validate(); err != nil {
		return "", err
	}

	/*
	По одному номеру телефона может быть только один пользователь.
	Если не указан номер телефона, то нельзя сохранить полного теску того кто уже есть в базе.

	Поэтому сначала метод проверяет указан ли номер телефона, если указан то ищет пользователя в базе.
	В случае находки, возвращает ошибку, в случае если не найден - записывает в базу. Если же номер телефона не указан, метод
	ищет пользователя по ФИО и в случае находки - возвращает ошибку, в противном случае записывает в базу.

	 */

	if len(user.PhoneNumber) != 0 {
		_, err := u.UserRepo.GetByPhoneNumber(user.PhoneNumber)
		if err == nil {
			return "", ErrUserWithSameNumberExist
		}
		if err != data.ErrorNoObjectsFound {
			return "", errors.WithStack(err)
		}
	} else {
		oldUsers, err := u.UserRepo.GetByFIO(user.Name, user.LastName, user.Patronymic)
		if err != nil {
			return "", errors.WithStack(err)
		}

		if len(oldUsers) != 0 {
			return "", errors.Errorf("user with the same name, lastName and patronymic already exists (%s, %s, %s). You should set phone number", user.Name, user.LastName, user.Patronymic)
		}
	}

	user.CreatedDateTime = time.Now().Unix()
	user.UpdatedDateTime = time.Now().Unix()

	user.UID = uuid.New().String()

	if err := u.UserRepo.Create(user); err != nil {
		return "", errors.WithStack(err)
	}

	return user.UID, nil
}
