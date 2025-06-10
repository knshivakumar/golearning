package models

type Order struct {
	OrderId  int     `json: "orderId"`
	Type     string  `json: "type"`
	Price    float64 `json: "price"`
	Quantity int     `json: "quantity`
}
