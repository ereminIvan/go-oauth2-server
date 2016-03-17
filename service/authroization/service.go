package authorization

import (
	"time"
	"errors"

	"github.com/ereminIvan/go-oauth2-server/token"
	"github.com/ereminIvan/go-oauth2-server/grants"
	"github.com/ereminIvan/go-oauth2-server/service/storage"
)

var (
	ErrorInvalidGrantType = errors.New("Invalid Grant type")
)


type Service struct {
	token *token.IToken
	//The registered grant types
	grantTypes map[string]grants.IGrant

	//Storage services
	storageService *storage.Service

	//The delimiter between scopes specified in the scope query string parameter
	//The OAuth 2 specification states it should be a space but most use a comma
	scopeDelimiter string //todo unused
	//The TTL (time to live) of an access token in seconds (default: 3600)
	accessTokenTTL time.Time //todo unused
	//The registered grant response types
	responseTypes []string //todo unused
	//Require the "scope" parameter to be in checkAuthoriseParams()
	requireScopeParam bool //todo unused
	//Default scope(s) to be used if none is provided
	defaultScope []string //todo unused
	//Require the "state" parameter to be in checkAuthoriseParams()
	requireStateParam bool //todo unused

}

type IService interface {
	AddGrantType(g grants.IGrant)
}

func NewService(t *token.IToken, ss *storage.Service) *Service {
	return &Service{
		token:t,
		storageService:ss,
	}
}

//AddGrantType Add range of grant type to authorization service
func (s *Service) AddGrantType(g ... grants.IGrant) {
	for _, grant := range g {
		if grant == nil {
			return ErrorInvalidGrantType
		}
		if len(s.grantTypes) {
			s.grantTypes = make(map[string]grants.IGrant, len(g))
		}
		if _, ok :=  s.grantTypes[grant.GetIdentifier()]; !ok {
			s.grantTypes[grant.GetIdentifier()] = grant
		}
	}
}