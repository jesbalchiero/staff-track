package product

var products = []Product{
	{ID: "1", Name: "Produto A", Price: 10.0},
	{ID: "2", Name: "Produto B", Price: 20.5},
	{ID: "3", Name: "Produto C", Price: 30.0},
}

func GetAllProducts() []Product {
	return products
}
