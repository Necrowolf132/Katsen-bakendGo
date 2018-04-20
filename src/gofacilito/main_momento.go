package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)
func main(){
	var clave string
	var HashHecho string
	fmt.Println("mete la mierda esa de clave aqui")
	fmt.Scanf("%s\n", &clave)
	if (clave == "0"){
		fmt.Println("mete la mierda esa de clave aqui")
		fmt.Scanf("%s\n", &clave)
		fmt.Println("mete la mierda esa de hash aqui")
		fmt.Scanf("%s\n", &HashHecho)
		err := bcrypt.CompareHashAndPassword([]byte(HashHecho),[]byte(clave))
		if err != nil {
			fmt.Println("contraseña invalida")
		}
	}else{
		hash , err := bcrypt.GenerateFromPassword([]byte(clave),bcrypt.DefaultCost)
		if err != nil {
			fmt.Println("paso elgo  con la clave", err)
		}
		fmt.Printf(string(hash))
		fmt.Printf("\n")
		}

}