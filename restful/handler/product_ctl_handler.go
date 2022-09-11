package handler

import (
	"net/http"
	"strconv"

	"github.com/kumin/GolangMaster/restful/entities"
	"github.com/kumin/GolangMaster/restful/services"
)

type ProductCtlHandler struct {
	ctlService *services.ProductCtlService
}

func NewProductCtlHandler(
	ctlService *services.ProductCtlService,
) *ProductCtlHandler {
	return &ProductCtlHandler{
		ctlService: ctlService,
	}
}

func (p *ProductCtlHandler) ListProducts(
	req *http.Request,
) (interface{}, error) {
	if req.Method != http.MethodGet {
		return nil, entities.MethodNotAllowErr
	}
	page, err1 := strconv.Atoi(req.URL.Query().Get("page"))
	limit, err2 := strconv.Atoi(req.URL.Query().Get("limit"))
	if err1 != nil || err2 != nil {
		return nil, entities.ParamInvalid
	}
	return p.ctlService.ListProducts(req.Context(), page, limit)
}

func (p *ProductCtlHandler) GetProduct(
	req *http.Request,
) (interface{}, error) {
	switch req.Method {
	case http.MethodGet:
		id, err := strconv.ParseInt(req.URL.Query().Get("id"), 10, 64)
		if err != nil {
			return nil, entities.ParamInvalid
		}
		return p.ctlService.GetProduct(req.Context(), id)
	default:
		return nil, entities.MethodNotAllowErr
	}
}
