package main

import (
	"net/http"

	r "github.com/beeceej/reader"
	"github.com/beeceej/reader/db"
	"github.com/beeceej/reader/puppies"
)

func API(appReader r.MonadReader) *http.ServeMux {
	mux := puppies.API(appReader, http.DefaultServeMux)
	return mux
}

func main() {
	appEnv := r.NewEnv()
	appReader := r.NewReader(appEnv).
		Bind(db.MySQLReader).
		Bind(puppies.RepositoryReader)
	mux := API(appReader)
	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(err.Error())
	}
}
