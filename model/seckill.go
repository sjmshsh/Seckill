package model

// User 用户
type User struct {
	ID int64 `json:"id"`
}

func (User) TableName() string {
	return "user"
}

// Product 产品
type Product struct {
	ID    int64 `json:"id"`
	Stock int   `json:"stock"` // 库存字段
}

func (Product) TableName() string {
	return "product"
}

// Order 订单
type Order struct {
	ID        int64 `json:"id"`
	UserID    int64 `json:"user_id"`
	ProductId int64 `json:"product_id"`
}

func (Order) TableName() string {
	return "order"
}
