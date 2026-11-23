# system-jav

Monorepo com frontend Next.js (Tailwind CSS) e backend Go (Echo).

## Estrutura

- `frontend/` — Next.js 16 + Tailwind CSS 4
- `backend/` — API em Go com Echo

## Rodando

```bash
# frontend
cd frontend
npm install
npm run dev

# backend
cd backend
go run .
```

O frontend e a API usam a porta `3000` por padrão, então rode os dois em portas diferentes (ex.: `npm run dev -- -p 3001` no frontend).
