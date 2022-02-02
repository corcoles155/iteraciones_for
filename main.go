package main

import "fmt"


func main()  {

	preguntarNumeroSecreto()

	forConContinue()

	forConGoto()

}	

func preguntarNumeroSecreto()  {
	numero:=1
	for {
		fmt.Println("Continuo")
		fmt.Println("Ingrese el número secreto")
		fmt.Scanln(&numero)
		if numero == 100 {
			break;
		}
	}
}

func forConContinue()  {
	var i = 0
	for i < 10 {
		fmt.Println("\n Valor de i: %d", i)
		if i == 5 {
			fmt.Println("Multiplicamos por 2 \n")
			i = i*2
			continue
		}
		fmt.Println("Paso por aquí \n")
		i++
	}
}

func forConGoto()  {
	var i int = 0

	//Al poner ":" estamos creando etiquetas
	RUTINA:
		for i < 10 {
			if i == 4 {
				i=i+2
				fmt.Println("Voy a RUTINA sumando 2 a i")
				goto RUTINA
			}
			fmt.Println("Valor de i: %d", i)
			i++
		}
}
