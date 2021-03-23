package server

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/kataras/iris/v12/mvc"
	"github.com/qibobo/webgo/controller"
)

func CreateServer() *iris.Application {
	app := iris.New()
	app.Logger().SetLevel("debug")

	basic := app.Party("/basic")
	{
		basic.UseRouter(recover.New())

		mvc.Configure(basic, basicMVC)
	}
	return app
}

func basicMVC(app *mvc.Application) {
	app.Party("/demo").
		Handle(new(controller.SampleControler))
	app.Party("/demo2").
		Handle(new(controller.SampleControler2))
}
