package settings

import (
	"errors"

	"gitee.com/plutoccc/devops_app/internal/middleware/log"
	"gitee.com/plutoccc/devops_app/internal/models"
	"gitee.com/plutoccc/devops_app/utils/query"
)

// CompileEnvReq ..
type CompileEnvReq struct {
	Name        string `json:"name,omitempty"`
	Image       string `json:"image,omitempty"`
	Command     string `json:"command,omitempty"`
	Args        string `json:"args,omitempty"`
	Description string `json:"description,omitempty"`
}

// GetCompileEnvs ..
func (pm *SettingManager) GetCompileEnvs(integrateType string) ([]*models.CompileEnv, error) {
	items, err := pm.model.GetCompileEnvs(integrateType)
	if err != nil {
		log.Log.Error("get interate settings error: %s", err.Error())
		return nil, err
	}

	return items, err
}

// GetCompileEnvByID ..
func (pm *SettingManager) GetCompileEnvByID(id int64) (*models.CompileEnv, error) {
	compileEnv, err := pm.model.GetCompileEnvByID(id)
	if err != nil {
		log.Log.Error("when GetCompileEnvBy id: %v, occur error: %s", id, err.Error())
		return nil, err
	}
	return compileEnv, err
}

// GetCompileEnvByID ..
func (pm *SettingManager) GetCompileEnvByName(name string) (*models.CompileEnv, error) {
	compileEnv, err := pm.model.GetCompileEnvByName(name)
	if err != nil {
		log.Log.Error("when GetCompileEnvBy name: %v, occur error: %s", name, err.Error())
		return nil, err
	}
	return compileEnv, err
}

// GetCompileEnvsByPagination ..
func (pm *SettingManager) GetCompileEnvsByPagination(filter *query.FilterQuery) (*query.QueryResult, error) {
	queryResult, settingsList, err := pm.model.GetCompileEnvsByPagination(filter)
	if err != nil {
		return nil, err
	}
	queryResult.Item = settingsList
	return queryResult, err
}

// resetEnv clear env config
func resetEnv(env *string) {
	*env = ""
}

func compileEnvNameUnique(pm *SettingManager, name string, stepId int64) error {
	if len(name) == 0 {
		return errors.New("param `Name` is not allowed empty")
	}

	exists, _ := pm.model.GetCompileEnvByName(name)

	if exists != nil && (stepId == 0 || exists.ID != stepId) {
		return errors.New("环境名称 `" + name + "` 已经存在")
	}

	return nil
}

// UpdateCompileEnv ..
func (pm *SettingManager) UpdateCompileEnv(request *CompileEnvReq, stepID int64) error {
	compileEnv, err := pm.model.GetCompileEnvByID(stepID)
	if err != nil {
		return err
	}

	if err := compileEnvNameUnique(pm, request.Name, stepID); err != nil {
		return err
	}

	if request.Name != "" {
		compileEnv.Name = request.Name
	}

	if request.Args != "" {
		compileEnv.Args = request.Args
	} else {
		resetEnv(&compileEnv.Args)
	}

	if request.Command != "" {
		compileEnv.Command = request.Command
	} else {
		resetEnv(&compileEnv.Command)
	}

	if request.Description != "" {
		compileEnv.Description = request.Description
	} else {
		resetEnv(&compileEnv.Description)
	}

	if request.Image != "" {
		compileEnv.Image = request.Image
	}

	return pm.model.UpdateCompileEnv(compileEnv)
}

// CreateCompileEnv ..
func (pm *SettingManager) CreateCompileEnv(request *CompileEnvReq, creator string) error {

	if err := compileEnvNameUnique(pm, request.Name, 0); err != nil {
		return err
	}

	// TODO: verify req struct is valid
	newCompileEnv := &models.CompileEnv{
		Name:        request.Name,
		Description: request.Description,
		Creator:     creator,
		Image:       request.Image,
		Command:     request.Command,
		Args:        request.Args,
	}

	return pm.model.CreateCompileEnv(newCompileEnv)
}

// DeleteCompileEnv ..
func (pm *SettingManager) DeleteCompileEnv(stageID int64) error {
	// TODO: add compile env delete verify
	return pm.model.DeleteCompileEnv(stageID)
}
