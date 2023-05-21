package users

import (
	"context"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/sultanfariz/synapsis-assignment/infrastructure/commons"
	"golang.org/x/crypto/bcrypt"
)

type UsersUsecase struct {
	UsersRepository UsersRepositoryInterface
	ContextTimeout  time.Duration
	jwtConfig       *commons.ConfigJWT
}

func NewUsersUsecase(ur UsersRepositoryInterface, timeout time.Duration, jwtConfig *commons.ConfigJWT) *UsersUsecase {
	return &UsersUsecase{
		UsersRepository: ur,
		ContextTimeout:  timeout,
		jwtConfig:       jwtConfig,
	}
}

func (uu *UsersUsecase) Register(ctx context.Context, user *User) (*User, error) {
	ctx, cancel := context.WithTimeout(ctx, uu.ContextTimeout)
	defer cancel()

	validator := validator.New()
	if err := validator.Struct(user); err != nil {
		return nil, commons.ErrValidationFailed
	}

	// hash password
	hashed, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashed)

	// check if email already exists
	data, err := uu.UsersRepository.GetByEmail(ctx, user.Email)
	if err != nil && err.Error() != "record not found" {
		return nil, err
	}
	if data != nil {
		return nil, commons.ErrUserAlreadyExists
	}

	// insert user to db
	user, err = uu.UsersRepository.Insert(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (uu *UsersUsecase) Login(ctx context.Context, user *User) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, uu.ContextTimeout)
	defer cancel()

	validator := validator.New()
	if err := validator.Struct(User{
		Email:    user.Email,
		Password: user.Password,
	}); err != nil {
		return "", commons.ErrValidationFailed
	}

	// check if user exists
	data, err := uu.UsersRepository.GetByEmail(ctx, user.Email)
	if err != nil && err.Error() != "record not found" {
		return "", err
	}
	if data == nil {
		return "", commons.ErrEmptyInput
	}

	// check if password is correct
	if err := bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(user.Password)); err != nil {
		return "", commons.ErrEmptyInput
	}

	fmt.Printf("user: %+v\n", data)
	// generate token
	token, err := uu.jwtConfig.GenerateToken(data.Id, data.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}
