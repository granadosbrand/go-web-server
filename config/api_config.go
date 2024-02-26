package config

import (
	"fmt"
	"net/http"
)

type ApiConfig struct {
	FileserverHits int
	JwtSecret      string
}

func (c *ApiConfig) MiddlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Incrementa el contador con cada petición.
		c.FileserverHits++
		fmt.Println("Hit added")

		// Continúa con el manejo de la petición.
		next.ServeHTTP(w, r)
	})
}
