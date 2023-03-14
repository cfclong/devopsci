package middleware

import (
	"gitee.com/plutoccc/devops_app/internal/middleware/log"
	mycasbin "github.com/go-start/start/internal/middleware/casbin"

	"github.com/astaxie/beego/context"
)

// Authorization 鉴权
func Authorization(c *context.Context, username string) (bool, error) {
	e, err := mycasbin.NewCasbin()
	if err != nil {
		log.Log.Error("casbin new occur error: %v", err.Error())
		return false, err
	}
	urlPath := c.Request.URL.Path
	urlMethod := c.Request.Method
	res, err := e.Enforce(username, urlPath, urlMethod)
	// TODO: user constraint permission
	// based on  urlpath get resource type, then get resource constraint
	log.Log.Debug("role key: %s, path: %s, method: %s, res: %v", username, urlPath, urlMethod, res)
	return res, err
}
