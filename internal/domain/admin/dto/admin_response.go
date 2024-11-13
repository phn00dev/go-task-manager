package dto

import (
	"time"

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
	turkmenistanLocation, err := time.LoadLocation("Asia/Ashgabat")
	if err != nil {

		turkmenistanLocation = time.UTC
	}
	return &AdminResponse{
		ID:        admin.ID,
		Firstname: admin.Firstname,
		Lastname:  admin.Lastname,
		Email:     admin.Email,
		AdminRole: admin.AdminRole,
		CreatedAt: admin.CreatedAt.In(turkmenistanLocation).Format("02-01-2006 15:04:05"),
		UpdatedAt: admin.UpdatedAt.In(turkmenistanLocation).Format("02-01-2006 15:04:05"),
		LastLogin: admin.LastLogin.In(turkmenistanLocation).Format("02-01-2006 15:04:05"),
	}
}
