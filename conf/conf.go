package conf

import (
	"bytes"
	"gopkg.in/yaml.v1"
	"io/ioutil"
	"log"
	"os"
)

type configuration struct {
	Listen       string            `yaml:"listen"`
	Host         string            `yaml:"host"`
	RedirectUrls map[string]string `yaml:"redirect_urls"`
	AppId        string            `yaml:"appid"`
	KeyParam     string            `yaml:"key_param"`
}

var Conf configuration

func ParseConfig(configFile string) {
	if fileInfo, err := os.Stat(configFile); err != nil {
		if os.IsNotExist(err) {
			log.Panicf("configuration file %v does not exist.", configFile)
		} else {
			log.Panicf("configuration file %v can not be stated. %v", configFile, err)
		}
	} else {
		if fileInfo.IsDir() {
			log.Panicf("%v is a directory name", configFile)
		}
	}

	content, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Panicf("read configuration file error. %v", err)
	}
	content = bytes.TrimSpace(content)

	if err = yaml.Unmarshal(content, &Conf); err != nil {
		log.Panicf("unmarshal yaml object error. %v", err)
		return
	}
}
