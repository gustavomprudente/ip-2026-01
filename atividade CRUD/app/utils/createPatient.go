package utils // Define o pacote utils para funções utilitárias

import ( // Bloco de importação de pacotes
	"database/sql" // Pacote padrão para interação com bancos de dados SQL
) // Fecha o bloco de importação

// CreatePatient insere um novo paciente no banco de dados
func CreatePatient(db *sql.DB, fullName, cpf, birthDate, phone, diagnosis string) error { // Função que recebe conexão e dados do paciente
	query := `INSERT INTO patients (full_name, cpf, birth_date, phone, diagnosis) 
	           VALUES ($1, $2, $3, $4, $5)` // Query SQL para inserir um novo paciente na tabela

	_, err := db.Exec(query, fullName, cpf, birthDate, phone, diagnosis) // Executa a query com os parâmetros recebidos
	if err != nil {                                                       // Verifica se houve erro ao executar a query
		return err                                                        // Retorna o erro encontrado
	} // Fecha o bloco de verificação de erro

	return nil // Retorna nil indicando que a inserção foi bem-sucedida
} // Fecha a função CreatePatient
