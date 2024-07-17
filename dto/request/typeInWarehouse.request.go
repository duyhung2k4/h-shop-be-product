package request

type InsertTypeInWarehouseReq struct {
	ProductId string   `json:"productId"`
	Hastag    string   `json:"hastag"`
	Name      string   `json:"name"`
	Price     *float64 `json:"price"`
	Count     uint     `json:"count"`
}

type UpdateTypeInWarehouseReq struct {
	Id     uint64   `json:"id"`
	Hastag string   `json:"hastag"`
	Name   string   `json:"name"`
	Price  *float64 `json:"price"`
	Count  uint     `json:"count"`
}

type DeleteTypeInWarehouseReq struct {
	Id uint64 `json:"id"`
}
