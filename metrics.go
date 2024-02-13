package main

import (
	"fmt"
	"net/http"
)


func (c *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Incrementa el contador con cada petición.
		c.fileserverHits++
		fmt.Println("Hit added")

		// Continúa con el manejo de la petición.
		next.ServeHTTP(w, r)
	})
}

func (c *apiConfig) resetMetrics(w http.ResponseWriter, r *http.Request) {
	c.fileserverHits = 0
}

func (c *apiConfig) handlerMetrics(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(`
		<html>

		<body>
			<h1>Welcome, Chirpy Admin</h1>
			<p>Chirpy has been visited %d times!</p>
		</body>

		</html>
			`, c.fileserverHits)))

}
