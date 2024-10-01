package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Variables
	var oracion string
	var letraBusq string
	lector := bufio.NewScanner(os.Stdin)
	// Codigo!
	fmt.Println("-CONTADOR DE CARACTERES-")
	fmt.Print("Ingrese una oracion!: ")
	lector.Scan()
	oracion = lector.Text()
	fmt.Print("Ingrese la letra que desea contar en la oracion!: ")
	lector.Scan()
	letraBusq = lector.Text()
	// LowerCase
	oracion = strings.ToLower(oracion)
	letraBusq = strings.ToLower(letraBusq)
	// Finder
	confirmador := strings.Contains(oracion, letraBusq)
	// Impresion
	fmt.Println("Su oracion posee un total de ", len(oracion), " caracteres! (Incluyendo espacios en blanco)")
	if confirmador == true {
		cont := strings.Count(oracion, letraBusq)
		if cont == 1 {
			fmt.Println("La oracion:", oracion, " Posee ", cont, " vez el caracter: ", letraBusq, ".")
		} else {
			if cont > 1 {
				fmt.Println("La oracion:", oracion, " Posee ", cont, " veces el caracter: ", letraBusq, ".")
			}
		}
	} else {
		if confirmador == false {
			fmt.Println("La oracion no contiene el caracter: ", letraBusq)
		}
	}
	//Otra opcion
	var opc int
	fmt.Println("Deseea examinar una palabra especifica de la oracion? (ADVERTENCIA: ESTAS SON MANEJADAS POR MODULOS VECTORIALES DE 0-n)")
	fmt.Println("1 = Si\n2 = No")
	fmt.Scan(&opc)
	if opc == 1 {
		var numVec int
		split := strings.Split(oracion, " ")
		fmt.Println("Ingrese el numero del vector correspondiente a la palabra")
		fmt.Scan(&numVec)
		fmt.Println("La palabra correspondiente al modulo n°", numVec, " del vector es: ", split[numVec])
		fmt.Println("La palabra posee ", len(split[numVec]), " caracteres!")
		confirmador = strings.Contains(split[numVec], letraBusq)
		if confirmador == true {
			cont := strings.Count(split[numVec], letraBusq)
			if cont == 1 {
				fmt.Println("La palabra:", split[numVec], " posee ", cont, " vez el caracter: ", letraBusq, ".")
			} else {
				if cont > 1 {
					fmt.Println("La palabra:", split[numVec], " Posee ", cont, " veces el caracter: ", letraBusq, ".")
				}
			}
		} else {
			if confirmador == false {
				fmt.Println("La palabra no contiene el caracter: ", letraBusq)
			}
		}
		fmt.Println("Ingrese un caracter cualquiera para finalizar el programa!")
		fmt.Scan(&numVec)
	}
}
