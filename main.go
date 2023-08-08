package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

const (
	readTimeout  = 5
	writeTimeout = 10
	idleTimeout  = 120
)

func getIndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	returnStatus := http.StatusOK
	w.WriteHeader(returnStatus)
	message := fmt.Sprintf("Good morning %s!", r.UserAgent())
	write, err := w.Write([]byte(message))
	if err != nil {
		return
	}
	print(write)
}

func main() {
	serverAddress := ":8080"
	l := log.New(os.Stdout, "sample-srv ", log.LstdFlags|log.Lshortfile)
	m := mux.NewRouter()

	m.HandleFunc("/", getIndexHandler)

	srv := &http.Server{
		Addr:         serverAddress,
		ReadTimeout:  readTimeout * time.Second,
		WriteTimeout: writeTimeout * time.Second,
		IdleTimeout:  idleTimeout * time.Second,
		Handler:      m,
	}

	l.Println("server started")
	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
