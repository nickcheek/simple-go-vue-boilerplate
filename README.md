# Go Vue Boilerplate

```
git clone https://github.com/yourusername/go-vue-boilerplate.git
cd go-vue-boilerplate
make install  # Install dependencies
make dev      # Start development
```

## What's Included

- **Backend**: Go with Gin framework, single `/api/welcome` endpoint
- **Frontend**: Vue 3 + TypeScript, calls the backend API on homepage
- **Development**: Hot reload for both frontend and backend

## Endpoints

- `GET /api/welcome` - Returns welcome message, version, and status

The frontend displays this data on the homepage with error handling and loading states.
