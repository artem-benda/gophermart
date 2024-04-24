package handler

import (
	"errors"
	"fmt"
	"github.com/artem-benda/gophermart/internal/domain/entity"
	"github.com/artem-benda/gophermart/internal/domain/service"
	"github.com/artem-benda/gophermart/internal/test/fake"
	appmock "github.com/artem-benda/gophermart/internal/test/mock"
	"github.com/gofiber/fiber/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io"
	"net/http"
	"testing"
)

var (
	testBalanceValue          float64 = 1.2345
	testTotalWithdrawalsValue float64 = 3.45667
	testRouteValue                    = "/test"
)

func Test_getUserBalance_get(t *testing.T) {
	type fields struct {
		balanceValue          float64
		balanceErr            error
		totalWithdrawalsValue *float64
		totalWithdrawalsErr   error
	}
	type args struct {
		userID int64
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
			name:          "on success should match balance and withdrawalsTotal values",
			fields:        fields{balanceValue: testBalanceValue, balanceErr: nil, totalWithdrawalsValue: &testTotalWithdrawalsValue, totalWithdrawalsErr: nil},
			args:          args{userID: 1},
			expectedError: false,
			expectedCode:  200,
			expectedBody:  `{"current":1.2345,"withdrawn":3.45667}`,
		},
		{
			name:          "on success should match balance and withdrawalsTotal=nil values",
			fields:        fields{balanceValue: testBalanceValue, balanceErr: nil, totalWithdrawalsValue: nil, totalWithdrawalsErr: nil},
			args:          args{userID: 1},
			expectedError: false,
			expectedCode:  200,
			expectedBody:  `{"current":1.2345}`,
		},
		{
			name:          "on unauthorized should return 401",
			fields:        fields{balanceValue: testBalanceValue, balanceErr: nil, totalWithdrawalsValue: &testTotalWithdrawalsValue, totalWithdrawalsErr: nil},
			args:          args{userID: 0},
			expectedError: false,
			expectedCode:  401,
			expectedBody:  ``,
		},
		{
			name:          "on GetUserByID error should return 500",
			fields:        fields{balanceValue: testBalanceValue, balanceErr: errors.New(""), totalWithdrawalsValue: &testTotalWithdrawalsValue, totalWithdrawalsErr: nil},
			args:          args{userID: 1},
			expectedError: false,
			expectedCode:  500,
			expectedBody:  ``,
		},
		{
			name:          "on GetWithdrawalsError should return 500",
			fields:        fields{balanceValue: testBalanceValue, balanceErr: nil, totalWithdrawalsValue: nil, totalWithdrawalsErr: errors.New("")},
			args:          args{userID: 1},
			expectedError: false,
			expectedCode:  500,
			expectedBody:  ``,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := fiber.New()
			fakeAuth := fake.NewAuthMiddleware()
			h := NewGetUserBalanceHandler(newTestUserBalanceService(tt.fields.balanceValue, tt.fields.balanceErr, tt.fields.totalWithdrawalsValue, tt.fields.totalWithdrawalsErr))
			app.Get(testRouteValue, h, fakeAuth)

			req, _ := http.NewRequest(
				"GET",
				testRouteValue,
				nil,
			)

			if tt.args.userID > 0 {
				req.Header.Set("Authorization", fmt.Sprintf("Bearer %d", tt.args.userID))
			}

			// Perform the request plain with the app.
			// The -1 disables request latency.
			res, err := app.Test(req, -1)
			defer res.Body.Close()

			// verify that no error occured, that is not expected
			assert.Equalf(t, tt.expectedError, err != nil, tt.name)

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

func newTestUserBalanceService(balance float64, balanceErr error, totalWithdrawn *float64, totalWithdrawnErr error) *service.User {
	userRepoMock := new(appmock.UserRepository)
	withdrawalRepoMock := new(appmock.WithdrawalRepository)

	svc := &service.User{
		UserRepository:       userRepoMock,
		WithdrawalRepository: withdrawalRepoMock,
		Salt:                 make([]byte, 32),
	}

	userRepoMock.On("GetUserByID", mock.Anything, int64(1)).Return(&entity.User{ID: 1, Login: "t", PasswordHash: "1", PointsBalance: balance}, balanceErr)
	withdrawalRepoMock.On("GetTotalByUserID", mock.Anything, int64(1)).Return(totalWithdrawn, totalWithdrawnErr)
	return svc
}
