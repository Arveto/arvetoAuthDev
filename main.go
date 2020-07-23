// Copyright (c) 2020, Arveto Ink. All rights reserved.
// Use of this source code is governed by a BSD
// license that can be found in the LICENSE file.

// from github.com/HuguesGuilleus/staticFile
//go:generate staticFile front/avatar.webp front/index.html

package main

import (
	"flag"
	"github.com/Arveto/arvetoAuth/pkg/public2"
	"log"
	"net/http"
	"strings"
)

var (
	provider  *public.Provider
	serverURL string
)

func main() {
	a := flag.String("a", "localhost:8000", "The listen address.")
	k := flag.String("k", "key.pem", "The private key file")
	flag.StringVar(&serverURL, "s", "http://localhost:8000", "The url of this server")
	flag.Parse()
	serverURL = strings.TrimSuffix(serverURL, "/")

	if p, err := public.NewProvider(*k); err != nil {
		log.Fatal(err)
	} else {
		provider = p
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Write(FrontIndex())
	})
	http.HandleFunc("/avatar", func(w http.ResponseWriter, r *http.Request) {
		log.Println("[AVATAR]")
		w.Header().Set("Content-Type", "image/webp")
		w.Write(FrontAvatar())
	})
	http.HandleFunc("/publickey", provider.PubHTTP)
	http.HandleFunc("/redirect", redirect)

	log.Println("[LISTEN]", *a, "for the server", serverURL)
	log.Fatal(http.ListenAndServe(*a, nil))
}

// Get parmas from URL Query, generate a JWT a redirect the user to it.
func redirect(w http.ResponseWriter, r *http.Request) {
	get := r.URL.Query().Get

	u := public.UserInfo{
		ID:     get("id"),
		Pseudo: get("pseudo"),
		Email:  get("email"),
		Avatar: serverURL + "/avatar",
	}
	if u.ID == "" || u.Pseudo == "" || u.Email == "" {
		http.Error(w, "paras id, pseudo or email are empty", http.StatusBadRequest)
		return
	}
	if err := u.Level.UnmarshalText([]byte(get("level"))); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	app := get("app")
	url := strings.TrimSuffix(get("u"), "/")

	if url == "" || app == "" {
		http.Error(w, "u or app are empty", http.StatusBadRequest)
		return
	}

	jwt, err := provider.CreateJWT(&u, app)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("[JWT]", url, u.ID)

	http.Redirect(w, r, url+"/login?jwt="+jwt+"&r="+get("r"), http.StatusTemporaryRedirect)
}
