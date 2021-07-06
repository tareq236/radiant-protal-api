package routes

import (
	"net/http"
	"radiant-protal-api/controllers"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {

	router := mux.NewRouter()

	router_version_1 := router.PathPrefix("/api/v1").Subrouter()
	router_version_1.PathPrefix("/images").Handler(http.StripPrefix("/api/v1/images", http.FileServer(http.Dir("./public/images"))))

	router_version_1.HandleFunc("/login", controllers.Login).Methods("POST")
	router_version_1.HandleFunc("/check_user", controllers.CheckUser).Methods("POST")

	// router_version_1.Use(auth.JwtAuthentication)

	// router_version_1.Use(auth.JwtAuthentication) //attach JWT auth middleware

	return router_version_1
}
