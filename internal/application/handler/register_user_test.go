package handler

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/artem-benda/gophermart/internal/domain/service"
	"github.com/artem-benda/gophermart/internal/infrastructure/dto"
	appmock "github.com/artem-benda/gophermart/internal/test/mock"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io"
	"net/http"
	"testing"
)

var (
	testUserID int64 = 1
)

func Test_registerUser_registerUser(t *testing.T) {
	type fields struct {
		login         string
		passwordHash  string
		insertUserErr error
	}
	type args struct {
		login    string
		password string
	}
	tests := []struct {
		name          string
		fields        fields
		args          args
		expectedError bool
		expectedCode  int
		expectedBody  string
	}{
		{
			name:          "on success should match password hash and return 200",
			fields:        fields{login: "test", passwordHash: "afd535a859b731dff667376f1bad148bec41a419ffbda681791843eb4e1e3b2f", insertUserErr: nil},
			args:          args{login: "test", password: "2fewfwe"},
			expectedError: false,
			expectedCode:  200,
			expectedBody:  ``,
		},
		{
			name:          "on InsertUser error should return 500",
			fields:        fields{login: "test", passwordHash: "afd535a859b731dff667376f1bad148bec41a419ffbda681791843eb4e1e3b2f", insertUserErr: errors.New("some error")},
			args:          args{login: "test", password: "2fewfwe"},
			expectedError: false,
			expectedCode:  500,
			expectedBody:  `Internal Server Error`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := fiber.New()
			validate := validator.New()
			err := validate.RegisterValidation("luhn", dto.LuhnStringValidator)
			if err != nil {
				log.Fatal(err)
			}
			h := NewRegisterUserHandler(newTestUserRegisterService(tt.fields.login, tt.fields.passwordHash, tt.fields.insertUserErr), validate)
			app.Post(testRouteValue, h)

			req, _ := http.NewRequest(
				"POST",
				testRouteValue,
				bytes.NewReader([]byte(fmt.Sprintf("{\"login\":\"%s\", \"password\":\"%s\"}", tt.args.login, tt.args.password))),
			)

			// Perform the request plain with the app.
			// The -1 disables request latency.
			res, err := app.Test(req, -1)

			// verify that no error occured, that is not expected
			assert.Equalf(t, tt.expectedError, err != nil, tt.name)

			defer res.Body.Close()

			// As expected errors lead to broken responses, the next
			// test case needs to be processed
			if tt.expectedError {
				return
			}

			// Verify if the status code is as expected
			assert.Equalf(t, tt.expectedCode, res.StatusCode, tt.name)

			// Read the response body
			body, err := io.ReadAll(res.Body)

			// Reading the response body should work everytime, such that
			// the err variable should be nil
			assert.Nilf(t, err, tt.name)

			// Verify, that the reponse body equals the expected body
			assert.Equalf(t, tt.expectedBody, string(body), tt.name)
		})
	}
}

func newTestUserRegisterService(login string, passwordHash string, insertUserErr error) *service.User {
	userRepoMock := new(appmock.UserRepository)
	withdrawalRepoMock := new(appmock.WithdrawalRepository)

	svc := &service.User{
		UserRepository:       userRepoMock,
		WithdrawalRepository: withdrawalRepoMock,
		Salt:                 []byte("1234567890asdfghj"),
	}

	if insertUserErr == nil {
		userRepoMock.On("Register", mock.Anything, login, passwordHash).Return(&testUserID, nil)
	} else {
		userRepoMock.On("Register", mock.Anything, login, passwordHash).Return(nil, insertUserErr)
	}

	return svc
}
