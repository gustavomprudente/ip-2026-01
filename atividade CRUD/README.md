## 🛠️ Pré-requisitos

Para rodar este projeto em sua máquina local, você precisará ter instalado:

1. **GoLang (Versão 1.20 ou superior)**: [Download do Go](https://go.dev/dl/)
2. **Docker Desktop** (para subir o banco de dados facilmente): [Download do Docker](https://www.docker.com/products/docker-desktop/)
3. **DBeaver Community** (ou qualquer gerenciador de banco de dados SQL): [Download do DBeaver](https://dbeaver.io/download/)

---

## 🚀 Passo a Passo para Execução

### 1. Clonar ou Acessar a Pasta do Repositório
Abra o seu terminal na pasta raiz onde o projeto está localizado:
```bash
cd "c:\Users\gusta\OneDrive\Documentos\Lista 01 - Jaqson\ip-2026-01"
```

### 2. Configurar o Arquivo `.env`
O arquivo de configuração do ambiente (`.env`) já foi criado automaticamente com as credenciais padrões do banco de dados na raiz do projeto. Caso precise alterar alguma porta ou usuário, basta abrir e editar as variáveis:
```env
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=hospital_teste
DB_HOST=localhost
DB_PORT=5432
```

### 3. Iniciar o Banco de Dados com Docker
Utilizaremos o Docker Compose para criar e rodar a instância do PostgreSQL 16 instantaneamente sem precisar instalar o banco localmente.

Execute o seguinte comando no terminal na raiz do projeto:
```bash
docker compose up -d
```
> O parâmetro `-d` roda o container em segundo plano.

### 4. Criar a Tabela no DBeaver
1. Abra o **DBeaver**.
2. Clique em **Nova Conexão** (ícone de tomada no canto superior esquerdo) e selecione **PostgreSQL**.
3. Insira as credenciais de acordo com o arquivo `.env`:
   - **Host**: `localhost`
   - **Port**: `5432`
   - **Database**: `hospital_teste`
   - **Username**: `postgres`
   - **Password**: `postgres`
4. Clique em **Testar Conexão**. Se solicitado, permita o download do driver PostgreSQL.
5. Após conectar, abra o console SQL (clique com o botão direito no banco `hospital_teste` -> **Editor SQL** -> **Novo Script SQL**).
6. Cole o seguinte código SQL para criar a tabela de pacientes e pressione `Alt + X` (ou clique no botão laranja de executar script):

```sql
CREATE TABLE patients (
    id SERIAL PRIMARY KEY,
    full_name VARCHAR(150) NOT NULL,
    cpf VARCHAR(14) NOT NULL UNIQUE,
    birth_date DATE NOT NULL,
    phone VARCHAR(20),
    diagnosis TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### 5. Instalar Dependências e Executar o Servidor Go
Com o banco de dados rodando e a tabela criada, instale os pacotes Go e inicie o servidor localmente:

```bash
# Atualiza os módulos e instala as dependências
go mod tidy

# Inicia o servidor local
go run app/main.go
```

Você verá a seguinte mensagem de sucesso no console:
```text
Servidor do Hospital Teste iniciado com sucesso em http://localhost:8080
```

### 6. Acessar o Sistema pelo Navegador
Abra o seu navegador web favorito e acesse:
👉 **[http://localhost:8080](http://localhost:8080)**

Pronto! Você pode usar a interface moderna para cadastrar, buscar, atualizar e deletar pacientes.

---

## 📁 Estrutura de Pastas Implementada

```text
app/
  main.go                      # Inicialização do servidor, roteamento e arquivos estáticos
  handlers/
    createPatientHandler.go    # Recebe POST /criar e insere paciente
    getPatientHandler.go       # Recebe GET /buscar e renderiza dados do paciente
    updatePatientHandler.go    # Recebe POST /atualizar e atualiza dados
    deletePatientHandler.go    # Recebe POST /deletar e remove paciente
  utils/
    connectToDB.go             # Estabelece conexão com o banco de dados PostgreSQL
    createPatient.go           # Função SQL de inserção
    getPatient.go              # Função SQL de busca e estruturação
    updatePatient.go           # Função SQL de atualização
    deletePatient.go           # Função SQL de deleção
static/
  index.html                   # Página inicial com cards de atalho das operações
  forms/
    createPatient.html         # Formulário moderno de cadastro
    getPatient.html            # Formulário de pesquisa
    updatePatient.html         # Formulário de edição
    deletePatient.html         # Formulário de exclusão
  styles/
    index.style.css            # Estilo premium da página inicial
    createPatient.style.css    # Estilo dos inputs e alertas do formulário de criação
    getPatient.style.css       # Estilo da busca e da ficha de dados retornada
    updatePatient.style.css    # Estilo do formulário de atualização
    deletePatient.style.css    # Estilo do formulário de exclusão com alerta vermelho
.env                           # Configurações reais do banco de dados local
.env.example                   # Modelo das variáveis de ambiente exigido
docker-compose.yml             # Arquivo real para subir o PostgreSQL instantaneamente
docker-compose.yml.example     # Modelo Docker Compose exigido
go.mod                         # Definição do módulo hospital-teste e dependências
go.sum                         # Checksums das dependências Go tidied
README.md                      # Este guia completo de instruções em português
```

---

## 🎨 Detalhes Visuais e de Usabilidade
- **Paleta de Cores**: Fundo moderno `#f5f5f5`, textos escuros legíveis `#1a1a1a`, destaque em azul hospitalar `#2a7ae2`.
- **Campos**: Bordas suaves arredondadas de `8px` com preenchimento generoso (`padding`), micro-transições interativas ao focar e botões com transição hover suave.
- **Mensagens dinâmicas**: Banners verdes de sucesso ou vermelhos de erro aparecem no topo dos formulários baseados nas ações efetuadas, mantendo o usuário informado em tempo real.
