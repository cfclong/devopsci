package models

import (
	"gitee.com/plutoccc/devops_app/utils/validate"
)

type Group struct {
	Addons
	Group       string              `orm:"column(group);unique;" json:"group"`
	Level       string              `orm:"column(level);null;" json:"level"`
	ParentId    int64               `orm:"column(parent_id);null;" json:"parentId"`
	Description string              `orm:"column(description);unique;" json:"description"`
	Users       []*GroupRoleUserRsp `orm:"-" json:"users"`
	Roles       []*GroupRole        `orm:"-" json:"roles"`
}

func (t *Group) TableName() string {
	return "sys_group"
}

type GroupReq struct {
	Group       string `json:"group"`
	Description string `json:"description"`
}

func (v *GroupReq) Verify() error {
	v.Group = validate.FormatString(v.Group)
	v.Description = validate.FormatString(v.Description)
	if err := validate.IsReservedBuName(v.Group); err != nil {
		return err
	}
	if err := validate.ValidateName(v.Group); err != nil {
		return err
	}
	if err := validate.ValidateDescription(v.Description); err != nil {
		return err
	}
	return nil
}

type GroupUserRel struct {
	Addons
	GroupId int64 `orm:"column(group_id);index;null" json:"groupId"`
	UserId  int64 `orm:"column(user_id);index;null" json:"userId"`
}

func (t *GroupUserRel) TableName() string {
	return "sys_group_user_rel"
}

func (t *GroupUserRel) TableUnique() [][]string {
	return [][]string{
		{"GroupId", "UserId"},
	}
}

type GroupUserConstraint struct {
	Addons
	Group      string `orm:"column(group);size(128)" json:"group"`
	User       string `orm:"column(user);size(128)" json:"user"`
	Constraint string `orm:"column(constraint);size(64)" json:"constraint"`
	Value      string `orm:"column(value)" json:"value"`
}

func (t *GroupUserConstraint) TableName() string {
	return "sys_group_user_constraint"
}

func (t *GroupUserConstraint) TableIndex() [][]string {
	return [][]string{
		{"Group", "User"},
		{"Group", "User", "Constraint"},
	}
}

// TableUnique ..
func (t *GroupUserConstraint) TableUnique() [][]string {
	return [][]string{
		{"Group", "User", "Constraint", "Value"},
	}
}

// Verify ..
func (t *GroupUserConstraint) Verify() error {
	return nil
}
