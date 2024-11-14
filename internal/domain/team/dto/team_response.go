package dto

import "github.com/phn00dev/go-task-manager/internal/models"

type TeamResponse struct {
	ID         int      `json:"id"`
	TeamName   string   `json:"team_name"`
	TeamStatus string   `json:"team_status"`
	AdminID    int      `json:"admin_id"`
	CreatedAt  string   `json:"created_at"`
	UpdatedAt  string   `json:"updated_at"`
	TeamLead   TeamLead `json:"team_lead"`
}

type TeamLead struct {
	ID        int    `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	AdminRole string `json:"admin_role"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func GetOneTeamResponse(team *models.Team) *TeamResponse {
	return &TeamResponse{
		ID:         team.ID,
		TeamName:   team.TeamName,
		TeamStatus: team.TeamStatus,
		AdminID:    team.AdminID,
		CreatedAt:  team.CreatedAt.Format("02-01-2006 15:04:05"),
		UpdatedAt:  team.UpdatedAt.Format("02-01-2006 15:04:05"),
		TeamLead: TeamLead{
			ID:        team.Admin.ID,
			Firstname: team.Admin.Firstname,
			Lastname:  team.Admin.Lastname,
			Email:     team.Admin.Email,
			AdminRole: team.Admin.AdminRole,
			CreatedAt: team.Admin.CreatedAt.Format("02-01-2006 15:04:05"),
			UpdatedAt: team.Admin.UpdatedAt.Format("02-01-2006 15:04:05"),
		},
	}
}

type AllTeamResponse struct {
	ID         int    `json:"id"`
	TeamName   string `json:"team_name"`
	TeamStatus string `json:"team_status"`
	AdminID    int    `json:"admin_id"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

func GetAllTeamResponses(teams []models.Team) []AllTeamResponse {
	var teamResponses []AllTeamResponse
	for _, team := range teams {
		teamResponse := AllTeamResponse{
			ID:         team.ID,
			TeamName:   team.TeamName,
			TeamStatus: team.TeamStatus,
			AdminID:    team.Admin.ID,
			CreatedAt:  team.CreatedAt.Format("02-01-2006 15:04:05"),
			UpdatedAt:  team.UpdatedAt.Format("02-01-2006 15:04:05"),
		}
		teamResponses = append(teamResponses, teamResponse)
	}
	return teamResponses
}
