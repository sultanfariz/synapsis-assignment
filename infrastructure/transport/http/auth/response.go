package auth

import (
	"time"

	"github.com/sultanfariz/synapsis-assignment/domain/users"
)

type UsersResponse struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phoneNumber"`
	Email       string `json:"email"`
	// Password    string `json:"password"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// func FromUserDomainList(domain []users.User) []UsersResponse {
// 	var response []UsersResponse
// 	for _, item := range domain {
// 		response = append(response, FromUserDomain(item))
// 	}
// 	return response
// }

func FromDomain(domain users.User) UsersResponse {
	return UsersResponse{
		Id:          domain.Id,
		Name:        domain.Name,
		Address:     domain.Address,
		PhoneNumber: domain.PhoneNumber,
		Email:       domain.Email,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
	}
}
