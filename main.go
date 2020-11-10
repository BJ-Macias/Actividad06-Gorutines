package main

import (
	"fmt"
	"time"
)

func Proceso(id uint64, imprimir chan bool, detener chan uint64) {
	i := uint64(0)
	for {
		select {
			case eliminar := <-detener:
				if id == eliminar {
					return
				}
			case <-imprimir:
				fmt.Printf("id %d: %d\n", id, i)
				i = i + 1		
			default:
				i = i + 1
			}
		time.Sleep(time.Millisecond * 500)
	}
}

func Print(imprimir chan bool, bandera chan bool) {
	for {
		select {
			case <-bandera:
				return
			default:
				imprimir <- true
		}
	}
}

func Stop(eliminarID uint64, detener chan uint64) {

	for i := uint64(0); i <= eliminarID; i++ {
		detener <-eliminarID
	}
	return
}

func main() {
	var op, ID, eliminarID uint64
	var pausa string
	imprimir := make(chan bool)
	bandera := make(chan bool)
	detener := make(chan uint64)

	for {
		fmt.Println("Elija una opcion")
		fmt.Println("1. Agregar proceso")
		fmt.Println("2. Mostrar procesos")
		fmt.Println("3. Eliminar un proceso")
		fmt.Println("4. Salir")
		fmt.Print("Opcion: ")
		fmt.Scanln(&op)

		switch op {
		case 1:
			go Proceso(ID, imprimir, detener)
			ID++
		case 2:
			go Print(imprimir, bandera)
			fmt.Scanln(&pausa)
			bandera <- true
		case 3:
			fmt.Println("Proceso a detener: ")
			fmt.Scanln(&eliminarID)
			go Stop(eliminarID, detener)
			fmt.Println("Proceso detenido exitosamente")

			fmt.Scanln(&pausa)
		case 4:
			fmt.Println("Saliendo")
			return
		}
	}
}