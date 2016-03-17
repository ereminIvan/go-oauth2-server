package token

import (
	"net/http"

	"github.com/ereminIvan/go-oauth2-server/api/model"
)

type CommonToken struct {
	tokenType		string `json:"token_type"`
	accessToken		string `json:"access_token"`
	expiredIn		string `json:"expired_in"`
	refreshToken	string `json:"refresh_token"`

	session model.ISession
}

type IToken interface {
	SetSession(session model.ISession) //Set Session
	DetermineAccessTokenInHeader(request http.Request) string //Determine the access token in the authorization header
}