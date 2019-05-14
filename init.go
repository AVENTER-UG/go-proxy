package main

import "os"

func init() {
	API_PROXYBIND = os.Getenv("API_PROXYBIND")
	API_PROXYPORT = os.Getenv("API_PROXYPORT")
	TARGET_URL = os.Getenv("TARGET_URL")
	SKIP_SSL = os.Getenv("SKIP_SSL")
}
