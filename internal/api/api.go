package api

import (
	"context"
	"task_manager/internal/service"
	"task_manager/platform/pkg/logger"
	apiV1 "task_manager/shared/pkg/openapi/tm/v1"

	"go.uber.org/zap"
)

type api struct {
	tmUserService service.UserService
}

func NewApi(tmUserService service.UserService) apiV1.Invoker {
	return &api{
		tmUserService: tmUserService,
	}
}

func (a *api) RegisterUser(ctx context.Context, request *apiV1.RegisterRequest) (apiV1.RegisterUserRes, error) {
	res, err := a.tmUserService.Register(ctx, request.Name, request.Email, request.Password)
	if err != nil {
		logger.Info(ctx, "apiV1.RegisterUser", zap.Error(err))
		return &apiV1.UserResponse{}, err
	}

	return &apiV1.UserResponse{
		ID:        res.ID,
		Email:     res.Email,
		Name:      res.Name,
		CreatedAt: res.CreatedAt,
	}, nil
}
func (a *api) LoginUser(ctx context.Context, request *apiV1.LoginRequest) (apiV1.LoginUserRes, error) {

	res, err := a.tmUserService.Login(ctx, request.Email, request.Password)
	if err != nil {
		logger.Info(ctx, "apiV1.LoginUser", zap.Error(err))
		return &apiV1.LoginResponse{}, err
	}

	return &apiV1.LoginResponse{Token: res}, nil
}

func (a *api) CreateTeam(ctx context.Context, request *apiV1.CreateTeamRequest) (apiV1.CreateTeamRes, error) {
	return nil, nil
}

func (a *api) GetTeams(ctx context.Context) (apiV1.GetTeamsRes, error) {
	return nil, nil
}

func (a *api) InviteToTeam(ctx context.Context, request *apiV1.InviteUserRequest, params apiV1.InviteToTeamParams) (apiV1.InviteToTeamRes, error) {
	return nil, nil
}

func (a *api) CreateTask(ctx context.Context, request *apiV1.CreateTaskRequest) (apiV1.CreateTaskRes, error) {
	return nil, nil
}

func (a *api) GetTasks(ctx context.Context, params apiV1.GetTasksParams) ([]apiV1.TaskResponse, error) {
	return nil, nil
}

func (a *api) UpdateTask(ctx context.Context, request *apiV1.UpdateTaskRequest, params apiV1.UpdateTaskParams) (apiV1.UpdateTaskRes, error) {
	return nil, nil
}

func (a *api) GetTaskHistory(ctx context.Context, params apiV1.GetTaskHistoryParams) ([]apiV1.HistoryResponse, error) {
	return nil, nil
}

func (a *api) AddTaskComment(ctx context.Context, request *apiV1.CommentRequest, params apiV1.AddTaskCommentParams) (*apiV1.CommentResponse, error) {
	return &apiV1.CommentResponse{}, nil
}

func (a *api) GetTaskComments(ctx context.Context, params apiV1.GetTaskCommentsParams) ([]apiV1.CommentResponse, error) {
	return nil, nil
}

func (a *api) GetTeamStats(ctx context.Context, params apiV1.GetTeamStatsParams) (*apiV1.StatsResponse, error) {
	return nil, nil
}
