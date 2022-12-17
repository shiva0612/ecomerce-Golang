package config

import (
	"log"
	"path/filepath"
	"time"

	"github.com/spf13/viper"
)

var (
	Prjconfig = new(PrjConfig)
)

type PrjConfig struct {
	Port string `mapstructure:"port",json:"port"`
	Jwt  struct {
		TokenTime   time.Duration `mapstructure:"token_time",json:"token_time"`
		RefreshTime time.Duration `mapstructure:"refresh_time",json:"refresh_time"`
		TokenKey    string        `mapstructure:"token_key",json:"token_key"`
	} `mapstructure:"jwt",json:"jwt"`
	Mongo struct {
		User   string `mapstructure:"user",json:"user"`
		Psw    string `mapstructure:"psw",json:"psw"`
		Auth   string `mapstructure:"auth",json:"auth"`
		Server string `mapstructure:"server",json:"server"`
		Port   string `mapstructure:"port",json:"port"`
	} `mapstructure:"mongo",json:"mongo"`
}

func LoadConfig(path string) error {
	return ParseConfigFile(path, Prjconfig)
}

func ParseConfigFile(path string, cls interface{}) error {
	viper.AddConfigPath(filepath.Dir(path))
	viper.SetConfigName(filepath.Base(path))
	viper.SetConfigType(filepath.Ext(path)[1:])

	err := viper.ReadInConfig()
	if err != nil {
		log.Println("error while reading config file: ", path, err.Error())
		return err
	}

	err = viper.Unmarshal(cls)
	if err != nil {
		log.Println("error while unmarshalling: ", err.Error())
		return err
	}
	log.Println("config loaded:----------------------------------------------------------------")
	log.Printf("%+v", Prjconfig)

	return nil

}
