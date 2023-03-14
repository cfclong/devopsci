package auth

// ExternalAccount contains queried information returned by an authenticate provider
// for an external account.
type ExternalAccount struct {
	// REQUIRED: The username of the account.
	User string
	// The nick name of the account.
	Name string
	// The email address of the account.
	Email string
	// Whether the user should be prompted as a site admin.
	Admin bool
}

// Provider defines an authenticate provider which provides ability to authentication against
// an external identity provider and query external account information.
type Provider interface {
	Authenticate(login, password string) (*ExternalAccount, error)
}
