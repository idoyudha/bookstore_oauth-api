package db

import (
	"bookstore_oauth-api/src/clients/cassandra"
	"bookstore_oauth-api/src/domain/access_token"
	"bookstore_oauth-api/src/utils/errors"
)

func NewRepository() DbRepository {
	return &dbRepository{}
}

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
}

type dbRepository struct {
}

func (r *dbRepository) GetById(id string) (*access_token.AccessToken, *errors.RestErr) {
	// get access token from CassandraDB
	session, err := cassandra.GetSession()
	if err != nil {
		panic(err)
	}
	defer session.Close()
	return nil, errors.NewInternalServerError("database connection not implemented yet!")
}
