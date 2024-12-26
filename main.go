package main

import (
	"github.com/lightningsdk/core"
	"github.com/lightningsdk/example/autogen"
)

func main() {
	core.Run("./config/lightning.yml", autogen.GetPlugins())
}
