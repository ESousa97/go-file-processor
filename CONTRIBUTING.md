# Contributing (Archived Project)

Thank you for your interest! This project is currently **archived** and no longer accepting new features or active maintenance.

## Project Purpose
The **Go File Processor** was created as a second learning project to explore Go's concurrency and streaming I/O. It remains available as a historical reference for:
- Worker Pool implementations.
- Channel-based pipelines.
- Middleware design patterns in Go.

## Exploring the Code
You are welcome to fork this project to use as a template or to experiment with its features. Key areas of interest:
- `internal/processor/csv_json.go`: The core engine using Worker Pools.
- `internal/processor/transformer.go`: The implementation of the Middleware pattern.
- `internal/processor/csv_json_bench_test.go`: Benchmarking logic to compare performance.

## License
The project remains under the **MIT License**, allowing you to use and modify it for your own purposes.
