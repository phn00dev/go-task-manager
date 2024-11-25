package dto

import "github.com/phn00dev/go-task-manager/internal/constants"

type CreateTeamRequest struct {
	TeamName   string           `json:"team_name" binding:"required,min=3,max=50"`
	TeamStatus constants.Status `json:"team_status" binding:"required"`
	AdminID    int              `json:"admin_id" binding:"required"`
}

type UpdateTeamRequest struct {
	TeamName   string           `json:"team_name" binding:"required,min=3,max=50"`
	TeamStatus constants.Status `json:"team_status" binding:"required"`
	AdminID    int              `json:"admin_id" binding:"required"`
}
