package main
import (
    "encoding/json"
    "fmt"
)

// Product 商品信息
type Product struct {
    Name      string `json:"name"`
    ProductID int64  `json:"productid"`
    Number    int    `json:"number"`
    Price     float64`json:"price"`
    IsOnSale  bool   `json:"isonsale"`
}


func defer_call() {
    defer func() { fmt.Println("打印前") }()
    defer func() { fmt.Println("打印中") }()
    defer func() { fmt.Println("打印后") }()

    panic("触发异常")
}

func main() {
    p := &Product{}
    p.Name = "Xiao mi 6"
    p.IsOnSale = true
    p.Number = 10000
    p.Price = 2499.00
    p.ProductID = 1
    if data, err := json.Marshal(p); err == nil {
		fmt.Println(string(data))
	} else {
		fmt.Println(err)
	}

	defer_call()
}