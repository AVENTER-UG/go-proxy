package main

import (
	"os"

	"git.aventer.biz/AVENTER/util"
)

func init() {
	APIProxyBind = util.Getenv("API_PROXYBIND", "0.0.0.0")
	APIProxyPort = util.Getenv("API_PROXYPORT", "10777")
	TargetURL = os.Getenv("TARGET_URL")
	SkipSSL = util.Getenv("SKIP_SSL", "false")
	BlockAgent = os.Getenv("BLOCK_USERAGENT")
	LogLevel = util.Getenv("LOGLEVEL", "info")
}
