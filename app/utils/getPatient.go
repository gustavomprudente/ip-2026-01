package utils // Define o pacote utils para funções utilitárias

import ( // Bloco de importação de pacotes
	"database/sql" // Pacote padrão para interação com bancos de dados SQL
) // Fecha o bloco de importação

// Patient representa a estrutura de dados de um paciente do hospital
type Patient struct { // Define a struct Patient com os campos da tabela
	ID        int    // Identificador único do paciente
	FullName  string // Nome completo do paciente
	CPF       string // CPF do paciente (documento único)
	BirthDate string // Data de nascimento do paciente
	Phone     string // Telefone de contato do paciente
	Diagnosis string // Diagnóstico médico do paciente
	CreatedAt string // Data e hora de criação do registro
} // Fecha a definição da struct Patient

// GetPatient busca um paciente no banco de dados pelo CPF informado
func GetPatient(db *sql.DB, cpf string) (Patient, error) { // Função que recebe conexão e CPF, retorna Patient e erro
	var p Patient // Declara variável para armazenar os dados do paciente encontrado

	query := `SELECT id, full_name, cpf, 
	           TO_CHAR(birth_date, 'YYYY-MM-DD'), 
	           COALESCE(phone, ''), 
	           COALESCE(diagnosis, ''), 
	           TO_CHAR(created_at, 'DD/MM/YYYY HH24:MI') 
	           FROM patients WHERE cpf = $1` // Query SQL para buscar paciente pelo CPF com formatação de datas

	err := db.QueryRow(query, cpf).Scan( // Executa a query e lê os valores retornados
		&p.ID,        // Armazena o ID do paciente
		&p.FullName,  // Armazena o nome completo do paciente
		&p.CPF,       // Armazena o CPF do paciente
		&p.BirthDate, // Armazena a data de nascimento formatada
		&p.Phone,     // Armazena o telefone do paciente
		&p.Diagnosis, // Armazena o diagnóstico do paciente
		&p.CreatedAt, // Armazena a data de criação formatada
	) // Fecha o Scan dos campos

	if err != nil { // Verifica se houve erro na busca
		return p, err // Retorna o paciente vazio e o erro encontrado
	} // Fecha o bloco de verificação de erro

	return p, nil // Retorna o paciente encontrado e nil indicando sucesso
} // Fecha a função GetPatient
