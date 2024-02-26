package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/granadosbrand/go-web-server/api/handlers"
	"github.com/granadosbrand/go-web-server/config"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbg := flag.Bool("debug", false, "Enable debug mode")
	flag.Parse() // Parsea los argumentos de la línea de comandos

	if *dbg {
		// Código para borrar la base de datos
		fmt.Println("Modo debug activado: la base de datos se eliminará")
		e := os.Remove("database.json")
		if e != nil {
			log.Fatal(e)
		}

	} else {
		// Código normal de inicio del servidor
		fmt.Println("Modo debug desactivado: inicio normal")
	}

	const filepathRoot = "."
	const port = "8080"
	apiCfg := &config.ApiConfig{
		FileserverHits: int(0),
		JwtSecret:      os.Getenv("JWT_SECRET"),
	}

	// Routers: ==================================

	router := chi.NewRouter()
	apiRouter := chi.NewRouter()
	adminRouter := chi.NewRouter()

	corsChi := middlewareCors(router)

	router.Mount("/api", apiRouter)
	router.Mount("/admin", adminRouter)

	fsHandler := apiCfg.MiddlewareMetricsInc(http.StripPrefix("/app", http.FileServer(http.Dir(filepathRoot))))
	router.Handle("/app", fsHandler)
	router.Handle("/app/*", fsHandler)

	apiRouter.Get("/healthz", handlerReadiness)
	apiRouter.Get("/reset", func(w http.ResponseWriter, r *http.Request) { handlers.ResetMetrics(apiCfg, w, r) })
	apiRouter.Post("/chirps", handlers.HandleChirpPost)
	apiRouter.Get("/chirps", handlers.HandleChirpGet)
	apiRouter.Get("/chirps/{chirpId}", handlers.HandleChirpByID)
	apiRouter.Post("/users", handlers.HandleCreateUser)
	apiRouter.Post("/login", func(w http.ResponseWriter, r *http.Request) { handlers.HandleLogin(apiCfg, w, r) })
	apiRouter.Put("/users", func(w http.ResponseWriter, r *http.Request) { handlers.HandleUpdateUser(apiCfg, w, r) })

	adminRouter.Get("/metrics", func(w http.ResponseWriter, r *http.Request) {handlers.HandlerMetrics(apiCfg, w, r)})



	// ===========================================

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: corsChi,
	}

	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
	log.Fatal(srv.ListenAndServe())
}
