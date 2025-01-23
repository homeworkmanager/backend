package admin

import (
	"context"

	adminService "homeworktodolist/internal/service/admin"
)

type AdminService interface {
	AddGroup(ctx context.Context, req adminService.AddGroup) error
	UpdateClasses(ctx context.Context) error
	RefreshAllData(ctx context.Context) error
	UpdateRole(ctx context.Context, req adminService.UpdateUserRole) error
}
