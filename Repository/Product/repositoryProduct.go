package Product

import (
	"context"
	"product-service/Config"
	"product-service/Controller/Dto/Request"
	"product-service/Controller/Dto/Response"
)

func (p *product) ListProduct(ctx context.Context, param Request.ListParam) (resp Response.RespDataListProduct, err error) {
	var data Response.RespListProduct

	query := `SELECT id, title, description, price, stock FROM t_product ORDER BY title ASC LIMIT ? OFFSET ?`
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
