package main

import (
	"fmt"
	"github.com/lightningsdk/core"
	"github.com/lightningsdk/example/autogen"
	"os"
)

func main() {
	app, err := core.BootstrapConfig("./config/lightning.yml", autogen.GetModules)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = core.StartService(app)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
