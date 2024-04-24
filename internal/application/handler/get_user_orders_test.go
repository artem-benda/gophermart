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

var (
	testOrderAccrualAmount float64 = 1.23456
)

func Test_getUserOrders_getList(t *testing.T) {
	type fields struct {
		ordersList []entity.Order
		ordersErr  error
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
			fields: fields{ordersList: []entity.Order{
				{
					Number:        "123456",
					Status:        entity.OrderStatusProcessing,
					UserID:        1,
					UploadedAt:    time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
					AccrualAmount: nil,
				},
				{
					Number:        "123457",
					Status:        entity.OrderStatusProcessed,
					UserID:        1,
					UploadedAt:    time.Date(2024, 1, 3, 0, 0, 0, 0, time.UTC),
					AccrualAmount: &testOrderAccrualAmount,
				},
			}},
			args:          args{userID: 1},
			expectedError: false,
			expectedCode:  200,
			expectedBody:  `[{"number":"123456","status":"PROCESSING","uploaded_at":"2024-01-01T00:00:00Z"},{"number":"123457","status":"PROCESSED","accrual":1.23456,"uploaded_at":"2024-01-03T00:00:00Z"}]`,
		},
		{
			name:          "on success for empty orders should return 204",
			fields:        fields{ordersList: make([]entity.Order, 0)},
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
			fields:        fields{ordersErr: errors.New("some error")},
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
			h := NewGetUserOrdersHandler(newTestGetOrdersService(tt.fields.ordersList, tt.fields.ordersErr))
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
			defer func(Body io.ReadCloser) {
				_ = Body.Close()
			}(res.Body)

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

func newTestGetOrdersService(ordersList []entity.Order, ordersErr error) *service.Order {
	orderRepoMock := new(appmock.OrderRepository)

	svc := &service.Order{
		OrderRepository: orderRepoMock,
	}

	orderRepoMock.On("GetByUserID", mock.Anything, int64(1)).Return(ordersList, ordersErr)
	return svc
}
