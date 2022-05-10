package orders

import "fmt"

func New() {
	fmt.Println("Order package!!")
	natOrd := NewNationalOrder()
	intOrd := NewInternationalOrder()
	ords := []Operations{natOrd, intOrd}
	ProcessOrder(ords)
}

// Product struct
type Product struct {
	name  string
	price int
}

// ProductDetail struct
type ProductDetail struct {
	Product
	amount int
	total  float32
}

// Summary struct
type Summary struct {
	total          float32
	subtotal       float32
	totalBeforeTax float32
}

// ShippingAddress struct
type ShippingAddress struct {
	street  string
	city    string
	country string
	cp      string
}

// Client struct
type Client struct {
	name     string
	lastName string
	email    string
	phone    string
}

// Order struct
type Order struct {
	products []*ProductDetail
	Summary
	ShippingAddress
	Client
}

// Processer interface
type Processer interface {
	FillOrderSummary()
}

// Printer interface
type Printer interface {
	PrintOrderDetails()
}

// Notifier interface
type Notifier interface {
	Notify()
}

// Operations interface
type Operations interface {
	Processer
	Printer
	Notifier
}

// ProcessOrder function
func ProcessOrder(orders []Operations) {
	for _, order := range orders {
		order.FillOrderSummary()
		order.Notify()
		order.PrintOrderDetails()
	}
}
