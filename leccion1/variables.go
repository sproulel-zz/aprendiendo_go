package main

import "fmt"

func cambiar(a, b string) (string, string) {
	return b, a
}

func main() {
	var nombre string = "lucas"

	fmt.Println("hola me llamo", nombre)

	var edad int = 23
	fmt.Printf("tengo %d años \n", edad)

	//tipos

	//sin definir todos tipos van a ser falso
	//string por ejemplo seria ""
	//int o numero seria 0
	var chevere bool
	fmt.Println("lucas es chevere?", chevere)

	//no tienes que poner el tipo cada vez
	str := "sin poner tipo"
	fmt.Println(str)

	//puedes definir variables a la vez
	peso, altura := 50, 1.64
	fmt.Printf("peso %d kilogramos y tengo %g metros \n", peso, altura)

	hola, mundo := cambiar("hola", "mundo")
	fmt.Println(hola, mundo)
}
