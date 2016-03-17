package model

import "github.com/ereminIvan/go-oauth2-server/api/service/storage"

type Scope struct {
	ID string `json:"id"`//Scope identifier
	Description string `json:"description"`//Scope description
	ScopeStorageImpl storage.IScope
}