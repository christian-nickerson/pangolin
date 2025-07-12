# Pangolin Agent Guidelines

## Build/Test Commands
- **Python**: `poetry run pytest` (tests in `models/tests/`), single test: `poetry run pytest models/tests/unit/test_specific.py::test_function`
- **Go**: `go test ./...` (control plane), single test: `go test ./control/internal/routes/health -run TestSpecific`
- **Rust**: `cargo test` (nodes), single test: `cargo test test_name`
- **Lint**: `poetry run ruff check` (Python), `golangci-lint run` (Go), `cargo clippy` (Rust)
- **Format**: `poetry run black .` (Python), `gofmt -w .` (Go), `cargo fmt` (Rust)
- **Proto**: `make proto` to regenerate protobuf files

## Code Style
- **Python**: Line length 119, double quotes, 4-space indent. Use type hints, structured logging, dynaconf for config
- **Go**: Standard gofmt, use charmbracelet/log, fiber framework, viper config, validator tags
- **Rust**: Standard rustfmt, use log crate, serde for serialization, config crate for settings
- **Imports**: Group stdlib, third-party, local (Go/Python). Use relative imports in proto generated code
- **Naming**: snake_case (Python/Rust), camelCase (Go), PascalCase for types/structs
- **Errors**: Return errors, don't panic. Use structured logging with context
- **Tests**: Use pytest (Python), testify (Go), built-in testing (Rust). Place in dedicated test directories

## Project Structure
Multi-language monorepo: `control/` (Go API), `models/` (Python gRPC), `nodes/` (Rust), `proto/` (shared protobuf)
