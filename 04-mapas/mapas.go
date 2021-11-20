package _4_mapas

import "fmt"

func Mapas(){

	dias:= make(map[int]string)
	// int se refiere al tipo de clave en este caso son enteros
	//string es al tipo de dato guardado en esa clave
	//los mapas no tienen un limite de datos

	dias[1] = "domingo"
	dias[2] = "lunes"

	fmt.Println(dias,len(dias))
	//modificamos
	dias[1]="sabado"
	//agregamos
	dias[8]="otro no existe"
	fmt.Println(dias,len(dias))
	//eliminamos
	delete(dias,8)

	fmt.Println(dias,len(dias))
	//buscamos
	fmt.Println(dias[1],len(dias[1]))

}
