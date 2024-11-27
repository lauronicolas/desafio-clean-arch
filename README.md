
# Order System

Este é um projeto que utiliza **Docker**, **migrations** e **Go** para gerenciar um sistema de pedidos. Siga os passos abaixo para configurar e rodar o projeto.

## Configuração e Execução

### 1. Iniciar os containers com Docker
Certifique-se de ter o Docker instalado e rodando em sua máquina. Depois, execute o seguinte comando para subir os serviços:

```bash
docker-compose up -d
```

### 2. Rodar as migrations
Execute as migrations usando o comando abaixo. Certifique-se de ajustar o caminho (`/path/to/migrations`) conforme necessário.

```bash
migrate -path=/path/to/migrations -database "mysql://root:root@tcp(localhost:3306)/orders" -verbose up
```

### 3. Navegar até o diretório do sistema
Acesse o diretório do sistema de pedidos:

```bash
cd cmd/ordersystem
```

### 4. Executar o sistema
Para rodar o sistema, use o seguinte comando:

```bash
go run main.go wire_gen.go
```

## Portas das Aplicações

O projeto utiliza as seguintes portas para as aplicações:

- **REST API:** `8000`
- **GraphQL:** `8080`
- **gRPC:** `50051`

## Pré-requisitos

Certifique-se de que você tem os seguintes pré-requisitos instalados:

- [Docker](https://www.docker.com/)
- [Go](https://go.dev/)
- [Migrate CLI](https://github.com/golang-migrate/migrate)

---

Sinta-se à vontade para abrir issues ou contribuir com melhorias para o projeto!
