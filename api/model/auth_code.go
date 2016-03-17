package model

import (
	"net/url"

	"github.com/ereminIvan/go-oauth2-server/api/service/storage"
)

type AuthCode struct {
	CommonToken

	redirectURI url.URL

	StorageImpl storage.IAuthCode
}

//SetRedirectURI Set the redirect URI for the authorization request
func (ac *AuthCode) SetRedirectURI(uri url.URL) error {
	ac.redirectURI = uri
	return nil
}

//GetRedirectURI Get the redirect URI
func (ac *AuthCode) GetRedirectURI() url.URL {
	return ac.redirectURI
}

//GenerateRedirectURI Generate a redirect URI
//   state The state parameter if set by the client
func (ac *AuthCode) GenerateRedirectURI(state string) string {
	ac.redirectURI.Query().Add("code", ac.GetID())
	ac.redirectURI.Query().Add("state", state)
	return ac.redirectURI
}

//Get session
func (ac *AuthCode) GetSession() ISession {
	if ac.session == nil {
		ac.session, _ =  ac.session.GetStorage().GetByAuthCode(ac);
	}
	return ac.session
}

//GetScopes Return all scopes associated with the session
func (ac *AuthCode) GetScopes() []Scope {
	if ac.scopes == nil {
		ac.scopes = ac.StorageImpl.GetScopes(ac)
	}
	return ac.scopes
}
//Save
func (ac *AuthCode) Save() error {
	err := ac.StorageImpl.Create(
		ac.GetID(),
		ac.GetExpiredTime(),
		ac.GetSession().GetID(),
		ac.GetRedirectURI(),
	)
	if err != nil {
		return err
	}
	// Associate the scope with the token
	err = ac.StorageImpl.AssociateScope(ac, ac.GetScopes())
	return err
}

func (ac * AuthCode) Expired() error {
	ac.StorageImpl.Delete(ac)
	return nil
}
