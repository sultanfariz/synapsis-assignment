package auth

import "github.com/sultanfariz/synapsis-assignment/domain/users"

type Auth struct {
	Name        string `json:"name" form:"name"`
	Address     string `json:"address" form:"address"`
	PhoneNumber string `json:"phoneNumber" form:"phoneNumber"`
	Email       string `json:"email" form:"email"`
	Password    string `json:"password" form:"password"`
}

func (a Auth) ToDomain() users.User {
	return users.User{
		Name:        a.Name,
		Address:     a.Address,
		PhoneNumber: a.PhoneNumber,
		Email:       a.Email,
		Password:    a.Password,
	}
}
