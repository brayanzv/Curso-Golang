package main

import (
	"fmt"
)


func main (){
	//indices definidos

	monedas:=[...]string{
		1: "pedro",
		3: "daniel",
		8: "gabriel",
	}

	fmt.Println(monedas, len(monedas))

	subMoneda:= monedas [3:]
	fmt.Println(subMoneda)

	subMoneda= monedas [3:6]
	fmt.Println(subMoneda)

	subMoneda= monedas [:]
	fmt.Println(subMoneda)

	subMoneda= monedas [:6]
	fmt.Println(subMoneda)
}