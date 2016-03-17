package model
import "github.com/ereminIvan/go-oauth2-server/api/service/storage"

type Client struct {
	id string //Client identifier
	secret string //Client secret
	name string //Client name
	redirectURI string //Client redirect URI
	StorageImpl storage.IClient //Authorization or resource server
}

type IClient interface {
	GetID() string //Return the client identifier
	GetSecret() string //Return the client secret
	GetName() string //Get the client name
	GetRedirectURI() string //Return the client redirect URI
}

//GetID
func (c *Client) GetID() string {
	return c.id
}

//GetSecret
func (c *Client) GetSecret() string {
	return c.secret
}

//GetName
func (c *Client) GetName() string {
	return c.name
}

//GetRedirectURI
func (c *Client) GetRedirectURI() string {
	return c.redirectURI
}