package request

type FileInfoRequest struct {
	Name      string `json:"name"`
	Format    string `json:"format"`
	DataBytes []byte `json:"dataBytes"`
}

type CreateProductRequest struct {
	InfoProduct map[string]interface{} `json:"infoProduct"`
	Avatar      *FileInfoRequest       `json:"avatar"`
	Files       []FileInfoRequest      `json:"files"`
}

type UpdateProductRequest struct {
	InfoProduct       map[string]interface{} `json:"infoProduct"`
	ListFieldDelete   []string               `json:"listFieldDelete"`
	Avatar            *FileInfoRequest       `json:"avatar"`
	Files             []FileInfoRequest      `json:"files"`
	ListFileIdDeletes []uint64               `json:"listFileIdDeletes"`
}

type DeleteProductRequest struct {
	ProductId string `json:"productId"`
	ShopId    uint64 `json:"shopId"`
}

type HeartRequest struct {
	ProductId string `json:"productId"`
}

type CartRequest struct {
	ProductId string `json:"productId"`
}
