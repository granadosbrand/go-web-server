package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
)

type apiConfig struct {
	fileserverHits int
}

func main() {

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
	apiCfg := apiConfig{
		fileserverHits: int(0),
	}

	// Routers: ==================================

	router := chi.NewRouter()
	apiRouter := chi.NewRouter()
	adminRouter := chi.NewRouter()

	corsChi := middlewareCors(router)

	router.Mount("/api", apiRouter)
	router.Mount("/admin", adminRouter)

	fsHandler := apiCfg.middlewareMetricsInc(http.StripPrefix("/app", http.FileServer(http.Dir(filepathRoot))))
	router.Handle("/app", fsHandler)
	router.Handle("/app/*", fsHandler)

	apiRouter.Get("/healthz", handlerReadiness)
	apiRouter.Get("/reset", apiCfg.resetMetrics)
	apiRouter.Post("/chirps", handleChirpPost)
	apiRouter.Get("/chirps", handleChirpGet)
	apiRouter.Get("/chirps/{chirpId}", handleChirpByID)
	apiRouter.Post("/users", handleCreateUser)

	adminRouter.Get("/metrics", apiCfg.handlerMetrics)

	// ===========================================

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: corsChi,
	}

	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
	log.Fatal(srv.ListenAndServe())
}
