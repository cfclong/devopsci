package dao

import (
	"gitee.com/plutoccc/devops_app/internal/models"

	"github.com/astaxie/beego/orm"
)

// UserRolesModel ...
type UserRolesModel struct {
	ormer              orm.Ormer
	GroupRoleTableName string
	UserName           string
}

// NewUserRolesModel ...
func NewUserRolesModel() (model *UserRolesModel) {
	return &UserRolesModel{
		ormer:              GetOrmer(),
		GroupRoleTableName: (&models.GroupRole{}).TableName(),
		UserName:           (&models.User{}).TableName(),
	}
}

// GetRoleByID ...
func (model *UserRolesModel) GetRoleByID(roleID int64) (*models.GroupRole, error) {
	role := models.GroupRole{}
	qs := model.ormer.QueryTable(model.GroupRoleTableName).
		Filter("id", roleID)
	if err := qs.One(&role); err != nil {
		return nil, err
	}
	return &role, nil
}

// GetRoleByName ...
func (model *UserRolesModel) GetRoleByName(group, name string) (*models.GroupRole, error) {
	role := models.GroupRole{}
	qs := model.ormer.QueryTable(model.GroupRoleTableName).
		Filter("deleted", false).
		Filter("group", group).
		Filter("role", name)

	if err := qs.One(&role); err != nil {
		return nil, err
	}
	return &role, nil
}

// GetUserByName ...
func (model *UserRolesModel) GetUserByName(name string) (*models.User, error) {
	role := models.User{}
	qs := model.ormer.QueryTable(model.UserName).
		Filter("user", name)
	if err := qs.One(&role); err != nil {
		return nil, err
	}
	return &role, nil
}

// GetUserByToken ...
func (model *UserRolesModel) GetUserByToken(token string) (*models.User, error) {
	role := models.User{}
	qs := model.ormer.QueryTable(model.UserName).
		Filter("token", token)
	if err := qs.One(&role); err != nil {
		return nil, err
	}
	return &role, nil
}
