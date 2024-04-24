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
	"time"
)

func Test_getWithdrawals_getList(t *testing.T) {
	type fields struct {
		withdrawalsList []entity.Withdrawal
		withdrawalsErr  error
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
			name: "on success should match orders",
			fields: fields{withdrawalsList: []entity.Withdrawal{
				{
					OrderNumber: "123456",
					Amount:      234.1243,
					UserID:      1,
					CreatedAt:   time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
					ProcessedAt: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
				},
				{
					OrderNumber: "123457",
					Amount:      44.4578,
					UserID:      1,
					CreatedAt:   time.Date(2024, 1, 2, 0, 0, 0, 0, time.UTC),
					ProcessedAt: time.Date(2024, 1, 2, 0, 0, 0, 0, time.UTC),
				},
			}},
			args:          args{userID: 1},
			expectedError: false,
			expectedCode:  200,
			expectedBody:  `[{"order":"123456","sum":234.1243,"processed_at":"2024-01-01T00:00:00Z"},{"order":"123457","sum":44.4578,"processed_at":"2024-01-02T00:00:00Z"}]`,
		},
		{
			name:          "on success for empty orders should return 204",
			fields:        fields{withdrawalsList: make([]entity.Withdrawal, 0)},
			args:          args{userID: 1},
			expectedError: false,
			expectedCode:  204,
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
			name:          "on GetByUserID error should return 500",
			fields:        fields{withdrawalsErr: errors.New("some error")},
			args:          args{userID: 1},
			expectedError: false,
			expectedCode:  500,
			expectedBody:  `some error`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := fiber.New()
			fakeAuth := fake.NewAuthMiddleware()
			h := NewGetWithdrawalsHandler(newTestGetWithdrawalsService(tt.fields.withdrawalsList, tt.fields.withdrawalsErr))
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

func newTestGetWithdrawalsService(withdrawalsList []entity.Withdrawal, withdrawalsErr error) *service.Withdrawal {
	withdrawalsRepoMock := new(appmock.WithdrawalRepository)

	svc := &service.Withdrawal{
		WithdrawalRepository: withdrawalsRepoMock,
	}

	withdrawalsRepoMock.On("GetListByUserID", mock.Anything, int64(1)).Return(withdrawalsList, withdrawalsErr)
	return svc
}
