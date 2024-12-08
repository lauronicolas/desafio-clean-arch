
# Order System

Este é um projeto que utiliza **Docker**, **migrations** e **Go** para gerenciar um sistema de pedidos. Siga os passos abaixo para configurar e rodar o projeto.

## Configuração e Execução

### 1. Iniciar os containers com Docker
Certifique-se de ter o Docker instalado e rodando em sua máquina. Depois, execute o seguinte comando para subir os serviços:

```bash
docker-compose up --build -d
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
