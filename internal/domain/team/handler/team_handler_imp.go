package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/phn00dev/go-task-manager/internal/domain/team/service"
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
	panic("team handler imp")
}

func (teamHandler teamHandlerImp) GetOne(ctx *gin.Context) {
	panic("team handler imp")
}

func (teamHandler teamHandlerImp) Create(ctx *gin.Context) {
	panic("team handler imp")
}

func (teamHandler teamHandlerImp) Update(ctx *gin.Context) {
	panic("team handler imp")
}

func (teamHandler teamHandlerImp) Delete(ctx *gin.Context) {
	panic("team handler imp")
}
