# go-layered — Modular REST API Boilerplate (Go)

![Go](https://img.shields.io/badge/Go-1.22+-00ADD8?logo=go) ![PostgreSQL](https://img.shields.io/badge/PostgreSQL-supported-4169E1?logo=postgresql)

Structured starter kit for building modular, production-ready REST APIs with Go. The project promotes clear separation of concerns so you can iterate quickly and scale safely.

## Highlights

- Modular architecture: each feature lives under `internal/features/<module>`
- Fast HTTP stack using Fiber
- DB support via `sqlx` and `pq` (Postgres)
- JSON encoding via `go-json`
- Request validation using `go-playground/validator`
- Generator scripts: `scripts/module.go` and `scripts/submodule.go` to bootstrap features

## Quickstart

1. Clone the repo and update module path in `go.mod` if you will publish under a new module name.
2. Copy the environment example and edit values:
   ```bash
   cp .env.example .env
   # edit .env
   ```
3. Run the app locally:
   ```bash
   go run cmd/app/main.go
   ```

The server will start on the port set in your `.env` (default 8080).

## Sample module: todo

The `todo` module is included to show a minimal, complete feature integration.

- Location: `internal/features/todo`
- Routes (registered at `/v1/todo`):
  - POST /v1/todo — Create a task
  - GET /v1/todo — List tasks
  - GET /v1/todo/:id — Get a single task
  - PUT /v1/todo/:id — Update a task
  - DELETE /v1/todo/:id — Delete a task

Example create request:

```bash
curl -X POST http://localhost:8080/v1/todo \
  -H 'Content-Type: application/json' \
  -d '{"title":"Buy milk"}'
```

The module uses an in-memory repository at `internal/features/todo/internal/repository/inmemory` for demonstration. Swap in a Postgres implementation when needed (templates and examples are present in `scripts/templates`).

## Project structure (key folders)

```
cmd/app/main.go               # Application entrypoint
internal/app                   # Core app config, provider, exception handling
internal/features/<module>      # Each feature module lives here
scripts/                       # Codegen templates and scripts
```

## Generators

- Create a module: `go run scripts/module.go <module>`
- Create a submodule: `go run scripts/submodule.go <module> <submodule>`

These scripts scaffold the module folders and boilerplate files using templates in `scripts/templates`.

## Testing & Development notes

- The sample `todo` module is intended for local development and demonstration. For production use, implement a persistent repository and add proper migrations.
- Keep `.env` values and secrets out of source control.
- Update `go.mod` module path to match your repository before publishing.

## Contributing

Contributions are welcome — open issues or PRs. When contributing back:

- Keep the original license header and copyright lines intact.
- Document breaking changes in PR descriptions.

## License

This repository is distributed under the MIT License. See `LICENSE` for details.

---
