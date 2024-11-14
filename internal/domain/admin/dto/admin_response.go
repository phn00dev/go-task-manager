package dto

import (
	"github.com/phn00dev/go-task-manager/internal/models"
)

type AdminResponse struct {
	ID        int    `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	AdminRole string `json:"admin_role"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	LastLogin string `json:"last_login"`
}

func GetAllAdminResponses(admins []models.Admin) []AdminResponse {
	var adminResponses []AdminResponse
	for _, admin := range admins {
		adminResponse := GetOneAdminResponse(&admin)
		adminResponses = append(adminResponses, *adminResponse)
	}
	return adminResponses
}

func GetOneAdminResponse(admin *models.Admin) *AdminResponse {
	var lastLogin string
	if admin.LastLogin != nil {
		lastLogin = admin.LastLogin.Format("02-01-2006 15:04:05")
	} else {
		lastLogin = "Not Logged In"
	}
	return &AdminResponse{
		ID:        admin.ID,
		Firstname: admin.Firstname,
		Lastname:  admin.Lastname,
		Email:     admin.Email,
		AdminRole: admin.AdminRole,
		CreatedAt: admin.CreatedAt.Format("02-01-2006 15:04:05"),
		UpdatedAt: admin.UpdatedAt.Format("02-01-2006 15:04:05"),
		LastLogin: lastLogin,
	}
}
