package utils // Define o pacote utils para funções utilitárias

import ( // Bloco de importação de pacotes
	"database/sql" // Pacote padrão para interação com bancos de dados SQL
) // Fecha o bloco de importação

// UpdatePatient atualiza os dados de um paciente existente no banco pelo CPF
func UpdatePatient(db *sql.DB, cpf, fullName, phone, diagnosis string) error { // Função que recebe conexão e dados a atualizar
	query := `UPDATE patients 
	           SET full_name = $1, phone = $2, diagnosis = $3 
	           WHERE cpf = $4` // Query SQL para atualizar os campos do paciente identificado pelo CPF

	result, err := db.Exec(query, fullName, phone, diagnosis, cpf) // Executa a query de atualização com os parâmetros
	if err != nil {                                                 // Verifica se houve erro ao executar a query
		return err                                                  // Retorna o erro encontrado
	} // Fecha o bloco de verificação de erro

	rowsAffected, err := result.RowsAffected() // Obtém o número de linhas afetadas pela atualização
	if err != nil {                             // Verifica se houve erro ao obter linhas afetadas
		return err                              // Retorna o erro encontrado
	} // Fecha o bloco de verificação de erro

	if rowsAffected == 0 { // Verifica se nenhuma linha foi atualizada
		return sql.ErrNoRows // Retorna erro indicando que o paciente não foi encontrado
	} // Fecha o bloco de verificação de linhas afetadas

	return nil // Retorna nil indicando que a atualização foi bem-sucedida
} // Fecha a função UpdatePatient
