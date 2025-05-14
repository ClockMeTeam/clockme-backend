package workdebt

import (
	"github.com/maevlava/ftf-clockify/internal/config"
)

/* List Api Clockify
* GetAllWorkspace	https://api.clockify.me/api/v1/workspaces
* GetALlUsers		https://api.clockify.me/api/v1/workspaces/{workspaceId}/users
* GetAllProjects	https://api.clockify.me/api/v1/workspaces/{workspaceId}/projects
* GetProjectById	https://api.clockify.me/api/v1/workspaces/{workspaceId}/projects/{projectId}
* Time Entries		https://api.clockify.me/api/v1/workspaces/{workspaceId}/user/{userId}/time-entries
 */

// TODO implement
type WorkDebtService interface {
	GetAllUserWorkDebt() (string, error)
	GetWorkDebtByProject(projectId string) (string, error)
}

type workDebtService struct {
	config *config.ApiConfig
}

func (w workDebtService) GetAllUserWorkDebt() (string, error) {
	//TODO implement me
	panic("implement me")
}

func (w workDebtService) GetWorkDebtByProject(projectId string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func NewService(cfg *config.ApiConfig) WorkDebtService {
	return &workDebtService{
		config: cfg,
	}
}
