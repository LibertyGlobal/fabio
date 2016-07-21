package admin

import (
	"fmt"
	"net/http"

	"github.com/LibertyGlobal/fabio/admin/api"
	"github.com/LibertyGlobal/fabio/admin/ui"
	"github.com/LibertyGlobal/fabio/config"
)

// ListenAndServe starts the admin api and ui server.
func ListenAndServe(cfg *config.Config, version string) error {
	ui.Version = version
	ui.Color = cfg.UI.Color
	ui.Title = cfg.UI.Title
	api.Cfg = cfg
	api.Version = version
	http.HandleFunc("/api/config", api.HandleConfig)
	http.HandleFunc("/api/manual", api.HandleManual)
	http.HandleFunc("/api/routes", api.HandleRoutes)
	http.HandleFunc("/api/version", api.HandleVersion)
	http.HandleFunc("/manual", ui.HandleManual)
	http.HandleFunc("/routes", ui.HandleRoutes)
	http.HandleFunc("/health", handleHealth)
	http.Handle("/", http.RedirectHandler("/routes", http.StatusSeeOther))
	return http.ListenAndServe(cfg.UI.Addr, nil)
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "OK")
}
