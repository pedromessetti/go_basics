package main

import "fmt"

type Printable interface {
	printInfo()
}

type Book struct {
	Title string
	Price float32
}

type Drink struct {
	Name string
	Price float32
}

func (b Book) printInfo() {
	fmt.Printf("Book: %s\nPrice: %.2f\n", b.Title, b.Price)
}

func (d Drink) printInfo() {
	fmt.Printf("Drink: %s\nPrice: %.2f\n", d.Name, d.Price)
}

func main() {
	
	harryPotter := Book {
		Title: "Harry Potter",
		Price: 19.99,
	}

	coffee := Drink {
		Name: "Coffee",
		Price: 0.5,
	}

	infos := []Printable {harryPotter, coffee}

	for _, info := range infos {
		info.printInfo()
	}

}
