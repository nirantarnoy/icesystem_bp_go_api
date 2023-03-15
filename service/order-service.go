package service

import (
	"log"

	"github.com/mashingan/smapping"
	"tarlek.com/icesystem/dto"
	"tarlek.com/icesystem/entity"
	"tarlek.com/icesystem/repository"
)

type OrderService interface {
	CreateOrder(oderDto dto.OrderCreateDto) entity.OrderCreate
	GetLastNo() string
}
type orderService struct {
	orderRepository repository.OrderRepository
}

// GetLastNo implements OrderService
func (db *orderService) GetLastNo() string {
	res := db.orderRepository.GetLastNo(1, 2, 123, "VP10")
	return res
}

func (db *orderService) CreateOrder(orderDto dto.OrderCreateDto) entity.OrderCreate {
	order := entity.OrderCreate{}
	err := smapping.FillStruct(&order, smapping.MapFields(&orderDto))
	if err != nil {
		log.Fatalf("Fail to mapping %v", err)
	}
	res := db.orderRepository.CreateOrder(order)
	return res
}

func NewOrderService(repo repository.OrderRepository) OrderService {
	return &orderService{orderRepository: repo}
}
