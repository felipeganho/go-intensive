//go run main.go
//GOOS=windows go build main.go

//go mod init
//go mod tidy

// package main

// import (
// 	"fmt"

// 	"github.com/google/uuid"
// )

// func main() {
// 	fmt.Println(uuid.New().String())
// }

package main

import "fmt"

type Order struct {
	ID       string
	Price    float64
	Quantity int
}

//func (o *Order)
func (o *Order) setPrice(price float64) {
	o.Price = price
	fmt.Println("Price dentro do setPrice:", o.Price)
}

func (o Order) getTotal() float64 {
	return o.Price * float64(o.Quantity)
}

// func NewOrder() Order {
// 	return Order{}
// }

func NewOrder() *Order {
	return &Order{}
}

func main() {
	println("Hello World!")
	name := "Felipe"
	println(name)

	// order := Order{
	// 	ID:       "123",
	// 	Price:    10.0,
	// 	Quantity: 5,
	// }

	order := NewOrder()
	order.ID = "123"
	order.Price = 10.0
	order.Quantity = 5

	order2 := order
	order.Price = 5
	fmt.Println(order2.Price)

	// fmt.Println(order.ID, order.Price, order.Quantity)
	// fmt.Println(order.getTotal())
	// order.setPrice(30)
	// fmt.Println(order.getTotal())
}
