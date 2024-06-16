package request

type InsertTypeInWarehouseReq struct {
	ProductId string   `json:"productId"`
	Hastag    string   `json:"hastag"`
	Name      string   `json:"name"`
	Price     *float64 `json:"price"`
	Count     uint     `json:"count"`
}
