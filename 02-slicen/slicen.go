package main

import "fmt"

func main(){
	numeros:=[]int{1,2,3}

	fmt.Println(numeros, len(numeros))

	//agregar datos
	numeros = append(numeros,5)
	numeros = append(numeros, 4)

	fmt.Println(numeros)

	//sub slicen
	subSlicen:= numeros[:2]

	numeros [0]= 100
	/*el subslicen sigue perteneciendo al slicen principal
	todo esto se da xq ambos slicen derivan de un array padre*/

	fmt.Println(numeros)
	fmt.Println(subSlicen,len(subSlicen))

	/*punteros
	longitud
	capacidad
	 */

	meses:= []string{"enero","febrero","marzo"}

	fmt.Printf("longitud: %v, capacidad: %v, %p \n", len(meses), cap(meses), meses)
	meses = append(meses, "abril", "mayo")

	fmt.Printf("longitud: %v, capacidad: %v,esta es la direccion de memoria%p \n", len(meses), cap(meses), meses)


}
