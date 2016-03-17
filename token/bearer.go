package token

import (
	"net/http"
	"strings"

	"github.com/ereminIvan/go-oauth2-server/model"
)

var _ = &IToken{}

var tokenType = "Bearer"

type Bearer struct {
	CommonToken
}

func (t *Bearer) DetermineAccessTokenInHeader(request http.Request) string {
	header := request.Header.Get("Authorization")
	if header != "" {
		return ""
	}

	if strings.HasPrefix(header, tokenType) {
		return strings.TrimRight(header, tokenType)
	}
	return ""
}

func (a *CommonToken) SetSession(session model.ISession) {
	a.session = session
}