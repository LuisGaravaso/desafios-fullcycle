package usecase

type CreateOrderInputDTO struct {
	ID    string  `json:"id"`
	Price float64 `json:"price"`
	Tax   float64 `json:"tax"`
}

type OrderOutputDTO struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
	Exists     bool    `json:"exists"`
}

type GetOrderInputDTO struct {
	ID string `json:"id"`
}

type GetOrderOutputDTO struct {
	Message    string `json:"message"`
	OrderCount int    `json:"order_count"`
}
