package payments

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sultanfariz/synapsis-assignment/domain/payments"
	"github.com/sultanfariz/synapsis-assignment/infrastructure/commons"
	httpControllers "github.com/sultanfariz/synapsis-assignment/infrastructure/transport/http"
)

type Controllers struct {
	PaymentsUsecase payments.PaymentsUsecase
}

func NewControllers(paymentsUC payments.PaymentsUsecase) *Controllers {
	return &Controllers{
		PaymentsUsecase: paymentsUC,
	}
}

func (controller *Controllers) Insert(c echo.Context) error {
	ctx := c.Request().Context()
	req := Payment{}
	if err := c.Bind(&req); err != nil {
		return httpControllers.ErrorResponse(c, http.StatusBadRequest, commons.ErrBadRequest)
	}
	domain := req.ToDomain()
	res, err := controller.PaymentsUsecase.Insert(ctx, domain)
	if err != nil {
		if errors.Is(err, commons.ErrPaymentMethodAlreadyExists) {
			return httpControllers.ErrorResponse(c, http.StatusConflict, commons.ErrPaymentMethodAlreadyExists)
		} else if errors.Is(err, commons.ErrValidationFailed) {
			return httpControllers.ErrorResponse(c, http.StatusConflict, commons.ErrValidationFailed)
		}
		return httpControllers.ErrorResponse(c, http.StatusInternalServerError, commons.ErrInternalServerError)
	}
	return httpControllers.SuccessResponse(c, http.StatusCreated, FromDomain(*res))
}

func (controller *Controllers) GetAll(c echo.Context) error {
	ctx := c.Request().Context()

	res, err := controller.PaymentsUsecase.GetAll(ctx)
	if err != nil {
		return httpControllers.ErrorResponse(c, http.StatusInternalServerError, commons.ErrInternalServerError)
	}
	return httpControllers.SuccessResponse(c, http.StatusOK, FromDomainList(res))
}
