package admin

import (
	"context"
	"homeworktodolist/internal/entity"

	adminService "homeworktodolist/internal/service/admin"
)

type AdminService interface {
	AddGroup(ctx context.Context, req adminService.AddGroup) error
	UpdateClasses(ctx context.Context) error
	RefreshAllData(ctx context.Context) error
	UpdateRole(ctx context.Context, req adminService.UpdateUserRole) error
	GetAllUsers(ctx context.Context) ([]entity.UserFullInfo, error)
}
