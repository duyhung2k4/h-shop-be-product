package request

type UpdateWarehouseReq struct {
	Id        uint   `json:"id"`
	ProductId string `json:"productId"`
	Count     uint   `json:"count"`
}
