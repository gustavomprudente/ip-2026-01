package main // Define o pacote principal para inicialização do servidor

import ( // Bloco de importação de dependências do servidor
	"fmt"                          // Pacote para formatação e saída de dados no console
	"hospital-teste/app/handlers"  // Importa os handlers de rotas do próprio projeto
	"log"                          // Pacote para registrar logs e erros de execução
	"net/http"                     // Pacote padrão para criar o servidor e tratar requisições HTTP
) // Fecha o bloco de importações

func main() { // Função de entrada principal que inicia o programa
	// Configura o servidor para servir arquivos estáticos da pasta "static"
	fileServer := http.FileServer(http.Dir("./static")) // Cria um manipulador de arquivos estáticos a partir do diretório static
	http.Handle("/static/", http.StripPrefix("/static/", fileServer)) // Registra a rota /static/ para servir os arquivos estáticos

	// Registra os manipuladores (handlers) para cada rota específica do sistema
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { // Rota raiz do servidor
		http.ServeFile(w, r, "./static/index.html") // Serve o arquivo da página inicial index.html
	}) // Fecha a definição da rota raiz

	http.HandleFunc("/criar", handlers.CreatePatientHandler) // Rota POST para criação de novos pacientes
	http.HandleFunc("/buscar", handlers.GetPatientHandler)   // Rota GET para busca de pacientes pelo CPF
	http.HandleFunc("/atualizar", handlers.UpdatePatientHandler) // Rota POST para atualização de dados dos pacientes
	http.HandleFunc("/deletar", handlers.DeletePatientHandler)   // Rota POST para remoção de pacientes pelo CPF

	port := ":8080" // Define a porta padrão em que o servidor irá escutar
	fmt.Printf("Servidor do Hospital Teste iniciado com sucesso em http://localhost%s\n", port) // Exibe mensagem informativa no console

	err := http.ListenAndServe(port, nil) // Inicia a escuta na porta e trata as conexões recebidas
	if err != nil { // Verifica se ocorreu algum erro crítico ao subir o servidor
		log.Fatal("Erro ao iniciar o servidor HTTP: ", err) // Registra o erro no log e interrompe a execução
	} // Fecha a verificação de erro do servidor
} // Fecha a função principal main
