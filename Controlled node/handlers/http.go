package handlers

import (
	"controlled-node/core"
	"controlled-node/observability"
	"encoding/json"
	"net/http"
)

func StatusHandler(sys *core.System) http.HandlerFunc { // Handler to return system status
	return func(w http.ResponseWriter, r *http.Request) {
		status := observability.Snapshot(sys)
		json.NewEncoder(w).Encode(status)
	}
}

func FaultHandler(sys *core.System) http.HandlerFunc { // Handler to set system mode based on query parameter
	return func(w http.ResponseWriter, r *http.Request) {

		mode := r.URL.Query().Get("mode")
		switch mode { // Validate and set mode
		case "normal", "overload", "degraded": // Valid modes
			sys.SetMode(mode)
			w.Write([]byte("mode set to " + mode))
			return

		default: // Handle invalid mode
			statusText := http.StatusText(http.StatusBadRequest)
			w.Write([]byte(statusText + ": Invalid mode: " + mode))
			return
		}
	}
}
