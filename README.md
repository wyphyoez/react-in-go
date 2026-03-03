# React in Go (v2)

A tiny full-stack starter that uses:

- **Go + Fiber** for backend and HTML rendering
- **React 18** for frontend UI
- **esbuild** for bundling TSX/CSS assets

## Run locally

Clone the project:

```bash
git clone https://github.com/waiyanphioe/react-in-go.git
cd react-in-go
```

Install dependencies:

```bash
npm install
```

Start app (builds frontend and runs Go server):

```bash
npm start
```

Open: `http://localhost:3000`

## v2 highlights

- Added `GET /api/v2/message` endpoint
- React app now fetches live data from backend API
- Upgraded UI with a small status dashboard
