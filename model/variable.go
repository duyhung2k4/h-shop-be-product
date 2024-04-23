package model

var (
	MapDefaultFieldProduct = map[string]string{
		"shopId": "shopId",
		"price":  "price",
	}
)

type QUEUE_PRODUCT string

const (
	PRODUCT_TO_ELASTIC QUEUE_PRODUCT = "product_to_elastic"
)
