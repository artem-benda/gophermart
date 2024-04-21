package api

import (
	"github.com/artem-benda/gophermart/internal/domain/entity"
	"github.com/artem-benda/gophermart/internal/infrastructure/dto"
	"github.com/gofiber/fiber/v3/client"
)

type AccrualAPI struct {
	Client *client.Client
}

func (api AccrualAPI) GetAccrualInfo(orderNumber string) (*entity.Accrual, error) {
	resp, err := api.Client.Get("/api/orders/{number}", client.Config{PathParam: map[string]string{"number": orderNumber}})
	if err != nil {
		return nil, err
	}
	d := new(dto.GetAccrualInfoResponse)
	err = resp.JSON(d)
	if err != nil {
		return nil, err
	}

	res := entity.Accrual{OrderNumber: d.Number, Status: entity.AccrualStatus(d.Status), AccrualAmount: d.Accrual}
	return &res, nil
}
