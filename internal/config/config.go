package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"sync"
)

type Config struct {
	Hostname string `json:"hostname"`
	Protocol string `json:"protocol"`
}

var (
	Cfg  *Config
	once sync.Once
)

func GetConfig(path string) {
	once.Do(func() {
		Cfg = &Config{}
		if err := cleanenv.ReadConfig(path, Cfg); err != nil {
			desc, _ := cleanenv.GetDescription(Cfg, nil)
			log.Fatalln(desc, err)
		}
	})
}

func Init(path string) {
	GetConfig(path)
}
