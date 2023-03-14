package initialize

import (
	"os"

	"gitee.com/plutoccc/devops_app/internal/middleware/log"
	"gitee.com/plutoccc/devops_app/utils/errors"
)

func Init() {

	// 注册/更新资源
	initResource()

	// 初始化/更新路由
	initRouterItems()

	// 更新所有用户权限策略
	// TODO: confirm
	// func initUsers(){
	// users, _ := dao.UserList()
	// for _, user := range users {
	// 	dao.InitSystemMember(user)
	// }
	//}()

	// 初始化系统组
	if err := InitAdminUserAndGroup(); err != nil {
		if !errors.OrmError1062(err) {
			log.Log.Error(err.Error())
			os.Exit(2)
		}
	}
}
