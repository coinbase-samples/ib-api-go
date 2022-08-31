package handlers

import (
	"github.com/cfluke-cb/ib-client-api/data"
)

func FetchProfile(id string) (data.ProfileResponse, error) {
	body := data.ProfileResponse{
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
