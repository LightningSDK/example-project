package autogen

import (
	"github.com/lightningsdk/core/model"
	"github.com/lightningsdk/example/local"
)

func GetPlugins() map[string]model.Plugin {
	plugins := map[string]model.Plugin{}
	plugins["local"] = local.NewPlugin()
	//plugins["github.com/lightningsdk/ui"] = ui.NewPlugin(app)
	//plugins["github.com/lightningsdk/blog"] = blog.NewPlugin(app)
	//plugins["github.com/lightningsdk/db"] = db.NewPlugin(app)
	return plugins
}
