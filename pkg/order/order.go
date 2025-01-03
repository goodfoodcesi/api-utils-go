package order

import "encoding/json"

const (
	Type         = "Order"
	OrderCreated = "OrderCreated"
)

type Order struct {
	OrderID    string `json:"order_id"`
	CustomerID string `json:"customer_id"`
	ShopID     string `json:"shop_id"`
	DriverID   string `json:"driver_id"`
}

func NewOrder(orderID, customerID, shopID, DriverID string) *Order {
	return &Order{
		OrderID:    orderID,
		CustomerID: customerID,
		ShopID:     shopID,
		DriverID:   DriverID,
	}
}

func (o *Order) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}

func (o *Order) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, o)
}
