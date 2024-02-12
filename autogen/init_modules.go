package autogen

import "github.com/lightningsdk/core"
import (
	"github.com/lightningsdk/blog"
	"github.com/lightningsdk/db"
)

func GetModules(app *core.App) map[string]core.Module {
	modules := map[string]core.Module{}
	modules["github.com/lightningsdk/blog"] = blog.NewModule(app)
	modules["github.com/lightningsdk/db"] = db.NewModule(app)
	return modules
}
