package main

import (
	"crypto/tls"
	"net/http"
	"net/http/httputil"
	"net/url"
	"regexp"

	"github.com/AVENTER-UG/util"
	"github.com/sirupsen/logrus"
)

// APIProxyPort is the Port where the service are listening
var APIProxyPort string

// APIProxyBind is the IP where the service are listening
var APIProxyBind string

// TargetURL is the Url to where the proxy will forward all access
var TargetURL string

// SkipSSL will disable the ssl check
var SkipSSL string

// BlockAgent include a regularexpression to denied access of specified user agents
var BlockAgent string

// BlockURL include a regularexpression to denied access of specified url
var BlockURL string

// LogLevel defines the loglevel
var LogLevel string

// MinVersion is just the version of this app, its set dynamic during compiling
var MinVersion string

var srv http.Server
var reAgent *regexp.Regexp
var reURL *regexp.Regexp

type handle struct {
	reverseProxy string
}

func (this *handle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if BlockAgent != "" {
		fi := reAgent.Find([]byte(r.UserAgent()))
		if len(fi) > 0 {
			logrus.Debug("Blocked: ", r.UserAgent())
			return
		}
	}

	if BlockURL != "" {
		fi := reURL.Find([]byte(r.URL.String()))
		if len(fi) > 0 {
			logrus.Debug("Blocked: ", r.URL.String())
			return
		}
	}

	remote, err := url.Parse(this.reverseProxy)
	if err != nil {
		logrus.Error(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(remote)

	if SkipSSL == "true" {
		proxy.Transport = &http.Transport{
			// #nosec G402
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	}

	r.Host = remote.Host
	r.Header.Set("X-Forwarded-Host", r.Header.Get("Host"))

	proxy.ServeHTTP(w, r)

	logrus.Info(this.reverseProxy + " " + r.Method + " " + r.URL.String() + " " + r.Proto + " " + r.UserAgent())
}

func listenLoop() {
	srv.Handler = &handle{reverseProxy: TargetURL}
	srv.Addr = APIProxyBind + ":" + APIProxyPort
	if err := srv.ListenAndServe(); err != nil {
		logrus.Fatal("ListenAndServe: ", err)
		srv.Close()
		listenLoop()
	}
}

func main() {
	util.SetLogging(LogLevel, false, "go-proxy")
	logrus.Infoln("GO-PROXY build"+MinVersion, APIProxyBind, APIProxyPort, TargetURL, SkipSSL)

	if BlockAgent != "" {
		logrus.Infoln("Block following Agents: ", BlockAgent)
		var err error
		reAgent, err = regexp.Compile(BlockAgent)

		if err != nil {
			logrus.Error(err)
		}
	}

	if BlockURL != "" {
		logrus.Infoln("Block following Url: ", BlockURL)
		var err error
		reURL, err = regexp.Compile(BlockURL)

		if err != nil {
			logrus.Error(err)
		}
	}

	listenLoop()
}
