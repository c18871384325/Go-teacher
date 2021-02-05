package auth

import (
	"net/http"

	"github.com/astaxie/beego"

	"cmdb/base/controllers/base"
	"cmdb/models"
	"cmdb/services"
)

// AuthorizationController 所有需要认证才能访问的基础控制器
type AuthorizationController struct {
	base.BaseController
	LoginUser *models.User
}

// Prepare 用户认证检查
func (c *AuthorizationController) Prepare() {
	sessionKey := beego.AppConfig.DefaultString("auth::SessionKey", "user")
	sessonValue := c.GetSession(sessionKey)
	c.Data["loginUser"] = nil

	if sessonValue != nil {
		if pk, ok := sessonValue.(int); ok {
			if user := services.UserService.GetByPk(pk); user != nil {
				c.Data["loginUser"] = user
				c.LoginUser = user
				return
			}
		}
	}

	action := beego.AppConfig.DefaultString("auth::LoginAction",
		"AthController.Login")
	c.Redirect(beego.URLFor(action), http.StatusFound)
}
