package grants

import (
	"net/http"
	"time"
)

type IGrant interface {
	CompleteFlow(request *http.Request) map[interface{}]interface{} //Complete the grant flow
}

type Grant struct {
	identifier     string
	responseType   string
	accessTokenTTL time.Duration
}

func (g *Grant) GetIdentifier() string {
	return g.identifier
}

func (g *Grant) SetIdentifier(id string) {
	g.identifier = id
}

func (g *Grant) GetResponseType() string {
	return g.responseType
}

func (g *Grant) GetAccessTokenTTl() time.Time {
	return g.accessTokenTTL
}
func (g *Grant) SetAccessTokenTTl(ttl time.Time) {
	g.accessTokenTTL = ttl
}
