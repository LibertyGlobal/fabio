package api

import (
	"net/http"

	"github.com/LibertyGlobal/fabio/config"
)

var Cfg *config.Config

func HandleConfig(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, r, Cfg)
}
