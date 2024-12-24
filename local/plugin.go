package local

import (
	"github.com/lightningsdk/core/model"
	"github.com/lightningsdk/example/cmd"
	"github.com/lightningsdk/example/routes"
)

type local struct {
	model.DefaultPlugin
}

func NewPlugin() model.Plugin {
	return &local{}
}

func (l *local) GetRoutes() []model.Route {
	return routes.GetHelloWorldRoutes()
}

func (l *local) GetCommands() map[string]model.Command {
	return map[string]model.Command{
		"hello-world": cmd.GetHelloWorldCommand(),
	}
}
