// API Key: SAgn9qi1Zt3OV8xfDaMRiBQvK
// API Key Secret: cTPCU5SZHOVupwkFJ694mYktRKn71uJA7NYyhegUxUFBKw5z5C
//Bearer Token: AAAAAAAAAAAAAAAAAAAAAJWMhAEAAAAA7JzOxzqXwPIXNflSVHaekkFhVRc%3DoG1M48ytfPHyKfzh5ACJNOCYQTTcLgABlBmxGHWnNZ6jCSN3yS

package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"sort"

	"github.com/gorilla/pat"
	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/twitter"
)

type ProviderIndex struct {
	Providers    []string
	ProvidersMap map[string]string
}

func main() {
	key := "SAgn9qi1Zt3OV8xfDaMRiBQvK"
	// maxAge := 86400 * 30 // 30 days
	// isProd := true

	store := sessions.NewCookieStore([]byte(key))
	// store.MaxAge(maxAge)
	// store.Options.Path = "/"
	// store.Options.HttpOnly = true // HttpOnly should always be enabled
	// store.Options.Secure = isProd

	gothic.Store = store

	goth.UseProviders(
		twitter.New("SAgn9qi1Zt3OV8xfDaMRiBQvK", "cTPCU5SZHOVupwkFJ694mYktRKn71uJA7NYyhegUxUFBKw5z5C", "http://127.0.0.1:3000/auth/twitter/callback"),
	)

	m := make(map[string]string)
	m["twitter"] = "Twitter"

	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	providerIndex := &ProviderIndex{
		Providers:    keys,
		ProvidersMap: m,
	}

	p := pat.New()
	p.Get("/auth/{provider}/callback", func(res http.ResponseWriter, req *http.Request) {

		user, err := gothic.CompleteUserAuth(res, req)
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		t, _ := template.ParseFiles("templates/success.html")
		t.Execute(res, user)
	})

	p.Get("/logout/{provider}", func(res http.ResponseWriter, req *http.Request) {
		gothic.Logout(res, req)
		res.Header().Set("Location", "/")
		res.WriteHeader(http.StatusTemporaryRedirect)
	})

	p.Get("/auth/{provider}", func(res http.ResponseWriter, req *http.Request) {
		// try to get the user without re-authenticating
		if gothUser, err := gothic.CompleteUserAuth(res, req); err == nil {
			t, _ := template.ParseFiles("templates/success.html")
			t.Execute(res, gothUser)
		} else {
			gothic.BeginAuthHandler(res, req)
		}
	})

	p.Get("/", func(res http.ResponseWriter, req *http.Request) {
		t, _ := template.ParseFiles("templates/index.html")
		t.Execute(res, providerIndex)
	})
	log.Println("listening on localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", p))
}
