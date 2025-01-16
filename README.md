# Projeto Final API Go

Este projeto é uma API RESTful construída com Go e o framework Gin. Ele fornece funcionalidades de gerenciamento de usuários, como registrar, atualizar, listar e remover usuários.

## Começando

### Pré-requisitos

- Docker
- Docker Compose

### Instalação

1. Clone o repositório:
    ```sh
    git clone https://github.com/hpaes/go-api-final-project.git
    cd go-api-final-project
    ```

2. Inicie a aplicação usando Docker Compose:
    ```sh
    docker-compose up --build
    ```

### Variáveis de Ambiente

A aplicação usa as seguintes variáveis de ambiente, que são definidas no arquivo `docker-compose.yml`:

- `APPLICATION_NAME`: Nome da aplicação
- `DB_HOST`: Host do banco de dados
- `DB_PORT`: Porta do banco de dados
- `DB_USER`: Usuário do banco de dados
- `DB_PASSWORD`: Senha do banco de dados
- `DB_NAME`: Nome do banco de dados
- `SERVER_TIMEOUT`: Timeout do servidor
- `SERVER_PORT`: Porta do servidor

### Endpoints da API

Os seguintes endpoints estão disponíveis na API:

- **Registrar Usuário**
    - **URL**: `/user`
    - **Método**: `POST`
    - **Descrição**: Registra um novo usuário.
    - **Corpo da Requisição**:
        ```json
        {
            "name": "string",
            "email": "string",
            "password": "string"
        }
        ```

- **Obter Detalhes do Usuário**
    - **URL**: `/user/:userId`
    - **Método**: `GET`
    - **Descrição**: Recupera detalhes de um usuário específico.
    - **Parâmetros da URL**:
        - `userId`: ID do usuário

- **Remover Usuário**
    - **URL**: `/user/:userId`
    - **Método**: `DELETE`
    - **Descrição**: Remove um usuário específico.
    - **Parâmetros da URL**:
        - `userId`: ID do usuário

- **Listar Usuários**
    - **URL**: `/users`
    - **Método**: `GET`
    - **Descrição**: Lista todos os usuários.
    - **Parâmetros de Consulta**:
        - `page`: Número da página para paginação

- **Atualizar Usuário**
    - **URL**: `/user`
    - **Método**: `PUT`
    - **Descrição**: Atualiza as informações do usuário.
    - **Corpo da Requisição**:
        ```json
        {
            "userId": "string",
            "name": "string",
            "email": "string",
            "password": "string"
        }
        ```

### Executando Testes

Para executar os testes, use o seguinte comando:
```sh
go test ./...
```

### Licença

Este projeto está licenciado sob a Licença MIT.
