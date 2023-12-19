package main

import (
	"fmt"

	"log"

	"os"

	"net/http"

	"github.com/go-chi/chi"

	"github.com/go-chi/cors"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hello world")

	//Load env variables from the .env file
	godotenv.Load(".env")

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("Default port is not set, please set it")
	}

	router := chi.NewRouter()
	router.Use( cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: false,
		MaxAge: 300,
	}) )

	// First version of my router
	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handlerError)

	router.Mount("/v1", v1Router)

	server :=  &http.Server{
		Handler: router,
		Addr: ":" + portString,
	}

	log.Printf("Server starting on port %v", portString)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}