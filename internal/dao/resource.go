package dao

import (
	"fmt"

	"gitee.com/plutoccc/devops_app/internal/models"
)

// 资源类型
func ResourceTypeList() ([]*models.ResourceType, error) {
	resourceTypeList := []*models.ResourceType{}
	if _, err := GetOrmer().QueryTable("sys_resource_type").OrderBy("-create_at").
		All(&resourceTypeList); err != nil {
		return nil, err
	}
	return resourceTypeList, nil
}

func ResourceTypeListFilter(resourceTypes []string) ([]*models.ResourceType, error) {
	resourceTypeList := []*models.ResourceType{}
	querySeter := GetOrmer().QueryTable("sys_resource_type").OrderBy("-create_at")
	if len(resourceTypes) > 0 {
		querySeter = querySeter.Filter("resource_type__in", resourceTypes)
	}
	if _, err := querySeter.All(&resourceTypeList); err != nil {
		return nil, err
	}
	return resourceTypeList, nil
}

func GetResourceTypeDetail(rt string, op, con []string) (*models.ResourceType, error) {
	var resourceType models.ResourceType
	if err := GetOrmer().QueryTable("sys_resource_type").Filter("resource_type", rt).
		One(&resourceType); err != nil {
		return nil, err
	}

	// 获取资源操作列表
	resourceOperations := []*models.ResourceOperation{}
	querySeter := GetOrmer().QueryTable("sys_resource_operation").Filter("resource_type", rt).OrderBy("-create_at")
	if len(op) > 0 {
		querySeter = querySeter.Filter("resource_operation__in", op)
	}
	if _, err := querySeter.All(&resourceOperations); err != nil {
		return nil, err
	}
	// 获取资源约束列表
	resourceConstraints := []*models.ResourceConstraint{}
	querySeter = GetOrmer().QueryTable("sys_resource_constraint").Filter("resource_type", rt).OrderBy("-create_at")
	if len(con) > 0 {
		querySeter = querySeter.Filter("resource_constraint__in", con)
	}
	if _, err := querySeter.All(&resourceConstraints); err != nil {
		return nil, err
	}
	resourceType.ResourceOperations = resourceOperations
	resourceType.ResourceConstraints = resourceConstraints

	return &resourceType, nil
}

func GetResourceType(rt string) (*models.ResourceType, error) {
	var resourceType models.ResourceType
	if err := GetOrmer().QueryTable("sys_resource_type").Filter("resource_type", rt).
		One(&resourceType); err != nil {
		return nil, err
	}
	return &resourceType, nil
}

func BatchCreateResourceType(req models.BatchResourceTypeReq) error {
	for _, resource := range req.Resources {
		resourceType := resource.ResourceType.ResourceType
		sql := `insert ignore into sys_resource_type(resource_type,description) values(?,?)`
		if _, err := GetOrmer().Raw(sql, resourceType, resource.ResourceType.Description).Exec(); err != nil {
			return err
		}

		if len(resource.ResourceOperations) > 0 {
			values := ""
			for index, op := range resource.ResourceOperations {
				if index == 0 {
					values = fmt.Sprintf("('%v','%v','%v')", resourceType, op.ResourceOperation, op.Description)
				} else {
					values = values + "," + fmt.Sprintf("('%v','%v','%v')", resourceType, op.ResourceOperation, op.Description)
				}
			}
			sql = `insert ignore into sys_resource_operation(resource_type,resource_operation,description) values` + values
			if _, err := GetOrmer().Raw(sql).Exec(); err != nil {
				return err
			}
		}

		if len(resource.ResourceConstraints) > 0 {
			values := ""
			for index, con := range resource.ResourceConstraints {
				if index == 0 {
					values = fmt.Sprintf("('%v','%v','%v')", resourceType, con.ResourceConstraint, con.Description)
				} else {
					values = values + "," + fmt.Sprintf("('%v','%v','%v')", resourceType, con.ResourceConstraint, con.Description)
				}
			}
			sql = `insert ignore into sys_resource_constraint(resource_type,resource_constraint,description) values` + values
			if _, err := GetOrmer().Raw(sql).Exec(); err != nil {
				return err
			}
		}
	}
	return nil
}

func CreateResourceType(resourceType, description string) (*models.ResourceType, error) {
	sql := `insert into sys_resource_type(resource_type,description) values(?,?)`
	if _, err := GetOrmer().Raw(sql, resourceType, description).Exec(); err != nil {
		return nil, err
	}

	res, err := GetResourceTypeDetail(resourceType, []string{}, []string{})
	if err != nil {
		return nil, err
	}
	return res, err
}

func UpdateResourceType(resourceType, description string) error {
	sql := `update sys_resource_type set description=? where resource_type=? and description<>?`
	if _, err := GetOrmer().Raw(sql, description, resourceType, description).Exec(); err != nil {
		return err
	}
	return nil
}

func DeleteResourceType(resourceType string) error {
	// 删除资源操作
	sql := `delete from sys_resource_operation where resource_type=?`
	if _, err := GetOrmer().Raw(sql, resourceType).Exec(); err != nil {
		return err
	}
	// 删除资源约束
	sql = `delete from sys_resource_constraint where resource_type=?`
	if _, err := GetOrmer().Raw(sql, resourceType).Exec(); err != nil {
		return err
	}
	// 删除资源类型
	sql = `delete from sys_resource_type where resource_type=?`
	if _, err := GetOrmer().Raw(sql, resourceType).Exec(); err != nil {
		return err
	}
	return nil
}

func GetResourceOperationList() ([]*models.ResourceOperation, error) {
	var op []*models.ResourceOperation
	if _, err := GetOrmer().QueryTable("sys_resource_operation").
		All(&op); err != nil {
		return nil, err
	}
	return op, nil
}

func GetResourceOperation(resourceType, resourceOperation string) (*models.ResourceOperation, error) {
	var op models.ResourceOperation
	if err := GetOrmer().QueryTable("sys_resource_operation").
		Filter("resource_type", resourceType).
		Filter("resource_operation", resourceOperation).
		One(&op); err != nil {
		return nil, err
	}
	return &op, nil
}

func GetResourceOperationByResourceTypes(resourceTypes []string) ([]*models.ResourceOperation, error) {
	var resItems []*models.ResourceOperation
	if _, err := GetOrmer().QueryTable("sys_resource_operation").
		Filter("resource_type__in", resourceTypes).
		All(&resItems); err != nil {
		return nil, err
	}
	return resItems, nil
}

func GetResourceOperationByResourceOperations(resourceOperations []string) ([]*models.ResourceOperation, error) {
	var resItems []*models.ResourceOperation
	if _, err := GetOrmer().QueryTable("sys_resource_operation").
		Filter("resource_operation__in", resourceOperations).
		All(&resItems); err != nil {
		return nil, err
	}
	return resItems, nil
}

func GetResourceOperationByIDs(resourceItemIDs []int64) ([]*models.ResourceOperation, error) {
	var resItems []*models.ResourceOperation
	if _, err := GetOrmer().QueryTable("sys_resource_operation").
		Filter("id__in", resourceItemIDs).
		All(&resItems); err != nil {
		return nil, err
	}
	return resItems, nil
}

func AddResourceOperation(resourceType, resourceOperation, description string) error {
	sql := `insert into sys_resource_operation(resource_type,resource_operation,description) values(?,?,?)`
	if _, err := GetOrmer().Raw(sql, resourceType, resourceOperation, description).Exec(); err != nil {
		return err
	}
	return nil
}

func UpdateResourceOperation(resourceType, resourceOperation, description string) error {
	sql := `update sys_resource_operation set description=? where resource_type=? and resource_operation=? and description<>?`
	if _, err := GetOrmer().Raw(sql, description, resourceType, resourceOperation, description).Exec(); err != nil {
		return err
	}
	return nil
}

func DeleteResourceOperation(resourceType, resourceOperation string) error {
	sql := `delete from sys_resource_operation where resource_type=? and resource_operation=?`
	if _, err := GetOrmer().Raw(sql, resourceType, resourceOperation).Exec(); err != nil {
		return err
	}
	return nil
}

func AddResourceConstraint(resourceType, resourceConstraint, description string) error {
	sql := `insert into sys_resource_constraint(resource_type,resource_constraint,description) values(?,?,?)`
	if _, err := GetOrmer().Raw(sql, resourceType, resourceConstraint, description).Exec(); err != nil {
		return err
	}
	return nil
}

func UpdateResourceConstraint(resourceType, resourceConstraint, description string) error {
	sql := `update sys_resource_constraint set description=? where resource_type=? and resource_constraint=? and description<>?`
	if _, err := GetOrmer().Raw(sql, description, resourceType, resourceConstraint, description).Exec(); err != nil {
		return err
	}
	return nil
}

func DeleteResourceConstraint(resourceType, resourceConstraint string) error {
	sql := `delete from sys_resource_constraint where resource_type=? and resource_constraint=?`
	if _, err := GetOrmer().Raw(sql, resourceType, resourceConstraint).Exec(); err != nil {
		return err
	}
	return nil
}

func GetResourceConstraintList(resourceType string) ([]string, error) {
	constraintList := []string{}
	sql := `select resource_constraint from sys_resource_constraint where resource_type=?`
	if _, err := GetOrmer().Raw(sql, resourceType).QueryRows(&constraintList); err != nil {
		return nil, err
	}
	return constraintList, nil
}

// GetUserConstraintByKey ..
func GetUserConstraintByKey(user string, constraintKey []string) (map[string][]string, error) {
	constraints := []*models.GroupUserConstraint{}

	if len(constraintKey) > 0 {
		if _, err := GetOrmer().QueryTable("sys_group_user_constraint").
			Filter("constraint__in", constraintKey).Filter("user", user).
			All(&constraints); err != nil {
			return nil, err
		}
	} else {
		return nil, nil
	}
	res := map[string][]string{}
	for _, con := range constraints {
		res[con.Constraint] = append(res[con.Constraint], con.Value)
	}
	return res, nil
}
