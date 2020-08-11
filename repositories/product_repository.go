package repositories

import (
	"github.com/go-xorm/xorm"
	"log"
	"stroe-server/datamodels"
)

type ProductRepository struct {
	engine *xorm.Engine
}

func NewProductRepository(engine *xorm.Engine) *ProductRepository {
	return &ProductRepository{
		engine: engine,
	}
}

func (repository ProductRepository) Insert(product datamodels.StoreProduct) (int64, error) {
	id, err := repository.engine.Insert(product)
	if err != nil {
		log.Println(err)
	}
	return id, err
}

func (repository ProductRepository) Delete(productID int64) bool {
	data := datamodels.StoreProduct{}
	_, err := repository.engine.ID(productID).Delete(&data)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func (repository ProductRepository) Update(product datamodels.StoreProduct) error {
	_, err := repository.engine.Id(product.ProductId).AllCols().Update(product)
	return err
}

func (repository ProductRepository) SelectByID(productID int64) (datamodels.StoreProduct, error) {
	data := datamodels.StoreProduct{}
	success, err := repository.engine.ID(productID).Get(&data)
	if err != nil && !success {
		log.Println(err, success)
		return data, err
	}
	return data, nil
}

func (repository ProductRepository) SelectAll() ([]datamodels.StoreProduct, error) {
	var dataList []datamodels.StoreProduct
	err := repository.engine.Desc("product_id").Find(&dataList)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return dataList, nil
}
