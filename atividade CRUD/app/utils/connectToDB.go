package utils // Define o pacote utils para funções utilitárias

import (                            // Bloco de importação de pacotes
	"database/sql"                  // Pacote padrão para interação com bancos de dados SQL
	"fmt"                           // Pacote para formatação de strings
	"os"                            // Pacote para acessar variáveis de ambiente do sistema

	"github.com/joho/godotenv"      // Biblioteca para carregar variáveis do arquivo .env
	_ "github.com/lib/pq"           // Driver PostgreSQL importado apenas pelo efeito colateral de registro
)                                   // Fecha o bloco de importação

// ConnectToDB estabelece e retorna uma conexão com o banco de dados PostgreSQL
func ConnectToDB() (*sql.DB, error) { // Função que retorna ponteiro para conexão e possível erro
	godotenv.Load()                   // Carrega as variáveis de ambiente do arquivo .env

	user := os.Getenv("DB_USER")         // Obtém o usuário do banco das variáveis de ambiente
	password := os.Getenv("DB_PASSWORD") // Obtém a senha do banco das variáveis de ambiente
	dbname := os.Getenv("DB_NAME")       // Obtém o nome do banco das variáveis de ambiente
	host := os.Getenv("DB_HOST")         // Obtém o host do banco das variáveis de ambiente
	port := os.Getenv("DB_PORT")         // Obtém a porta do banco das variáveis de ambiente

	// Monta a string de conexão com os parâmetros do banco de dados
	connStr := fmt.Sprintf( // Formata a string de conexão com os valores obtidos
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", // Template da connection string
		host, port, user, password, dbname,                              // Valores das variáveis de ambiente
	) // Fecha a chamada do Sprintf

	db, err := sql.Open("postgres", connStr) // Abre conexão com o banco PostgreSQL
	if err != nil {                          // Verifica se houve erro ao abrir a conexão
		return nil, err                      // Retorna nil e o erro encontrado
	} // Fecha o bloco de verificação de erro

	err = db.Ping()   // Testa se a conexão com o banco está ativa
	if err != nil {    // Verifica se houve erro ao testar a conexão
		return nil, err // Retorna nil e o erro encontrado
	} // Fecha o bloco de verificação de erro

	return db, nil // Retorna a conexão estabelecida e nil indicando sucesso
} // Fecha a função ConnectToDB
