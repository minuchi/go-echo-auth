package lib

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
)

type Config struct {
	Database DatabaseConfig `json:"database" yaml:"database"`
	Keys     KeysConfig     `json:"keys" yaml:"keys"`
}

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DbName   string `yaml:"dbname"`
	Port     string `yaml:"port"`
	Sslmode  string `yaml:"sslmode"`
	Timezone string `yaml:"timezone"`
}

type KeysConfig struct {
	AccessTokenSecret  string `yaml:"accessTokenSecret"`
	RefreshTokenSecret string `yaml:"refreshTokenSecret"`
}

var config Config

func init() {
	configFile, err := os.Open("config.yml")
	if err != nil {
		fmt.Println(err)
	}
	byteValue, _ := ioutil.ReadAll(configFile)

	err = yaml.Unmarshal(byteValue, &config)
	if err != nil {
		fmt.Println(err)
	}

	defer func(configFile *os.File) {
		err := configFile.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(configFile)
}

func LoadConfig() *Config {
	return &config
}
