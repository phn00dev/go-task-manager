package dto

type CreateTeamRequest struct {
	TeamName   string `json:"team_name" binding:"required,min=3,max=50"`
	TeamStatus string `json:"team_status" binding:"required"`
}

type UpdateTeamRequest struct {
	TeamName   string `json:"team_name" binding:"required,min=3,max=50"`
	TeamStatus string `json:"team_status" binding:"required"`
}
