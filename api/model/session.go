package model
import "github.com/ereminIvan/go-oauth2-server/api/service/storage"

type Session struct {
	ID           uint64
	client       Client                  //Session Client identifier
	ownerID      uint64                  //Session owner identifier
	ownerType    string                  //Session owner type (e.g. "user")
	authCode     AuthCode                //Auth code
	accessToken  AccessToken             //Access token
	refreshToken RefreshToken            //Refresh token
	scopes       map[uint64]Scope        //Session scopes

	StorageImpl storage.ISession
	ClientStorageImpl storage.IClient
}

type ISession interface {
	SetID(id uint64) error
	GetID() uint64
	AssociateScope(scope Scope) error
	HasScope(scope Scope) bool
	GetScopes() map[uint64]Scope
	AssociateAccessToken(token AccessToken) error
	GetStorage() storage.ISession
}

//AssociateScope Associate a scope
func (s *Session) AssociateScope(scope Scope) {
	if _, ok := s.scopes[scope.ID]; !ok {
		s.scopes[scope.ID] = scope
	}
	return nil
}
//HasScope Check if access token has an associated scope
func (s *Session) HasScope(scope Scope) bool {
	_, ok := s.scopes[scope.ID]
	return ok
}

//getScopes Return all scopes associated with the session
func (s *Session) GetScopes() map[uint64]Scope {
	return s.scopes
}

//AssociateAccessToken Associate an access token with the session
func (s *Session) AssociateAccessToken(token ICommonToken) error {
	s.accessToken = token
	return nil
}

//AssociateRefreshToken Associate a refresh token with the session
func (s *Session) AssociateRefreshToken(token IRefreshToken) error {
	s.refreshToken = token
	return nil
}

//AssociateClient Associate a client with the session
func (s *Session) AssociateClient(client IClient) error {
	s.client = client
	return nil
}

//GetClient Return the session client
func (s *Session) GetClient() IClient {
	if s.client != nil {
		return s.client
	}
	s.client = s.ClientStorageImpl.GetBySession(s)
	return s.client
}

//SetOwner Set the session owner
func (s *Session) SetOwner(ownerType string, ownerID string) error {
	s.ownerID = ownerID
	s.ownerType = ownerType
	//s.server.getEventEmitter().emit(SessionOwnerEvent{Session:s})

	return nil
}

//GetOwnerID Return session owner identifier
func (s *Session) GetOwnerID() string {
	return s.ownerID
}

//GetOwnerType Return session owner type
func (s *Session) GetOwnerType() string {
	return s.ownerType
}

//Save the session
func (s *Session) Save() error {
	// Save the session and get an identifier
	s.ID = s.StorageImpl.Create(
		s.GetOwnerType(),
		s.GetOwnerID(),
		s.GetClient().GetID(),
		s.GetClient().GetRedirectURI(),
	)
	for _, scope := range s.GetScopes() {
		s.StorageImpl.AssociateScope(&s, &scope)
	}
	return nil
}
