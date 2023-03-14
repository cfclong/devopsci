package main

import (
	"runtime"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql" // import your used driver

	"gitee.com/plutoccc/devops_app/internal/cronjob"
	"gitee.com/plutoccc/devops_app/internal/initialize"
	"gitee.com/plutoccc/devops_app/internal/migrations"
	"gitee.com/plutoccc/devops_app/internal/models"
	"gitee.com/plutoccc/devops_app/internal/routers"
)

func main() {
	models.InitDB()
	migrations.Migrate()
	// TODO: resource items migrate later
	initialize.Init()

	cronjob.RunPublishJobServer()

	routers.RegisterRoutes()
	beego.Info("Beego version:", beego.VERSION)
	beego.Info("Golang version:", runtime.Version())
	beego.Run()
}
