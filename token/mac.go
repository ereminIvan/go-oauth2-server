package token

import (
	"net/http"

	"github.com/ereminIvan/go-oauth2-server/model"
)

var _ = &IToken{}

//MAC Token type
type MAC struct {
	CommonToken
}

func (t *MAC) DetermineAccessTokenInHeader(request http.Request) string {
	return ""
}

func (t *MAC) SetSession(session model.ISession) {
	t.session = session
}