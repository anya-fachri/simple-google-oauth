package google_oauth

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/pat"
	"github.com/markbates/goth/gothic"
)

type RouterImpl struct {
	Router *pat.Router
}

func NewRouter() *RouterImpl {
	return &RouterImpl{Router: pat.New()}
}

func (*RouterImpl) Route() *pat.Router {
	r := NewRouter()
	r.Router.Get("/auth/{provider}/callback", func(res http.ResponseWriter, req *http.Request) {

		user, err := gothic.CompleteUserAuth(res, req)
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		t, _ := template.ParseFiles("public/success.html")
		t.Execute(res, user)
	})

	r.Router.Get("/auth/{provider}", func(res http.ResponseWriter, req *http.Request) {
		gothic.BeginAuthHandler(res, req)
	})

	r.Router.Get("/", func(res http.ResponseWriter, req *http.Request) {
		t, _ := template.ParseFiles("public/index.html")
		t.Execute(res, false)
	})
	return r.Router
}
