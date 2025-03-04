package user

import "homeworktodolist/internal/entity"

type credsData struct {
	UserID  entity.UserID  `json:"userID"`
	GroupID entity.GroupID `json:"groupID"`
}

func (d *credsData) toCreds() entity.UserCreds {
	return entity.UserCreds{
		UserID:  d.UserID,
		GroupID: d.GroupID,
	}
}
