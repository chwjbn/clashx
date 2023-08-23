package auth

import (
	"github.com/chwjbn/clashx/component/auth"
)

var authenticator auth.Authenticator

func Authenticator() auth.Authenticator {
	return authenticator
}

func SetAuthenticator(au auth.Authenticator) {
	authenticator = au
}
