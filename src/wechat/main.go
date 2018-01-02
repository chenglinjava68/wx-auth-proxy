package wechat

import (
	"conf"
	"flag"
	"web"
)

func main() {
	cfgFile := flag.String("c", "config.yml", "configuration file")

	// parse config
	conf.ParseConfig(*cfgFile)

	web.Start()
}
