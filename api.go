package main

import (
	"errors"
	"log"
	"net/http"
	"runtime/debug"

	"github.com/gorilla/mux"
)

func getBuildInfo() (string, error) {
	buildInfo, ok := debug.ReadBuildInfo()
	if !ok {
		return "", errors.New("unable to read build info")
	}
	return buildInfo.String(), nil
}

func index(w http.ResponseWriter, _ *http.Request) {
	if _, err := w.Write([]byte("hello world")); err != nil {
		log.Println(err)
	}
}

func info(w http.ResponseWriter, _ *http.Request) {
	buildInfo, err := getBuildInfo()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if _, err := w.Write([]byte(err.Error())); err != nil {
			log.Println(err)
		}
		return
	}
	if _, err := w.Write([]byte(buildInfo)); err != nil {
		log.Println(err)
	}
}

func handler() http.Handler {
	r := mux.NewRouter()
	r.Handle("/", http.HandlerFunc(index)).Methods(http.MethodGet)
	r.Handle("/info", http.HandlerFunc(info)).Methods(http.MethodGet)
	return r
}

func main() {
	s := http.Server{
		Addr:    ":8998",
		Handler: handler(),
	}
	log.Println(s.Addr)
	log.Println(getBuildInfo())
	log.Fatal(s.ListenAndServe())
}
