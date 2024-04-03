package Repository

import "product-service/Repository/Product"

type Repository struct {
	Product Product.ProductRepository
}

var ApplicationRepository = Repository{
	Product: Product.NewRepository(),
}
