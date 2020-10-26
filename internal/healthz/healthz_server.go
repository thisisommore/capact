package healthz

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"

	"projectvoltron.dev/voltron/pkg/httputil"
)

func NewHTTPServer(log *zap.Logger, healthzAddr, appName string) httputil.StartableServer {
	router := mux.NewRouter()
	router.HandleFunc("/healthz", func(w http.ResponseWriter, _ *http.Request) {
		if _, err := fmt.Fprintf(w, "%s - OK", appName); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	return httputil.NewStartableServer(
		log.Named(appName).With(zap.String("server", "healthz")),
		healthzAddr,
		router,
	)
}