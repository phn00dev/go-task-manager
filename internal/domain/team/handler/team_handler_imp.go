package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/phn00dev/go-task-manager/internal/domain/team/dto"
	"github.com/phn00dev/go-task-manager/internal/domain/team/service"
	bindandvalidate "github.com/phn00dev/go-task-manager/internal/utils/bind_and_validate"
	"github.com/phn00dev/go-task-manager/internal/utils/response"
)

type teamHandlerImp struct {
	teamService service.TeamService
}

func NewTeamHandler(service service.TeamService) TeamHandler {
	return teamHandlerImp{
		teamService: service,
	}
}

func (teamHandler teamHandlerImp) GetAll(ctx *gin.Context) {
	teams, err := teamHandler.teamService.GetAllTeams()
	if err != nil {
		response.Error(ctx, http.StatusNotFound, "teams not found", err.Error())
		return
	}
	teamResponses := dto.GetAllTeamResponses(teams)
	response.Success(ctx, http.StatusOK, "all teams", teamResponses)
}

func (teamHandler teamHandlerImp) GetOne(ctx *gin.Context) {
	teamIDStr := ctx.Param("teamID")
	teamID, _ := strconv.Atoi(teamIDStr)
	team, err := teamHandler.teamService.GetOneTeam(teamID)
	if err != nil {
		response.Error(ctx, http.StatusNotFound, "user not found", err.Error())
		return
	}
	teamResponse := dto.GetOneTeamResponse(team)
	response.Success(ctx, http.StatusOK, "get team", teamResponse)
}

func (teamHandler teamHandlerImp) Create(ctx *gin.Context) {
	var createTeamRequest dto.CreateTeamRequest
	if !bindandvalidate.BindAndValidate(ctx, &createTeamRequest) {
		return
	}

	// create team
	if err := teamHandler.teamService.CreateTeam(createTeamRequest); err != nil {
		response.Error(ctx, http.StatusInternalServerError, "team creation error", err.Error())
		return
	}
	response.Success(ctx, http.StatusCreated, "team created successfully", nil)
}

func (teamHandler teamHandlerImp) Update(ctx *gin.Context) {
	teamIdStr := ctx.Param("teamID")
	teamID, _ := strconv.Atoi(teamIdStr)

	var updateteamRequest dto.UpdateTeamRequest
	// bind and validate
	if !bindandvalidate.BindAndValidate(ctx, &updateteamRequest) {
		return
	}
	// update team
	if err := teamHandler.teamService.UpdateTeam(teamID, updateteamRequest); err != nil {
		response.Error(ctx, http.StatusInternalServerError, "team updated error", err.Error())
		return
	}
	response.Success(ctx, http.StatusOK, "team updated successfully", nil)
}

func (teamHandler teamHandlerImp) Delete(ctx *gin.Context) {
	teamIdStr := ctx.Param("teamID")
	teamID, _ := strconv.Atoi(teamIdStr)
	// delete team
	if err := teamHandler.teamService.DeleteTeam(teamID); err != nil {
		response.Error(ctx, http.StatusInternalServerError, "team deleted error", err.Error())
		return
	}
	response.Success(ctx, http.StatusOK, "team deleted successfully", nil)
}
