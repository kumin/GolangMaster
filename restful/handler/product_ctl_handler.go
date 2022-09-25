package handler

import (
	"net/http"

	"github.com/kumin/GolangMaster/restful/entities"
	"github.com/kumin/GolangMaster/restful/services"
)

type ProductCtlHandler struct {
	ctlService *services.ProductCtlServices
}

func NewProductCtlHandler(
	ctlService *services.ProductCtlServices,
) *ProductCtlHandler {
	return &ProductCtlHandler{
		ctlService: ctlService,
	}
}

func (p *ProductCtlHandler) AddProduct(
	req *http.Request,
) (interface{}, error) {
	if req.Method != http.MethodPost {
		return nil, entities.MethodNotAllowErr
	}

	return p.ctlService.AddProduct(req.Context(), req)
}

func (p *ProductCtlHandler) ListProducts(
	req *http.Request,
) (interface{}, error) {
	if req.Method != http.MethodGet {
		return nil, entities.MethodNotAllowErr
	}

	return p.ctlService.ListProducts(req.Context(), req)
}

func (p *ProductCtlHandler) GetProduct(
	req *http.Request,
) (interface{}, error) {
	switch req.Method {
	case http.MethodGet:
		return p.ctlService.GetProduct(req.Context(), req)
	default:
		return nil, entities.MethodNotAllowErr
	}
}
