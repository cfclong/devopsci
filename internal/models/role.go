package models

import (
	"gitee.com/plutoccc/devops_app/utils/validate"
)

// GroupRole ..
type GroupRole struct {
	Addons
	Group       string                   `orm:"column(group)" json:"group"`
	Role        string                   `orm:"column(role)" json:"role"`
	Description string                   `orm:"column(description)" json:"description"`
	Users       []*GroupRoleBundlingUser `orm:"-" json:"users"`
}

// TableName ..
func (t *GroupRole) TableName() string {
	return "sys_group_role"
}

// TableIndex ..
func (t *GroupRole) TableIndex() [][]string {
	return [][]string{
		{"Group"},
		{"Group", "Role"},
	}
}

// TableUnique ..
func (t *GroupRole) TableUnique() [][]string {
	return [][]string{
		{"Group", "Role"},
	}
}

// GroupRoleReq ..
type GroupRoleReq struct {
	Group       string  `json:"group"`
	Role        string  `json:"role"`
	Description string  `json:"description"`
	Operations  []int64 `json:"operations"`
}

// Verify ..
func (v *GroupRoleReq) Verify() error {
	v.Role = validate.FormatString(v.Role)
	v.Description = validate.FormatString(v.Description)
	if err := validate.ValidateName(v.Role); err != nil {
		return err
	}
	if err := validate.ValidateDescription(v.Description); err != nil {
		return err
	}
	return nil
}

type RoleRsp struct {
	Addons
	Role        string `json:"role"`
	Description string `json:"description"`
}

type GroupRoleOperation struct {
	Addons
	Group       string `orm:"column(group)" json:"group"`
	Role        string `orm:"column(role)" json:"role"`
	PolicyName  string `orm:"column(policy_name)" json:"policy_name"`
	OperationID int64  `orm:"column(operation_id)" json:"operation_id"`
}

// TableName ..
func (t *GroupRoleOperation) TableName() string {
	return "sys_group_role_operation"
}

// TableIndex ..
func (t *GroupRoleOperation) TableIndex() [][]string {
	return [][]string{
		{"Group", "Role"},
	}
}

// TableUnique ..
func (t *GroupRoleOperation) TableUnique() [][]string {
	return [][]string{
		{"Group", "Role", "OperationID"},
	}
}

// GroupRoleOperationReq ..
type GroupRoleOperationReq struct {
	Group      string  `json:"group"`
	Role       string  `json:"role"`
	Operations []int64 `json:"operations"`
}

type GroupRoleBundlingUser struct {
	User string `json:"user"`
	Name string `json:"name"`
	Addons
}

type GroupRoleUser struct {
	Addons
	Group string `orm:"column(group);index" json:"group"`
	User  string `orm:"column(user)" json:"user"`
	Role  string `orm:"column(role)" json:"role"`
}

// TableName ..
func (t *GroupRoleUser) TableName() string {
	return "sys_group_role_user"
}

// TableIndex ..
func (t *GroupRoleUser) TableIndex() [][]string {
	return [][]string{
		{"Group", "User"},
		{"Group", "Role"},
	}
}

// TableUnique ..
func (t *GroupRoleUser) TableUnique() [][]string {
	return [][]string{
		{"Group", "User", "Role"},
	}
}

type GroupRoleUserReq struct {
	Group string   `json:"group"`
	Users []string `json:"users"`
	Roles []string `json:"roles"`
}

type GroupRoleBundlingReq struct {
	Group string   `json:"group"`
	Role  string   `json:"role"`
	Users []string `json:"users"`
}

type GroupRoleUserRsp struct {
	Addons
	User  string     `json:"user"`
	Name  string     `json:"name"`
	Email string     `json:"email"`
	Phone string     `json:"phone"`
	Roles []*RoleRsp `json:"roles"`
}
