package app

import (
	"io/ioutil"
	"sync"

	"gopkg.in/yaml.v2"
)

var (
	once   sync.Once
	kernel *App
)

func Kernel() *App {
	once.Do(func() {
		kernel = &App{}
	})
	return kernel
}

type App struct {
	config *config
}

func (a *App) SetConfig(path string) {
	data, _ := ioutil.ReadFile(path)
	if err := yaml.Unmarshal(data, &a.config); err != nil {
		panic(err)
	}
}

func (a *App) GetConfig() *config {
	return a.config
}
