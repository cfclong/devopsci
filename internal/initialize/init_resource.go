package initialize

import (
	"fmt"

	"gitee.com/plutoccc/devops_app/internal/dao"
	"gitee.com/plutoccc/devops_app/internal/middleware/log"
	"gitee.com/plutoccc/devops_app/internal/models"
	"gitee.com/plutoccc/devops_app/utils/errors"
)

type ResourceReq struct {
	Resources []BatchResourceTypeSpec `json:"resources"`
}

type BatchResourceTypeSpec struct {
	ResourceType       []string   `json:"resource_type"`
	ResourceOperation  [][]string `json:"resource_operation"`
	ResourceConstraint [][]string `json:"resource_constraint"`
}

func ToBatchResourceTypeReq(specs []BatchResourceTypeSpec) models.BatchResourceTypeReq {
	var req models.BatchResourceTypeReq
	for _, spec := range specs {
		var resourceType models.ResourceTypeReq
		var resourceOperations []models.ResourceOperationReq
		var resourceConstraints []models.ResourceConstraintReq

		if len(spec.ResourceType) == 2 {
			resourceType = models.ResourceTypeReq{
				ResourceType: spec.ResourceType[0],
				Description:  spec.ResourceType[1],
			}
		}
		if len(spec.ResourceOperation) > 0 {
			for _, op := range spec.ResourceOperation {
				if len(op) == 2 {
					resourceOperations = append(resourceOperations, models.ResourceOperationReq{
						ResourceOperation: op[0],
						Description:       op[1],
					})
				}
			}
		}
		if len(spec.ResourceConstraint) > 0 {
			for _, con := range spec.ResourceConstraint {
				if len(con) == 2 {
					resourceConstraints = append(resourceConstraints, models.ResourceConstraintReq{
						ResourceConstraint: con[0],
						Description:        con[1],
					})
				}
			}
		}
		req.Resources = append(req.Resources, models.ResourceReq{
			ResourceType:        resourceType,
			ResourceOperations:  resourceOperations,
			ResourceConstraints: resourceConstraints,
		})
	}
	return req
}

// init resource
func initResource() error {
	if err := dao.BatchCreateResourceType(ToBatchResourceTypeReq(resourceReq.Resources)); err != nil {
		log.Log.Error("Init resource error: %s", err.Error())
		return err
	}
	return nil
}

type RouterReq struct {
	Routers [][]string `json:"routers"`
}

// initRouterItems ..
func initRouterItems() error {

	// TODO: fix
	// if err := dao.DeleteGatewayRouteByBackend("start"); err != nil {
	// 	log.Log.Error("Init gateway error: %s", err.Error())
	// 	return err
	// }
	for _, route := range gaetwayReq.Routers {
		if len(route) != 5 {
			err := fmt.Errorf("invalid router parameter: %v", route)
			return err
		}
		// TODO: Does validation already exist?
		if err := dao.CreateGatewayRoute(route[0], route[1], route[2], route[3], route[4]); err != nil {
			if !errors.OrmError1062(err) {
				log.Log.Error("Init gateway error: %s", err.Error())
				return err
			}
		}
	}
	return nil
}
