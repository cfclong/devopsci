package kuberes

import (
	"fmt"

	"gitee.com/plutoccc/devops_app/internal/core/settings"
	"gitee.com/plutoccc/devops_app/internal/dao"
	"gitee.com/plutoccc/devops_app/internal/middleware/log"
	"gitee.com/plutoccc/devops_app/internal/models"

	"gitee.com/plutoccc/devops_app/utils/query"

	"github.com/astaxie/beego/orm"
)

type ExtensionParam struct {
	Force   bool //when user deploy its app and the app is existed in other namespace, the old app will be deleted
	Patcher PatcherFunction
}

type DeployWorker struct {
	Name      string
	arHandle  *AppRes
	kubeRes   *KubeAppRes
	extension *ExtensionParam
	template  AppTemplate
}

func NewDeployWorker(name, namespace, kind string, ar *AppRes, eparam *ExtensionParam, tpl AppTemplate) *DeployWorker {
	return &DeployWorker{
		Name:      name,
		arHandle:  ar,
		kubeRes:   NewKubeAppRes(ar.Client, ar.Cluster, namespace, kind),
		extension: eparam,
		template:  tpl,
	}
}

func (wk *DeployWorker) Start(templateName string, param AppParam) error {
	log.Log.Info("deploying application: ", wk.Name)
	app, err := wk.arHandle.Appmodel.GetAppByName(wk.arHandle.Cluster, wk.kubeRes.Namespace, param.Name)
	if err == nil {
		return wk.updateAppRes(*app)
	}
	if err != orm.ErrNoRows {
		return err
	}
	return wk.createAppRes(templateName, param)

}

func (wk *DeployWorker) updateAppRes(app models.CaasApplication) error {
	//delete possible resource
	log.Log.Info("delete possible deploy and pod resource: ", wk.arHandle.Cluster, wk.kubeRes.Namespace, app.Name, app.Kind)
	wk.deleteApplication(app.Name)
	_, err := wk.arHandle.ReconfigureApp(app, wk.template)
	if err != nil {
		return err
	}
	return nil
}

func (wk *DeployWorker) createAppRes(templateName string, param AppParam) error {
	// create new app resource
	app, err := wk.createKubeAppRes(templateName, param)
	if err != nil {
		return err
	}
	err = wk.arHandle.Appmodel.CreateApp(*app)
	if err != nil {
		wk.kubeRes.DeleteAppResource(wk.template)
		wk.arHandle.Appmodel.DeleteApp(*app)
		return err
	}
	if wk.extension != nil {
		if wk.extension.Patcher != nil {
			wk.extension.Patcher(*app)
		}
	}
	return nil
}

func (wk *DeployWorker) createKubeAppRes(templateName string, param AppParam) (*models.CaasApplication, error) {
	app, err := wk.template.GenerateAppObject(wk.arHandle.Cluster, wk.kubeRes.Namespace, templateName, wk.arHandle.ProjectID)
	if err != nil {
		return nil, err
	}
	//delete possible resource
	log.Log.Info("delete possible deploy and pod resource: ", wk.arHandle.Cluster, wk.kubeRes.Namespace, param.Name, app.Kind)
	wk.deleteApplication(param.Name)
	log.Log.Info("create resource: ", wk.arHandle.Cluster, wk.kubeRes.Namespace, param.Name, app.Kind)
	if err := wk.kubeRes.CreateAppResource(wk.template); err != nil {
		return nil, err
	}
	return app, nil
}

func (wk *DeployWorker) deleteApplication(appname string) {
	filter := query.NewFilterQuery(false)
	filter.FilterKey = "name"
	filter.FilterVal = appname
	res, err := wk.arHandle.Appmodel.GetAppList(filter, wk.arHandle.ProjectID, wk.kubeRes.cluster, wk.kubeRes.Namespace)
	if err != nil {
		log.Log.Error("deleteApplication error: ", err.Error())
		return
	}
	applist := res.Item.([]models.CaasApplication)
	ar := *wk.arHandle
	for _, app := range applist {
		exist, err := wk.kubeRes.CheckAppIsExisted(app.Name)
		if err == nil && exist && wk.arHandle.Cluster != app.Cluster {
			ar.Cluster = app.Cluster
			err = (&ar).DeleteApp(app.Namespace, app.Name)
			if err != nil {
				log.Log.Info(fmt.Sprintf("delete unsuitable application(%s/%s) failed: %v!", app.Cluster, app.Name, err))
			} else {
				log.Log.Info(fmt.Sprintf("delete unsuitable application(%s/%s) successfully!", app.Cluster, app.Name))
			}
		}
	}
}

func getDefaultPullSecretAndRegistryAddr(envID int64) (string, string, error) {
	projectEnv, err := dao.NewProjectModel().GetProjectEnvByID(envID)
	if err != nil {
		log.Log.Error("when create registry secret get project env by id: %v, error: %s", envID, err.Error())
		return "", "", err
	}
	integrateSettingRegistry, err := dao.NewSysSettingModel().GetIntegrateSettingByID(projectEnv.Registry)
	if err != nil {
		log.Log.Error("when create registry secret get integrate setting by id: %v, error: %s", projectEnv.Registry, err.Error())
		return "", "", err
	}
	config := settings.Config{}
	configJSON, err := config.Struct(integrateSettingRegistry.DecryptConfig(), integrateSettingRegistry.Type)
	if err != nil {
		log.Log.Error("when parse registry config error: %s", err.Error())
		return "", "", err
	}
	var url string
	if registryConf, ok := configJSON.(*settings.RegistryConfig); ok {
		url = registryConf.URL
	}

	return fmt.Sprintf("registry-%x", integrateSettingRegistry.ID), url, nil
}
