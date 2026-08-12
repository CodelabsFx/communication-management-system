# Communication System

Scaffolded workspace for a modular monolith communication platform.

## Structure

- `frontend/` - Vue 3 + TypeScript frontend application
- `backend/` - Go backend with REST, WebSocket, PostgreSQL, Redis, and S3-compatible storage support
- `docker-compose.yml` - local development services

## Quick start

1. Start services:
   ```bash
   docker compose up --build
   ```

2. Frontend:
   - `http://localhost:5173`

3. Backend:
   - `http://localhost:8080`
