package dao

import (
	"fmt"
	"strings"

	"gitee.com/plutoccc/devops_app/internal/models"
	"github.com/isbrick/tools"
)

func CreateGatewayRoute(router, method, backend, resourceType, resourceOperation string) error {
	sql := `insert ignore into sys_resource_router(router,method,backend,resource_type,resource_operation) values(?,?,?,?,?)`
	if _, err := GetOrmer().Raw(sql, router, method, backend, resourceType, resourceOperation).Exec(); err != nil {
		return err
	}
	return nil
}

// GetResourceRouterItems ..
func GetResourceRouterItems(resourceType string, resourceOperations []string) ([]*models.GatewayRouter, error) {
	routerItems := []*models.GatewayRouter{}
	query := GetOrmer().QueryTable("sys_resource_router")
	if len(resourceOperations) > 0 {
		if tools.IsSliceContainsStr(resourceOperations, "*") {
			query = query.Filter("resource_type", resourceType)
		} else {
			query = query.Filter("resource_operation__in", resourceOperations)
		}
	}
	if _, err := query.All(&routerItems); err != nil {
		return nil, err
	}
	return routerItems, nil
}

func generateCasbinRules(resourceRouter []*models.GatewayRouter, roleName string) [][]string {
	res := make([][]string, 0, len(resourceRouter))
	for _, item := range resourceRouter {

		if !strings.HasPrefix(item.Router, "/") {
			item.Router = fmt.Sprintf("/%s", item.Router)
		}
		// casbinPolicy := []string{"p", roleName, item.Router, item.Method}
		casbinPolicy := []string{roleName, item.Router, item.Method}
		res = append(res, casbinPolicy)
	}
	return res
}

func GetGatewayRoute(router, method string) (*models.GatewayRouter, error) {
	var routerInspect models.GatewayRouter
	if err := GetOrmer().QueryTable("sys_resource_router").Filter("router", router).
		Filter("method", method).One(&routerInspect); err != nil {
		return nil, err
	}
	return &routerInspect, nil
}

func GatewayRouteListByBackend(backend string) ([]*models.GatewayRouter, error) {
	var routers []*models.GatewayRouter
	if _, err := GetOrmer().QueryTable("sys_resource_router").Filter("backend", backend).All(&routers); err != nil {
		return nil, err
	}
	return routers, nil
}

func DeleteGatewayRouteByBackend(backend string) error {
	sql := `delete from sys_resource_router where backend=?`
	if _, err := GetOrmer().Raw(sql, backend).Exec(); err != nil {
		return err
	}
	return nil
}
