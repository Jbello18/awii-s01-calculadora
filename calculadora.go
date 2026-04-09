package main

import "fmt"

func main(){
	//Encabezado decorativo
	fmt.Println("==== CALCULADORA CIENTÍFICA v1.0 ====")

	//Variables fijas
	a := 14
	b := 62

	//Operaciones
	suma := a + b
	resta := a - b
	multiplicacion := a * b 
	division := float64(a) / float64(b)

	fmt.Printf("a + b = %d\n", suma)
	fmt.Printf("a - b = %d\n", resta)
	fmt.Printf("a * b = %d\n", multiplicacion)
	fmt.Printf("a / b = %.2f\n", division)

}