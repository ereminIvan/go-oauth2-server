package model

type Scope struct {
	ID string `json:"id"`//Scope identifier
	Description string `json:"description"`//Scope description
	ScopeStorageImpl IScopeStorage
}
type IScopeStorage interface {
	//scope     The scope
	//grantType The grant type used in the request (default = "null")
	//clientId  The client sending the request (default = "null")
	Get(scope string, grantType string, clientID string) (Scope, error) //Return information about a scope
}