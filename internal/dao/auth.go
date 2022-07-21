package dao

import "github.com/linxbin/corn-service/internal/model"

func (d *Dao) GetAuth(appKey, appSecret string) (model.Auth, error) {
	auth := model.Auth{AppKey: appKey, AppSecret: appSecret}
	return auth.Get(d.engine)
}

func (d *Dao) GetAdminAuth() (model.Auth, error) {
	auth := model.Auth{AppKey: "eddycjy", AppSecret: "go-corn-admin"}
	return auth.Get(d.engine)
}
