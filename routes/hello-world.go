package routes

import (
	"github.com/lightningsdk/core/model"
	"net/http"
)

func GetHelloWorldRoutes() []model.Route {
	return []model.Route{{
		Path:   "hello-world",
		Weight: 0,
		Handler: func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write([]byte("Hello World!"))
		},
	}}
}
