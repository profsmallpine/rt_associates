package main

import (
	"log"
	"net/http"
	"os"
)

// adapter function for use in setting up middlewares.
type adapter func(http.Handler) http.Handler

// chain function for running middleware code before web request.
func chain(handler http.Handler, adapters ...adapter) http.Handler {
	for i := len(adapters) - 1; i >= 0; i-- {
		handler = adapters[i](handler)
	}

	return handler
}

// logRequest is simply for logging each hit to the web server.
func logRequest(logger *log.Logger) adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger.Println(r.Method, r.RequestURI)
			h.ServeHTTP(w, r)
		})
	}
}

func redirectToHTTPS() adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if os.Getenv("ENVIRONMENT") != "development" {
				if r.Header.Get("x-forwarded-proto") != "https" {
					url := "https://" + r.Host + r.RequestURI
					http.Redirect(w, r, url, http.StatusPermanentRedirect)
					return
				}
			}

			h.ServeHTTP(w, r)
		})
	}
}
