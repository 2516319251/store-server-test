package service

import (
	"stroe-server/datamodels"
	"stroe-server/datasource"
	"stroe-server/repositories"
)

type IOrderService interface {
	Insert(datamodels.StoreOrder) (int64, error)
	Delete(int64) bool
	Update(datamodels.StoreOrder) error
	GetByID(int64) (datamodels.OrderInformation, error)
	GatAll() ([]datamodels.OrderInformation, error)
}

type OrderService struct {
	repository *repositories.OrderRepository
}

func NewOrderService() IOrderService {
	return &OrderService{
		repository: repositories.NewOrderRepository(datasource.InstanceMaster()),
	}
}

func (service *OrderService) Insert(order datamodels.StoreOrder) (int64, error) {
	return service.repository.Insert(order)
}

func (service *OrderService) Delete(orderID int64) bool {
	return service.repository.Delete(orderID)
}

func (service *OrderService) Update(order datamodels.StoreOrder) error {
	return service.repository.Update(order)
}

func (service *OrderService) GetByID(orderID int64) (datamodels.OrderInformation, error) {
	return service.repository.SelectByID(orderID)
}

func (service *OrderService) GatAll() ([]datamodels.OrderInformation, error) {
	return service.repository.SelectAll()
}
