# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.0.0] - 2026-03-12

### Added
- Initial implementation of `CSVToJSONProcessor` with Worker Pool.
- Transformation layer (Middleware) with Email, Role, and Masking filters.
- Atomic metrics system for progress tracking.
- Benchmark suite comparing sequential vs. parallel performance.
- Automation via `Makefile` (build, test, bench, generate-data).
- Full professional documentation (Godoc, README, CONTRIBUTING, etc.).

### Changed
- Refactored processing logic for better resilience against line errors.

### Fixed
- Linting and spelling warnings reported by CSpell and gopls.

---
[1.0.0]: https://github.com/ESousa97/go-file-processor/releases/tag/v1.0.0
