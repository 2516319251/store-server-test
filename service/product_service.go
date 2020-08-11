package service

import (
	"stroe-server/datamodels"
	"stroe-server/datasource"
	"stroe-server/repositories"
)

type IProductService interface {
	Insert(datamodels.StoreProduct) (int64, error)
	Delete(int64) bool
	Update(datamodels.StoreProduct) error
	GetByID(int64) (datamodels.StoreProduct, error)
	GatAll() ([]datamodels.StoreProduct, error)
}

type ProductService struct {
	repository *repositories.ProductRepository
}

func NewProductService() IProductService {
	return &ProductService{
		repository: repositories.NewProductRepository(datasource.InstanceMaster()),
	}
}

func (service *ProductService) Insert(product datamodels.StoreProduct) (int64, error) {
	return service.repository.Insert(product)
}

func (service *ProductService) Delete(productID int64) bool {
	return service.repository.Delete(productID)
}

func (service *ProductService) Update(product datamodels.StoreProduct) error {
	return service.repository.Update(product)
}

func (service *ProductService) GetByID(productID int64) (datamodels.StoreProduct, error) {
	return service.repository.SelectByID(productID)
}

func (service *ProductService) GatAll() ([]datamodels.StoreProduct, error) {
	return service.repository.SelectAll()
}
