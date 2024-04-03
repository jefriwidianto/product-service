package Product

import (
	"context"
	"product-service/Controller/Dto/Request"
	"product-service/Controller/Dto/Response"
)

type ProductRepository interface {
	ListProduct(ctx context.Context, param Request.ListParam) (resp Response.RespDataListProduct, err error)
}

type product struct{}

func NewRepository() ProductRepository {
	return &product{}
}