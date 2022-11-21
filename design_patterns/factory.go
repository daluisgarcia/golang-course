package main

import "fmt"

type IProduct interface {
	setStock(stock int)
	getStock() int
	setName(name string)
	getName() string
}

type Computer struct {
	stock int
	name  string
}

func (c *Computer) setStock(stock int) {
	c.stock = stock
}

func (c *Computer) getStock() int {
	return c.stock
}

func (c *Computer) setName(name string) {
	c.name = name
}

func (c *Computer) getName() string {
	return c.name
}

type Laptop struct {
	Computer
}

func NewLaptop() IProduct { // Struct constructor
	return &Laptop{
		Computer{
			stock: 25,
			name:  "Laptop computer",
		},
	}
}

type Desktop struct {
	Computer
}

func NewDesktop() IProduct { // Struct constructor
	return &Desktop{
		Computer{
			stock: 35,
			name:  "Desktop computer",
		},
	}
}

func GetComputerFactory(name string) (IProduct, error) { // Factory constructor
	switch name {
	case "laptop":
		return NewLaptop(), nil
	case "desktop":
		return NewDesktop(), nil
	default:
		return nil, fmt.Errorf("Computer type not recognized")
	}
}

// func main() {
// 	computer, err := GetComputerFactory("laptop")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Println(computer.getName())
// 	fmt.Println(computer.getStock())

// 	computer, err = GetComputerFactory("desktop")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Println(computer.getName())
// 	fmt.Println(computer.getStock())
// }
