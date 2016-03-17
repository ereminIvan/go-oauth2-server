package model

import (
"time"
"github.com/ereminIvan/go-oauth2-server/api/service/storage"
)


//RefreshToken token entity
type RefreshToken struct {
	accessToken AccessToken //Access token associated to refresh token

	StorageImpl storage.IRefreshToken

	CommonToken
}

type IRefreshToken interface {
	SetAccessToken(token AccessToken) error //Associate an access token
	GetAccessToken() (AccessToken, error) //Return access token
}

//SetAccessToken Associate an access token
func (rt *RefreshToken) SetAccessToken(token AccessToken) error {
	rt.accessToken = token
	return nil
}

//GetAccessToken Return access token
func (rt *RefreshToken) GetAccessToken() (AccessToken, error) {
	if rt.accessToken == nil {
		rt.accessToken.StorageImpl.Get(&rt.accessToken)
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