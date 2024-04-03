package Response

type RespListProduct struct {
	Id          string  `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int32   `json:"stock"`
}

type RespDataListProduct struct {
	List      []RespListProduct `json:"list"`
	Page      int32             `json:"page"`
	TotalData int64             `json:"totalData"`
}
