# Contribuindo para o Go File Processor

Obrigado por seu interesse em contribuir! Este projeto segue padrões rigorosos de qualidade e concorrência em Go.

## Setup de Desenvolvimento

1.  **Requisitos**: Go 1.22+ e `make`.
2.  **Clone**: `git clone https://github.com/ESousa97/go-file-processor.git`.
3.  **Testes**: Use `make test` para garantir que tudo está ok.

## Convenções de Código

- Siga o [Effective Go](https://golang.org/doc/effective_go.html).
- Rode `go fmt` antes de cada commit.
- Todos os itens exportados devem ter Godoc comments profissional em Português.
- Mantenha a modularização extrema: cada arquivo com uma responsabilidade única.

## Processo de Pull Request

1.  Crie uma branch descritiva (`feature/`, `fix/`, `perf/`).
2.  Garanta que os benchmarks não regrediram via `make bench`.
3.  Atualize o `CHANGELOG.md` na seção `[Unreleased]`.
4.  Solicite revisão de código.

## Áreas de Contribuição

- Suporte a novos formatos (XML, Avro).
- Otimização do Consumer para reduzir ainda mais o overhead de serialização.
- Melhorias na CLI (ex: progress bar mais detalhada).
