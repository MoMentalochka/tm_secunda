package api

import (
	"context"
	"database/sql"
	apiV1 "task_manager/shared/pkg/openapi/tm/v1"
)

type TaskManagerApi struct {
	db *sql.DB
}

func NewTaskManagerApi(db *sql.DB) apiV1.Invoker {
	return &TaskManagerApi{
		db: db,
	}
}

func (a *TaskManagerApi) RegisterUser(ctx context.Context, request *apiV1.RegisterRequest) (apiV1.RegisterUserRes, error) {
	return nil, nil
}
func (a *TaskManagerApi) LoginUser(ctx context.Context, request *apiV1.LoginRequest) (apiV1.LoginUserRes, error) {
	return nil, nil
}

func (a *TaskManagerApi) CreateTeam(ctx context.Context, request *apiV1.CreateTeamRequest) (apiV1.CreateTeamRes, error) {
	return nil, nil
}

func (a *TaskManagerApi) GetTeams(ctx context.Context) (apiV1.GetTeamsRes, error) {
	return nil, nil
}

func (a *TaskManagerApi) InviteToTeam(ctx context.Context, request *apiV1.InviteUserRequest, params apiV1.InviteToTeamParams) (apiV1.InviteToTeamRes, error) {
	return nil, nil
}

func (a *TaskManagerApi) CreateTask(ctx context.Context, request *apiV1.CreateTaskRequest) (apiV1.CreateTaskRes, error) {
	return nil, nil
}

func (a *TaskManagerApi) GetTasks(ctx context.Context, params apiV1.GetTasksParams) ([]apiV1.TaskResponse, error) {
	return nil, nil
}

func (a *TaskManagerApi) UpdateTask(ctx context.Context, request *apiV1.UpdateTaskRequest, params apiV1.UpdateTaskParams) (apiV1.UpdateTaskRes, error) {
	return nil, nil
}

func (a *TaskManagerApi) GetTaskHistory(ctx context.Context, params apiV1.GetTaskHistoryParams) ([]apiV1.HistoryResponse, error) {
	return nil, nil
}

func (a *TaskManagerApi) AddTaskComment(ctx context.Context, request *apiV1.CommentRequest, params apiV1.AddTaskCommentParams) (*apiV1.CommentResponse, error) {
	return &apiV1.CommentResponse{}, nil
}

func (a *TaskManagerApi) GetTaskComments(ctx context.Context, params apiV1.GetTaskCommentsParams) ([]apiV1.CommentResponse, error) {
	return nil, nil
}

func (a *TaskManagerApi) GetTeamStats(ctx context.Context, params apiV1.GetTeamStatsParams) (*apiV1.StatsResponse, error) {
	return nil, nil
}
