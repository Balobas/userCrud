package repository

import (
	"context"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"userCrud/app/data"
)

type User struct {
	Collection *mongo.Collection
}

func NewUserRepositoryMongo(collection *mongo.Collection) data.UserRepository {
	return User{Collection: collection}
}

func (u User) GetByUid(uid string) (data.User, error) {
	var user data.User
	filter := bson.D{{"_id", uid}}

	if err := u.Collection.FindOne(context.TODO(), filter).Decode(&user); err != nil {
		if err == mongo.ErrNoDocuments {
			return data.User{}, data.ErrorNoObjectsFound
		}
		return data.User{}, errors.WithStack(err)
	}
	return user, nil
}

func (u User) GetByPhoneNumber(phone string) (data.User, error) {
	var user data.User
	filter := bson.D{{"phoneNumber", phone}}

	if err := u.Collection.FindOne(context.TODO(), filter).Decode(&user); err != nil {
		if err == mongo.ErrNoDocuments {
			return data.User{}, data.ErrorNoObjectsFound
		}
		return data.User{}, errors.WithStack(err)
	}
	return user, nil
}

func (u User) GetByFIO(name, lastName, patronymic string) ([]data.User, error) {
	filter := bson.D{{"name", name}, {"lastName", lastName}, {"patronymic", patronymic}}

	users := []data.User{}
	cur, err := u.Collection.Find(context.TODO(), filter, options.Find())
	if err != nil {
		return []data.User{}, errors.WithStack(err)
	}

	for cur.Next(context.TODO()) {
		var user data.User
		if err := cur.Decode(&user); err != nil {
			return []data.User{}, errors.WithStack(err)
		}

		users = append(users, user)
	}

	if err := cur.Err(); err != nil {
		return []data.User{}, errors.WithStack(err)
	}

	return users, nil
}

func (u User) Create(user data.User) error {
	if _, err := u.Collection.InsertOne(context.TODO(), user); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (u User) Update(user data.User) error {
	filter := bson.D{{"_id", user.UID}}

	update := bson.D{
		{"$set", bson.D{
			{"lastName", user.LastName},
			{"name", user.Name},
			{"patronymic", user.Patronymic},
			{"phoneNumber", user.PhoneNumber},
			{"createdDateTime", user.CreatedDateTime},
			{"updatedDateTime", user.UpdatedDateTime},
			{"isArchived", user.IsArchived},
		},},
	}

	if _, err := u.Collection.UpdateOne(context.TODO(), filter, update); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (u User) Delete(uid string) error {
	filter := bson.D{{"_id", uid}}
	if _, err := u.Collection.DeleteOne(context.TODO(), filter); err != nil {
		return errors.WithStack(err)
	}
	return nil
}
