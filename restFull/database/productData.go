package database

var ProductDataList []ProductData

type ProductData struct {
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}

func init() {
	ProductDataList = []ProductData{
		{Id: 1, Name: "Product 1", Price: 10.00},
		{Id: 2, Name: "Product 2", Price: 20.00},
		{Id: 3, Name: "Product 3", Price: 30.00},
		{Id: 4, Name: "Product 4", Price: 40.00},
		{Id: 5, Name: "Product 5", Price: 50.00},
		{Id: 6, Name: "Product 6", Price: 60.00},
	}
}
