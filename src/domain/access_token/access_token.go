// domain definition

package access_token

import (
	"bookstore_oauth-api/src/utils/crypto_utils"
	"bookstore_oauth-api/src/utils/errors"
	"fmt"
	"strings"
	"time"
)

const (
	expirationTime             = 24
	grantTypePassword          = "password"
	grantTypeClientCredentials = "client_credentials"
)

type AccessTokenRequest struct {
	GrantType string `json:"grant_type"`
	Scope     string `json:"scope"`

	// used for password grant type
	Username string `json:"username"`
	Password string `json:"password"`

	// used for client credentials grant type
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

func (token *AccessTokenRequest) Validate() *errors.RestErr {
	switch token.GrantType {
	case grantTypePassword:
		break

	case grantTypeClientCredentials:
		break

	default:
		return errors.NewBadRequestError("invalid grant_type parameter")
	}
	return nil
}

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId      int64  `json:"user_id"`
	ClientId    int64  `json:"client_id"` // client_id coming from platform (ex: android or web)
	Expires     int64  `json:"expires"`
}

func (token *AccessToken) Validate() *errors.RestErr {
	token.AccessToken = strings.TrimSpace(token.AccessToken)
	if token.AccessToken == "" {
		return errors.NewBadRequestError("invalid access token id")
	}
	if token.UserId <= 0 {
		return errors.NewBadRequestError("invalid user id")
	}
	if token.ClientId <= 0 {
		return errors.NewBadRequestError("invalid client id")
	}
	if token.Expires <= 0 {
		return errors.NewBadRequestError("invalid expiration time")
	}
	return nil
}

func GetNewAcessToken(userId int64) AccessToken {
	return AccessToken{
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

func (token AccessToken) IsExpired() bool {
	return time.Unix(token.Expires, 0).Before(time.Now().UTC())
}

func (token *AccessToken) Generate() {
	token.AccessToken = crypto_utils.GetMd5(fmt.Sprintf("at-%d-%d-ran", token.UserId, token.Expires))
}
