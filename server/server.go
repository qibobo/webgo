package server

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/kataras/iris/v12/mvc"
	"github.com/qibobo/webgo/controller"
	"github.com/qibobo/webgo/db"
	"go.uber.org/zap"
)

func CreateServer(logger *zap.Logger, demoDB db.DemoDB) *iris.Application {
	app := iris.New()
	app.Logger().SetLevel("debug")

	basic := app.Party("/basic")
	{
		basic.UseRouter(recover.New())

		mvc.Configure(basic, func(app *mvc.Application) {
			app.Party("/demo").
				Handle(controller.NewSampleController(logger, demoDB))
			app.Party("/demo2").
				Handle(new(controller.SampleControler2))
		})
	}
	return app
}

func basicMVC(app *mvc.Application) {

}
