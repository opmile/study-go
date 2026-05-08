# study-go

Estudo dirigido de Go via [`learn-go-with-tests`](https://github.com/quii/learn-go-with-tests) (Chris James).

Trilha única, escolhida para evitar dispersão entre múltiplas fontes. Objetivo: construir modelo mental sólido da linguagem — não cobertura superficial de sintaxe.

## Estrutura

```
study-go/
├── CLAUDE.md                      # contexto e diretrizes para sessões com Claude Code
├── learn-go-with-tests/           # fork do livro, organizado por fases
│   ├── study-plan.md              # plano de estudos (gerado em 2026-04-22)
│   ├── fase-1-fundamentos/        # Hello World → Maps
│   ├── fase-2-idiomas/            # DI, Mocking, Concorrência
│   ├── fase-3-avancado/           # Reflection, Generics, Math
│   ├── fase-4-testing/            # Acceptance tests, Mocks
│   ├── fase-5-aplicacao/          # HTTP server, JSON, WebSockets
│   └── fase-6-qa/                 # Q&A de consulta
└── exercs/                        # implementações próprias por capítulo
    └── <N-cap>/
        ├── go.mod                 # cada capítulo é módulo Go isolado
        ├── <impl>.go
        ├── <impl>_test.go
        └── NNN_*.md               # notas de aprendizado
```

## Convenções

- **Prefixo numérico**: capítulos do livro com implementação própria em `exercs/` recebem prefixo (`1-hello-world`, `2-integers`, ...) também na pasta correspondente em `learn-go-with-tests/fase-N/`. Mantém ordem de estudo explícita.
- **Módulo por capítulo**: cada `exercs/<cap>/` tem próprio `go.mod`. Isolamento total — apagar a pasta não afeta os outros.
- **Notas `NNN_*.md`**: registros de aprendizado/decisão por capítulo, gerados durante explicações via Claude Code. Estrutura: `# Título`, `## Contexto`, `## Por que`.

## Comandos

```bash
# Rodar testes de um capítulo (a partir da pasta do exerc)
cd exercs/1-hello-world
go test

# Rodar todos os testes do livro upstream (estado original)
go test ./learn-go-with-tests/...

# Formatar código (idiomático antes de commit)
gofmt -w .

# Cobertura de testes
go test -cover
go test -coverprofile=c.out && go tool cover -html=c.out
```

## Progresso

Fase 1 — Fundamentos:

- [x] 1. hello-world
- [x] 2. integers
- [x] 3. iteration
- [x] 4. arrays-and-slices
- [ ] 5. structs-methods-and-interfaces
- [ ] 6. pointers-and-errors
- [ ] 7. maps

Fase 2 — Idiomas:

- [ ] dependency-injection
- [ ] mocking
- [ ] concurrency
- [ ] select
- [ ] sync
- [ ] context

## Licença e atribuição

`learn-go-with-tests/` é fork de [quii/learn-go-with-tests](https://github.com/quii/learn-go-with-tests) (MIT). Código próprio em `exercs/` é estudo pessoal, sem licença pública definida.
