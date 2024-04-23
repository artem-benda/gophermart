package api

import (
	"errors"
	"github.com/artem-benda/gophermart/internal/domain/entity"
	"github.com/artem-benda/gophermart/internal/infrastructure/dto"
	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v3/log"
)

var ErrTemporary = errors.New("temporary error")

type AccrualAPI struct {
	Client *resty.Client
}

func (api AccrualAPI) GetAccrualInfo(orderNumber string) (*entity.Accrual, error) {
	log.Debug("making accrual api call for order number: ", orderNumber)
	result := new(dto.GetAccrualInfoResponse)

	resp, err := api.Client.R().
		// SetPathParam("number", orderNumber).
		SetPathParams(map[string]string{
			"number": orderNumber,
		}).
		SetResult(result).
		Get("/api/orders/{number}")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() == 204 {
		return nil, nil
	}
	if resp.StatusCode() != 200 {
		log.Debug("unexpected status code: ", resp.StatusCode())
		return nil, ErrTemporary
	}

	res := entity.Accrual{OrderNumber: result.Number, Status: entity.AccrualStatus(result.Status), AccrualAmount: result.Accrual}
	return &res, nil
}
