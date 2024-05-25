package model

var (
	MapDefaultFieldProduct = map[string]string{
		"categoryId": "categoryId",
		"price":      "price",
		"name":       "name",
	}
)

type QUEUE_PRODUCT string

const (
	PRODUCT_TO_ELASTIC        QUEUE_PRODUCT = "product_to_elastic"
	UPDATE_PRODUCT_TO_ELASTIC QUEUE_PRODUCT = "update_product_to_elastic"
	DELETE_PRODUCT_TO_ELASTIC QUEUE_PRODUCT = "delete_product_to_elastic"
)

var (
	TRUE_VALUE  = true
	FALSE_VALUE = false

	MapDataEmpty    map[string]interface{}
	ArrMapDataEmpty []map[string]interface{}
)
