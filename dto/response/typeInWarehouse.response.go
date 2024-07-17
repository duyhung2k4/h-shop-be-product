package response

type UpdateTypeInWarehouseRes struct {
	Id     uint64   `json:"id"`
	Hastag string   `json:"hastag"`
	Name   string   `json:"name"`
	Price  *float64 `json:"price"`
	Count  uint     `json:"count"`
}

type DeleteTypeInWarehouseRes struct {
	Success bool `json:"success"`
}
