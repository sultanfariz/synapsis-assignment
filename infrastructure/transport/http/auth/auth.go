package auth

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sultanfariz/synapsis-assignment/domain/users"
	"github.com/sultanfariz/synapsis-assignment/infrastructure/commons"
	httpControllers "github.com/sultanfariz/synapsis-assignment/infrastructure/transport/http"
)

type Controllers struct {
	UsersUsecase users.UsersUsecase
}

func NewControllers(usersUC users.UsersUsecase) *Controllers {
	return &Controllers{
		UsersUsecase: usersUC,
	}
}

func (controller *Controllers) Register(c echo.Context) error {
	ctx := c.Request().Context()
	req := Auth{}
	if err := c.Bind(&req); err != nil {
		return httpControllers.ErrorResponse(c, http.StatusBadRequest, commons.ErrBadRequest)
	}
	domain := req.ToDomain()
	res, err := controller.UsersUsecase.Register(ctx, &domain)
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

func (controller *Controllers) Login(c echo.Context) error {
	ctx := c.Request().Context()
	req := Auth{}
	if err := c.Bind(&req); err != nil {
		return httpControllers.ErrorResponse(c, http.StatusBadRequest, commons.ErrBadRequest)
	}
	domain := req.ToDomain()
	res, err := controller.UsersUsecase.Login(ctx, &domain)
	if err != nil {
		if errors.Is(err, commons.ErrInvalidCredentials) {
			return httpControllers.ErrorResponse(c, http.StatusConflict, commons.ErrInvalidCredentials)
		}
		if errors.Is(err, commons.ErrValidationFailed) {
			return httpControllers.ErrorResponse(c, http.StatusConflict, commons.ErrValidationFailed)
		}
		return httpControllers.ErrorResponse(c, http.StatusInternalServerError, commons.ErrInternalServerError)
	}
	cookie := commons.CreateCookie(res)
	c.SetCookie(cookie)
	return httpControllers.SuccessResponse(c, http.StatusOK, res)
}

// func (controller *Controllers) Logout(c echo.Context) error {
// 	cookie := commons.DeleteCookie()
// 	c.SetCookie(cookie)
// 	return httpControllers.SuccessResponse(c, http.StatusOK, nil)
// }
