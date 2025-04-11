# 💼 GOpportunities

Uma API RESTful desenvolvida em Go para divulgação de vagas de emprego. Usuários podem cadastrar, listar, buscar e remover oportunidades de forma simples e rápida.

---
## 🚀 Funcionalidades

- 📋 Listar todas as vagas
- 🔍 Buscar uma vaga específica por ID
- ➕ Criar uma nova vaga
- ❌ Deletar uma vaga

---

## 🛠️ Tecnologias Utilizadas

- [Go](https://golang.org/)
- [Gin Gonic](https://github.com/gin-gonic/gin)
- [GORM](https://gorm.io/)
- [PostgreSQL](https://www.postgresql.org/)
- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

---

## 🐳 Como Rodar o Projeto

1. Clone o repositório:

   ```bash
   git clone https://github.com/pedrodese/gopportunities.git
   cd gopportunities

2. Suba os containers com Docker Compose:
   
    ```bash
    docker-compose up
    
4. Acesse a API no navegador ou via Postman:
    http://localhost:8080

📦 Instalação de Dependências
    instale as dependências do Go com:
    
    ```bash
    go mod tidy

⚙️ Configuração do Banco de Dados
O projeto já inclui um docker-compose.yml com o banco PostgreSQL configurado com os seguintes dados padrão:

    DB_HOST=localhost
    DB_PORT=5432
    DB_USER=postgres
    DB_PASSWORD=postgres
    DB_NAME=gopportunities
    Você pode alterar essas variáveis diretamente no docker-compose.yml ou configurar um .env com suporte no código.
