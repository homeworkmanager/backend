package homework

import (
	"context"
	"go.uber.org/mock/gomock"
	"homeworktodolist/internal/entity"
	"homeworktodolist/internal/tx_manager"
	"reflect"
	"testing"
	"time"
)

func TestService_GetByGroupAndTime(t *testing.T) {
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		ctx      context.Context
		userID   entity.UserID
		groupID  entity.GroupID
		fromTime time.Time
		toTime   time.Time
	}

	type want struct {
		group entity.HomeworkID
		err   error
	}

	tests := []struct {
		name    string
		mockFn  func(*MockHomeworkRepo)
		args    args
		want    []entity.Homework
		wantErr bool
	}{
		{
			name:    "Successful case",
			mockFn: func(m *MockHomeworkRepo) {
				m.EXPECT().GetByGroupAndTime(ctx, gomock.Any(),gomock.Any(),gomock.Any(),gomock.Any()).Times(1).Return([]entity.Homework{}, nil)
			},
			args:    args{
				ctx:      ctx,
				userID:   1,
				groupID:  1,
				fromTime: time.Now(),
				toTime:   time.Now(),
			},
			want:    nil,
			wantErr: false,
		},
	},
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				homeworkRepo:          tt.fields.homeworkRepo,
				homeworkStatusService: tt.fields.homeworkStatusService,
				manager:               tt.fields.manager,
			}
			got, err := s.GetByGroupAndTime(tt.args.ctx, tt.args.userID, tt.args.groupID, tt.args.fromTime, tt.args.toTime)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByGroupAndTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetByGroupAndTime() got = %v, want %v", got, tt.want)
			}
		})
	}
}
