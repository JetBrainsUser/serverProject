/**
 * @fileOverview
 * This is a sample Go code file demonstrating file documentation header.
 * It provides a brief overview of the file's purpose and usage.
 *
 * @module SampleModule
 */

// Copyright 2023 Your Name
// All rights reserved.
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

// Constants for defining timeout values in seconds
const (
	readTimeout  = 5
	writeTimeout = 10
	idleTimeout  = 120
)

// getIndexHandler handles requests to the root path ("/") and responds with a simple message.
func getIndexHandler(w http.ResponseWriter, r *http.Request) {
	// Set the response content type to plain text
	w.Header().Set("Content-Type", "text/plain")
	returnStatus := http.StatusOK
	// Set the HTTP status code to 200 (OK)
	w.WriteHeader(returnStatus)
	message := fmt.Sprintf("Hi  %s!", r.UserAgent())
	// Write a greeting message to the response writer
	_, err := w.Write([]byte(message))
	if err != nil {
		return
	}
}

func main() {
	// Define the server's address and port
	serverAddress := ":8080"
	// Create a new logger for server logs
	l := log.New(os.Stdout, "sample-srv ", log.LstdFlags|log.Lshortfile)
	// Create a new Gorilla Mux router
	m := mux.NewRouter()

	// Register the getIndexHandler function to handle requests to the root path ("/")
	m.HandleFunc("/", getIndexHandler)

	// Create an HTTP server with specified settings
	srv := &http.Server{
		Addr:         serverAddress,
		ReadTimeout:  readTimeout * time.Second,
		WriteTimeout: writeTimeout * time.Second,
		IdleTimeout:  idleTimeout * time.Second,
		Handler:      m, // Set the router as the request handler
	}

	// Log a message indicating that the server has started
	l.Println("server started")
	// Start the HTTP server, and if there's an error, panic
	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
