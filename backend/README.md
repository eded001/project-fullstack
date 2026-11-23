# jav-back - Go REST API

API REST mínima em Go, sem banco de dados, com um recurso `Item` armazenado em memória (os dados são perdidos ao reiniciar o servidor).

## Stack

- **Framework**: Echo (`github.com/labstack/echo/v5`)
- **Armazenamento**: map em memória protegido por `sync.RWMutex`
- **Porta**: `:3000`

## Requisitos

- Go 1.25+

## Estrutura

```shell
main.go       # Rotas, handlers e armazenamento em memória
go.mod        # Módulo e dependências
go.sum        # Checksums das dependências
.gitignore    # Arquivos ignorados pelo git
README.md     # Documentação
```

## Como rodar

```bash
cd backend
go run .
# Server runs on http://localhost:3000
```

Para gerar um binário:

```bash
go build -o api .
./api
```

## Endpoints

| Method | Endpoint     | Description          | Sucesso |
| ------ | ------------ | -------------------- | ------- |
| GET    | `/health`    | Health check         | `200`   |
| GET    | `/items`     | Lista todos os itens | `200`   |
| GET    | `/items/:id` | Busca item por ID    | `200`   |
| POST   | `/items`     | Cria item            | `201`   |
| PUT    | `/items/:id` | Atualiza item        | `200`   |
| DELETE | `/items/:id` | Remove item          | `204`   |

## Exemplos

```bash
# Health check
curl http://localhost:3000/health

# Criar item
curl -X POST http://localhost:3000/items \
  -H "Content-Type: application/json" \
  -d '{"name": "meu item"}'

# Listar itens
curl http://localhost:3000/items

# Buscar item
curl http://localhost:3000/items/1

# Atualizar item
curl -X PUT http://localhost:3000/items/1 \
  -H "Content-Type: application/json" \
  -d '{"name": "item atualizado"}'

# Remover item
curl -X DELETE http://localhost:3000/items/1
```

## Formato

Item:

```json
{ "id": 1, "name": "meu item" }
```

Erros (`400` id/nome inválido, `404` item não encontrado):

```json
{ "error": "mensagem" }
```
