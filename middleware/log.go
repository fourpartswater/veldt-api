package middleware

import (
	"bytes"
	"net/http"
	"strings"
	"time"

	"github.com/vova616/xxhash"
	"github.com/zenazn/goji/web"
	"github.com/zenazn/goji/web/mutil"

	"github.com/unchartedsoftware/prism/log"
)

// Logger is a middleware that logs each request. When standard output is a
// TTY, Logger will print in color, otherwise it will print in black and white.
func Log(c *web.C, h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		lw := mutil.WrapWriter(w)
		t1 := time.Now()
		h.ServeHTTP(lw, r)
		if lw.Status() == 0 {
			lw.WriteHeader(http.StatusOK)
		}
		t2 := time.Now()
		printResp(r, lw, t2.Sub(t1))
	}
	return http.HandlerFunc(fn)
}

func printResp(r *http.Request, w mutil.WriterProxy, dt time.Duration) {
	var buf bytes.Buffer
	cW(&buf, bMagenta, "%s ", r.Method)
	urlsplit := strings.Split(r.URL.String(), "?")
	url := urlsplit[0]
	cs := strings.Split(url, "/")
	if len(cs) == 2 && cs[0] == "" && cs[1] == "" {
		cW(&buf, bBlack, "/")
	} else {
		for _, c := range cs {
			if c != "" {
				cW(&buf, bBlack, "/")
				hash := xxhash.Checksum32([]byte(c))
				cW(&buf, randColor(hash), c)
			}
		}
	}
	if len(urlsplit) > 1 {
		cW(&buf, bBlack, "?")
		hash := xxhash.Checksum32([]byte(urlsplit[1]))
		cW(&buf, randColor(hash), "%#x ", hash)
	} else {
		buf.WriteString(" ")
	}
	status := w.Status()
	if status < 200 {
		cW(&buf, bBlue, "%03d", status)
	} else if status < 300 {
		cW(&buf, bGreen, "%03d", status)
	} else if status < 400 {
		cW(&buf, bCyan, "%03d", status)
	} else if status < 500 {
		cW(&buf, bYellow, "%03d", status)
	} else {
		cW(&buf, bRed, "%03d", status)
	}
	buf.WriteString(" in ")
	if dt < 500*time.Millisecond {
		cW(&buf, nGreen, "%s", dt)
	} else if dt < 5*time.Second {
		cW(&buf, nYellow, "%s", dt)
	} else {
		cW(&buf, nRed, "%s", dt)
	}
	buf.WriteString(" to ")
	buf.WriteString(r.RemoteAddr)
	log.Debug(buf.String())
}
