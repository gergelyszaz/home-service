package service

import (
	"net/http"
	"os"

	api "github.com/gergelyszaz/home-service/gen"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// NewRouter creates a new router for any number of api routers
func NewRouter(routers ...api.Router) chi.Router {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.HandleFunc("/", index)
	auth := middleware.BasicAuth("", users)
	for _, api := range routers {
		for _, route := range api.Routes() {
			var handler http.Handler
			handler = route.HandlerFunc
			router.With(auth).Method(route.Method, route.Pattern, handler)
		}
	}

	return router
}

func index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/index.html")
}

var users = map[string]string{
	os.Getenv("USERNAME"): os.Getenv("PASSWORD"),
}
