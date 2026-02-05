package main
import (
	"fmt"
	"math/rand/v2" 
)


func produtor(ch chan int){
	for{
	randNum := rand.IntN(100)
	
	ch <- randNum
	}
}
func consumidor(ch chan int, numero int){


	for{
		consome := <- ch
		if consome > numero{
			fmt.Println(consome)
		}
	}
}
func main(){

	ch := make(chan int)
	var numero int
	fmt.Scanln(&numero)

	go produtor(ch)
	go consumidor(ch, numero)

	select{}
}