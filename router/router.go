package router

import (
	"minusblog/controllers"
	"minusblog/minus"
)

func init() {
	minus.Mapp.Handlers.Add("/", &controllers.MinusController{})
}
