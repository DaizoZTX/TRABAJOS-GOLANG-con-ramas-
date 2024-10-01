//CONVERSOR DE TEMPERATURA (Celcius a Farenheit)
package main
import "fmt"
func main(){
	//Inicializacion de la variables
	celcius := 0.00
	fmt.Println("CONVERSOR DE GRADOS CELCIUS A FARENHEIT!")
	fmt.Println("Ingrese el n° de grados celcius a transformar...")
	fmt.Scanln(&celcius)
	gradosTransform:=(celcius*9/5)+32
	fmt.Println(celcius," Grados Celcius son equivalentes a", gradosTransform, "Grados Farenheit\nIngrese cualquier tecla para finalizar el programa!")
	fmt.Scan(&celcius)
}
