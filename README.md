# go-layered — Modular REST API Boilerplate (Go)

![Go](https://img.shields.io/badge/Go-1.22+-00ADD8?logo=go) ![PostgreSQL](https://img.shields.io/badge/PostgreSQL-supported-4169E1?logo=postgresql)

A minimal, practical starting point for building modular REST APIs in Go.

## Key points

- Features are organized under `internal/modules/<name>` so each unit contains its own entity, models, repository, usecase, controller, and router.
- Uses Fiber for HTTP, `sqlx` + `pq` for Postgres, `go-json` for JSON, and `validator` for request validation.
- Includes small generator scripts to scaffold modules and submodules.

## Quickstart

1. Clone the repo and set your module path in `go.mod` if you plan to publish under a new path.
2. Copy the env example and edit values:

   ```bash
   cp .env.example .env
   # edit .env to set APP_PORT and DB_DSN
   ```

3. Ensure dependencies are up to date and `go.sum` is generated:

   ```bash
   go mod tidy
   ```

4. Run the app:

   ```bash
   go run cmd/app/main.go
   ```

   Server listens on the port configured in `.env` (defaults to 8080).

## Included example: todo

A basic `todo` feature is provided as a reference implementation under `internal/modules/todo`.

Routes (mounted at `/v1/todo`):

- POST /v1/todo — create a task
- GET /v1/todo — list tasks
- GET /v1/todo/:id — get task by id
- PUT /v1/todo/:id — update a task
- DELETE /v1/todo/:id — delete a task

The example uses an in-memory repository for easy local testing. Replace it with the Postgres implementation when you need persistence — templates for repository implementations are in `scripts/templates`.

Example request:

```bash
curl -X POST http://localhost:8080/v1/todo \
  -H 'Content-Type: application/json' \
  -d '{"title":"Buy milk"}'
```

## Project layout (important files)

```
cmd/app/main.go               # application bootstrap
internal/app/                  # app config, provider, exception handling
internal/modules/<feature>/    # each feature lives here
scripts/                       # scaffolding templates and scripts
go.mod, go.sum                 # module files
README.md, LICENSE             # repo docs and license
```

## Scaffolding

- Create a new feature: `go run scripts/module.go <name>`
- Create a submodule: `go run scripts/submodule.go <module> <submodule>`

These scripts generate folders and boilerplate based on templates in `scripts/templates`.

## Development notes

- The sample repository implementations are intentionally simple.
- For production work you should:
  - Implement a persistent repository (Postgres) and migrations
  - Add logging, metrics, and observability
  - Harden configs and secrets management
  - Add automated tests
- Keep sensitive values out of source control; use the `.env` or other secret managers.

## Contributing

Contributions are welcome. If you fork and publish changes, keep the original license text and attribution for existing code. You may add your own copyright line for new work.

## License

This project is released under the MIT License. See `LICENSE` for details.

---
