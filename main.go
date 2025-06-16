package main

import (
	"fmt"
	"strings"
)

var Entero int = 10

func main() {
	entero := 10
	fmt.Println(entero + Entero)

	//Entero
	decimal := 3.14
	suma := entero + int(decimal)
	fmt.Println(suma)

	//texto
	mensaje := "Hola"
	concatenado := mensaje + "Maykol"
	enMayuscula := strings.ToUpper(concatenado)
	fmt.Println(enMayuscula)

	//bOOLEANOS
	esVerdadero := true
	fmt.Println(esVerdadero)

	// arrya fijiocon valor 1 en la posicion 3
	arrayFijo1 := [3]int{2: 1}
	fmt.Println(arrayFijo1) // Salida: [0 0 1]

	arrayFijo := [4]int{2: 5, 3: 1}
	fmt.Println(arrayFijo) // Salida: [0 0 5 1]ç

	//Arrays y slices
	arrayfijo := [3]int{1, 2, 3}
	fmt.Println(arrayfijo)
	sliceVariable := []int{4, 5, 6}
	sliceVariable = append(sliceVariable, 7)
	fmt.Println(sliceVariable)

	//MAPAS
	diccionario := map[string]int{
		"clave1": 1,
		"clave2": 2,
	}
	fmt.Println(diccionario["clave2"])
	diccionario1 := map[string]int{"clave1": 1, "clave2": 2}
	fmt.Println(diccionario1["clave1"])

	//struct
	type Person struct {
		nombre string
		edad   int
	}

	//creacion de persona
	persona1 := Person{nombre: "luis", edad: 20}

	// reasignacion
	persona1 = Person{nombre: "mateo"}

	fmt.Println(persona1)

}
