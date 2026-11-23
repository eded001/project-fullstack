# jav-back

API mínima em Go com Echo, sem banco de dados: uma lista de itens mantida em memória (os dados são perdidos ao reiniciar o servidor).

## Stack

- **Framework**: Echo (`github.com/labstack/echo/v5`)
- **Armazenamento**: slice em memória protegido por `sync.RWMutex`
- **Porta**: `:3000`

## Como rodar

```bash
cd backend
go run .
```

Para gerar um binário:

```bash
go build -o api .
./api
```

## Endpoints

| Método | Endpoint                | Descrição                     |
| ------ | ----------------------- | ----------------------------- |
| GET    | `/`                     | Hello World                   |
| GET    | `/list`                 | Lista todos os itens          |
| POST   | `/list/add/:item`       | Adiciona um item              |
| DELETE | `/list/delete/:id`      | Remove o item pelo índice     |

## Exemplos

```bash
curl http://localhost:3000/

curl http://localhost:3000/list

curl -X POST http://localhost:3000/list/add/maçã

curl -X DELETE http://localhost:3000/list/delete/0
```

## Formato

Sucesso:

```json
{ "status": "ok", "list": ["maçã"] }
```

Erro (`400` id inválido, `404` índice inexistente):

```json
{ "status": "error", "message": "ID inválido" }
```
