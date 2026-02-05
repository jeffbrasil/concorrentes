package main
import (
	"fmt"
	"math/rand/v2" 
)

func monitora(ch chan <- int, monitor chan int){
	count := 0
	for {

		count += <- monitor
		if count == 2{
			close(ch)
			break
		}
	}

}
func produtor(ch chan <-int, monitor chan int){

	for i := 0; i < 10; i ++{
		randNum := rand.N(100)
		
		ch <- randNum	
	}

	monitor <- 1
}
func consumidor(ch <- chan int, numero int, done chan bool){

	
	for consome := range ch{
		if consome > numero{
			fmt.Println(consome)
		}
		
	}

	done <- true
}
func main(){

	ch := make(chan int, 100)
	done := make(chan bool)
	monitor := make(chan int)

	var numero int
	fmt.Scanln(&numero)

	go monitora(ch, monitor)
	go produtor(ch, monitor)
	go produtor(ch, monitor)
	go consumidor(ch, numero, done)

	select{
	case <- done:
		fmt.Println("fim")
	}
}