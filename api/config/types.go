package config

type Response struct {
	Message interface{} `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

type AdminAuth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AdminUpdate struct {
	Password string `json:"password"`
}

type AdminOutput struct {
	Id       string `json:"id"`
	Username string `json:"username"`
}

type CategoryInput struct {
	Name string `json:"name"`
}
type CategoryOutput struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type UserAuth struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserCreate struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Address  string `json:"address"`
}

type UserUpdatePassword struct {
	Password string `json:"password"`
}

type UserUpdateData struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

type UserOutput struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Address string `json:"address"`
}

type ProductInput struct {
	Name        string `json:"name"`
	Price       uint   `json:"price"`
	Description string `json:"description"`
	Stock       uint   `json:"stock"`
	Category_id string `json:"category_id"`
}

type ProductOutput struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Price       uint   `json:"price"`
	Description string `json:"description"`
	Stock       uint   `json:"stock"`
	Category    string `json:"category,omitempty"`
	Seller      string `json:"seller,omitempty"`
}

type InvoiceItemInput struct {
	Product_id string `json:"product_id"`
	Qty        uint   `json:"qty"`
}

type InvoiceInput struct {
	Buyer_id string             `json:"buyer_id"`
	Items    []InvoiceItemInput `json:"items"`
}

type InvoiceItemOutput struct {
	Id           string `json:"id"`
	Product_name string `json:"product_name"`
	Qty          uint   `json:"qty"`
	Price        uint   `json:"price"`
}

type InvoiceOutput struct {
	Id          string              `json:"id"`
	Buyer_name  string              `json:"buyer_name"`
	Items       []InvoiceItemOutput `json:"items"`
	Total_price uint                `json:"total_price"`
	Bought_at   string              `json:"bought_at"`
}
