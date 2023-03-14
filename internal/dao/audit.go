package dao

import (
	"gitee.com/plutoccc/devops_app/internal/models"
)

func AuditInsert(audit *models.Audit) error {
	_, err := GetOrmer().Insert(audit)
	return err
}

func AuditList() ([]*models.Audit, error) {
	auditList := []*models.Audit{}
	if _, err := GetOrmer().QueryTable("sys_audit").OrderBy("-create_at").All(&auditList); err != nil {
		return nil, err
	}
	return auditList, nil
}
