package http

import (
	"net/http"

	"github.com/unchartedsoftware/plog"
	"github.com/unchartedsoftware/prism"
	"github.com/unchartedsoftware/prism/util/json"
)

const (
	// MetaRoute represents the HTTP route for the resource.
	MetaRoute = "/meta"
)

// MetaHandler represents the HTTP route response handler.
func MetaHandler(w http.ResponseWriter, r *http.Request) {
	// set content type response header
	w.Header().Set("Content-Type", "application/json")
	// parse meta req from URL
	req, err := parseRequestJSON(r.Body)
	if err != nil {
		log.Warn(err)
		handleErr(w, err)
		return
	}
	// get pipeline id
	pipeline, ok := json.GetString(req, "pipeline")
	if !ok {
		// send error response
		log.Warn(err)
		handleErr(w, err)
		return
	}
	// ensure it's generated
	err = prism.GenerateMeta(pipeline, req)
	if err != nil {
		log.Warn(err)
		handleErr(w, err)
		return
	}
	// get meta data from store
	meta, err := prism.GetMetaFromStore(pipeline, req)
	if err != nil {
		log.Warn(err)
		handleErr(w, err)
		return
	}
	// send response
	w.WriteHeader(200)
	w.Write(meta)
}
