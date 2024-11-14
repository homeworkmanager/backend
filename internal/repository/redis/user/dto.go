package user

import "homeworktodolist/internal/entity"

type credsData struct {
	UserID  entity.UserID  `json:"userID"`
	Role    entity.Role    `json:"role"`
	GroupID entity.GroupID `json:"groupID"`
}

func (d *credsData) toCreds() entity.UserCreds {
	return entity.UserCreds{
		UserID:  d.UserID,
		Role:    d.Role,
		GroupID: d.GroupID,
	}
}
