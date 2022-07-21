package service

import "github.com/linxbin/corn-service/pkg/app"

type UserLoginRequest struct {
	Username string `form:"username" binding:"required,min=0,max=32"`
	Password string `form:"password" binding:"required,min=0,max=255"`
}

func (svc *Service) Login(request *UserLoginRequest) (string, error) {
	if err := svc.dao.MatchUser(request.Username, request.Password); err != nil {
		return "", err
	}

	auth, err := svc.dao.GetAdminAuth()
	if err != nil {
		return "", err
	}
	token, err := app.GenerateToken(auth.AppKey, auth.AppSecret)
	if err != nil {
		return "", err
	}
	return token, nil
}
