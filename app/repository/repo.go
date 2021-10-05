package repository

import (
	"github.com/pkg/errors"
	"os"
	"userCrud/app/data"
	"userCrud/app/database"
)

type RepositoryDatabase string

const (
	RepoDatabaseMongo = RepositoryDatabase("mongo")
)

type DatabaseOptions struct {
	Database                   RepositoryDatabase
	DatabaseName               string
	UsersTableOrCollectionName string
	DatabaseURI                string
}

func GetUserRepository() (data.UserRepository, error) {
	options := DatabaseOptions{}

	options.Database = RepositoryDatabase(os.Getenv("REPOSITORY_DATABASE"))
	options.DatabaseName = os.Getenv("DATABASE_NAME")
	options.UsersTableOrCollectionName = os.Getenv("USERS_TABLE_OR_COLLECTION_NAME")
	options.DatabaseURI = os.Getenv("DATABASE_URI")


	//// test
	//options.Database = RepoDatabaseMongo
	//options.DatabaseName = "usercrud"
	//options.UsersTableOrCollectionName = "users"
	//options.DatabaseURI = "mongodb://127.0.0.1:27017"
	////

	if options.Database == RepoDatabaseMongo {
		client := database.GetMongoConnection(options.DatabaseURI)
		collection := client.Database(options.DatabaseName).Collection(options.UsersTableOrCollectionName)
		return NewUserRepositoryMongo(collection), nil
	}

	return nil, errors.Errorf("invalid database: %s", options.Database)
}
