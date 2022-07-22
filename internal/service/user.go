package service

import "github.com/linxbin/corn-service/pkg/app"

type UserLoginRequest struct {
	Username string `form:"username" binding:"required,min=0,max=32"`
	Password string `form:"password" binding:"required,min=0,max=255"`
}

type UserInfo struct {
	Token    string `json:"token"`
	Username string `json:"username"`
	Id       uint32 `json:"id"`
}

func (svc *Service) Login(request *UserLoginRequest) (*UserInfo, error) {
	user, err := svc.dao.MatchUser(request.Username, request.Password)
	if err != nil {
		return nil, err
	}

	auth, err := svc.dao.GetAdminAuth()
	if err != nil {
		return nil, err
	}
	token, err := app.GenerateToken(auth.AppKey, auth.AppSecret)
	if err != nil {
		return nil, err
	}
	u := &UserInfo{
		Username: user.Username,
		Id:       user.ID,
		Token:    token,
	}
	return u, nil
}
