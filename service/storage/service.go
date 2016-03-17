package storage

import (
	"time"
	"net/url"

	"github.com/ereminIvan/go-oauth2-server/model"
)

type Service struct {
	//Storage List
	clientStorageImpl IClient
	scopeStorageImpl IScope
	sessionStorageImpl ISession
	refreshTokenStorageImpl IRefreshToken
	accessTokenStorageImpl IAccessToken
	authCodeStorageImpl IAuthCode
}

type IClient interface {
	GetBySession(s *model.Session) (model.IClient, error)
}

type IScope interface {
	//scope     The scope
	//grantType The grant type used in the request (default = "null")
	//clientId  The client sending the request (default = "null")
	Get(scope string, grantType string, clientID string) (model.Scope, error) //Return information about a scope
}

type ISession interface {
	GetByAccessToken(token model.AccessToken) (model.Session, error) //Get a session from an access token
	GetByAuthCode(authCode model.AuthCode) (model.Session, error) //Get a session from an auth code
	GetScopes(session model.Session) []model.Scope //Get a session's scopes
	AssociateScope(session *model.Session, scope *model.Scope) error //Associate a scope with a session
	//ownerType - Session owner's type (user, client)
	//ownerID - Session owner's ID
	//clientID - Client ID
	//clientRedirectURI - Client redirect URI (default = null)
	Create(ownerType string, ownerID uint64, clientID, clientRedirectURI string) (uint64, error) //Create a new session
}

type IRefreshToken interface {
	Create(sessionID string, expired time.Time, accessTokenId string)
	Delete(accessToken string)
}

//IAccessToken
type IAccessToken interface {
	Get(token string, *model.AccessToken) //Get an instance of AccessToken
	AssociateScope(token model.AccessToken, scope model.Scope)//Associate a scope with an access token
	GetScopes(token model.AccessToken) //Get the scopes for an access token

	Create(token string, expireTime time.Time, sessionID uint64) error//Creates a new access token
	Delete(model.AccessToken) //Delete an access token
}

type IAuthCode interface {
	GetScopes(ac model.AuthCode) ([]model.Scope, error)
	AssociateScope(ac model.AuthCode, scopes []model.Scope) error
	Create(acID uint64, expiredTime time.Time, sessionID uint64, redirectURI url.URL) error
	Delete(ac model.AuthCode) error
}