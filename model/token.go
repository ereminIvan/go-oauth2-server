package model

import "time"

type CommonToken struct {
	id			string //Token identifier
	session		ISession //Associated session
	scopes		map[uint64]Scope //Session scopes
	expiredTime	time.Time //Token expire time
}

type ICommonToken interface {
	//Session manipulation
	SetSession(session ISession) error
	GetSession() (ISession, error)

	AssociateScope(scope Scope) error
	FormatScopes(unformatted []Scope) error

	//Expired time
	SetExpiredTime(duration time.Time) error
	GetExpiredTime() time.Time
	IsExpired() bool

	SetID(id string) error
	GetID() string

	Expire() error
	Save() error
}

func (t *CommonToken) SetSession(s ISession) error {
	t.session = s
	return nil
}

func (t *CommonToken) GetSession() (ISession, error) {
	return t.session, nil
}

func (t *CommonToken) SetExpiredTime(time time.Time) error {
	t.expiredTime = time
	return nil
}

func (t *CommonToken) GetExpiredTime() time.Time {
	return t.expiredTime
}

func (t *CommonToken) IsExpired() bool {
	return time.Since(t.expiredTime) > 0
}

func (t *CommonToken) SetID(id string) error {
	t.id = id
	return nil
}

func (t *CommonToken) GetID() string {
	if t.id == "" {

	}
	return t.id
}

func (t *CommonToken) AssociateScope(scope Scope) error {
	if _, ok := t.scopes[scope.ID]; ok {
		t.scopes[scope.ID] = scope
	}
	return nil
}

func (t *CommonToken) FormatScopes(unformatted []Scope) error {
	return nil
}

func (t *CommonToken) Expired() error {
	return nil
}

func (t *CommonToken) Save() error {
	return nil
}