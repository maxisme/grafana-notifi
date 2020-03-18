package main

import (
	"github.com/coreos/go-systemd/activation"
	"github.com/go-chi/chi"
	"github.com/tylerb/graceful"
	"log"
	"net/http"
	"os"
	"time"
)

var f *os.File
func main() {
	var err error
	f, err = os.OpenFile("/var/log/grafana-notifi.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)

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
