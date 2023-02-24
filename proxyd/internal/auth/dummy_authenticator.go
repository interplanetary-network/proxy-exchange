package auth

type dummyAuthenticator struct {
}

func NewDummyAuthenticator() Authenticator {
	return &dummyAuthenticator{}
}

func (a *dummyAuthenticator) Authenticate(username, password string) error {
	return nil
}
