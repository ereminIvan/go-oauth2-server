package grants

type IGrant interface {
	GetIdentifier() string //Return the identifier
	SetIdentifier(id string) //Return the identifier
	GetResponseType() string //Return the response type
	CompleteFlow() map[interface{}]interface{} //Complete the grant flow
}

type Common struct {
	identifier string
	resourceType string
}