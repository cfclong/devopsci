package ldap

import (
	"fmt"

	"gitee.com/plutoccc/devops_app/pkg/auth"

	"github.com/astaxie/beego"
	ldap "github.com/colynn/go-ldap-client/v3"
)

// Provider a ldap authentication provider.
// TODO: support configuration later
type Provider struct {
	baseDN       string
	host         string
	port         int
	bindDN       string
	bindPassword string
	userFilter   string
}

// NewProvider creates a new ldap authentication provider.
func NewProvider(opts ...Option) auth.Provider {
	provider := &Provider{}
	for _, opt := range opts {
		opt(provider)
	}
	return provider
}

// Authenticate ..
func (p *Provider) Authenticate(user, password string) (*auth.ExternalAccount, error) {
	port, _ := beego.AppConfig.Int("ldap::port")
	client := &ldap.Client{
		Base:               beego.AppConfig.String("ldap::baseDN"),
		Host:               beego.AppConfig.String("ldap::host"),
		Port:               port,
		UseSSL:             false,
		BindDN:             beego.AppConfig.String("ldap::bindDN"),
		BindPassword:       beego.AppConfig.String("ldap::bindPassword"),
		UserFilter:         beego.AppConfig.String("ldap::userFilter"),
		GroupFilter:        "(memberUid=%s)",
		Attributes:         []string{"givenName", "sn", "mail", "sAMAccountName"},
		SkipTLS:            true,
		InsecureSkipVerify: true,
	}
	defer client.Close()

	authVerify, resp, err := client.Authenticate(user, password)
	if !authVerify {
		return nil, fmt.Errorf("authVerify error: %v", err)
	}

	// TODO: resp add verification
	return &auth.ExternalAccount{
		Name:  resp["sn"] + resp["givenName"],
		User:  resp["sAMAccountName"],
		Email: resp["mail"],
	}, nil
}
