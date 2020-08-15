package routes

import (
	"github.com/kataras/iris/v12/mvc"
	"stroe-server/bootstrap"
	"stroe-server/service"
	"stroe-server/web/controllers"
)

func OrderHandler(b *bootstrap.Bootstrapper) {
	orderService := service.NewOrderService()
	order := mvc.New(b.Party("/order"))
	order.Register(orderService)
	order.Handle(new(controllers.OrderController))
}
