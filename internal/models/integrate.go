package models

import (
	"encoding/base64"
	"gitee.com/plutoccc/devops_app/utils"
)

// IntegrateSetting the Basic Data of stages based on commpany
type IntegrateSetting struct {
	Addons
	Name        string `orm:"column(name);size(64)" json:"name"`
	Type        string `orm:"column(type);size(64)" json:"type"`
	Config      string `orm:"column(config);type(text)" json:"config"`
	Description string `orm:"column(description);size(256)" json:"description"`
	Creator     string `orm:"column(creator);size(64)" json:"creator"`
}

// TableName ...
func (t *IntegrateSetting) TableName() string {
	return "sys_integrate_setting"
}

func (t *IntegrateSetting) CryptoConfig(raw string) {
	t.Config = t.crypto(raw)
}

func (t *IntegrateSetting) DecryptConfig() string {
	return t.decrypt()
}

func (t *IntegrateSetting) crypto(raw string) string {
	plainText := []byte(raw)
	return base64.StdEncoding.EncodeToString(utils.AesEny(plainText))
}

func (t *IntegrateSetting) decrypt() string {
	cfg, _ := base64.StdEncoding.DecodeString(t.Config)
	return string(utils.AesEny(cfg))
}
