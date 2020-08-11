package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"stroe-server/service"
)

type OrderController struct {
	Ctx            iris.Context
	OrderService service.IOrderService
}

func (o *OrderController) Get() mvc.View {
	orderArr, err := o.OrderService.GatAll()
	if err != nil {
		o.Ctx.Application().Logger().Debug("查询订单信息失败", err, orderArr)
	}

	return mvc.View{
		Name: "order/view.html",
		Data: iris.Map{
			"order": orderArr,
		},
	}
}



