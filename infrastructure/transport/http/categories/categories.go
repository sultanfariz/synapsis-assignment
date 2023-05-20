package categories

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sultanfariz/synapsis-assignment/domain/categories"
	"github.com/sultanfariz/synapsis-assignment/infrastructure/commons"
	httpControllers "github.com/sultanfariz/synapsis-assignment/infrastructure/transport/http"
)

type Controllers struct {
	CategoriesUsecase categories.CategoriesUsecase
}

func NewControllers(categoriesUC categories.CategoriesUsecase) *Controllers {
	return &Controllers{
		CategoriesUsecase: categoriesUC,
	}
}

func (controller *Controllers) Insert(c echo.Context) error {
	ctx := c.Request().Context()
	req := Category{}
	if err := c.Bind(&req); err != nil {
		return httpControllers.ErrorResponse(c, http.StatusBadRequest, commons.ErrBadRequest)
	}
	domain := req.ToDomain()
	res, err := controller.CategoriesUsecase.Insert(ctx, domain)
	if err != nil {
		if errors.Is(err, commons.ErrUserAlreadyExists) {
			return httpControllers.ErrorResponse(c, http.StatusConflict, commons.ErrUserAlreadyExists)
		} else if errors.Is(err, commons.ErrValidationFailed) {
			return httpControllers.ErrorResponse(c, http.StatusConflict, commons.ErrValidationFailed)
		}
		return httpControllers.ErrorResponse(c, http.StatusInternalServerError, commons.ErrInternalServerError)
	}
	return httpControllers.SuccessResponse(c, http.StatusCreated, FromDomain(*res))
}

func (controller *Controllers) GetAll(c echo.Context) error {
	ctx := c.Request().Context()

	res, err := controller.CategoriesUsecase.GetAll(ctx)
	if err != nil {
		return httpControllers.ErrorResponse(c, http.StatusInternalServerError, commons.ErrInternalServerError)
	}
	return httpControllers.SuccessResponse(c, http.StatusOK, FromDomainList(res))
}
