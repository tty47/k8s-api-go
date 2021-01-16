package v1

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/jrmanes/k8s-api-go/internal/data"
)

// New returns the API V1 Handler with configuration.
func New() http.Handler {
	r := chi.NewRouter()

	// Mount user routes
	ur := &UserRouter{
		Repository: &data.UserRepository{
			Data: data.New(),
		},
	}
	r.Mount("/", ur.Routes())

	return r
}
