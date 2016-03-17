package model

import "time"


//RefreshToken token entity
type RefreshToken struct {
	accessToken AccessToken //Access token associated to refresh token
	accessTokenID string //Id of the access token

	StorageImpl IRefreshTokenStorage

	CommonToken
}

type IRefreshToken interface {
	SetAccessTokenID(id string) error //Set the ID of the associated access token
	SetAccessToken(token AccessToken) error //Associate an access token
	GetAccessToken() (AccessToken, error) //Return access token
}

type IRefreshTokenStorage interface {
	Create(sessionID string, expired time.Time, accessTokenId string)
	Delete(accessToken string)
}

//SetAccessTokenID Set the ID of the associated access token
func (rt *RefreshToken) SetAccessTokenID(id string) error {
	rt.accessTokenID = id
	return nil
}

//SetAccessToken Associate an access token
func (rt *RefreshToken) SetAccessToken(token AccessToken) error {
	rt.accessToken = token
	return nil
}

//GetAccessToken Return access token
func (rt *RefreshToken) GetAccessToken() (AccessToken, error) {
	if rt.accessToken == nil {
		rt.accessToken.StorageImpl.Get(rt.accessTokenID, &rt.accessToken)
	}
	return rt.accessToken, nil
}

func (rt *RefreshToken) Save() error {
	at, _ := rt.GetAccessToken()
	return rt.StorageImpl.Create(
		rt.GetID(),
		rt.GetExpiredTime(),
		at.GetID(),
	)
}

func (rt *RefreshToken) Expired() error {
	rt.StorageImpl.Delete(rt);
	return nil
}