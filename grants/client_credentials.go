package grants

import (
	"errors"
	"net/http"
	"encoding/json"
	"strings"
	"time"

	"github.com/ereminIvan/go-oauth2-server/model"
	"github.com/ereminIvan/go-oauth2-server"
	"github.com/ereminIvan/go-oauth2-server/service/authroization"
)

type ClientCredentials struct {
	Grant
}

func NewClientCredentials() *ClientCredentials {
	return &ClientCredentials{
		Grant.identifier:"client_credentials",
	}
}

type params struct {
	ClientID int `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Scopes string `json:"scope"`
}

func (g *ClientCredentials) CompleteFlow(request http.Request) (map[interface{}]interface{}, error) {
	p := &params{}
	err := json.Unmarshal(request.Body, p)

	if err != nil {
		return err
	}

	if p.ClientID == 0 {
		return errors.New("Invalid Client ID")
	}

	if p.ClientSecret == "" {
		return errors.New("Invalid Client Secret")
	}

	// Validate any scopes that are in the request
	scopeDelimiter := " " //todo get from service
	scopeParameterRequired := true //todo get from service
	defaultScope := []string{} //todo get from service
	var scopes []string
	for _, s := range strings.Split(strings.Trim(p.Scopes, " "), scopeDelimiter) {
		if len(s) > 0 {
			scopes = append(scopes, s)
		}
	}
	if scopeParameterRequired && len(defaultScope) == 0 && len(scopes) == 0 {
		return errors.New("Invalid Scopes count")
	}
	if len(scopes) == 0 && len(defaultScope) > 0 {
		scopes = defaultScope
	}
	//todo ... a lot validation action with Client, Service Auth

	// Create a new session
	session := model.Session{Client:Client, ownerType: "client", ownerID: Client.GetID()}

	// Generate an access token
	accessToken := &model.AccessToken{}
	accessToken.SetID(go_oauth2_server.Generator{10}.Generate())
	accessToken.Server = server
	accessToken.SetExpiredTime(time.Now().Add(g.accessTokenTTL))

	// Associate scopes with the session and access token
	session.AssociateScope(scopes)

	// Save everything
	session.Save()
	accessToken.SetSession(session)
	accessToken.Save()

	service := authorization.Service{}
	service.GetTokenType().SetSession(session)
	service.GetTokenType().SetParam("access_token", accessToken.GetID())
	service.GetTokenType().SetParam("expires_in", accessToken.CommonToken.GetExpiredTime())

return $this->server->getTokenType()->generateResponse();


return nil
}
