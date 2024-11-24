# API de Produtos em Golang

Uma API REST para gerenciar produtos, desenvolvida em Golang com o framework Gin e utilizando MySQL como banco de dados.

## Requisitos

Antes de começar, certifique-se de ter os seguintes itens instalados:

Golang (versão 1.23.3 ou superior)
MySQL (versão 8 ou superior)
Git
Ferramenta para gerenciar o banco (opcional): MySQL Workbench ou phpMyAdmin
Configuração do Banco de Dados
Inicie o serviço MySQL no seu computador.

Acesse o MySQL pelo terminal ou por uma ferramenta gráfica (como Workbench ou phpMyAdmin).

Execute o seguinte comando no MySQL para criar o banco de dados:

## Instalação

```bash
CREATE DATABASE produtos_db;
```

## Configuração do Projeto

```bash

Clone este repositório:
git clone https://github.com/andersonfae/ramada-go
cd api-produtos
```

## Inicialize o módulo Go e instale as dependências:

```bash
go mod tidy
Verifique e atualize as configurações de conexão ao banco de dados no arquivo database/connection.go:

```

## BANCO DE DADOS

```bash
dsn := "root:senha_do_mysql@tcp(127.0.0.1:3306)/produtos_db?charset=utf8mb4&parseTime=True&loc=Local"
Substitua root pelo seu usuário do MySQL.
Substitua senha_do_mysql pela senha do seu MySQL.
```

## Execução

Execute a aplicação:

```bash
go run main.go
Se tudo estiver correto, você verá a seguinte mensagem no terminal:

[GIN-debug] Listening and serving HTTP on :8080
Acesse os endpoints da API em http://localhost:8080.
```

## Endpoints Disponíveis

```bash

GET /produtos Lista todos os produtos com paginação.
GET /produtos/{id} Retorna os detalhes de um produto específico.
GET /produtos?nome={nome}&categoria={categoria} Busca produtos por nome, categoria ou faixa de preço.
POST /produtos Adiciona um novo produto.
POST /produtos/importar Importa produtos em massa via JSON ou CSV.
PUT /produtos/{id} Atualiza os dados de um produto.
DELETE /produtos/{id} Exclui um produto.
```
#   r a m a d a - g o 
 
 
