package autogen

import (
	"github.com/lightningsdk/core"
	"github.com/lightningsdk/core/model"
	"github.com/lightningsdk/example/local"
)

// TODO: does this even need to autogenerate? should it just be manual? the list is not pulled from anywhere else
func GetPlugins() map[string]model.Plugin {
	plugins := map[string]model.Plugin{}
	plugins["github.com/lightningsdk/example/local"] = local.NewPlugin()
	plugins["github.com/lightningsdk/core"] = core.NewPlugin()
	return plugins
}
