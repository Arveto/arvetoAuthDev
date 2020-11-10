// Copyright (c) 2020, Arveto Ink. All rights reserved.
// Use of this source code is governed by a BSD
// license that can be found in the LICENSE file.

// from github.com/HuguesGuilleus/staticFile
//go:generate staticFile front/avatar.webp front/r2d2.webp front/index.html

package main

import (
	"flag"
	"github.com/Arveto/auth-go"
	"github.com/HuguesGuilleus/static.v3"
	"log"
	"net/http"
	"strings"
)

var provider *auth.Provider

func main() {
	a := flag.String("a", "localhost:8000", "The listen address.")
	k := flag.String("k", "key.pem", "The private key file")
	flag.Parse()

	var err error
	provider, err = auth.NewProvier(*k)
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/", static.Html().Func(FrontIndex))
	http.HandleFunc("/avatar", func(w http.ResponseWriter, r *http.Request) {
		log.Println("[AVATAR]")
		w.Header().Set("Content-Type", "image/webp")
		if strings.Contains(r.URL.Query().Get("u"), "bot") {
			w.Write(FrontR2d2())
		} else {
			w.Write(FrontAvatar())
		}
	})
	http.HandleFunc("/publickey", provider.ServerKey)
	http.HandleFunc("/redirect", redirect)
	http.HandleFunc("/generate", generate)

	log.Println("[LISTEN]", *a)
	log.Fatal(http.ListenAndServe(*a, nil))
}

func generate(w http.ResponseWriter, r *http.Request) {
	jwt := getJWT(w, r)
	if jwt == "" {
		return
	}
	w.Header().Set("Content-Type", "application/jwt")
	w.Write([]byte(jwt))
	w.Write([]byte("\r\n"))
}

// Get parmas from URL Query and generate a JWT and a redirect the user to it.
func redirect(w http.ResponseWriter, r *http.Request) {
	jwt := getJWT(w, r)
	if jwt == "" {
		return
	}
	url := strings.TrimSuffix(r.URL.Query().Get("u"), "/")
	http.Redirect(w, r, url+"/login?jwt="+jwt+"&r="+r.URL.Query().Get("r"), http.StatusTemporaryRedirect)
}

// Get parmas from URL Query and generate a JWT.
func getJWT(w http.ResponseWriter, r *http.Request) string {
	query := r.URL.Query()
	get := r.URL.Query().Get

	u := auth.User{
		ID:     get("id"),
		Pseudo: get("pseudo"),
		Email:  get("email"),
		Bot:    get("bot") == "true",
	}
	teams := query["teams"]
	u.Teams = make(auth.Teams, len(teams))
	for _, k := range teams {
		u.Teams[k] = true
	}

	if u.ID == "" || u.Pseudo == "" || u.Email == "" {
		http.Error(w, "Query id, pseudo or email are empty", http.StatusBadRequest)
		return ""
	}

	if err := u.Level.UnmarshalText([]byte(get("level"))); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return ""
	}

	app := get("app")
	if app == "" {
		http.Error(w, "app are empty", http.StatusBadRequest)
		return ""
	}

	jwt, err := provider.JWT(app, &u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return ""
	}

	log.Println("[JWT]", app, u.ID)

	return jwt
}
