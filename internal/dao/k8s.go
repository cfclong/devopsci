package dao

import (
	"gitee.com/plutoccc/devops_app/internal/models"

	"github.com/astaxie/beego/orm"
)

// K8sClusterModel ...
type K8sClusterModel struct {
	ormer                orm.Ormer
	ApplicationTableName string
	NamespaceTableName   string
}

// NewK8sClusterModel ...
func NewK8sClusterModel() (model *K8sClusterModel) {
	return &K8sClusterModel{
		ormer:                GetOrmer(),
		ApplicationTableName: (&models.CaasApplication{}).TableName(),
	}
}

// GetApplicationsByProjectID ...
func (model *K8sClusterModel) GetApplicationsByProjectID(projectID int64) ([]*models.CaasApplication, error) {
	apps := []*models.CaasApplication{}
	qs := model.ormer.QueryTable(model.ApplicationTableName).
		Filter("project_id", projectID).
		Filter("deleted", false)
	_, err := qs.All(&apps)
	return apps, err
}

// GetApplication ...
func (model *K8sClusterModel) GetApplication(cluster, department, svcName string) ([]*models.CaasApplication, error) {
	apps := []*models.CaasApplication{}
	qs := model.ormer.QueryTable(model.ApplicationTableName).
		Filter("cluster", cluster).
		Filter("namespace", department).
		Filter("name", svcName).
		Filter("deleted", false)
	_, err := qs.All(&apps)
	if err != nil {
		return nil, err
	}
	return apps, err
}
