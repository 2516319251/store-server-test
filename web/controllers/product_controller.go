package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"strconv"
	"stroe-server/datamodels"
	"stroe-server/service"
	"time"
)

type ProductController struct {
	Ctx            iris.Context
	ProductService service.IProductService
}

func (p *ProductController) GetAll() mvc.View {
	productArr, _ := p.ProductService.GatAll()
	return mvc.View{
		Name: "product/view.html",
		Data: iris.Map{
			"productArray": productArr,
		},
	}
}

func (p *ProductController) GetAdd() mvc.View {
	return mvc.View{
		Name: "product/add.html",
	}
}

func (p *ProductController) PostAdd() {
	product := datamodels.StoreProduct{
		ProductName:  p.Ctx.FormValue("ProductName"),
		ProductNum:   p.Ctx.PostValueInt64Default("ProductNum", 0),
		ProductImage: p.Ctx.FormValue("ProductImage"),
		ProductUrl:   p.Ctx.FormValue("ProductUrl"),
		CreateTime:   time.Now(),
	}

	_, err := p.ProductService.Insert(product)
	if err != nil {
		p.Ctx.Application().Logger().Debug(err)
	}
	p.Ctx.Redirect("/product/all")
}

func (p *ProductController) GetDelete() {
	idString := p.Ctx.URLParam("id")
	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		p.Ctx.Application().Logger().Debug(err)
	}

	isDelete := p.ProductService.Delete(id)
	if isDelete {
		p.Ctx.Application().Logger().Debug("删除商品成功，ID为：" + idString)
	} else {
		p.Ctx.Application().Logger().Debug("删除商品失败，ID为：" + idString)
	}
	p.Ctx.Redirect("/product/all")
}

func (p *ProductController) GetManager() mvc.View {
	idString := p.Ctx.URLParam("id")
	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		p.Ctx.Application().Logger().Debug(err)
	}

	product, err := p.ProductService.GetByID(id)
	if err != nil {
		p.Ctx.Application().Logger().Debug(err)
	}

	return mvc.View{
		Name: "product/manager.html",
		Data: iris.Map{
			"product": product,
		},
	}
}

func (p *ProductController) PostUpdate() {
	product := datamodels.StoreProduct{
		ProductId:    p.Ctx.PostValueInt64Default("ProductId", 0),
		ProductName:  p.Ctx.FormValue("ProductName"),
		ProductNum:   p.Ctx.PostValueInt64Default("ProductNum", 0),
		ProductImage: p.Ctx.FormValue("ProductImage"),
		ProductUrl:   p.Ctx.FormValue("ProductUrl"),
		CreateTime:   time.Now(),
	}

	err := p.ProductService.Update(product)
	if err != nil {
		p.Ctx.Application().Logger().Debug(err)
	}
	p.Ctx.Redirect("/product/all")
}
