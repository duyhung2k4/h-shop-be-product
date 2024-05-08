package request

type FileInfoRequest struct {
	Name      string `json:"name"`
	Format    string `json:"format"`
	DataBytes []byte `json:"dataBytes"`
}

type CreateProductRequest struct {
	InfoProduct map[string]interface{} `json:"infoProduct"`
	Files       []FileInfoRequest      `json:"files"`
}

type UpdateProductRequest struct {
	InfoProduct       map[string]interface{} `json:"infoProduct"`
	Files             []FileInfoRequest      `json:"files"`
	ListFileIdDeletes []uint64               `json:"listFileIdDeletes"`
}

type DeleteProductRequest struct {
	ProductId string `json:"productId"`
	ShopId    uint64 `json:"shopId"`
}
