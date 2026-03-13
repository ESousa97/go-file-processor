# Changelog

Todas as mudanças notáveis neste projeto serão documentadas neste arquivo.

O formato é baseado em [Keep a Changelog](https://keepachangelog.com/pt-BR/1.1.0/),
e este projeto segue [Versionamento Semântico](https://semver.org/lang/pt-BR/spec/v2.0.0.html).

## [1.0.0] - 2026-03-12

### Added
- Implementação inicial do `CSVToJSONProcessor` com Worker Pool.
- Camada de Transformação (Middleware) com filtros de Email, Role e Masking.
- Sistema de métricas atômicas para acompanhamento de progresso.
- Suíte de Benchmarks comparando performance sequencial vs paralela.
- Automação via `Makefile` (build, test, bench, generate-data).
- Documentação profissional completa (Godoc, README, CONTRIBUTING, etc.).

### Changed
- Refatoração da lógica de processamento para maior resiliência a erros de linha.

### Fixed
- Avisos de linting e ortografia reportados pelo CSpell e gopls.

---
[1.0.0]: https://github.com/ESousa97/go-file-processor/releases/tag/v1.0.0
