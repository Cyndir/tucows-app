package model

type Customer struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	ID        string `json:"id"`
}

type Product struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Order struct {
	ID         string `json:"id" sql:"id"`
	CustomerID string `json:"customer_id" sql:"customerid"`
	ProductID  string `json:"product_id" sql:"productid"`
	Status     string `json:"status" sql:"status"`
	Total      int    `json:"total"`
}
