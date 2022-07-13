package config

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"sync"
)

const pathConfigFile string = "./config.json"

// Config структура, которая содержит необходимую конфигурацию

type Config struct {
	Database struct {
		Host     string `json:"host"`
		Port     string `json:"port"`
		Name     string `json:"name"`
		Username string `json:"username"`
		Password string `json:"password"`
	} `json:"database" ini:"database"`
}

var (
	config     = &Config{}
	configOnce = sync.Once{}
)

func GetConfig() *Config {
	configOnce.Do(func() {
		file, err := os.Open(pathConfigFile)
		if err != nil {
			log.Fatal(err)
		}
		config.load(file)
	})

	return config
}

func (c *Config) load(r io.Reader) {
	if err := json.NewDecoder(r).Decode(&c); err != nil {
		log.Panic("Read config file: ", err)
	}
}
