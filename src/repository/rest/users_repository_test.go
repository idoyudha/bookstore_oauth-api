package rest

import (
	"net/http"
	"os"
	"testing"

	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	rest.StartMockupServer()
	os.Exit(m.Run())
}

func TestLoginUserTimeoutFromApi(t *testing.T) {
	rest.FlushMockups() // removing previous mockup and starting with fresh mockup system
	rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          "http://localhost:8080/users/login",
		ReqBody:      `{"email":"test@mail.com", "password":"password"}`,
		RespHTTPCode: -1,
		RespBody:     `{}`,
	})

	repository := usersRepository{}

	user, err := repository.LoginUser("test@mail.com", "password")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "invalid restclient response when trying to login user", err.Message)
}

func TestLoginUserInvalidErrorInterface(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          "http://localhost:8080/users/login",
		ReqBody:      `{"email":"test@mail.com","password":"password"}`,
		RespHTTPCode: http.StatusNotFound,
		RespBody:     `{"message": "invalid login credentials", "status": "404", "error": "not_found"}`,
	})

	repository := usersRepository{}

	user, err := repository.LoginUser("test@mail.com", "password")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "invalid error interface when trying to login user", err.Message)
}

func TestLoginUserInvalidLoginCredentials(t *testing.T) {
	//TODO
}

func TestLoginUserInvalidUserJsonResponse(t *testing.T) {
	//TODO
}

func TestLoginUserNoError(t *testing.T) {
	//TODO
}
