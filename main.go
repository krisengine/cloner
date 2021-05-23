package main

import (
	"./app"
	"log"
	"os"
)

const DirectoryPermission = 0755

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	config := app.Config{}
	config.Read("config.json")
	initialize(config.RepositoryDir, config.Projects)
	application := app.App{Config: &config}
	application.CloneAll()
	application.RunServer()
}

// создание всех необходимых директорий
func initialize(ropositoryDir string, projects []string) {
	err := os.MkdirAll(ropositoryDir, DirectoryPermission)
	if err != nil {
		panic(err)
	}
	for _, project := range projects {
		err := os.MkdirAll(ropositoryDir+"/"+project, DirectoryPermission)
		if err != nil {
			panic(err)
		}
	}
}
