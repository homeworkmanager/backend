package group

import (
	"context"
	"go.uber.org/mock/gomock"
	"homeworktodolist/internal/entity"
	"homeworktodolist/internal/errs"
	"testing"
)

func TestService_Create(t *testing.T) {
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		ctx   context.Context
		group entity.Group
	}
	tests := []struct {
		name    string
		mockFn  func(*MockGroupRepo)
		args    args
		want    entity.GroupID
		wantErr bool
	}{
		{
			name: "Successful case",
			mockFn: func(m *MockGroupRepo) {
				m.EXPECT().GetByName(ctx, "БСБО-01-23").Times(1).Return(entity.Group{}, errs.GroupNotFound)
				m.EXPECT().Create(ctx, gomock.Any()).Times(1).Return(entity.GroupID(1), nil)
			},
			args: args{
				ctx: ctx,
				group: entity.Group{
					Name: "БСБО-01-23",
				},
			},
			want:    1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			repo := NewMockGroupRepo(ctrl)
			tt.mockFn(repo)

			s := &Service{
				groupRepo: repo,
			}
			got, err := s.Create(tt.args.ctx, tt.args.group)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Create() got = %v, want %v", got, tt.want)
			}
		})
	}
}
