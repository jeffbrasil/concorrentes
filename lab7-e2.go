package main
import (
	"fmt"
	"math/rand/v2" 
)


func produtor(ch chan int){
	for i := 0; i < 10000; i ++{
		randNum := rand.IntN(100)
		
		ch <- randNum	
	}
	close(ch)
}
func consumidor(ch chan int, numero int, done chan bool){

	i := 0
	for i < 10000{
		consome := <- ch
		if consome > numero{
			fmt.Println(consome)
		}
		i++
	}

	done <- true
}
func main(){

	ch := make(chan int)
	done := make(chan bool)

	var numero int
	fmt.Scanln(&numero)

	go produtor(ch)
	go consumidor(ch, numero, done)

	select{
	case <- done:
		fmt.Println("fim")
	}
}