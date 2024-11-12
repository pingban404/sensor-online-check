package config

import (
    "fmt"
    "gopkg.in/yaml.v3"
    "io/ioutil"
    //  "log"
)

type ElasticsearchConfig struct {
    URL      string `yaml:"url"`
    AuthToken string `yaml:"auth_token"`
}

type Config struct {
    Elasticsearch ElasticsearchConfig `yaml:"elasticsearch"`
}

func LoadConfig(filePath string) (*Config, error) {
    data, err := ioutil.ReadFile(filePath)
    if err != nil {
        return nil, fmt.Errorf("could not read config file: %v", err)
    }

    var config Config
    err = yaml.Unmarshal(data, &config)
    if err != nil {
        return nil, fmt.Errorf("could not unmarshal config: %v", err)
    }

    return &config, nil
}

