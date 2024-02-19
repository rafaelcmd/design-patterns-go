package main

import "fmt"

/* Factory Method pattern encapsulates the object creation logic within a method,
allowing subclasses to provide different implementations of that method to create
different types of objects while keeping the client code independent of the actual
object creation process. */

func main() {
	ak47, _ := getGun("ak47")
	musket, _ := getGun("musket")
	printDetails(ak47)
	printDetails(musket)
}

func printDetails(g IGun) {
	fmt.Printf("Gun: %s", g.getName())
	fmt.Println()
	fmt.Printf("Power: %d", g.getPower())
	fmt.Println()
}
