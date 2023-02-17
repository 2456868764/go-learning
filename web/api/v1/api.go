package v1

import (
	"fmt"
	"io"
	"net/http"
)

func GetUserAgent(w http.ResponseWriter, r *http.Request) {
	ip := r.Header.Get("User-Agent")
	io.WriteString(w, fmt.Sprintf("User-Agent=%s\n", ip))
}

func GetIP(w http.ResponseWriter, r *http.Request) {
	ip := r.Header.Get("REMOTE-ADDR")
	io.WriteString(w, fmt.Sprintf("ip=%s\n", ip))
}

func GetHeaders(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.Header {
		io.WriteString(w, fmt.Sprintf("%s=%s\n", k, v))
	}
}
