package main

import (
	"net/http"
)

func oauthAuthorize(w http.ResponseWriter, r *http.Request) {}
func oauthAccessToken(w http.ResponseWriter, r *http.Request) {}

func main() {
	http.HandleFunc("/oauth2/authorize", oauthAuthorize)
	http.HandleFunc("/oauth2/access_token", oauthAccessToken)

	http.ListenAndServe(":8080", nil)
}
