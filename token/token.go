package token

import (
	"net/http"

	"github.com/ereminIvan/go-oauth2-server/entity"
)

type CommonToken struct {
	tokenType		string `json:"token_type"`
	accessToken		string `json:"access_token"`
	expiredIn		string `json:"expired_in"`
	refreshToken	string `json:"refresh_token"`

	session entity.ISession
}

type IToken interface {
	SetSession(session entity.ISession) //Set Session
	DetermineAccessTokenInHeader(request http.Request) string //Determine the access token in the authorization header
}