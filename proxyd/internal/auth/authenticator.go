package auth

// Authenticator is an interface for authenticating users.
type Authenticator interface {
	// Authenticate authenticates the user with the given username and password.
	// If the authentication is successful, it returns nil.
	Authenticate(username, password string) error
}
