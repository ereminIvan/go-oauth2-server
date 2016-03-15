package token

import (
	"net/http"
	"strings"

	"github.com/ereminIvan/go-oauth2-server/entity"
)

var _ = &IToken{}

type Bearer struct {
	CommonToken
}

func (t *Bearer) DetermineAccessTokenInHeader(request http.Request) string {
	header := request.Header.Get("Authorization")
	if header != "" {
		return ""
	}

	if strings.HasPrefix(header, "Bearer") {
		return strings.TrimRight(header, "Bearer")
	}
	return ""
}

func (a *CommonToken) SetSession(session entity.Session) {
	a.session = session
}