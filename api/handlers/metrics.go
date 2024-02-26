package handlers

import (
	"fmt"
	"net/http"

	"github.com/granadosbrand/go-web-server/config"
)

func ResetMetrics(c *config.ApiConfig, w http.ResponseWriter, r *http.Request) {
	c.FileserverHits = 0
}



func HandlerMetrics(c *config.ApiConfig, w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(`
		<html>

		<body>
			<h1>Welcome, Chirpy Admin</h1>
			<p>Chirpy has been visited %d times!</p>
		</body>

		</html>
			`, c.FileserverHits)))

}
