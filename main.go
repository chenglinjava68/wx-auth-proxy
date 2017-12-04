package main

import (
	"wx-auth-proxy/conf"
	"wx-auth-proxy/web"
	"flag"
)


func main() {
	cfgFile := flag.String("c", "config.yml", "configuration file")

	// parse config
	conf.ParseConfig(*cfgFile)

	web.Start()
}

