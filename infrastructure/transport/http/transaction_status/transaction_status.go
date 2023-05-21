package transaction_status

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sultanfariz/synapsis-assignment/domain/transaction_status"
	"github.com/sultanfariz/synapsis-assignment/infrastructure/commons"
	httpControllers "github.com/sultanfariz/synapsis-assignment/infrastructure/transport/http"
)

type Controllers struct {
	TransactionStatusUsecase transaction_status.TransactionStatusUsecase
}

func NewControllers(trxStatusUC transaction_status.TransactionStatusUsecase) *Controllers {
	return &Controllers{
		TransactionStatusUsecase: trxStatusUC,
	}
}

func (controller *Controllers) Insert(c echo.Context) error {
	ctx := c.Request().Context()
	req := TransactionStatus{}
	if err := c.Bind(&req); err != nil {
		return httpControllers.ErrorResponse(c, http.StatusBadRequest, commons.ErrBadRequest)
	}
	domain := req.ToDomain()
	res, err := controller.TransactionStatusUsecase.Insert(ctx, domain)
	if err != nil {
		if errors.Is(err, commons.ErrTransactionStatusAlreadyExists) {
			return httpControllers.ErrorResponse(c, http.StatusConflict, commons.ErrTransactionStatusAlreadyExists)
		} else if errors.Is(err, commons.ErrValidationFailed) {
			return httpControllers.ErrorResponse(c, http.StatusConflict, commons.ErrValidationFailed)
		}
		return httpControllers.ErrorResponse(c, http.StatusInternalServerError, commons.ErrInternalServerError)
	}
	return httpControllers.SuccessResponse(c, http.StatusCreated, FromDomain(*res))
}

func (controller *Controllers) GetAll(c echo.Context) error {
	ctx := c.Request().Context()

	res, err := controller.TransactionStatusUsecase.GetAll(ctx)
	if err != nil {
		return httpControllers.ErrorResponse(c, http.StatusInternalServerError, commons.ErrInternalServerError)
	}
	return httpControllers.SuccessResponse(c, http.StatusOK, FromDomainList(res))
}
