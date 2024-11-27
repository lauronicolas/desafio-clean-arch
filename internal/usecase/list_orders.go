package usecase

import (
	"github.com/lauronicolas/curso-go/Clean-Arch-APP/internal/entity"
)

type ListOrderOutputDTO []OrderOutputDTO

type ListOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewListOrdersUseCase(
	OrderRepository entity.OrderRepositoryInterface,
) *ListOrdersUseCase {
	return &ListOrdersUseCase{
		OrderRepository: OrderRepository,
	}
}

func (c *ListOrdersUseCase) Execute() (ListOrderOutputDTO, error) {
	orders, err := c.OrderRepository.List()
	if err != nil {
		return nil, err
	}

	var listOutputDTO ListOrderOutputDTO

	for _, order := range orders {
		outputDto := &OrderOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.Price + order.Tax,
		}
		listOutputDTO = append(listOutputDTO, *outputDto)
	}

	return listOutputDTO, nil
}
