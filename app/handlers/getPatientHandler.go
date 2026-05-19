package handlers // Define o pacote handlers para funções de tratamento de rotas

import ( // Bloco de importação de pacotes
	"fmt"                      // Pacote para formatação de strings e saída
	"hospital-teste/app/utils" // Importa o pacote utils com funções utilitárias do projeto
	"net/http"                 // Pacote padrão para funcionalidades HTTP
) // Fecha o bloco de importação

// GetPatientHandler trata as requisições HTTP para buscar um paciente pelo CPF
func GetPatientHandler(w http.ResponseWriter, r *http.Request) { // Handler que recebe writer e request HTTP
	if r.Method != http.MethodGet { // Verifica se o método HTTP é GET
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed) // Retorna erro 405 se não for GET
		return                                                             // Encerra a execução do handler
	} // Fecha o bloco de verificação de método

	cpf := r.URL.Query().Get("cpf") // Obtém o CPF da query string da URL
	if cpf == "" {                   // Verifica se o CPF foi informado
		http.Redirect(w, r, "/static/forms/getPatient.html", http.StatusSeeOther) // Redireciona para o formulário de busca
		return                                                                     // Encerra a execução do handler
	} // Fecha o bloco de verificação de CPF vazio

	db, err := utils.ConnectToDB() // Conecta ao banco de dados PostgreSQL
	if err != nil {                // Verifica se houve erro na conexão
		http.Redirect(w, r, "/static/forms/getPatient.html?status=error&msg=Erro+ao+conectar+ao+banco", http.StatusSeeOther) // Redireciona com erro
		return // Encerra a execução do handler
	} // Fecha o bloco de verificação de erro de conexão
	defer db.Close() // Garante que a conexão será fechada ao final da função

	patient, err := utils.GetPatient(db, cpf) // Busca o paciente no banco pelo CPF informado
	if err != nil {                            // Verifica se houve erro na busca
		http.Redirect(w, r, "/static/forms/getPatient.html?status=error&msg=Paciente+n%C3%A3o+encontrado", http.StatusSeeOther) // Redireciona com erro de paciente não encontrado
		return // Encerra a execução do handler
	} // Fecha o bloco de verificação de erro de busca

	w.Header().Set("Content-Type", "text/html; charset=utf-8") // Define o tipo de conteúdo como HTML UTF-8

	// Escreve a página HTML com os dados do paciente encontrado
	fmt.Fprintf(w, `<!DOCTYPE html> <!-- Declaração do tipo de documento HTML5 -->
<html lang="pt-BR"> <!-- Abre a tag HTML com idioma português do Brasil -->
<head> <!-- Abre a seção de metadados da página -->
    <meta charset="UTF-8"> <!-- Define a codificação de caracteres como UTF-8 -->
    <meta name="viewport" content="width=device-width, initial-scale=1.0"> <!-- Configura viewport para responsividade -->
    <title>Resultado da Busca — Hospital Teste</title> <!-- Título exibido na aba do navegador -->
    <link rel="preconnect" href="https://fonts.googleapis.com"> <!-- Pré-conecta ao servidor de fontes Google -->
    <link href="https://fonts.googleapis.com/css2?family=Inter:wght@400;500;600;700&display=swap" rel="stylesheet"> <!-- Importa a fonte Inter do Google Fonts -->
    <link rel="stylesheet" href="/static/styles/getPatient.style.css"> <!-- Linka o arquivo CSS de estilos da busca -->
</head> <!-- Fecha a seção de metadados -->
<body> <!-- Abre o corpo da página -->
    <div class="container"> <!-- Container principal centralizado -->
        <a href="/" class="back-link">← Voltar ao Início</a> <!-- Link para retornar à página inicial -->
        <h1>📋 Resultado da Busca</h1> <!-- Título principal da página -->
        <div class="result-card"> <!-- Card com os dados do paciente -->
            <div class="field"> <!-- Campo individual de dado -->
                <span class="label">ID</span> <!-- Rótulo do campo -->
                <span class="value">%d</span> <!-- Valor do ID do paciente -->
            </div> <!-- Fecha o campo -->
            <div class="field"> <!-- Campo individual de dado -->
                <span class="label">Nome Completo</span> <!-- Rótulo do campo -->
                <span class="value">%s</span> <!-- Valor do nome do paciente -->
            </div> <!-- Fecha o campo -->
            <div class="field"> <!-- Campo individual de dado -->
                <span class="label">CPF</span> <!-- Rótulo do campo -->
                <span class="value">%s</span> <!-- Valor do CPF do paciente -->
            </div> <!-- Fecha o campo -->
            <div class="field"> <!-- Campo individual de dado -->
                <span class="label">Data de Nascimento</span> <!-- Rótulo do campo -->
                <span class="value">%s</span> <!-- Valor da data de nascimento -->
            </div> <!-- Fecha o campo -->
            <div class="field"> <!-- Campo individual de dado -->
                <span class="label">Telefone</span> <!-- Rótulo do campo -->
                <span class="value">%s</span> <!-- Valor do telefone -->
            </div> <!-- Fecha o campo -->
            <div class="field"> <!-- Campo individual de dado -->
                <span class="label">Diagnóstico</span> <!-- Rótulo do campo -->
                <span class="value">%s</span> <!-- Valor do diagnóstico -->
            </div> <!-- Fecha o campo -->
            <div class="field"> <!-- Campo individual de dado -->
                <span class="label">Cadastrado em</span> <!-- Rótulo do campo -->
                <span class="value">%s</span> <!-- Valor da data de cadastro -->
            </div> <!-- Fecha o campo -->
        </div> <!-- Fecha o card de resultado -->
        <a href="/static/forms/getPatient.html" class="btn-back">Nova Busca</a> <!-- Botão para realizar nova busca -->
    </div> <!-- Fecha o container principal -->
</body> <!-- Fecha o corpo da página -->
</html> <!-- Fecha o documento HTML -->`,
		patient.ID,        // Insere o ID do paciente no template
		patient.FullName,  // Insere o nome completo no template
		patient.CPF,       // Insere o CPF no template
		patient.BirthDate, // Insere a data de nascimento no template
		patient.Phone,     // Insere o telefone no template
		patient.Diagnosis, // Insere o diagnóstico no template
		patient.CreatedAt, // Insere a data de cadastro no template
	) // Fecha o Fprintf que escreve o HTML na resposta
} // Fecha a função GetPatientHandler
