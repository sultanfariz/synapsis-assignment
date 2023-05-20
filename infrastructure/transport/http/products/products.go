package products

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sultanfariz/synapsis-assignment/domain/products"
	"github.com/sultanfariz/synapsis-assignment/infrastructure/commons"
	httpControllers "github.com/sultanfariz/synapsis-assignment/infrastructure/transport/http"
)

type Controllers struct {
	ProductsUsecase products.ProductsUsecase
}

func NewControllers(productsUC products.ProductsUsecase) *Controllers {
	return &Controllers{
		ProductsUsecase: productsUC,
	}
}

func (controller *Controllers) Insert(c echo.Context) error {
	ctx := c.Request().Context()
	req := Product{}
	if err := c.Bind(&req); err != nil {
		return httpControllers.ErrorResponse(c, http.StatusBadRequest, commons.ErrBadRequest)
	}
	domain := req.ToDomain()
	res, err := controller.ProductsUsecase.Insert(ctx, &domain)
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

func (controller *Controllers) GetAll(c echo.Context) error {
	ctx := c.Request().Context()
	category := c.QueryParam("category")
	if category != "" {
		res, err := controller.ProductsUsecase.GetByCategory(ctx, category)
		if err != nil {
			if errors.Is(err, commons.ErrCategoryNotFound) {
				return httpControllers.ErrorResponse(c, http.StatusConflict, commons.ErrCategoryNotFound)
			}
			return httpControllers.ErrorResponse(c, http.StatusInternalServerError, commons.ErrInternalServerError)
		}
		return httpControllers.SuccessResponse(c, http.StatusOK, FromDomainList(res))
	}

	res, err := controller.ProductsUsecase.GetAll(ctx)
	if err != nil {
		return httpControllers.ErrorResponse(c, http.StatusInternalServerError, commons.ErrInternalServerError)
	}
	return httpControllers.SuccessResponse(c, http.StatusOK, FromDomainList(res))
}
