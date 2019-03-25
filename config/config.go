package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Repository struct {
	Type   string
	Name   string
	Url    string
	Path   string
	Branch string
	Script string
	Secret string
}

var Config struct {
	Host         string
	Port         string
	Git          string
	Log          string
	Repositories [] *Repository
}

func Init(file string) error {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(b, &Config)
	if err != nil {
		return err
	}
	return nil
}

func FindRepositoryConfig(t string, url string) *Repository {
	for k, v := range Config.Repositories {
		if v.Type == t && v.Url == url {
			return Config.Repositories[k]
		}
	}
	return nil
}
