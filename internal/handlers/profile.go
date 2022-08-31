package handlers

import (
	"github.com/cfluke-cb/ib-client-api/model"
)

func FetchProfile(id string) (model.ProfileResponse, error) {
	body := model.ProfileResponse{
		UserId:      id,
		Email:       "jay.parisi@coinbase.com",
		Name:        "Jay Parisi",
		LegalName:   "Jay Parisi",
		UserName:    "jprix",
		Roles:       []string{"user"},
		Address:     "123 Happy Canyon Way, Denver",
		DateOfBirth: "1/23/2001",
	}
	return body, nil
}
