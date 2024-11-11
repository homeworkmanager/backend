package user

import "homeworktodolist/internal/entity"

type credsData struct {
	UserID  entity.UserID `json:"userID"`
	Role    entity.Role   `json:"role"`
	GroupId int64         `json:"groupID"`
}
