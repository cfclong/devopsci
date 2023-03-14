package api

import (
	"gitee.com/plutoccc/devops_app/internal/dao"
	"gitee.com/plutoccc/devops_app/internal/middleware/log"
)

type AuditController struct {
	BaseController
}

func (ac *AuditController) AuditList() {
	res, err := dao.AuditList()
	if err != nil {
		ac.HandleInternalServerError(err.Error())
		log.Log.Error("Get audit list error: %s", err.Error())
		return
	}
	ac.Data["json"] = NewResult(true, res, "")
	ac.ServeJSON()
}
