package dto

type CreateTeamRequest struct {
	TeamName string `json:"team_name" binding:"required,min=3,max=50"`
}

type UpdateTeamRequest struct {
	TeamName string `json:"team_name" binding:"required,min=3,max=50"`
}
