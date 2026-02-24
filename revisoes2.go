package main
import (
	"fmt"
	"sync"
)

func tarefas(i int, wg *sync.WaitGroup){

	wg.Done() //sinaliza que a tarefa terminou
	fmt.Println("Tarefa %d concluída \n", i) 
}
func main(){

	var wg sync.WaitGroup

	for i:= 1; i< 4; i++ {
		wg.Add(1) // incrementa 1
		tarefas(i, &wg)

	}

	wg.Wait()
	fmt.Println("Tarefas finalizadas")
}