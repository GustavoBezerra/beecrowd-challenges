package main

import "fmt"

func calculateArea(raio float64) string {
	return fmt.Sprintf("%.4f", 3.14159*raio*raio)
}

func main() {
	var raio float64
	fmt.Scan(&raio)

	area := calculateArea(raio)
	fmt.Printf("A=%s\n", area)
}
