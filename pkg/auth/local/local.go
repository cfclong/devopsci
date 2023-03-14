package local

import (
	"fmt"

	"gitee.com/plutoccc/devops_app/pkg/auth"

	"golang.org/x/crypto/bcrypt"
)

// Provider a local authentication provider.
// TODO: support configuration later
type Provider struct {
	name     string
	email    string
	user     string
	password string
}

// NewProvider creates a new local authentication provider.
func NewProvider(opts ...Option) auth.Provider {
	provider := &Provider{}
	for _, opt := range opts {
		opt(provider)
	}
	return provider
}

// Authenticate ..
func (p *Provider) Authenticate(loginUser, password string) (*auth.ExternalAccount, error) {
	_, err := CompareHashAndPassword(p.password, password)
	if err != nil {
		return nil, fmt.Errorf("comparehas password, error: %v", err.Error())
	}
	return &auth.ExternalAccount{
		Name:  p.name,
		Email: p.email,
		User:  p.user,
	}, nil

}

// CompareHashAndPassword ..
func CompareHashAndPassword(e string, p string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(e), []byte(p))
	if err != nil {
		return false, err
	}
	return true, nil
}
