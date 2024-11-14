package admin

import (
	"context"
	adminService "homeworktodolist/internal/service/admin"
)

type AdminService interface {
	AddGroup(ctx context.Context, req adminService.AddGroup) error
}
