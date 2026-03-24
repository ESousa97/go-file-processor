# Contributing to Go File Processor

Thank you for your interest in contributing! This project follows rigorous standards for quality and concurrency in Go.

## Development Setup

1.  **Requirements**: Go 1.22+ and `make`.
2.  **Clone**: `git clone https://github.com/ESousa97/go-file-processor.git`.
3.  **Tests**: Use `make test` to ensure everything is OK.

## Code Conventions

- Follow [Effective Go](https://golang.org/doc/effective_go.html).
- Run `go fmt` before each commit.
- All exported items must have professional Godoc comments in English.
- Maintain extreme modularization: each file with a single responsibility.

## Pull Request Process

1.  Create a descriptive branch (`feature/`, `fix/`, `perf/`).
2.  Ensure benchmarks haven't regressed via `make bench`.
3.  Update `CHANGELOG.md` in the `[Unreleased]` section.
4.  Request a code review.

## Areas for Contribution

- Support for new formats (XML, Avro).
- Consumer optimization to further reduce serialization overhead.
- CLI improvements (e.g., more detailed progress bar).
