package handler

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/artem-benda/gophermart/internal/domain/service"
	"github.com/artem-benda/gophermart/internal/infrastructure/dto"
	"github.com/artem-benda/gophermart/internal/test/fake"
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

func Test_uploadOrder_upload(t *testing.T) {
	type fields struct {
		uploadErr error
	}
	type args struct {
		userID      int64
		orderNumber string
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
			name:          "on success should return 202",
			fields:        fields{uploadErr: nil},
			args:          args{userID: 1, orderNumber: "12345678903"},
			expectedError: false,
			expectedCode:  202,
			expectedBody:  ``,
		},
		{
			name:          "on invalid order number should return 422",
			fields:        fields{uploadErr: nil},
			args:          args{userID: 1, orderNumber: "12345678900"},
			expectedError: false,
			expectedCode:  422,
			expectedBody:  ``,
		},
		{
			name:          "on unauthorized should return 401",
			fields:        fields{},
			args:          args{userID: 0},
			expectedError: false,
			expectedCode:  401,
			expectedBody:  ``,
		},
		{
			name:          "on upload error should return 500",
			fields:        fields{uploadErr: errors.New("some error")},
			args:          args{userID: 1, orderNumber: "12345678903"},
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
			fakeAuth := fake.NewAuthMiddleware()
			h := NewUploadOrderHandler(newTestUploadOrderService(tt.fields.uploadErr), validate)
			app.Post(testRouteValue, h, fakeAuth)

			req, _ := http.NewRequest(
				"POST",
				testRouteValue,
				bytes.NewReader([]byte(tt.args.orderNumber)),
			)

			if tt.args.userID > 0 {
				req.Header.Set("Authorization", fmt.Sprintf("Bearer %d", tt.args.userID))
			}

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

func newTestUploadOrderService(uploadErr error) *service.Order {
	orderRepoMock := new(appmock.OrderRepository)

	svc := &service.Order{
		OrderRepository: orderRepoMock,
	}

	orderRepoMock.On("Upload", mock.Anything, int64(1), mock.Anything).Return(uploadErr)
	return svc
}
