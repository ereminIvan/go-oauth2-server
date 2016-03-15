package token

import (
	"net/http"

	"github.com/ereminIvan/go-oauth2-server/entity"
)

var _ = &IToken{}

//MAC Token type
type MAC struct {
	CommonToken
}

func (t *MAC) DetermineAccessTokenInHeader(request http.Request) string {
	return ""
}

func (t *MAC) SetSession(session entity.Session) {
	t.session = session
}