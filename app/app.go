package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
)

type App struct {
	Config *Config
}

// Запуск вебсервера
func (app *App) RunServer() {
	router := mux.NewRouter()
	router.HandleFunc("/", app.Update)
	log.Fatal(http.ListenAndServe(app.Config.Host+":"+app.Config.Port, router))
}

// лонирование всех репозиториев из конфига.
func (app *App) CloneAll() {
	for _, project := range app.Config.Projects {
		stat, _ := os.Stat(app.Config.RepositoryDir + "/" + project + "/config")
		if stat == nil {
			app.Clone(project)
		}
	}
}

// клонирование репозитория
func (app *App) Clone(project string) {
	log.Print("clone " + project)
	path := app.Config.RepositoryDir + "/" + project
	url := fmt.Sprintf(
		"https://%s:%s@%s/%s.git",
		app.Config.GitlabUser,
		app.Config.GitlabToken,
		app.Config.GitlabHost,
		project,
	)
	cmd := exec.Command("git", "clone", "--mirror", url, path)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Print(err)
	}
}

// обработка запроса (обновление из репозитория)
func (app *App) Update(writer http.ResponseWriter, request *http.Request) {
	bodyBytes, err := ioutil.ReadAll(request.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)
	project := gjson.Get(bodyString, "project.path_with_namespace").String()
	path := app.Config.RepositoryDir + "/" + project
	_ = os.Chdir(path)
	log.Print("fetch " + project)
	cmd := exec.Command("git", "fetch", "-q", "--all", "-p")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		log.Print(err)
	}
}
