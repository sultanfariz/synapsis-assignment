package carts

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sultanfariz/synapsis-assignment/domain/carts"
	"github.com/sultanfariz/synapsis-assignment/infrastructure/commons"
	httpControllers "github.com/sultanfariz/synapsis-assignment/infrastructure/transport/http"
	"github.com/sultanfariz/synapsis-assignment/infrastructure/transport/http/products"
)

type Controllers struct {
	CartsUsecase carts.CartsUsecase
}

func NewControllers(cartsUC carts.CartsUsecase) *Controllers {
	return &Controllers{
		CartsUsecase: cartsUC,
	}
}

func (controller *Controllers) Insert(c echo.Context) error {
	ctx := c.Request().Context()
	req := Cart{}
	if err := c.Bind(&req); err != nil {
		return httpControllers.ErrorResponse(c, http.StatusBadRequest, commons.ErrBadRequest)
	}
	userId := commons.GetUserId(c)
	domain := req.ToDomain()
	res, err := controller.CartsUsecase.Insert(ctx, userId, domain.ProductId)
	if err != nil {
		if errors.Is(err, commons.ErrCategoryNotFound) {
			return httpControllers.ErrorResponse(c, http.StatusConflict, commons.ErrCategoryNotFound)
		} else if errors.Is(err, commons.ErrValidationFailed) {
			return httpControllers.ErrorResponse(c, http.StatusConflict, commons.ErrValidationFailed)
		}
		return httpControllers.ErrorResponse(c, http.StatusInternalServerError, commons.ErrInternalServerError)
	}
	return httpControllers.SuccessResponse(c, http.StatusCreated, FromDomain(*res))
}

func (controller *Controllers) GetByUser(c echo.Context) error {
	ctx := c.Request().Context()
	// extract user email from jwt token
	userId := commons.GetUserId(c)
	res, err := controller.CartsUsecase.GetByUser(ctx, userId)
	if err != nil {
		if errors.Is(err, commons.ErrCartIsEmpty) {
			return httpControllers.SuccessResponse(c, http.StatusOK, []products.Product{})
		}
		return httpControllers.ErrorResponse(c, http.StatusInternalServerError, commons.ErrInternalServerError)
	}
	return httpControllers.SuccessResponse(c, http.StatusOK, products.FromDomainList(res))
}
