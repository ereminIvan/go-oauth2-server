package model

import "github.com/ereminIvan/go-oauth2-server/api/service/storage"

//AccessToken
type AccessToken struct {
	StorageImpl storage.IAccessToken

	CommonToken
}


//GetSession
func (t *AccessToken) GetSession() (ISession, error) {
	var err error

	if t.session != nil {
		return t.session
	}
	t.session, err = t.session.GetStorage().GetByAccessToken(t)

	return t.session, err
}

//HasScope Check if access token has an associated scope
func (t *AccessToken) HasScope(scopeID string) bool {
	if t.scopes == nil {
		t.GetScopes()
	}
	_, ok := t.scopes[scopeID];
	return ok
}

//GetScopes Return all scopes associated with the access token
func (t *AccessToken) GetScopes() []Scope {
	t.scopes = t.StorageImpl.GetScopes(t)
	return t.scopes
}

//Save
func (t *AccessToken) Save() error {
	session, _ := t.GetSession()
	t.StorageImpl.Create(t.GetID(), t.GetExpiredTime(), session.GetID())

	for _, scope := range t.GetScopes() {
		t.StorageImpl.AssociateScope(t, scope)
	}
	return nil
}

func (t *AccessToken) Expired() error{
	t.StorageImpl.Delete(t)
	return nil
}