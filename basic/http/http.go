package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/pprof"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "index page")

}

func healthz(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "ok\n")
}

func getUserAgent(w http.ResponseWriter, r *http.Request) {
	ip  := r.Header.Get("User-Agent")
	io.WriteString(w, fmt.Sprintf("User-Agent=%s\n", ip))
}

func getIP(w http.ResponseWriter, r *http.Request) {
	ip  := r.Header.Get("REMOTE-ADDR")
	io.WriteString(w, fmt.Sprintf("ip=%s\n", ip))
}

func getForm(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "before parse form %v\n", r.Form)
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "parse form error %v\n", r.Form)
	}
	fmt.Fprintf(w, "before parse form %v\n", r.Form)
}

func getHeaders(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.Header {
		io.WriteString(w, fmt.Sprintf("%s=%s\n", k, v))
	}
}

func getUrl(w http.ResponseWriter, r *http.Request)  {
	data, _ := json.Marshal(r.URL)
	fmt.Fprintf(w, string(data))
}

func queryParams(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	fmt.Fprintf(w, "query : %+v\n", q)
}



func getBodyOnce(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err!= nil {
		fmt.Fprintf(w, "read body fail: %v", err)
		return
	}
	//将 []byte 转换为 string
	fmt.Fprintf(w, "body content: [%s]\n", string(body))
	//尝试再次读取，读不到任何东西，但不报错
	body, err = io.ReadAll(r.Body)
	fmt.Fprintf(w, "read content again: [%s] and data length %d \n", string(body), len(body))

}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/body/once", getBodyOnce)
	mux.HandleFunc("/headers", getHeaders)
	mux.HandleFunc("/form", getForm)
	mux.HandleFunc("/ip", getIP)
	mux.HandleFunc("/user-agent", getUserAgent)
	mux.HandleFunc("/url", getUrl)
	mux.HandleFunc("/query", queryParams)

	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(err)
	}
}