package autogen

import (
	"github.com/lightningsdk/blog"
	"github.com/lightningsdk/core"
	"github.com/lightningsdk/db"
	"github.com/lightningsdk/ui"
)

func GetModules(app *core.App) map[string]core.Module {
	modules := map[string]core.Module{}
	modules["github.com/lightningsdk/blog"] = blog.NewModule(app)
	modules["github.com/lightningsdk/db"] = db.NewModule(app)
	modules["github.com/lightningsdk/ui"] = ui.NewModule(app)
	return modules
}
