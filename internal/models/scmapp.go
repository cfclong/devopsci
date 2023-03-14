package models

// ScmApp ...
type ScmApp struct {
	Addons
	Creator           string   `orm:"column(creator);size(64);null" json:"creator"`
	Name              string   `orm:"column(name);size(64)" json:"name"`
	FullName          string   `orm:"column(full_name);size(64)" json:"full_name"`
	Language          string   `orm:"column(language);size(64)" json:"language"`
	BranchName        string   `orm:"column(branch_name);size(64)" json:"branch_name"`
	Path              string   `orm:"column(path);size(255)" json:"path"`
	RepoID            int64    `orm:"column(repo_id)" json:"repo_id"`
	CompileEnvID      int64    `orm:"column(compile_env_id);size(64)" json:"compile_env_id"`
	BuildPath         string   `orm:"column(build_path);size(64)" json:"build_path"`
	Dockerfile        string   `orm:"column(dockerfile);size(256)" json:"dockerfile"`
	BranchHistoryList []string `orm:"-" json:"branch_history_list"`
}

// TableName ..
func (t *ScmApp) TableName() string {
	return "pub_scm_app"
}

// AppBranch ...
type AppBranch struct {
	Addons
	AppID      int64  `orm:"column(app_id);" json:"app_id"`
	BranchName string `orm:"column(branch_name);size(64)" json:"branch_name"`
	Path       string `orm:"column(path);size(256)" json:"path"`
}

// TableName ...
func (t *AppBranch) TableName() string {
	return "pub_app_branch"
}
