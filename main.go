package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// Req: http://localhost:1234/upper?word=abc
// Res: ABC
func upperCaseHandler(w http.ResponseWriter, r *http.Request) {
	query, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "invalid request")
		return
	}

	word := query.Get("word")
	if len(word) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "missing word")
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, strings.ToUpper(word))
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/upper", upperCaseHandler)

	log.Fatal(http.ListenAndServe(":1234", r))
}
