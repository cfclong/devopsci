package publish

import (
	"gitee.com/plutoccc/devops_app/internal/middleware/log"
	"gitee.com/plutoccc/devops_app/internal/models"
	"gitee.com/plutoccc/devops_app/utils"

	"github.com/astaxie/beego/logs"
)

// GetCanAddedApps ...
func (pm *PublishManager) GetCanAddedApps(publishID int64) ([]*models.ProjectApp, error) {
	publish, err := pm.model.GetPublishByID(publishID)
	if err != nil {
		log.Log.Error("when GetCanAddedApps, get publish by id: %v error: %s", publishID, err.Error())
		return nil, err
	}
	publishApps, err := pm.model.GetPublishAppsByID(publishID)
	if err != nil {
		log.Log.Error("when GetCanAddedApps, get publish apps by publishId: %v error: %s", publishID, err.Error())
		return nil, err
	}
	appIDByAdded := []int64{}
	for _, pa := range publishApps {
		appIDByAdded = append(appIDByAdded, pa.ProjectAppID)
	}

	projectApps, err := pm.projectModel.GetProjectApps(publish.ProjectID)
	if err != nil {
		log.Log.Error("when GetCanAddedApps, get project Apps by project id: %v, occur errro: %s", publish.ProjectID, err.Error())
		return nil, err
	}
	rsp := []*models.ProjectApp{}

	for _, app := range projectApps {
		if !utils.IntContains(appIDByAdded, app.ID) {
			branchList := []string{}
			branches, err := pm.gitAppModel.GetAppBranches(app.ScmID)
			if err == nil {
				for _, item := range branches {
					branchList = append(branchList, item.BranchName)
				}
			} else {
				logs.Warn("when get can added apps, get app branches by id: %v, error: %v", app.ID, err.Error())
			}
			if len(branchList) == 0 {
				branchList = []string{"master"}
			}
			// TODO: branch history list get need commbine
			app.BranchHistoryList = branchList
			rsp = append(rsp, app)
		}
	}
	return rsp, nil
}

// AddPublishApps ..
func (pm *PublishManager) AddPublishApps(publishID int64, req *PublishAddApps) error {
	apps := req.Apps
	return pm.createPublishApps(apps, publishID)
}

// DeletePublishApp ..
func (pm *PublishManager) DeletePublishApp(publishAppID int64) error {
	app, err := pm.model.GetPublishApp(publishAppID)
	if err != nil {
		log.Log.Error("when DeletePublishApp, get publish app occur error: %s", err.Error())
		return err
	}
	app.MarkDeleted()
	return pm.model.UpdatePublishApp(app)
}
