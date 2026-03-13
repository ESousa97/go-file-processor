# Go File Processor

> Processamento paralelo e resiliente de arquivos massivos com Worker Pool em Go.

[![CI](https://github.com/ESousa97/go-file-processor/actions/workflows/ci.yml/badge.svg)](https://github.com/ESousa97/go-file-processor/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/ESousa97/go-file-processor)](https://goreportcard.com/report/github.com/ESousa97/go-file-processor)
[![CodeFactor](https://www.codefactor.io/repository/github/ESousa97/go-file-processor/badge)](https://www.codefactor.io/repository/github/ESousa97/go-file-processor)
[![Go Reference](https://pkg.go.dev/badge/github.com/ESousa97/go-file-processor.svg)](https://pkg.go.dev/github.com/ESousa97/go-file-processor)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Version](https://img.shields.io/github/go-mod/go-version/ESousa97/go-file-processor)](https://github.com/ESousa97/go-file-processor)
[![Last Commit](https://img.shields.io/github/last-commit/ESousa97/go-file-processor)](https://github.com/ESousa97/go-file-processor/commits/main)

---

O **Go File Processor** é uma ferramenta de linha de comando e biblioteca de alto desempenho projetada para converter arquivos CSV massivos (milhões de registros) em JSON estruturado de forma eficiente. Utilizando o padrão Worker Pool e processamento via canais (channels), garante o uso otimizado de CPU e memória constante, independentemente do tamanho do arquivo de entrada.

## Demonstração

### Como Biblioteca

```go
proc := processor.NewCSVToJSONProcessor()
config := processor.Config{WorkerCount: 8}

// Adicione transformadores (Chain of Responsibility)
config.AddTransformer(processor.EmailFilter(`@company.com$`))
config.AddTransformer(processor.FieldMasker("email"))

metrics, err := proc.Process("input.csv", "output.json", config)
```

### Como CLI (Output)

```text
[INFO] Iniciando processamento...
[INFO] Progresso: 100000 linhas processadas
[RESUMO] EXECUÇÃO CONCLUÍDA EM 1.2s
- Total de linhas lidas: 100000
- Processadas com sucesso: 98500
- Erros/Ignoradas: 1500
```

## Stack Tecnológico

| Tecnologia          | Papel                                             |
| ------------------- | ------------------------------------------------- |
| **Go**              | Linguagem principal (Concorrência nativa)         |
| **Worker Pool**     | Gerenciamento de paralelismo e carga              |
| **slog**            | Structured logging para observabilidade           |
| **Atomic Counters** | Coleta de métricas sem lock                       |
| **Channels**        | Comunicação segura entre Producer/Worker/Consumer |

## Pré-requisitos

- Go >= 1.22
- Make (opcional, para automação)

## Instalação e Uso

### A partir do source

```bash
git clone https://github.com/ESousa97/go-file-processor.git
cd go-file-processor
make build
./fileproc -input data.csv -output data.json -workers 4
```

### Geração de Dados de Teste

Para testar performance com arquivos gigantes (100k+ linhas):

```bash
make generate-data
make bench
```

## Makefile Targets

| Target          | Descrição                                           |
| --------------- | --------------------------------------------------- |
| `build`         | Compila o binário `fileproc` na raiz do projeto     |
| `test`          | Executa todos os testes unitários                   |
| `bench`         | Roda a suíte de benchmarks (Sequencial vs Paralelo) |
| `generate-data` | Gera arquivo `large_test.csv` com 100.000 registros |
| `clean`         | Remove binários e arquivos temporários de teste     |

## Arquitetura

O projeto segue uma arquitetura modular focada em streaming de dados:

1.  **Producer**: Lê o arquivo CSV linha a linha (bufio) e despacha para o canal de jobs.
2.  **Worker Pool**: Conjunto de goroutines que consomem os jobs, aplicam transformações e validam tipos.
3.  **Consumer**: Coleta os resultados processados e serializa o JSON final via streaming.
4.  **Transformation Layer**: Padrão Middleware que permite injetar lógica de filtro e alteração em tempo de execução.

## API Reference

Veja a documentação completa e exemplos em [pkg.go.dev](https://pkg.go.dev/github.com/ESousa97/go-file-processor).

## Configuração (CLI Flags)

| Flag       | Descrição                         | Tipo   | Padrão        |
| ---------- | --------------------------------- | ------ | ------------- |
| `-input`   | Caminho do arquivo CSV de entrada | string | `input.csv`   |
| `-output`  | Caminho do arquivo JSON de saída  | string | `output.json` |
| `-workers` | Número de workers simultâneos     | int    | `4`           |

## Roadmap

- [x] Implementação core (Worker Pool)
- [x] Camada de Transformação (Middleware)
- [x] Sistema de Métricas e Logs Estruturados
- [x] Benchmarking e Otimização

## Contribuindo

Contribuições são bem-vindas! Veja o [CONTRIBUTING.md](CONTRIBUTING.md) para detalhes sobre nosso processo de desenvolvimento.

## Licença

Distribuído sob a licença MIT. Veja `LICENSE` para mais informações.

## Autor

**Enoque Sousa**

- [Portfólio](https://enoquesousa.vercel.app)
- [GitHub](https://github.com/ESousa97)
