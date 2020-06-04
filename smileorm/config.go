package smileorm

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	DB DBConfig
}

type DBConfig struct {
	Host         string
	DB           string
	Password     string
	Username     string
	Protocol     string
	Port         int
	MaxOpenConns int
	MaxIdleConns int
}

func GetConfig() (*Config, error) {
	viper.AddConfigPath(getConfigDir())
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("viper read config failed")
	}

	var cfg Config
	err = viper.Unmarshal(&cfg)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("unmarshal config failed")
	}
	return &cfg, nil
}

func getConfigDir() string {
	dir, err := filepath.Abs(filepath.Dir("../"))
	if err != nil {
		panic("read config dir error: " + err.Error())
	}
	dir = strings.Replace(dir, "\\", "/", -1)
	return strings.Join([]string{dir, "config"}, "/")
}
