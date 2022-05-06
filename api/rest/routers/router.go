package routers

import (
	"fmt"
	"log"
	"net/http"

	"template-service/services"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var router *mux.Router

// func middleware(next http.Handler) http.Handler{
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
// 		// do stuff
// 		next.ServeHTTP(w,r)
// 	})
// }

// ListenREST creates the routes and connects them to service functions
func ListenREST(s services.Services, port int) error {
	// Base router
	router = mux.NewRouter().StrictSlash(true).PathPrefix("").Subrouter()
	// router.Use(middleware)

	// Router for v1 of the API
	addV1Router(s)

	// Handle CORS
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"DELETE", "GET", "HEAD", "POST", "PUT", "PATCH", "OPTIONS"})
	handler := handlers.CORS(originsOk, headersOk, methodsOk)(router)

	// Get the port
	portStr := fmt.Sprintf(":%d", port)

	// Start the server and listen on the specified port
	log.Fatal(http.ListenAndServe(portStr, handler))
	
	return nil
}

// addV1Router adds all routes for the v1 API
func addV1Router(s services.Services) {
	v1Router := router.PathPrefix("/v1").Subrouter()

	addRoutes(
		WithHealthRoutes(v1Router, s),
	)
}

// RoutesConfig function definition
type RoutesConfig func(*Routers)

// Routers defines all API resource routers
type Routers struct {
	Health *mux.Router
}

func addRoutes(cfgs ...RoutesConfig) *Routers {
	var r Routers
	for _, cfg := range cfgs {
		cfg(&r)
	}
	return &r
}
