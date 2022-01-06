package main

import (
	"google_oauth/app"
	"google_oauth/service"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
)

func main() {

	// session config
	key := "ThIs iS my kEY"
	maxAge := 120 // 2 seconds
	isProd := false

	// store config
	store := sessions.NewCookieStore([]byte(key))
	store.MaxAge(maxAge)
	store.Options.Path = "/"
	store.Options.HttpOnly = true
	store.Options.Secure = isProd

	gothic.Store = store

	// oauth credential
	goth.UseProviders(
		google.New("504511383397-683g45fakmm2rhmh72m0pirii3s2hlbl.apps.googleusercontent.com", "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", "http://localhost:3000/auth/google/callback", "email", "profile"),
	)

	db, err := app.NewDB()
	if err != nil {
		log.Fatal(err)
	}
	s := service.NewUserAuthServiceImpl(db)
	r := app.NewRouter(s)
	r.Route()
	log.Println("listening on localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", r.Router))
}
