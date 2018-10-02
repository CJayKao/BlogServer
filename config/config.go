package config

import (
	"errors"
	"flag"
	"io/ioutil"
	"log"
	"sync"
	yaml "gopkg.in/yaml.v2"
)

var ServerVersion = "1.0"

var (
	Config struct {
		Version  string `yaml:"version"`
		MysqlURL string `yaml:"mysql_url"`
		MongoURL string `yaml:"Mongo_url"`
	}
	once sync.Once
)

func init() {

	once.Do(func() {
		var filepath string
		flag.StringVar(&filepath, "conf", "config.yml", "config file path")
		flag.Parse()
		data, err := ioutil.ReadFile(filepath)
		if err != nil {
			log.Fatalln(err)
		}

		err = yaml.Unmarshal(data, &Config)
		if err != nil {
			log.Fatalln(err)
		}

		// if Config.Version != CurrentConfigVersion {
		// 	log.Fatalln(errors.New("config.yml 版本已更新"))
		// }
	
	})
}
