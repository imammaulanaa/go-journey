package config

import (
    "gopkg.in/yaml.v2"
    "os"
)

type Config struct {
    App struct {
        Name string `yaml:"name"`
        Env  string `yaml:"env"`
    } `yaml:"app"`
}

func LoadConfig(filePath string) (*Config, error) {
    file, err := os.ReadFile(filePath)
    if err != nil {
        return nil, err
    }

    var cfg Config
    err = yaml.Unmarshal(file, &cfg)
    return &cfg, err
}
