package app

import (
	"context"
	"fmt"
	"google_oauth/service"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/pat"
	"github.com/markbates/goth/gothic"
)

type RouterImpl struct {
	Router          *pat.Router
	UserAuthService service.UserAuthService
}

func NewRouter(u service.UserAuthService) *RouterImpl {
	return &RouterImpl{
		Router:          pat.New(),
		UserAuthService: u,
	}
}

func (r *RouterImpl) Route() {

	r.Router.Get("/auth/{provider}/callback", func(res http.ResponseWriter, req *http.Request) {

		user, err := gothic.CompleteUserAuth(res, req)
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		ctx := context.Background()
		err = r.UserAuthService.AddUser(ctx, &user)
		if err != nil {
			log.Fatal(err)
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
}
