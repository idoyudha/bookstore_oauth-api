// entire business logic to handle this domain

package access_token

import (
	"strings"

	"github.com/idoyudha/bookstore_utils-go/rest_errors"
)

type Repository interface {
	GetById(string) (*AccessToken, rest_errors.RestErr)
	Create(AccessToken) rest_errors.RestErr
	UpdateExpirationTime(AccessToken) rest_errors.RestErr
}

type Service interface {
	GetById(string) (*AccessToken, rest_errors.RestErr)
	Create(AccessToken) rest_errors.RestErr
	UpdateExpirationTime(AccessToken) rest_errors.RestErr
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	// access token service need repository
	return &service{
		repository: repo,
	}
}

func (s *service) GetById(accessTokenId string) (*AccessToken, rest_errors.RestErr) {
	accessTokenId = strings.TrimSpace(accessTokenId)
	if len(accessTokenId) == 0 {
		return nil, rest_errors.NewBadRequestError("invalid access token id")
	}

	accessToken, err := s.repository.GetById(accessTokenId) // taking access token and pass to the database (repository)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}

func (s *service) Create(token AccessToken) rest_errors.RestErr {
	if err := token.Validate(); err != nil {
		return err
	}
	return s.repository.Create(token)
}

func (s *service) UpdateExpirationTime(token AccessToken) rest_errors.RestErr {
	if err := token.Validate(); err != nil {
		return err
	}
	return s.repository.UpdateExpirationTime(token)
}
