package project

import (
	"fmt"

	"gitee.com/plutoccc/devops_app/internal/middleware/log"
	"gitee.com/plutoccc/devops_app/internal/models"
	"gitee.com/plutoccc/devops_app/utils/query"
)

// CreateProjectApp ...
func (pm *ProjectManager) CreateProjectApp(projectID int64, item *ProjectAppReq, creator string) error {
	log.Log.Debug("request params: %+v", item)

	projectAppModel := models.ProjectApp{
		Addons:    models.NewAddons(),
		Creator:   creator,
		ProjectID: projectID,
		ScmID:     item.SCMID,
	}

	_, err := pm.model.CreateProjectAppIfNotExist(&projectAppModel)
	if err != nil {
		log.Log.Error("create project app error: %s", err)
		return err
	}

	return nil
}

// GetProjectApps ..
func (pm *ProjectManager) GetProjectApps(projectID int64) ([]*ProjectAppRsp, error) {
	modelProjectApps, err := pm.model.GetProjectApps(projectID)
	if err != nil {
		return nil, err
	}
	return pm.formatProjectAppsResp(modelProjectApps)
}

// GetProjectAppsByPagination ..
func (pm *ProjectManager) GetProjectAppsByPagination(projectID int64, filter *models.ProejctAppFilterQuery) (*query.QueryResult, error) {
	apps, modelDatas, err := pm.model.GetProjectAppsList(projectID, filter)
	if err != nil {
		return nil, err
	}

	projectAppsRsp, err := pm.formatProjectAppsResp(modelDatas)
	if err != nil {
		return nil, err
	}
	apps.Item = projectAppsRsp

	return apps, nil
}

// DeleteProjectApp ...
func (pm *ProjectManager) DeleteProjectApp(projectAppID int64) error {
	log.Log.Debug("delete project app, projectAppID: %v", projectAppID)

	_, err := pm.model.GetProjectApp(projectAppID)
	if err != nil {
		log.Log.Error("when delete project app, get project app occur error: %s", err.Error())
		return fmt.Errorf("当前代码库可能已经删除，请你刷新页面后重试")
	}

	// TODO: add publish order verify
	err = pm.model.DeleteProjectApp(projectAppID)
	if err != nil {
		return err
	}
	// TODO: delete app service constraint
	return nil
}

// UpdateProjectApp ..
func (pm *ProjectManager) UpdateProjectApp(projectID, projectAppID int64, req *ProjectAppUpdateReq) error {
	_, err := pm.model.GetProjectAppByScmID(projectID, req.ScmID)
	if err == nil {
		return fmt.Errorf("already exist scmid: %v register", req.ScmID)
	}
	projectApp, err := pm.model.GetProjectApp(projectAppID)
	if err != nil {
		return err
	}
	projectApp.ScmID = req.ScmID
	return pm.model.UpdateProjectApp(projectApp)
}
