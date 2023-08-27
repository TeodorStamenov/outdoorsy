package util

import "github.com/spf13/viper"

type DbConfig struct {
	Addr string `mapstructure:"address"`
	Port string `mapstructure:"port"`
	User string `mapstructure:"user"`
	Pass string `mapstructure:"password"`
	Name string `mapstructure:"dbname"`
}

type Config struct {
	Db DbConfig `mapstructure:"db"`
}

var vp *viper.Viper

func LoadConfig() (Config, error) {
	vp = viper.New()

	var config Config

	vp.SetConfigName("config")
	vp.SetConfigType("json")
	vp.AddConfigPath("..")
	vp.AddConfigPath(".")

	err := vp.ReadInConfig()
	if err != nil {
		return Config{}, err
	}

	err = vp.Unmarshal(&config)
	if err != nil {
		return Config{}, err
	}

	return config, err
}
