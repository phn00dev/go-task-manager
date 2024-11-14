package dto

import (
	"github.com/phn00dev/go-task-manager/internal/models"
)

type LoginAdminResponse struct {
	AdminData   AdminResponse `json:"admin_data"`
	AccessToken string        `json:"access_token"`
}

func NewLoginAdminResponse(admin *models.Admin, accessToken string) *LoginAdminResponse {
	adminData := GetOneAdminResponse(admin)
	return &LoginAdminResponse{
		AdminData:   *adminData,
		AccessToken: accessToken,
	}
}
