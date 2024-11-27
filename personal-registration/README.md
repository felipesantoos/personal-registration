
Cadastro Pessoal - Aplicação Multicontainer
Este projeto é uma aplicação multicontainer baseada em Go para cadastro pessoal, utilizando Docker e PostgreSQL.

Pré-requisitos
Antes de começar, certifique-se de ter o seguinte instalado em sua máquina:

Docker
Docker Compose
Estrutura do Projeto
.docker/: Contém o Dockerfile e outras configurações para construir o container da API.
.env: Arquivo de ambiente para armazenar informações sensíveis, como credenciais do banco de dados.
main.go: Ponto de entrada da aplicação Go.
src/: Código-fonte da aplicação, incluindo conexão e lógica do banco de dados.
Configuração e Execução
Clone o Repositório:

bash
Copy code
git clone <repository-url>
cd personal-registration
Configure as Variáveis de Ambiente:

Atualize o arquivo .env com os seguintes valores:

env
Copy code
DB_HOST=database
DB_PORT=5432
DB_USER=seu_usuario
DB_PASSWORD=sua_senha
DB_NAME=nome_do_seu_banco
Construa e Execute os Containers:

Use o Docker Compose para construir e iniciar os containers:

bash
Copy code
docker-compose up --build
Este comando irá:

Construir o container da API a partir do Dockerfile localizado em .docker/Dockerfile.api.
Iniciar um container PostgreSQL com o banco de dados inicializado a partir de src/connection/init.sql.
Acesse a API:

A API estará acessível em http://localhost:8000.

Solução de Problemas
Porta Já em Uso: Se encontrar um erro de alocação de porta, certifique-se de que as portas 5432 (PostgreSQL) e 8000 (API) não estão sendo usadas por outras aplicações.
Problemas de Conexão com o Banco de Dados: Verifique se as variáveis de ambiente no arquivo .env correspondem à configuração do banco de dados.
Desenvolvimento
Para executar a aplicação Go localmente sem Docker:

Configure uma instância local do PostgreSQL e atualize o arquivo .env conforme necessário.

Execute a aplicação utilizando:

bash
Copy code
go run main.go