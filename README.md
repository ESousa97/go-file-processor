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
[![Last Commit](https://img.shields.io/github/last-commit/ESousa97/go-file-processor)](https://github.com/ESousa97/go-file-processor)

</div>

---

> **Note: Archival Project**  
> This was my second major project in Go, built as a deep dive into the language's idiomatic concurrency patterns and high-performance I/O. It is now archived but serves as a solid reference for ETL (Extract, Transform, Load) implementations in Golang.

**Go File Processor** is a high-performance command-line tool and library designed to efficiently convert massive CSV files (millions of records) into structured JSON. It demonstrates the power of Go's concurrency primitives to achieve maximum throughput with minimal memory overhead.

## 🚀 Core Learning Objectives

This project was a hands-on laboratory to master several Go concepts:

*   **Concurrency via Worker Pool:** Leveraging `goroutines` and `channels` to process data in parallel without overwhelming the system.
*   **Memory Efficiency (Streaming):** Using `io.Reader` and `io.Writer` to process gigabytes of data with a constant, tiny memory footprint.
*   **The Middleware Pattern:** Implementing a "Chain of Responsibility" for data transformation that is both flexible and type-safe.
*   **Atomic Operations:** Using `sync/atomic` for high-speed metrics tracking, avoiding the overhead of mutexes.
*   **Idiomatic Project Layout:** Following standard Go folder structures (`cmd/`, `internal/`) and build automation with `Makefile`.

## Demonstration

### As a Library

```go
proc := processor.NewCSVToJSONProcessor()
config := processor.Config{WorkerCount: 8}

// Fluent transformation chain
config.AddTransformer(processor.EmailFilter(`@company.com$`))
config.AddTransformer(processor.FieldMasker("email"))

metrics, err := proc.Process("input.csv", "output.json", config)
```

### As a CLI

```bash
./fileproc -input data.csv -output data.json -workers 4
```

## Tech Stack & Architecture

| Technology          | What I Learned                                                      |
| ------------------- | ------------------------------------------------------------------- |
| **Worker Pool**     | How to orchestrate multiple goroutines for parallel work.           |
| **Channels**        | Managing safe communication and backpressure between stages.        |
| **Streaming I/O**   | Processing files record-by-record instead of loading to RAM.        |
| **Atomic Counters** | Implementing thread-safe counters with maximum performance.         |
| **Structured Logs** | Using `slog` for modern, machine-readable observability.            |

### Pipeline Flow

The system uses a streaming model to maintain low memory usage:
`Input CSV -> Producer -> Job Channel -> [Workers + Transformers] -> Result Channel -> Consumer -> Output JSON`

## Makefile Targets

| Target               | Description                                               |
| -------------------- | --------------------------------------------------------- |
| `make build`         | Compiles the `fileproc` binary.                           |
| `make test`          | Runs the full unit test suite.                            |
| `make bench`         | Runs benchmarks to see the speed of Parallel vs Sequential. |
| `make generate-data` | Generates a 100k row test file for performance testing.   |

## 📚 Final Thoughts

Building this project taught me that Go isn't just about syntax; it's about a philosophy of simplicity and performance. The transition from sequential processing to a parallel worker pool showed me how Go empowers developers to build tools that scale effortlessly.

---

<div align="center">

## Author

**Enoque Sousa**

[![LinkedIn](https://img.shields.io/badge/LinkedIn-0077B5?style=flat&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/enoque-sousa-bb89aa168/)
[![GitHub](https://img.shields.io/badge/GitHub-100000?style=flat&logo=github&logoColor=white)](https://github.com/ESousa97)
[![Portfolio](https://img.shields.io/badge/Portfolio-FF5722?style=flat&logo=target&logoColor=white)](https://enoquesousa.vercel.app)

**[⬆ Back to top](#go-file-processor)**

Made with ❤️ by [Enoque Sousa](https://github.com/ESousa97)

**Project Status:** Archived — Educational Milestone

</div>
