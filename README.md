<div align="center">
  <h1>Go File Processor</h1>
  <p>Parallel and resilient processing of massive files with Worker Pool in Go.</p>

  <img src="assets/github-go.png" alt="Go File Processor Banner" width="600px">

  <br>

[![CI](https://github.com/ESousa97/go-file-processor/actions/workflows/ci.yml/badge.svg)](https://github.com/ESousa97/go-file-processor/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/ESousa97/go-file-processor)](https://goreportcard.com/report/github.com/ESousa97/go-file-processor)
[![CodeFactor](https://www.codefactor.io/repository/github/ESousa97/go-file-processor/badge)](https://www.codefactor.io/repository/github/ESousa97/go-file-processor)
[![Go Reference](https://pkg.go.dev/badge/github.com/ESousa97/go-file-processor.svg)](https://pkg.go.dev/github.com/ESousa97/go-file-processor)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Version](https://img.shields.io/github/go-mod/go-version/ESousa97/go-file-processor)](https://github.com/ESousa97/go-file-processor)
[![Last Commit](https://img.shields.io/github/last-commit/ESousa97/go-file-processor)](https://github.com/ESousa97/go-file-processor/commits/main)

</div>

---

**Go File Processor** is a high-performance command-line tool and library designed to efficiently convert massive CSV files (millions of records) into structured JSON. Using the Worker Pool pattern and channel-based processing, it ensures optimized CPU usage and constant memory consumption, regardless of the input file size.

## Demonstration

### As a Library

Add transformers and configure the execution pool fluently:

```go
proc := processor.NewCSVToJSONProcessor()
config := processor.Config{WorkerCount: 8}

// Add transformers (Chain of Responsibility)
config.AddTransformer(processor.EmailFilter(`@company.com$`))
config.AddTransformer(processor.FieldMasker("email"))

metrics, err := proc.Process("input.csv", "output.json", config)
```

### As a CLI

Run massive processing with real-time metrics:

```bash
./fileproc -input data.csv -output data.json -workers 4
```

Output:

```text
[INFO] Starting processing...
[INFO] Progress: 100000 rows processed
[SUMMARY] EXECUTION COMPLETED IN 1.2s
- Total lines read: 100000
- Successfully processed: 98500
- Errors/Ignored: 1500
```

## Tech Stack

| Technology          | Role                                                                |
| ------------------- | ------------------------------------------------------------------- |
| **Go 1.22+**        | Core language with high-performance native concurrency              |
| **Worker Pool**     | Parallelism management and load control                             |
| **slog**            | Structured logging for observability and traceability               |
| **Atomic Counters** | High-performance metrics collection without contention (lock-free)  |
| **Channels**        | Secure and decoupled communication between Producer, Workers, and Consumer |

## Prerequisites

- **Go >= 1.22**
- **Make** (for build automation and benchmarks)

## Installation and Usage

### From Source

```bash
git clone https://github.com/ESousa97/go-file-processor.git
cd go-file-processor
make build
```

### Data Generation and Benchmark

To validate performance with 100k+ row files:

```bash
make generate-data
make bench
```

## Makefile Targets

| Target               | Description                                               |
| -------------------- | --------------------------------------------------------- |
| `make build`         | Compiles the `fileproc` binary at the project root        |
| `make test`          | Runs the unit test suite                                  |
| `make bench`         | Runs performance comparisons (Sequential vs Parallel)     |
| `make generate-data` | Generates a massive test file (100,000 records)           |
| `make clean`         | Removes binaries and temporary files                      |

## Architecture

The project uses a channel-based streaming model to process data without loading the entire file into memory.

```mermaid
graph LR
    Input[CSV Input] --> Producer[Producer]
    Producer --> Jobs{Job Channel}
    Jobs --> W1[Worker 1]
    Jobs --> W2[Worker 2]
    Jobs --> WN[Worker N]
    W1 & W2 & WN --> Transformers[Transformation Layer]
    Transformers --> Results{Result Channel}
    Results --> Consumer[Consumer]
    Consumer --> Output[JSON Output]

    subgraph "Worker Pool"
    W1
    W2
    WN
    end
```

## API Reference

Detailed technical documentation available at [pkg.go.dev/github.com/ESousa97/go-file-processor](https://pkg.go.dev/github.com/ESousa97/go-file-processor).

## Configuration (CLI Flags)

| Flag       | Description                       | Type     | Default       |
| ---------- | --------------------------------- | -------- | ------------- |
| `-input`   | Input CSV file path               | `string` | `input.csv`   |
| `-output`  | Output JSON file path             | `string` | `output.json` |
| `-workers` | Number of concurrent workers       | `int`    | `4`           |

## Roadmap

Follow the project's evolution stages:

- [x] **Phase 1: Foundation** — Worker Pool and streaming core implementation.
- [x] **Phase 2: Transformation** — Middleware layer (Chain of Responsibility).
- [x] **Phase 3: Observability** — Atomic metrics and structured logs (`slog`).
- [x] **Phase 4: Governance** — CI/CD, Professional documentation, and Badges.

## Contributing

Interested in collaborating? Check our [CONTRIBUTING.md](CONTRIBUTING.md) for code standards and PR process.

## License

This project is licensed under the **MIT License** — see the [LICENSE](LICENSE) file for details.

<div align="center">

## Author

**Enoque Sousa**

[![LinkedIn](https://img.shields.io/badge/LinkedIn-0077B5?style=flat&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/enoque-sousa-bb89aa168/)
[![GitHub](https://img.shields.io/badge/GitHub-100000?style=flat&logo=github&logoColor=white)](https://github.com/ESousa97)
[![Portfolio](https://img.shields.io/badge/Portfolio-FF5722?style=flat&logo=target&logoColor=white)](https://enoquesousa.vercel.app)

**[⬆ Back to top](#go-file-processor)**

Made with ❤️ by [Enoque Sousa](https://github.com/ESousa97)

**Project Status:** Active — Constantly updated

</div>
