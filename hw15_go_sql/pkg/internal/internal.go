package internal

type User struct {
	ID       int    `json:"id,string,omitempty"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Order struct {
	ID          int         `json:"id,string,omitempty"`
	UserName    string      `json:"userName"`
	OrderDate   string      `json:"orderDate,string,omitempty"`
	TotalAmount interface{} `json:"totalAmount"`
}

type Product struct {
	ID    int         `json:"id,string,omitempty"`
	Name  string      `json:"name"`
	Price interface{} `json:"price"`
}
