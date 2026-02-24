package main

import (
	"fmt"
	"time"

)
func exibirMensagem(texto string){
	fmt.Println(texto)
}

func main(){

	go exibirMensagem("Oi, eu sou o Goku")

	go func(){
		fmt.Println("Eae Tafarel")
	}()

	time.Sleep(time.Second)
}