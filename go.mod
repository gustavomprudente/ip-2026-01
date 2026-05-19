module hospital-teste // Define o nome do módulo Go do projeto

go 1.21 // Define a versão mínima do Go necessária

require ( // Bloco de dependências externas do projeto
	github.com/joho/godotenv v1.5.1 // Biblioteca para carregar variáveis do arquivo .env
	github.com/lib/pq v1.10.9        // Driver PostgreSQL para Go
) // Fecha o bloco de dependências