package main

import "fmt"

func main() {
	// Encabezado
	fmt.Println("==== CALCULADORA CIENTÍFICA v1.0 ====")

	var n1 float64
	var n2 float64
	var operacion string

	// 1. Pedir los datos al usuario
	fmt.Print("Ingresa el primer número: ")
	fmt.Scan(&n1)
	fmt.Print("Ingresa el segundo número: ")
	fmt.Scan(&n2)

	fmt.Print("Ingresa la operación (+, -, *, /, ^, !): ")
	fmt.Scan(&operacion)

	// 3.
	switch operacion {
	case "+":
		fmt.Printf("Resultado: %.2f + %.2f = %.2f\n", n1, n2, n1+n2)
	case "-":
		fmt.Printf("Resultado: %.2f - %.2f = %.2f\n", n1, n2, n1-n2)
	case "*":
		fmt.Printf("Resultado: %.2f * %.2f = %.2f\n", n1, n2, n1*n2)
	case "/":
		if n2 == 0 {
			fmt.Println("Error: no se puede dividir entre cero")
		} else {
			fmt.Printf("Resultado: %.2f / %.2f = %.2f\n", n1, n2, n1/n2)
		}
	case "^":
		resultadoPot := 1.0
		for i := 0; i < int(n2); i++ {
			resultadoPot *= n1
		}
		fmt.Printf("Resultado: %.2f ^ %.0f = %.2f\n", n1, n2, resultadoPot)
	case "!":
		if n1 < 0 {
			fmt.Println("Error: No existe factorial de números negativos.")
		} else {
			resultadoFact := 1.0
			for i := 1; i <= int(n1); i++ {
				resultadoFact *= float64(i)
			}
			fmt.Printf("Resultado: %.0f! = %.2f\n", n1, resultadoFact)
		}
	default:
		fmt.Println("Error: operación no reconocida")
	}
}
