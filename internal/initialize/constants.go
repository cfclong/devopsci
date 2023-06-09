package initialize

var resourceReq = ResourceReq{
	Resources: []BatchResourceTypeSpec{
		BatchResourceTypeSpec{
			ResourceType: []string{"*", "所有资源"},
			ResourceOperation: [][]string{
				[]string{"*", "所有操作"},
			},
			ResourceConstraint: [][]string{
				[]string{"*", "所有约束"},
			},
		},
		BatchResourceTypeSpec{
			ResourceType: []string{"auth", "认证"},
			ResourceOperation: [][]string{
				[]string{"*", "认证所有操作"},
				[]string{"UserLogin", "用户登录"},
				[]string{"UserLogout", "用户登出"},
				[]string{"GetCurrentUser", "获取当前用户信息"},
			},
			ResourceConstraint: [][]string{},
		},
		BatchResourceTypeSpec{
			ResourceType: []string{"audit", "操作审计"},
			ResourceOperation: [][]string{
				[]string{"*", "操作审计所有操作"},
				[]string{"AuditList", "获取操作审计列表"},
			},
			ResourceConstraint: [][]string{},
		},
		BatchResourceTypeSpec{
			ResourceType: []string{"user", "用户"},
			ResourceOperation: [][]string{
				[]string{"*", "用户所有操作"},
				[]string{"UserList", "获取用户列表"},
				[]string{"CreateUser", "创建用户"},
				[]string{"GetUser", "获取用户详情"},
				[]string{"UpdateUser", "更新用户"},
				[]string{"DeleteUser", "删除用户"},
				[]string{"GetUserResourceConstraintValues", "获取用户资源约束的值"},
			},
			ResourceConstraint: [][]string{
				[]string{"user", "用户账号"},
			},
		},
		BatchResourceTypeSpec{
			ResourceType: []string{"repository", "我的应用"},
			ResourceOperation: [][]string{
				[]string{"*", "我的应用所有操作"},
				[]string{"GetAppsByPagination", "获取应用分页列表"},
				[]string{"GetAllApps", "获取应用列表"},
				[]string{"GetSCMIntegrateSettings", "代码源集成列表"},
				[]string{"CreateSCMApp", "创建代码应用"},
				[]string{"VerifySCMAppConnetion", "验证代码应用"},
				[]string{"GetScmApp", "获取代码应用信息"},
				[]string{"UpdateScmApp", "更新代码应用信息"},
				[]string{"DeleteScmApp", "删除代码应用"},

				[]string{"GetAppBranches", "获取应用分支"},
				[]string{"SyncAppBranches", "同步远程分支"},
				[]string{"GetGitProjectsByRepoID", "获取代码仓库项目列表"},
			},
			ResourceConstraint: [][]string{
				[]string{"project_id", "项目ID"},
			},
		},

		BatchResourceTypeSpec{
			ResourceType: []string{"project", "项目"},
			ResourceOperation: [][]string{
				[]string{"*", "项目所有操作"},
				[]string{"ProjectList", "获取项目列表"},
				[]string{"CreateProject", "创建项目"},
				[]string{"UpdateProject", "更新项目信息"},
				[]string{"DeleteProject", "删除项目"},
				[]string{"GetProject", "获取项目信息"},
				[]string{"GetprojectMemberByConstraint", "获取项目成员信息"},

				[]string{"CreateProjectApp", "项目添加应用"},
				[]string{"UpdateProjectApp", "更新项目应用"},
				[]string{"GetProjectApps", "获取项目应用列表"},
				[]string{"GetProjectApp", "获取项目应用详情"},
				[]string{"GetProjectAppsByPagination", "获取项目应用分页列表"},
				[]string{"GetArrange", "获取应用编排"},
				[]string{"SetArrange", "设置应用编排"},
				[]string{"DeleteProjectApp", "删除项目应用"},
				[]string{"ParserAppArrange", "应用编排解析"},
				[]string{"GetJenkinsConfig", "获取Jenkins配置"},

				[]string{"GetProjectEnvs", "项目环境列表"},
				[]string{"GetProjectPipelinesByPagination", "项目流程分页列表"},

				// project app service
				[]string{"GetProjectAppServices", "应用服务列表"},
				[]string{"GetAppServiceInspect", "应用服务详情"},
				[]string{"GetAppServiceLog", "应用服务日志"},
				[]string{"GetAppServiceEvent", "应用服务事件"},
				[]string{"AppServiceRestart", "重启应用服务"},
				[]string{"AppServiceScale", "水平扩展应用服务"},
				[]string{"AppServiceTerminal", "应用服务终端调试"},
				[]string{"DeleteAppService", "删除应用服务"},

				// project pipeline
				[]string{"PipelineCreate", "创建项目流程"},
				[]string{"PipelineUpdate", "更新流程基础信息"},
				[]string{"ProjectPipelineInfo", "获取项目流程信息"},
				[]string{"PipelineDelete", "删除项目流程"},
				[]string{"FlowStepList", "获取任务模板列表"},

				[]string{"GetProjectEnvsByPagination", "项目环境分页列表"},
				[]string{"CreateProjectEnv", "新建项目环境"},
				[]string{"UpdateProjectEnv", "更新项目环境"},
				[]string{"ProjectAppServiceStats", "获取项目应用统计"},
			},
			ResourceConstraint: [][]string{
				[]string{"project_id", "项目ID"},
			},
		},
		BatchResourceTypeSpec{
			ResourceType: []string{"publish", "流水线"},
			ResourceOperation: [][]string{
				[]string{"*", "流水线所有操作"},
				[]string{"GetProjectPipelines", "项目流程列表"},
				[]string{"PublishList", "流水线列表"},
				[]string{"CreatePublishOrder", "创建流水线"},
				[]string{"GetPublish", "流水线详情"},
				[]string{"ClosePublish", "关闭流水线"},
				[]string{"DeletePublish", "删除流水线"},
				[]string{"GetCanAddedApps", "获取可添加应用列表"},
				[]string{"AddPublishApp", "版本添加应用"},
				[]string{"DeletePublishApp", "版本删除应用"},
				[]string{"GetOpertaionLogByPagination", "获取流水线操作日志"},
				[]string{"GetBackTo", "获取回退列表"},
				[]string{"TriggerBackTo", "触发流水线回退操作"},
				[]string{"GetNextStage", "获取流转列表"},
				[]string{"TriggerNextStage", "触发流水线流转操作"},
				[]string{"GetStepInfo", "获取步骤执行信息"},
				[]string{"RunStep", "触发步骤执行"},
				[]string{"RunStepCallback", "步骤执行回调"},
			},
			ResourceConstraint: [][]string{
				[]string{"project_id", "项目ID"},
				[]string{"publishID", "发布单ID"},
				[]string{"envID", "环境ID"},
			},
		},
		BatchResourceTypeSpec{
			ResourceType: []string{"system", "系统设置"},
			ResourceOperation: [][]string{
				[]string{"*", "系统设置所有操作"},
				[]string{"GetCompileEnvs", "编译环境列表"},
				[]string{"GetIntegrateClusters", "获取集成的集群列表"},
				[]string{"GetIntegrateSettings", "获取集成配置列表"},

				[]string{"FlowComponentList", "获取基础组件列表"},
				[]string{"FlowStepListByPagination", "获取任务模板分页列表"},
				[]string{"FlowStepCreate", "创建任务模板"},
				[]string{"FlowStepUpdate", "更新任务模板"},
				[]string{"FlowStepDelete", "删除任务模板"},
			},
			ResourceConstraint: [][]string{},
		},
	},
}

var gaetwayReq = RouterReq{
	Routers: [][]string{
		[]string{"start/api/v1/login", "POST", "start", "auth", "UserLogin"},
		[]string{"start/api/v1/logout", "GET", "start", "auth", "UserLogout"},
		[]string{"start/api/v1/getCurrentUser", "GET", "start", "auth", "GetCurrentUser"},
		[]string{"start/api/v1/audit", "GET", "start", "audit", "AuditList"},
		[]string{"start/api/v1/users", "GET", "start", "user", "UserList"},
		[]string{"start/api/v1/users", "POST", "start", "user", "CreateUser"},
		[]string{"start/api/v1/users/:user", "GET", "start", "user", "GetUser"},
		[]string{"start/api/v1/users/:user", "PUT", "start", "user", "UpdateUser"},
		[]string{"start/api/v1/users/:user", "DELETE", "start", "user", "DeleteUser"},
		[]string{"start/api/v1/users/:user/resources/:resourceType/constraints/values", "GET", "start", "user", "GetUserResourceConstraintValues"},
		[]string{"start/api/v1/groups", "GET", "start", "group", "GroupList"},
		[]string{"start/api/v1/groups/:group", "GET", "start", "group", "GetGroup"},
		[]string{"start/api/v1/groups/:group", "PUT", "start", "group", "UpdateGroup"},
		[]string{"start/api/v1/groups/:group", "DELETE", "start", "group", "DeleteGroup"},
		[]string{"start/api/v1/groups/:group/users", "GET", "start", "group", "GroupUserList"},
		[]string{"start/api/v1/groups/:group/users", "POST", "start", "group", "AddGroupUsers"},
		[]string{"start/api/v1/groups/:group/users/:user", "PUT", "start", "group", "UpdateGroupUser"},
		[]string{"start/api/v1/groups/:group/users/:user", "DELETE", "start", "group", "RemoveGroupUser"},

		// app repo
		[]string{"start/api/v1/integrate/settings/scms", "GET", "start", "repository", "GetSCMIntegrateSettings"},
		[]string{"start/api/v1/apps/create", "POST", "start", "repository", "CreateSCMApp"},
		[]string{"start/api/v1/apps/verifyapp", "POST", "start", "repository", "VerifySCMAppConnetion"},
		[]string{"start/api/v1/apps", "POST", "start", "repository", "GetAppsByPagination"},
		[]string{"start/api/v1/apps", "GET", "start", "repository", "GetAllApps"},
		[]string{"start/api/v1/repos/:repo_id/projects", "POST", "start", "repository", "GetGitProjectsByRepoID"},
		[]string{"start/api/v1/apps/:app_id/branches", "POST", "start", "repository", "GetAppBranches"},
		[]string{"start/api/v1/apps/:app_id/syncBranches", "POST", "start", "repository", "SyncAppBranches"},
		[]string{"start/api/v1/apps/:app_id", "GET", "start", "repository", "GetScmApp"},
		[]string{"start/api/v1/apps/:app_id", "PUT", "start", "repository", "UpdateScmApp"},
		[]string{"start/api/v1/apps/:app_id", "DELETE", "start", "repository", "DeleteScmApp"},

		// project
		[]string{"start/api/v1/projects", "POST", "start", "project", "ProjectList"},
		[]string{"start/api/v1/users/:project_id/projectMemberByConstraint", "GET", "start", "project", "GetprojectMemberByConstraint"},
		[]string{"start/api/v1/projects/create", "POST", "start", "project", "CreateProject"},
		[]string{"start/api/v1/projects/:project_id", "PUT", "start", "project", "UpdateProject"},
		[]string{"start/api/v1/projects/:project_id", "DELETE", "start", "project", "DeleteProject"},
		[]string{"start/api/v1/projects/:project_id", "GET", "start", "project", "GetProject"},
		[]string{"start/api/v1/projects/:project_id/pipelines", "GET", "start", "project", "GetProjectPipelines"},
		[]string{"start/api/v1/projects/:project_id/pipelines", "POST", "start", "project", "GetProjectPipelinesByPagination"},
		[]string{"start/api/v1/pipelines/flow/steps", "GET", "start", "project", "FlowStepList"},
		[]string{"start/api/v1/projects/:project_id/pipelines/create", "POST", "start", "project", "PipelineCreate"},
		[]string{"start/api/v1/projects/:project_id/pipelines/:id", "GET", "start", "project", "ProjectPipelineInfo"},
		[]string{"start/api/v1/projects/:project_id/pipelines/:id", "PUT", "start", "project", "PipelineUpdate"},
		[]string{"start/api/v1/projects/:project_id/pipelines/:id", "DELETE", "start", "project", "PipelineDelete"},
		[]string{"start/api/v1/projects/:project_id/apps/create", "POST", "start", "project", "CreateProjectApp"},
		[]string{"start/api/v1/projects/:project_id/apps", "GET", "start", "project", "GetProjectApps"},
		[]string{"start/api/v1/projects/:project_id/apps/:project_app_id", "GET", "start", "project", "GetProjectApp"},
		[]string{"start/api/v1/projects/:project_id/apps", "POST", "start", "project", "GetProjectAppsByPagination"},
		[]string{"start/api/v1/projects/:project_id/apps/:app_id/:arrange_env/arrange", "GET", "start", "project", "GetArrange"},
		[]string{"start/api/v1/projects/:project_id/apps/:app_id/:arrange_env/arrange", "POST", "start", "project", "SetArrange"},
		[]string{"start/api/v1/arrange/yaml/parser", "POST", "start", "project", "ParserAppArrange"},
		[]string{"start/api/v1/pipelines/stages/:stage_id/jenkins-config", "GET", "start", "project", "GetJenkinsConfig"},

		[]string{"start/api/v1/projects/:project_id/apps/:project_app_id", "PUT", "start", "project", "UpdateProjectApp"},
		[]string{"start/api/v1/projects/:project_id/apps/:project_app_id", "DELETE", "start", "project", "DeleteProjectApp"},
		[]string{"start/api/v1/projects/:project_id/clusters/:cluster/apps", "POST", "start", "project", "GetProjectAppServices"},
		[]string{"start/api/v1/clusters/:cluster/namespaces/:namespace/apps/:app", "GET", "start", "project", "GetAppServiceInspect"},
		[]string{"start/api/v1/clusters/:cluster/namespaces/:namespace/apps/:app", "DELETE", "start", "project", "DeleteAppService"},
		[]string{"start/api/v1/clusters/:cluster/namespaces/:namespace/apps/:app/log", "GET", "start", "project", "GetAppServiceLog"},
		[]string{"start/api/v1/clusters/:cluster/namespaces/:namespace/apps/:app/event", "GET", "start", "project", "GetAppServiceEvent"},
		[]string{"start/api/v1/clusters/:cluster/namespaces/:namespace/apps/:app/restart", "POST", "start", "project", "AppServiceRestart"},
		[]string{"start/api/v1/clusters/:cluster/namespaces/:namespace/apps/:app/scale", "POST", "start", "project", "AppServiceScale"},
		[]string{"start/api/v1/clusters/:cluster/namespaces/:namespace/pods/:podname/containernames/:containername", "GET", "start", "project", "AppServiceTerminal"},

		[]string{"start/api/v1/projects/:project_id/publish/stats", "POST", "start", "project", "ProjectPublishStats"},
		[]string{"start/api/v1/projects/:project_id/envs", "GET", "start", "project", "GetProjectEnvs"},
		[]string{"start/api/v1/projects/:project_id/envs", "POST", "start", "project", "GetProjectEnvsByPagination"},
		[]string{"start/api/v1/projects/:project_id/envs/create", "POST", "start", "project", "CreateProjectEnv"},
		[]string{"start/api/v1/projects/:project_id/envs/:env_id", "PUT", "start", "project", "UpdateProjectEnv"},

		// publish
		[]string{"start/api/v1/projects/:project_id/publishes", "POST", "start", "publish", "PublishList"},
		[]string{"start/api/v1/projects/:project_id/publishes/create", "POST", "start", "publish", "CreatePublishOrder"},
		[]string{"start/api/v1/projects/:project_id/publishes/:publish_id", "GET", "start", "publish", "GetPublish"},
		[]string{"start/api/v1/projects/:project_id/publishes/:publish_id", "PUT", "start", "publish", "ClosePublish"},
		[]string{"start/api/v1/projects/:project_id/publishes/:publish_id", "DELETE", "start", "publish", "DeletePublish"},
		[]string{"start/api/v1/projects/:project_id/publishes/:publish_id/apps/can_added", "GET", "start", "publish", "GetCanAddedApps"},
		[]string{"start/api/v1/projects/:project_id/publishes/:publish_id/apps/create", "POST", "start", "publish", "AddPublishApp"},
		[]string{"start/api/v1/projects/:project_id/publishes/:publish_id/apps/:publish_app_id", "DELETE", "start", "publish", "DeletePublishApp"},
		[]string{"start/api/v1/projects/:project_id/publishes/:publish_id/audits", "POST", "start", "publish", "GetOpertaionLogByPagination"},
		[]string{"start/api/v1/projects/:project_id/publishes/:publish_id/stages/:stage_id/back-to", "GET", "start", "publish", "GetBackTo"},
		[]string{"start/api/v1/projects/:project_id/publishes/:publish_id/stages/:stage_id/back-to", "POST", "start", "publish", "TriggerBackTo"},
		[]string{"start/api/v1/projects/:project_id/publishes/:publish_id/stages/:stage_id/next-stage", "GET", "start", "publish", "GetNextStage"},
		[]string{"start/api/v1/projects/:project_id/publishes/:publish_id/stages/:stage_id/next-stage", "POST", "start", "publish", "TriggerNextStage"},
		[]string{"start/api/v1/pipelines/:project_id/publishes/:publish_id/stages/:stage_id/steps/:step_name", "GET", "start", "publish", "GetStepInfo"},
		[]string{"start/api/v1/pipelines/:project_id/publishes/:publish_id/stages/:stage_id/steps/:step_name", "POST", "start", "publish", "RunStep"},
		[]string{"start/api/v1/pipelines/:project_id/publishes/:publish_id/stages/:stage_id/steps/:step_name/callback", "POST", "start", "publish", "RunStepCallback"},

		// integrate
		[]string{"start/api/v1/integrate/compile_envs", "GET", "start", "system", "GetCompileEnvs"},
		[]string{"start/api/v1/integrate/clusters", "GET", "start", "system", "GetIntegrateClusters"},
		[]string{"start/api/v1/integrate/settings", "GET", "start", "system", "GetIntegrateSettings"},

		// task template
		[]string{"start/api/v1/pipelines/flow/components", "GET", "start", "system", "FlowComponentList"},
		[]string{"start/api/v1/pipelines/flow/steps", "POST", "start", "system", "FlowStepListByPagination"},
		[]string{"start/api/v1/pipelines/flow/steps/create", "POST", "start", "system", "FlowStepCreate"},
		[]string{"start/api/v1/pipelines/flow/steps/:step_id", "PUT", "start", "system", "FlowStepUpdate"},
		[]string{"start/api/v1/pipelines/flow/steps/:step_id", "DELETE", "start", "system", "FlowStepDelete"},
	},
}
