package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"stroe-server/service"
	"stroe-server/web/controllers"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")

	template := iris.HTML("./web/views", ".html").
		Layout("shared/layout.html").
		Reload(true)
	app.RegisterView(template)

	app.HandleDir("/assets", "./web/public")

	app.OnAnyErrorCode(func(ctx iris.Context) {
		ctx.ViewData("Message", ctx.Values().
			GetStringDefault("message", "The page you're looking for doesn't exist"))
		ctx.View("shared/error.html")
	})

	productService := service.NewProductService()
	product := mvc.New(app.Party("/product"))
	product.Register(productService)
	product.Handle(new(controllers.ProductController))

	orderService := service.NewOrderService()
	order := mvc.New(app.Party("/order"))
	order.Register(orderService)
	order.Handle(new(controllers.OrderController))

	app.Listen(":8080", iris.WithOptimizations)
}
