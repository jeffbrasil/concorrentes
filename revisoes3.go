package main
import(

	"fmt"
	"time"
)

func produtor(ch chan string){

	ch <- "Processando mensagem..."
	time.Sleep(2 * time.Second)
	ch <- "Finalizando mensagem"

	close(ch)
}
func main(){

	mensagens := make(chan string)
	
	go produtor(mensagens)

	for msg := range mensagens {
		fmt.Println("mensagem recebida:  ", msg)
	}

	fmt.Println("fluxo encerrado")

}