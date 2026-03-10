package java;

import java.util.concurrent.*;

// 1. Criamos uma classe que implementa Callable<T>
// O <String> define o tipo de dado que a tarefa vai devolver
class MinhaTarefaComRetorno implements Callable<String> {
    private final String nome;

    public MinhaTarefaComRetorno(String nome) {
        this.nome = nome;
    }

    // 2. No Callable, sobrescrevemos o método call() em vez de run()
    // O método call() permite retornar um valor e lançar exceções [cite: 66, 75]
    @Override
    public String call() throws Exception {
        System.out.println("Iniciando cálculo na tarefa " + nome + "...");
        Thread.sleep(2000); // Simulando um processamento demorado
        return "Resultado da tarefa " + nome + " processado com sucesso!";
    }
}

public class PrincipalCallable {
    public static void main(String[] args) {
        // 3. Criamos o ExecutorService [cite: 53]
        ExecutorService executor = Executors.newFixedThreadPool(1);

        // 4. Enviamos a tarefa usando submit()
        // O submit() nos devolve um objeto Future imediatamente [cite: 64, 73, 97]
        System.out.println("Enviando tarefa...");
        Future<String> futuroResultado = executor.submit(new MinhaTarefaComRetorno("Cálculo Pesado"));

        // 5. Aqui poderíamos fazer outras coisas enquanto a tarefa roda...
        System.out.println("Aguardando o resultado no balcão... [cite: 98]");

        try {
            // 6. O método get() bloqueia a execução até que o resultado esteja pronto [cite: 90, 91]
            String resultadoFinal = futuroResultado.get(); 
            System.out.println("Mensagem recebida: " + resultadoFinal); 
        } catch (InterruptedException | ExecutionException e) {
            e.printStackTrace();
        }

        // 7. Finalizando o executor [cite: 34, 99]
        executor.shutdown();
    }
}