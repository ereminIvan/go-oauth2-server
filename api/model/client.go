package model

type Client struct {
	id string //Client identifier
	secret string //Client secret
	name string //Client name
	redirectURI string //Client redirect URI
	StorageImpl IClientStorage //Authorization or resource server
}

type IClient interface {
	GetID() string //Return the client identifier
	GetSecret() string //Return the client secret
	GetName() string //Get the client name
	GetRedirectURI() string //Return the client redirect URI
}

type IClientStorage interface {
	GetBySession(s *Session) (IClient, error)
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