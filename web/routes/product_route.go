package routes

import (
	"github.com/kataras/iris/v12/mvc"
	"stroe-server/bootstrap"
	"stroe-server/service"
	"stroe-server/web/controllers"
)

func ProductHandler(b *bootstrap.Bootstrapper) {
	productService := service.NewProductService()
	product := mvc.New(b.Party("/product"))
	product.Register(productService)
	product.Handle(new(controllers.ProductController))
}
