package google_oauth

import "github.com/gorilla/pat"

type Router interface {
	Route() *pat.Router
}
