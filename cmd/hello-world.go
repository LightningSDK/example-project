package cmd

import (
	"fmt"
	"github.com/lightningsdk/core/model"
)

func GetHelloWorldCommand() model.Command {
	return model.Command{
		Help: "prints hello world",
		Function: func(app model.App) error {
			fmt.Println("Hello World")
			return nil
		},
	}
	//
	//return func() {
	//	app, err := service.BootstrapConfig("./config/lightning.yml", autogen.GetPlugins)
	//	if err != nil {
	//		fmt.Println(err)
	//		os.Exit(1)
	//	}
	//
	//	err = cmd.RunCommand(app)
	//	if err != nil {
	//		fmt.Println(err)
	//		os.Exit(1)
	//	}
	//}
}
