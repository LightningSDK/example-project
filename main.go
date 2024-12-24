package main

import (
	"github.com/lightningsdk/core"
	"github.com/lightningsdk/example/autogen"
)

func main() {
	app := core.NewApp()
	// this loads any included plugins
	app.RegisterPlugins(autogen.GetPlugins())
	app.Bootstrap()
	app.Run()
}
