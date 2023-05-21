package transactions

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sultanfariz/synapsis-assignment/domain/transactions"
	"github.com/sultanfariz/synapsis-assignment/infrastructure/commons"
	httpControllers "github.com/sultanfariz/synapsis-assignment/infrastructure/transport/http"
)

type Controllers struct {
	TransactionsUsecase transactions.TransactionsUsecase
}

func NewControllers(transactionsUC transactions.TransactionsUsecase) *Controllers {
	return &Controllers{
		TransactionsUsecase: transactionsUC,
	}
}

func (controller *Controllers) Insert(c echo.Context) error {
	ctx := c.Request().Context()
	req := Transaction{}
	if err := c.Bind(&req); err != nil {
		return httpControllers.ErrorResponse(c, http.StatusBadRequest, commons.ErrBadRequest)
	}
	userId := commons.GetUserId(c)
	domain := req.ToDomain()
	domain.UserId = userId
	res, err := controller.TransactionsUsecase.Insert(ctx, domain)
	if err != nil {
		if errors.Is(err, commons.ErrProductNotFound) {
			return httpControllers.ErrorResponse(c, http.StatusNotFound, commons.ErrProductNotFound)
		} else if errors.Is(err, commons.ErrProductOutOfStock) {
			return httpControllers.ErrorResponse(c, http.StatusConflict, commons.ErrProductOutOfStock)
		} else if errors.Is(err, commons.ErrValidationFailed) {
			return httpControllers.ErrorResponse(c, http.StatusBadRequest, commons.ErrValidationFailed)
		}
		return httpControllers.ErrorResponse(c, http.StatusInternalServerError, commons.ErrInternalServerError)
	}
	return httpControllers.SuccessResponse(c, http.StatusCreated, FromDomain(*res))
}

func (controller *Controllers) GetByUser(c echo.Context) error {
	ctx := c.Request().Context()
	userId := commons.GetUserId(c)
	res, err := controller.TransactionsUsecase.GetByUser(ctx, userId)
	if err != nil {
		if errors.Is(err, commons.ErrTransactionNotFound) {
			return httpControllers.ErrorResponse(c, http.StatusNotFound, commons.ErrTransactionNotFound)
		}
		return httpControllers.ErrorResponse(c, http.StatusInternalServerError, commons.ErrInternalServerError)
	}
	return httpControllers.SuccessResponse(c, http.StatusOK, FromDomainList(res))
}

func (controller *Controllers) UpdateStatus(c echo.Context) error {
	ctx := c.Request().Context()
	req := UpdateStatus{}
	id := c.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		return httpControllers.ErrorResponse(c, http.StatusBadRequest, commons.ErrBadRequest)
	}
	if err := c.Bind(&req); err != nil {
		return httpControllers.ErrorResponse(c, http.StatusBadRequest, commons.ErrBadRequest)
	}
	userId := commons.GetUserId(c)
	res, err := controller.TransactionsUsecase.UpdateStatus(ctx, intId, userId, req.StatusId)
	if err != nil {
		if errors.Is(err, commons.ErrTransactionNotFound) {
			return httpControllers.ErrorResponse(c, http.StatusNotFound, commons.ErrTransactionNotFound)
		} else if errors.Is(err, commons.ErrValidationFailed) {
			return httpControllers.ErrorResponse(c, http.StatusBadRequest, commons.ErrValidationFailed)
		}
		return httpControllers.ErrorResponse(c, http.StatusInternalServerError, commons.ErrInternalServerError)
	}
	return httpControllers.SuccessResponse(c, http.StatusOK, FromDomain(*res))
}
