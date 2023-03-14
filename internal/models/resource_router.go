package models

type GatewayRouter struct {
	Addons
	Router            string `orm:"column(router)" json:"router"`
	Method            string `orm:"column(method)" json:"method"`
	Backend           string `orm:"column(backend)" json:"backend"`
	ResourceType      string `orm:"column(resource_type)" json:"resource_type"`
	ResourceOperation string `orm:"column(resource_operation)" json:"resource_operation"`
}

func (t *GatewayRouter) TableName() string {
	return "sys_resource_router"
}

func (t *GatewayRouter) TableIndex() [][]string {
	return [][]string{
		{"Backend"},
		{"Router", "Method"},
	}
}

func (u *GatewayRouter) TableUnique() [][]string {
	return [][]string{
		{"Router", "Method"},
		{"ResourceType", "ResourceOperation"},
	}
}
