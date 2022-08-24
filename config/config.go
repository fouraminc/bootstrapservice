package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	App   `yaml:"app"`
	HTTP  `yaml:"http"`
	Log   `yaml:"logger"`
	MONGO `yaml:"mongo"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}

	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath("config")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Error reading config file:", err)
		return nil, err
	}

	fmt.Println("Application is:", viper.GetString("app.name"))
	err = viper.Unmarshal(&cfg)
	if err != nil {
		fmt.Println("Error unmarshalling config:", err)
		return nil, err
	}
	fmt.Println(cfg)

	return cfg, nil
}

type App struct {
	Name    string
	Version string
}

type HTTP struct {
	Port string
}

type Log struct {
	Level string
}

type MONGO struct {
	HOST string
	PORT string
	DB   string
}
