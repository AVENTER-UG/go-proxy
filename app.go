package main

import (
	"crypto/tls"
	"log"
	"net/http"
	"net/http/httputil"
	_ "net/http/pprof"
	"net/url"
)

var API_PROXYPORT string
var API_PROXYBIND string
var TARGET_URL string
var SKIP_SSL string

var MinVersion string

var srv http.Server

type handle struct {
	reverseProxy string
}

func (this *handle) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	log.Println(this.reverseProxy + " " + r.Method + " " + r.URL.String() + " " + r.Proto + " " + r.UserAgent())
	remote, err := url.Parse(this.reverseProxy)
	if err != nil {
		log.Fatalln(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(remote)

	if SKIP_SSL == "true" {
		proxy.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	}

	r.Host = remote.Host

	proxy.ServeHTTP(w, r)
}

func main() {
	log.Println("GO-PROXY build"+MinVersion, API_PROXYBIND, API_PROXYPORT, TARGET_URL)
	srv.Handler = &handle{reverseProxy: TARGET_URL}
	srv.Addr = API_PROXYBIND + ":" + API_PROXYPORT
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalln("ListenAndServe: ", err)
	}
}
