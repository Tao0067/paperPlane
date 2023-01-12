package config

import (
	"encoding/json"
	"log"
	"os"
	"sync"
)

type Database struct {
	User         string
	Password     string
	LocalAddress string
	Database     string
}

type Configuration struct {
	Db Database
}

var config *Configuration
var once sync.Once //保证只执行一次，配置存再内存中

func LoadConfig() *Configuration {
	once.Do(func() {
		file, err := os.Open("config.json")
		if err != nil {
			log.Fatalln("Cannot open config file", err)
		}

		decoder := json.NewDecoder(file)
		config = &Configuration{}
		err = decoder.Decode(config)

		if err != nil {
			log.Fatalln("Cannot get configuration from file", err)
		}
	})

	return config
}
