package dao

import "github.com/linxbin/corn-service/internal/model"

func (d *Dao) GetAdminAuth() (model.Auth, error) {
	auth := model.Auth{AppKey: "eddycjy", AppSecret: "go-corn-admin"}
	return auth.Get(d.engine)
}
