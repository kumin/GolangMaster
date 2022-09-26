package services

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/kumin/GolangMaster/restful/entities"
	"github.com/kumin/GolangMaster/restful/repos"
	"github.com/kumin/GolangMaster/restful/repos/mysql"
)

type ProductCtlServices struct {
	repo repos.ProductRepo
}

func NewProductCtlServices(
	repo *mysql.ProductMysqlRepo,
) *ProductCtlServices {
	return &ProductCtlServices{
		repo: repo,
	}
}

func (p *ProductCtlServices) AddProduct(
	ctx context.Context,
	req *http.Request,
) (*entities.Product, error) {
	defer req.Body.Close()
	body, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}
	var prod *entities.Product
	if err := json.Unmarshal(body, &prod); err != nil {
		return nil, err
	}

	return p.repo.AddProduct(ctx, prod)
}

func (p *ProductCtlServices) GetProduct(
	ctx context.Context,
	req *http.Request,
) (*entities.Product, error) {
	id, err := strconv.ParseInt(req.URL.Query().Get("id"), 10, 64)
	if err != nil {
		return nil, entities.ParamInvalid
	}

	return p.repo.GetProduct(ctx, id)
}

func (p *ProductCtlServices) ListProducts(
	ctx context.Context,
	req *http.Request,
) ([]*entities.Product, error) {
	page, err1 := strconv.Atoi(req.URL.Query().Get("page"))
	limit, err2 := strconv.Atoi(req.URL.Query().Get("limit"))
	if err1 != nil || err2 != nil {
		return nil, entities.ParamInvalid
	}

	return p.repo.ListProducts(ctx, page, limit)
}
