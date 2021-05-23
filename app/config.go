package app

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Config struct {
	GitlabToken   string   `json:"gitlab_token"`
	RepositoryDir string   `json:"repository_dir"`
	Projects      []string `json:"projects"`
	Host          string   `json:"server_host"`
	Port          string   `json:"server_port"`
	GitlabUser    string   `json:"gitlab_user"`
	GitlabHost    string   `json:"gitlab_host"`
}

func (config *Config) Read(path string) *Config {
	jsonFile, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, config)
	if err != nil {
		panic(err)
	}
	return config
}
