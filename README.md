# React in Go

A small React frontend served by Go Fiber, with the API deployed as a Vercel Serverless Function.

## Local development

Clone the project:

```bash
git clone https://github.com/wyphyoez/react-in-go.git
cd react-in-go
```

Install dependencies:

```bash
npm ci
go mod download
```

Build the frontend:

```bash
npm run build
```

Run the traditional local Fiber server:

```bash
npm start
```

The local server is available at `http://localhost:3000`.

## API endpoints

The Vercel Function is defined in `api/index.go` and is available under `/api` after deployment.

| Endpoint | Description |
|---|---|
| `GET /api` | Returns the service status |
| `GET /api/health` | Returns `{ "status": "ok" }` |

## Deploy to Vercel

Import this repository into Vercel. Vercel detects the root `go.mod`, builds the Go function under `api/`, and uses `vercel.json` for the frontend build and API routing.

The build command is:

```bash
npm run build
```

The generated static frontend is placed in `public/`, while requests to `/api/*` are handled by the Go Fiber serverless function.

You can also deploy with the Vercel CLI:

```bash
npm install --global vercel
vercel login
vercel
```
