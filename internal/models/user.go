package models

import (
	"time"

	"gitee.com/plutoccc/devops_app/utils/validate"
)

// LoginType defined
const (
	LocalAuth = iota + 1
	LDAPAuth
)

type User struct {
	Addons
	User  string `orm:"column(user);unique" json:"user"`
	Name  string `orm:"column(name)" json:"name"`
	Email string `orm:"column(email)" json:"email"`
	Token string `orm:"column(token);unique;" json:"token"`

	LoginType int    `orm:"column(login_type);" json:"login_type"`
	Password  string `json:"-" gorm:"type:varchar(128);comment:密码"`

	LastLoginTime time.Time        `orm:"column(last_login_time);null;type(datetime);" json:"lastLoginTime"`
	Admin         int              `orm:"-" json:"admin"`
	GroupAdmin    int              `orm:"-" json:"groupAdmin"`
	UserGroups    []*UserGroupRole `orm:"-" json:"roles"`
	Groups        []*Group         `orm:"-" json:"groups"`
}

type UserGroupRole struct {
	Group       string `json:"group"`
	GroupAdmin  bool   `json:"group_admin"`
	Role        string `json:"role"`
	Description string `json:"description"`
}

func (t *User) TableName() string {
	return "sys_user"
}

// UserReq ..
type UserReq struct {
	User     string `json:"user"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (v *UserReq) Verify() error {
	v.User = validate.FormatString(v.User)
	v.Name = validate.FormatString(v.Name)
	v.Email = validate.FormatString(v.Email)

	if err := validate.ValidateName(v.User); err != nil {
		return err
	}
	if err := validate.ValidateDescription(v.Name); err != nil {
		return err
	}

	if len(v.Email) > 0 {
		if err := validate.ValidateEmail(v.Email); err != nil {
			return err
		}
	}
	return nil
}
