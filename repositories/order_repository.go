package repositories

import (
	"github.com/go-xorm/xorm"
	"log"
	"stroe-server/datamodels"
)

type OrderRepository struct {
	engine *xorm.Engine
}

func NewOrderRepository(engine *xorm.Engine) *OrderRepository {
	return &OrderRepository{engine: engine}
}

func (repository OrderRepository) Insert(order datamodels.StoreOrder) (int64, error) {
	id, err := repository.engine.Insert(order)
	if err != nil {
		log.Println(err)
	}
	return id, err
}

func (repository OrderRepository) Delete(orderID int64) bool {
	data := datamodels.StoreOrder{}
	_, err := repository.engine.ID(orderID).Delete(&data)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func (repository OrderRepository) Update(order datamodels.StoreOrder) error {
	_, err := repository.engine.Id(order.ProductId).AllCols().Update(order)
	return err
}

func (repository OrderRepository) SelectByID(orderID int64) (datamodels.OrderInformation, error) {
	data := datamodels.OrderInformation{}
	err := repository.engine.Table("store_order").
		Select("store_order.order_id,store_user.nickname,store_product.product_name,store_order.status,store_order.create_time").
		Join("INNER", "store_user", "store_order.user_id = store_user.user_id").
		Join("INNER", "store_product", "store_order.product_id = store_product.product_id").
		Where("store_order.order_id = ?", orderID).
		Find(&data)
	if err != nil {
		log.Println(err)
	}
	return data, err
}

func (repository OrderRepository) SelectAll() ([]datamodels.OrderInformation, error) {
	var dataList []datamodels.OrderInformation
	err := repository.engine.Table("store_order").
		Select("store_order.order_id,store_user.nickname,store_product.product_name,store_order.status,store_order.create_time").
		Join("INNER", "store_user", "store_order.user_id = store_user.user_id").
		Join("INNER", "store_product", "store_order.product_id = store_product.product_id").
		Desc("store_order.order_id").
		Find(&dataList)
	if err != nil {
		log.Println(err)
	}
	return dataList, err
}
