package main
import "fmt"
func main(){
	numeroIngreso := 0
	salida := ""
	fmt.Println("Ingrese un numero entero!")
	fmt.Scan(&numeroIngreso)
	if numeroIngreso%2==0{
		fmt.Printf("El numero %d es PRIMO! \n", numeroIngreso)
	}
	if numeroIngreso%2!=0 {
		fmt.Printf("El numero %d no es PRIMO! \n", numeroIngreso)
	}
	fmt.Println("Ingrese una tecla para finalizar el programa")
	fmt.Scan(&salida)
}