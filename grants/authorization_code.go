package grants

import (
	"net/http"
	"time"
)

type AuthorizationCode struct {
	Grant
	authorizationTokenTTL time.Time
	requireClientSecret   bool
}

func NewAuthorizationCode() *AuthorizationCode {
	return &AuthorizationCode{
		Grant.identifier:      "authorization_code",
		Grant.responseType:    "code",
		authorizationTokenTTL: time.Now().Add(600),
		requireClientSecret:   true,
	}
}

func (g *AuthorizationCode) checkAuthorizeParams() {

}

func (g *AuthorizationCode) newAuthorizeRequest() {

}

func (g *AuthorizationCode) completeFlow(request *http.Request) {
	request.PostForm.Get("client_id")

}
