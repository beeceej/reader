package puppies

import (
	"encoding/json"
	"net/http"
	"strconv"

	r "github.com/beeceej/reader"
)

func API(appReader r.MonadReader, mux *http.ServeMux) *http.ServeMux {
	appReader.With(func(env r.Env) {
		repo := GetRepository.Run(env).(Repository)
		mux.HandleFunc("/puppies", HandleGetPuppy(repo))
	})
	return mux
}

func HandleGetPuppy(repo Repository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		rawWantedID := r.URL.Query().Get("id")
		wantedID, _ := strconv.Atoi(rawWantedID)
		puppy, _ := repo.GetPuppy(wantedID)
		b, _ := json.Marshal(puppy)
		w.Write(b)
	}
}
