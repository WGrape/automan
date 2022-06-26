package config

import (
	"automan/pkg/model"
	"fmt"
	"github.com/BurntSushi/toml"
)

func New() (*model.Config, error) {
	var (
		config model.Config
		path   = "config/config.toml"
	)
	_, err := toml.DecodeFile(path, &config)
	if err != nil {
		fmt.Println("failed to decode config file:", err)
		return &config, err
	}
	return &config, nil
}
