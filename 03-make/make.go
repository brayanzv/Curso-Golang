package main

import "fmt"

func main(){
	numeros:= make([]int,0)

	for i := 100; i<110 ;i++  {
		numeros=append(numeros,i)
	}
	fmt.Println(numeros, len(numeros), cap(numeros))
}
