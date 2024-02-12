package main

import (
	"fmt"
	"log"
	"net/http"
)

type apiConfig struct {
	fileserverHits int
}

func (c *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Incrementa el contador con cada petición.
        c.fileserverHits++

        // Continúa con el manejo de la petición.
        next.ServeHTTP(w, r)
    })
}




func (c *apiConfig) handlerMetrics(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	result := fmt.Sprint("Hits: ", c.fileserverHits)
	w.Write([]byte(result))
}
func (c *apiConfig) resetMetrics(w http.ResponseWriter, r *http.Request) {

	c.fileserverHits = 0
}

func main() {
	const filepathRoot = "."
	const port = "8080"

	apiCfg := apiConfig{
		fileserverHits: int(0),
	}

	mux := http.NewServeMux()
	mux.Handle("/app/", http.StripPrefix("/app/", apiCfg.middlewareMetricsInc(http.FileServer(http.Dir(filepathRoot)))))

	mux.HandleFunc("/healthz", handlerReadiness)
	mux.HandleFunc("/app/assets", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filepathRoot+"/assets/index.html")

	})
	mux.HandleFunc("/metrics", apiCfg.handlerMetrics)
	mux.HandleFunc("/reset", apiCfg.resetMetrics)
	corsMux := middlewareCors(mux)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: corsMux,
	}

	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
	log.Fatal(srv.ListenAndServe())
}


