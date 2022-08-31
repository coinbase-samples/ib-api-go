package model

type ProfileResponse struct {
	UserId      string   `json:"userId"`
	Email       string   `json:"email"`
	Name        string   `json:"name"`
	LegalName   string   `json:"description"`
	UserName    string   `json:"username"`
	Roles       []string `json:"roles"`
	Address     string   `json:"address"`
	DateOfBirth string   `json:"dateOfBirth"`
}

type OrderRequest struct {
	ProductId   string  `json:"productId"`
	Side        string  `json:"side"`
	OrderType   string  `json:"orderType"`
	Quantity    float32 `json:"quantity"`
	LimitPrice  float32 `json:"limitPrice"`
	TimeInForce string  `json:"timeInForce"`
}

type OrderResponse struct {
	OrderId string       `json:"orderId"`
	Order   OrderRequest `json:"order"`
}
