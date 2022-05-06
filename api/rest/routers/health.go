package routers

import (
	v1 "template-service/api/rest/v1"
	"template-service/services"

	"github.com/gorilla/mux"
)

// WithHealthRoutes defines all the Health check routes
// These routes don't trace and don't require authentication
// health
func WithHealthRoutes(router *mux.Router, s services.Services) RoutesConfig {
	return func(r *Routers) {
		healthRouter := router.PathPrefix("/health").Subrouter()
		healthController := v1.NewHealthController(s)

		healthRouter.HandleFunc("", healthController.GetHealth).Methods("GET")
		r.Health = healthRouter
	}
}
