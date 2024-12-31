package order

import "github.com/google/uuid"

const (
	Type         = "Order"
	OrderCreated = "OrderCreated"
)

type Order struct {
	ID         string `json:"id"`
	OrderID    string `json:"order_id"`
	CustomerID string `json:"customer_id"`
	ShopID     string `json:"shop_id"`
	DeliveryID string `json:"delivery_id"`
}

func NewOrder(orderID, customerID, shopID, deliveryID string) *Order {
	return &Order{
		ID:         uuid.NewString(),
		OrderID:    orderID,
		CustomerID: customerID,
		ShopID:     shopID,
		DeliveryID: deliveryID,
	}
}
