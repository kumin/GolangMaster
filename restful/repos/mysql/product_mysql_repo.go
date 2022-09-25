package mysql

import (
	"context"
	"log"

	"github.com/kumin/GolangMaster/restful/entities"
	"github.com/kumin/GolangMaster/restful/infras"
	"github.com/kumin/GolangMaster/restful/repos"
	"go.uber.org/zap"
)

var _ repos.ProductRepo = &ProductMysqlRepo{}

type ProductMysqlRepo struct {
	db     *infras.MysqlConnector
	logger *zap.Logger
}

func NewProductMysqlRepo(
	db *infras.MysqlConnector,
) *ProductMysqlRepo {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}
	return &ProductMysqlRepo{
		db:     db,
		logger: logger,
	}
}

func (p *ProductMysqlRepo) AddProduct(
	ctx context.Context,
	prod *entities.Product,
) (*entities.Product, error) {
	if err := p.db.Client.WithContext(ctx).Create(prod).Error; err != nil {
		return nil, err
	}
	return prod, nil
}

func (p *ProductMysqlRepo) GetProduct(
	ctx context.Context,
	id int64,
) (*entities.Product, error) {
	p.logger.Info("get product from mysql", zap.Int64("id", id))
	var prod *entities.Product
	if err := p.db.Client.WithContext(ctx).First(&prod, id).Error; err != nil {
		return nil, err
	}
	return prod, nil
}

func (p *ProductMysqlRepo) ListProducts(
	ctx context.Context,
	page int,
	limit int,
) ([]*entities.Product, error) {
	p.logger.Info("list products from mysql", zap.Int("page", 0), zap.Int("limit", limit))
	var prods []*entities.Product
	if err := p.db.Client.WithContext(ctx).
		Offset((page - 1) * limit).
		Limit(limit).
		Find(&prods).Error; err != nil {
		return nil, err
	}

	return prods, nil
}
