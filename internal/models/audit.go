package models

type Audit struct {
	Addons
	User            string `orm:"column(user)" json:"user"`
	Method          string `orm:"column(method)" json:"method"`
	Operation       string `orm:"column(operation)" json:"operation"`
	OperationObject string `orm:"column(operation_object)" json:"operation_object"`
	OperationBody   string `orm:"column(operation_body);type(text)" json:"operation_body"`
	OperationStatus int    `orm:"column(operation_status)" json:"operation_status"`
}

func (t *Audit) TableName() string {
	return "sys_audit"
}
