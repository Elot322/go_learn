package orders

var (
	subtotal       float32
	total          float32
	totalBeforeTax float32
	totalTaxes     float32
	totalExtraFee  float32
)

// функция GetProductDetail, получающая в качестве аргументов поля,
// необходимые для создания новой структуры ProductDetail,           // и возвращающая ее.
func GetProductDetail(name string, price, amount int, total float32) (pd *ProductDetail) {
	pd = &ProductDetail{
		amount: amount,
		total:  total,
		Product: Product{
			name:  name,
			price: price,
		},
	}
	return
}

// функция SetClient, получающая в качестве аргументов поля,
// необходимые для создания новой структуры Client, и возвращающая
// ее.
func SetClient(name, lastName, email, phone string) (cl Client) {
	cl = Client{
		name:     name,
		lastName: lastName,
		email:    email,
		phone:    phone,
	}
	return
}

// функция SetShippingAddress, получающая в качестве аргументов
// поля, необходимые для создания новой структуры ShippingAddress, и
// возвращающая ее.
func SetShippingAddress(street, city, country, cp string) (spa ShippingAddress) {
	spa = ShippingAddress{
		street,
		city,
		country,
		cp,
	}
	return
}

// функция CalculateSubTotal
func CalculateSubTotal(products []*ProductDetail) (subtotal float32) {
	for _, v := range products {
		subtotal += v.total
	}
	return
}
