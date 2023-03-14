package api

import (
	"gitee.com/plutoccc/devops_app/internal/dao"
	"gitee.com/plutoccc/devops_app/internal/middleware/log"
	"gitee.com/plutoccc/devops_app/internal/models"
)

type RoleController struct {
	BaseController
}

// RoleList ..
func (r *RoleController) RoleList() {
	rsp, err := dao.GroupRoleList("system")
	if err != nil {
		r.HandleInternalServerError(err.Error())
		log.Log.Error("Get role list error: %s", err.Error())
		return
	}
	r.Data["json"] = NewResult(true, rsp, "")
	r.ServeJSON()
}

func (r *RoleController) GetRole() {
	roleName := r.GetStringFromPath(":role")

	rsp, err := dao.GetGroupRoleByName("system", roleName)
	if err != nil {
		r.HandleInternalServerError(err.Error())
		log.Log.Error("Get role error: %s", err.Error())
		return
	}
	r.Data["json"] = NewResult(true, rsp, "")
	r.ServeJSON()
}

func (r *RoleController) CreateRole() {
	var req models.GroupRoleReq
	r.DecodeJSONReq(&req)
	// group use system
	req.Group = "system"

	if err := req.Verify(); err != nil {
		r.HandleBadRequest(err.Error())
		log.Log.Error("Create role error: %s", err.Error())
		return
	}

	rsp, err := dao.CreateGroupRole(&req)
	if err != nil {
		r.HandleInternalServerError(err.Error())
		log.Log.Error("Create role error: %s", err.Error())
		return
	}

	r.Data["json"] = NewResult(true, rsp, "")
	r.ServeJSON()
}

func (r *RoleController) UpdateRole() {
	roleName := r.GetStringFromPath(":role")
	var req models.GroupRoleReq
	r.DecodeJSONReq(&req)
	req.Group = "system"
	req.Role = roleName

	if err := req.Verify(); err != nil {
		r.HandleBadRequest(err.Error())
		log.Log.Error("Update role error: %s", err.Error())
		return
	}

	if err := dao.UpdateGroupRole(&req); err != nil {
		r.HandleInternalServerError(err.Error())
		log.Log.Error("Update role error: %s", err.Error())
		return
	}

	r.Data["json"] = NewResult(true, nil, "")
	r.ServeJSON()
}

func (r *RoleController) DeleteRole() {
	roleName := r.GetStringFromPath(":role")
	if err := dao.DeleteGroupRole("system", roleName); err != nil {
		r.HandleInternalServerError(err.Error())
		log.Log.Error("Delete role error: %s", err.Error())
		return
	}

	r.Data["json"] = NewResult(true, nil, "")
	r.ServeJSON()
}

func (r *RoleController) RoleBundlingList() {
	groupName := r.GetStringFromPath(":group")
	roleName := r.GetStringFromPath(":role")
	rsp, err := dao.GroupRoleBundlingList(groupName, roleName)
	if err != nil {
		r.HandleInternalServerError(err.Error())
		log.Log.Error("Get role list error: %s", err.Error())
		return
	}
	r.Data["json"] = NewResult(true, rsp, "")
	r.ServeJSON()
}

func (r *RoleController) RoleBundling() {
	groupName := r.GetStringFromPath(":group")
	roleName := r.GetStringFromPath(":role")
	var req models.GroupRoleBundlingReq
	r.DecodeJSONReq(&req)
	req.Group = groupName
	req.Role = roleName

	if err := dao.GroupRoleBundling(&req); err != nil {
		r.HandleInternalServerError(err.Error())
		log.Log.Error("role bundling error: %s", err.Error())
		return
	}

	r.Data["json"] = NewResult(true, nil, "")
	r.ServeJSON()
}

func (r *RoleController) RoleUnbundling() {
	groupName := r.GetStringFromPath(":group")
	roleName := r.GetStringFromPath(":role")
	var req models.GroupRoleBundlingReq
	r.DecodeJSONReq(&req)
	req.Group = groupName
	req.Role = roleName

	if err := dao.GroupRoleUnbundling(&req); err != nil {
		r.HandleInternalServerError(err.Error())
		log.Log.Error("role unbundling error: %s", err.Error())
		return
	}

	r.Data["json"] = NewResult(true, nil, "")
	r.ServeJSON()
}

func (r *RoleController) RoleOperationList() {
	roleName := r.GetStringFromPath(":role")
	rolesOperations, err := dao.GetRoleOperationsByRoleName(roleName)
	if err != nil {
		r.HandleInternalServerError(err.Error())
		log.Log.Error("get role operations by role name error: %s", err.Error())
		return
	}
	resIDs := []int64{}
	for _, item := range rolesOperations {
		resIDs = append(resIDs, item.OperationID)
	}
	rsp, err := dao.GetResourceOperationByIDs(resIDs)
	if err != nil {
		r.HandleInternalServerError(err.Error())
		log.Log.Error("get role operations by ids error: %s", err.Error())
		return
	}

	r.Data["json"] = NewResult(true, rsp, "")
	r.ServeJSON()
}

func (r *RoleController) AddRoleOperation() {
	roleName := r.GetStringFromPath(":role")
	var req models.GroupRoleOperationReq
	r.DecodeJSONReq(&req)
	req.Role = roleName
	req.Group = "system"

	if err := dao.AddRoleOperation(&req); err != nil {
		r.HandleInternalServerError(err.Error())
		log.Log.Error("Add role operation error: %s", err.Error())
		return
	}

	r.Data["json"] = NewResult(true, nil, "")
	r.ServeJSON()
}

func (r *RoleController) RemoveRoleOperation() {
	roleName := r.GetStringFromPath(":role")
	operationID, _ := r.GetInt64FromPath(":operationID")
	req := models.GroupRoleOperationReq{}
	req.Role = roleName
	req.Operations = []int64{operationID}

	if err := dao.DeleteGroupRolePolicy(&req); err != nil {
		r.HandleInternalServerError(err.Error())
		log.Log.Error("Remove role operation error: %s", err.Error())
		return
	}

	r.Data["json"] = NewResult(true, nil, "")
	r.ServeJSON()
}
