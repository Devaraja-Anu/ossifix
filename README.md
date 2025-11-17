# ossifix

Lightweight Go scaffolding CLI and templates for web/server projects.

ossifix provides a small CLI-based project scaffold and template collection for several Go web frameworks. It includes generator logic, reusable templates, and small CLI UI helpers useful when building or bootstrapping web services.

**Quick facts:**

- **Language:** Go
- **Platform:** Cross-platform (examples use Bash; works on Windows with Git Bash / WSL)
- **Module:** Go modules (see `go.mod`)

## Quickstart

Build the tool (from the repository root):

```bash
# build a binary named `ossifix`
go build -o ossifix .
```

Run the CLI directly during development:

```bash
go run ./cmd
# or run the top-level main package
go run main.go
```

Run tests (when present):

```bash
go test ./...
```

Notes: on Windows use Git Bash, WSL, or another POSIX-like shell for the same commands.

## Purpose

The repo is a small scaffolding toolkit that can generate starter projects and server code using templates for multiple Go frameworks. It contains:

- generator code (in `internal/scaffold`)
- framework templates (in `templates/` and `fullstack/`)
- example/sub-project (`squish/`) with its own module
- small CLI UI helpers under `ui/`

## Repository layout

- `main.go` — optional top-level runner for quick iteration
- `cmd/` — CLI command implementations (e.g., `init`, `root`)
- `internal/` — internal packages not intended for external use
  - `models/` — domain models (e.g., `project.go`)
  - `scaffold/` — generator and file-creation helpers
- `templates/` — scaffolding templates
- `fullstack/` — framework-specific server templates (Echo, Gin, Fiber, Chi, ...)
- `squish/` — sub-project example with its own `go.mod`
- `ui/` — small interactive helpers (`selector`, `spinner`, `textInput`)

## Usage notes

- Inspect `cmd/` to see the available CLI commands and flags.
- Templates in `templates/` and `fullstack/` are meant to be processed/copied by the scaffold generator.
- `squish/` can be used as a runnable example or starting point for an API server.

## Contributing

Contributions are welcome. Recommended workflow:

1. Fork the repository and create a branch.
2. Run `go test ./...` and linting tools locally.
3. Open a pull request with a clear description of your changes.

If you'd like starter issues (tests, docs, example flows), open an issue and I'll help scope it.

## License

This repository is licensed under the MIT License — see the `LICENSE` file.
