package Product

import (
	"context"
	"product-service/Config"
	pb "product-service/Controller/Dto/Proto"
	"product-service/Controller/Dto/Request"
	"product-service/Controller/Dto/Response"
)

func (p *product) ListProduct(ctx context.Context, param Request.ListParam) (resp Response.RespDataListProduct, err error) {
	var data Response.RespListProduct

	query := `SELECT id, title, description, price, stock FROM t_product WHERE stock > 0 ORDER BY title ASC LIMIT ? OFFSET ?`
	rows, err := Config.DATABASE_MAIN.Get().QueryContext(ctx, query, param.Limit, param.Offset)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&data.Id, &data.Title, &data.Description, &data.Price, &data.Stock)
		if err != nil {
			return
		}

		resp.List = append(resp.List, data)
	}

	resp.Page = int32(param.Page + 1)
	resp.TotalData, err = countDataListProduct(ctx)
	return
}

func countDataListProduct(ctx context.Context) (count int64, err error) {
	query := `SELECT COUNT(id) FROM t_product`
	err = Config.DATABASE_MAIN.Get().QueryRowContext(ctx, query).Scan(&count)
	return
}

func (p *product) CheckExistsProductId(ctx context.Context, id string) (exists bool, err error) {
	query := `SELECT EXISTS (SELECT 1 FROM t_product WHERE id = ? AND stock > 0) AS "exists"`
	err = Config.DATABASE_MAIN.Get().QueryRowContext(ctx, query, id).Scan(&exists)
	return
}

func (p *product) DetailProduct(ctx context.Context, id string) (res pb.Product, err error) {
	query := `SELECT id, title, description, price, stock FROM t_product WHERE id = ? AND stock > 0`
	if err = Config.DATABASE_MAIN.Get().QueryRowContext(ctx, query, id).Scan(&res.Id, &res.Title, &res.Description,
		&res.Price, &res.Stock); err != nil {
		return
	}
	return
}
