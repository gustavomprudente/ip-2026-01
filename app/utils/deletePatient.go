package utils // Define o pacote utils para funções utilitárias

import ( // Bloco de importação de pacotes
	"database/sql" // Pacote padrão para interação com bancos de dados SQL
) // Fecha o bloco de importação

// DeletePatient remove um paciente do banco de dados pelo CPF informado
func DeletePatient(db *sql.DB, cpf string) error { // Função que recebe conexão e CPF do paciente a deletar
	query := `DELETE FROM patients WHERE cpf = $1` // Query SQL para deletar o paciente pelo CPF

	result, err := db.Exec(query, cpf) // Executa a query de exclusão com o CPF informado
	if err != nil {                    // Verifica se houve erro ao executar a query
		return err                     // Retorna o erro encontrado
	} // Fecha o bloco de verificação de erro

	rowsAffected, err := result.RowsAffected() // Obtém o número de linhas afetadas pela exclusão
	if err != nil {                             // Verifica se houve erro ao obter linhas afetadas
		return err                              // Retorna o erro encontrado
	} // Fecha o bloco de verificação de erro

	if rowsAffected == 0 { // Verifica se nenhuma linha foi deletada
		return sql.ErrNoRows // Retorna erro indicando que o paciente não foi encontrado
	} // Fecha o bloco de verificação de linhas afetadas

	return nil // Retorna nil indicando que a exclusão foi bem-sucedida
} // Fecha a função DeletePatient
