//package main

import (
	"github.com/coreos/go-systemd/activation"
	"github.com/go-chi/chi"
	"github.com/tylerb/graceful"
	"net/http"
	"time"
)

func main() {
	m := chi.NewRouter()
	m.HandleFunc("/api", ApiProxyHandler)

	listeners, err := activation.Listeners()
	if err != nil{
		panic(err)
	}
	err = graceful.Serve(&http.Server{Handler: m}, listeners[0], 5*time.Second)
	if err != nil{
		panic(err)
	}
}
