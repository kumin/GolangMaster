package services

import (
	"context"

	"github.com/kumin/GolangMaster/restful/entities"
)

var prods = []*entities.Product{
	{
		ID:   123456,
		Name: "Iphone 13",
		Properties: &entities.Properties{
			Price:    30000000.04,
			Category: "Smart Phone",
		},
	},
	{
		ID:   123458,
		Name: "Tivi Sony",
		Properties: &entities.Properties{
			Price:    15000000.04,
			Category: "Tivi",
		},
	},
	{
		ID:   123457,
		Name: "Tu Lanh Panasonic",
		Properties: &entities.Properties{
			Price:    20000000.08,
			Category: "Tu Lanh",
		},
	},
}

type ProductCtlService struct {
}

func NewProductCtlHandler() *ProductCtlService {
	return &ProductCtlService{}
}

func (p *ProductCtlService) ListProducts(
	ctx context.Context,
	page int,
	limit int) (
	[]*entities.Product, error) {
	return prods, nil
}

func (p *ProductCtlService) GetProduct(
	ctx context.Context,
	id int64,
) (*entities.Product, error) {
	for _, p := range prods {
		if p.ID == id {
			return p, nil
		}
	}

	return nil, nil
}
