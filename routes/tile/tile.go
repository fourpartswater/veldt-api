package tile

import (
	"fmt"
	"net/http"

	log "github.com/unchartedsoftware/plog"
	"github.com/unchartedsoftware/prism/generation/tile"
	"github.com/zenazn/goji/web"

	"github.com/unchartedsoftware/prism-server/routes"
)

const (
	// Route represents the HTTP route for the resource.
	Route = "/tile" +
		"/:" + routes.TileType +
		"/:" + routes.TileIndex +
		"/:" + routes.StoreType +
		"/:" + routes.TileZ +
		"/:" + routes.TileX +
		"/:" + routes.TileY
)

func handleTileErr(w http.ResponseWriter) {
	// send error
	w.WriteHeader(500)
	fmt.Fprint(w, `{"status": "error"}`)
}

// Handler represents the HTTP route response handler.
func Handler(c web.C, w http.ResponseWriter, r *http.Request) {
	// set content type response header
	w.Header().Set("Content-Type", "application/json")
	// parse tile req from URL and body
	tileReq, err := routes.NewTileRequest(c.URLParams, r.Body)
	if err != nil {
		log.Warn(err)
		handleTileErr(w)
		return
	}
	// ensure it's generated
	err = tile.GenerateTile(tileReq)
	if err != nil {
		log.Warn(err)
		handleTileErr(w)
		return
	}
	// get tile data from store
	tileData, err := tile.GetTileFromStore(tileReq)
	if err != nil {
		log.Warn(err)
		handleTileErr(w)
		return
	}
	// send response
	w.WriteHeader(200)
	w.Write(tileData)
}
