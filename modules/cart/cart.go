package cart

type Cart struct {
	ID         int `json:"id"`
	ProductID  int `json:"product_id"`
	TotalPrice int `json:"total_price"`
}
