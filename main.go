package main

import (
	"net/http"

	"github.com/ereminIvan/go-oauth2-server/api/service/authroization"
	"github.com/ereminIvan/go-oauth2-server/api/token"
	"github.com/ereminIvan/go-oauth2-server/api/service/storage"
)

func oauthAuthorize(w http.ResponseWriter, r *http.Request) {}
func oauthAccessToken(w http.ResponseWriter, r *http.Request) {}

func main() {
	authorization.NewService(&token.Bearer{}, &storage.Service{})

	http.HandleFunc("/oauth2/authorize", oauthAuthorize)
	http.HandleFunc("/oauth2/access_token", oauthAccessToken)

	http.ListenAndServe(":8080", nil)
}
