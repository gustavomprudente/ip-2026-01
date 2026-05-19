package handlers // Define o pacote handlers para funções de tratamento de rotas

import ( // Bloco de importação de pacotes
	"hospital-teste/app/utils" // Importa o pacote utils com funções utilitárias do projeto
	"net/http"                 // Pacote padrão para funcionalidades HTTP
	"net/url"                  // Pacote para manipulação de URLs e query strings
) // Fecha o bloco de importação

// UpdatePatientHandler trata as requisições HTTP para atualizar dados de um paciente
func UpdatePatientHandler(w http.ResponseWriter, r *http.Request) { // Handler que recebe writer e request HTTP
	if r.Method != http.MethodPost { // Verifica se o método HTTP é POST
		http.Redirect(w, r, "/static/forms/updatePatient.html", http.StatusSeeOther) // Redireciona para o formulário se não for POST
		return                                                                        // Encerra a execução do handler
	} // Fecha o bloco de verificação de método

	cpf := r.FormValue("cpf")            // Obtém o CPF enviado pelo formulário
	fullName := r.FormValue("full_name")  // Obtém o nome completo enviado pelo formulário
	phone := r.FormValue("phone")         // Obtém o telefone enviado pelo formulário
	diagnosis := r.FormValue("diagnosis") // Obtém o diagnóstico enviado pelo formulário

	db, err := utils.ConnectToDB() // Conecta ao banco de dados PostgreSQL
	if err != nil {                // Verifica se houve erro na conexão
		redirectURL := "/static/forms/updatePatient.html?status=error&msg=" + url.QueryEscape("Erro ao conectar ao banco de dados") // Monta URL com mensagem de erro
		http.Redirect(w, r, redirectURL, http.StatusSeeOther)                                                                       // Redireciona com erro de conexão
		return                                                                                                                       // Encerra a execução do handler
	} // Fecha o bloco de verificação de erro de conexão
	defer db.Close() // Garante que a conexão será fechada ao final da função

	err = utils.UpdatePatient(db, cpf, fullName, phone, diagnosis) // Chama a função para atualizar o paciente no banco
	if err != nil {                                                 // Verifica se houve erro ao atualizar
		redirectURL := "/static/forms/updatePatient.html?status=error&msg=" + url.QueryEscape("Erro ao atualizar. Verifique se o CPF está cadastrado.") // Monta URL com mensagem de erro
		http.Redirect(w, r, redirectURL, http.StatusSeeOther)                                                                                           // Redireciona com erro de atualização
		return                                                                                                                                           // Encerra a execução do handler
	} // Fecha o bloco de verificação de erro de atualização

	redirectURL := "/static/forms/updatePatient.html?status=success&msg=" + url.QueryEscape("Paciente atualizado com sucesso!") // Monta URL com mensagem de sucesso
	http.Redirect(w, r, redirectURL, http.StatusSeeOther)                                                                       // Redireciona com mensagem de sucesso
} // Fecha a função UpdatePatientHandler
