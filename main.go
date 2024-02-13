package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type apiConfig struct {
	fileserverHits int
}


func main() {
	const filepathRoot = "."
	const port = "8080"

	router := chi.NewRouter()
	apiRouter := chi.NewRouter()
	adminRouter := chi.NewRouter()

	corsChi := middlewareCors(router)

	router.Mount("/api", apiRouter)
	router.Mount("/admin", adminRouter)

	apiCfg := apiConfig{
		fileserverHits: int(0),
	}

	fsHandler := apiCfg.middlewareMetricsInc(http.StripPrefix("/app", http.FileServer(http.Dir(filepathRoot))))
	router.Handle("/app", fsHandler)
	router.Handle("/app/*", fsHandler)

	apiRouter.Get("/healthz", handlerReadiness)
	apiRouter.Get("/reset", apiCfg.resetMetrics)

	adminRouter.Get("/metrics", apiCfg.handlerMetrics)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: corsChi,
	}

	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
	log.Fatal(srv.ListenAndServe())
}
