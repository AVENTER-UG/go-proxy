package main

import (
	"os"

	"git.aventer.biz/AVENTER/util"
)

func init() {
	APIProxyBind = os.Getenv("API_PROXYBIND")
	APIProxyPort = os.Getenv("API_PROXYPORT")
	TargetURL = os.Getenv("TARGET_URL")
	SkipSSL = os.Getenv("SKIP_SSL")
	BlockAgent = os.Getenv("BLOCK_USERAGENT")
	LogLevel = util.Getenv("LOGLEVEL", "info")
}
