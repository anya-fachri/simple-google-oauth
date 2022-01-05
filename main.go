package main

import (
	google_oauth "google_oauth/router"
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
	maxAge := 86400 * 30 // 30 days
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

	// router and server
	r := google_oauth.NewRouter().Route()
	log.Println("listening on localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
