package dao

import (
	"gitee.com/plutoccc/devops_app/internal/models"

	"github.com/astaxie/beego/orm"
)

type TemplateModel struct {
	tOrmer    orm.Ormer
	TableName string
}

func NewTemplateModel() *TemplateModel {
	return &TemplateModel{
		tOrmer:    GetOrmer(),
		TableName: (&models.CaasTemplate{}).TableName(),
	}
}

func (tm *TemplateModel) CreateTemplate(template models.CaasTemplate) (*models.CaasTemplate, error) {
	template.Addons = models.NewAddons()
	_, err := tm.tOrmer.Insert(&template)
	if err != nil {
		return nil, err
	}

	return tm.GetTemplate(template.Namespace, template.Name)
}

func (tm *TemplateModel) UpdateTemplate(template models.CaasTemplate) error {
	template.Addons = template.Addons.UpdateAddons()
	_, err := tm.tOrmer.Update(&template)

	return err
}

func (tm *TemplateModel) DeleteTemplate(namespace, name string) error {
	sql := "update " + tm.TableName + " set deleted=1, delete_at=now() where namespace=? and name=? and deleted=0"
	_, err := tm.tOrmer.Raw(sql, namespace, name).Exec()

	return err
}

func (tm *TemplateModel) GetTemplate(namespace, name string) (*models.CaasTemplate, error) {
	var template models.CaasTemplate

	if err := tm.tOrmer.QueryTable(tm.TableName).
		Filter("namespace", namespace).
		Filter("name", name).
		Filter("deleted", 0).One(&template); err != nil {
		return nil, err
	}

	return &template, nil
}

func (tm *TemplateModel) GetTemplateByID(templateId int64) (*models.CaasTemplate, error) {
	var template models.CaasTemplate

	if err := tm.tOrmer.QueryTable(tm.TableName).
		Filter("id", templateId).
		Filter("deleted", 0).One(&template); err != nil {
		return nil, err
	}

	return &template, nil
}
