package repos

import (
	"context"

	"github.com/kumin/GolangMaster/restful/entities"
)

type ProductRepo interface {
	AddProduct(ctx context.Context, prod *entities.Product) (*entities.Product, error)
	GetProduct(ctx context.Context, id int64) (*entities.Product, error)
	ListProducts(ctx context.Context, page int, limit int) ([]*entities.Product, error)
}
