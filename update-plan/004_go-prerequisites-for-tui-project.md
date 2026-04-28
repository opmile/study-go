# Pré-requisitos de Go para o projeto TUI

## Contexto

Decisão de escopo de aprendizado antes de começar o projeto Notas TUI (bloco de notas em terminal, ecossistema Charm), descrito em `update-plan/002_tui-project-learning-outcomes.md`. Pergunta: até qual fase do `learn-go-with-tests` precisa ter sido estudada antes de iniciar o TUI?

## Por que

### Resposta curta

**Fase 1 completa + parte da Fase 2** (DI + mocking). Pula concorrência da Fase 2 — o TUI não usa goroutines/channels crus; esses ficam reservados para o projeto seguinte (ETL).

### Mapeamento de outcome do TUI para capítulo do livro

| Outcome do TUI | Onde aprende |
|---|---|
| Structs, zero values, métodos | Fase 1 `structs-methods-and-interfaces` |
| Interfaces implícitas (`NoteStore`) | Fase 1 `structs-methods-and-interfaces` + **Fase 2 DI/mocking** (padrão "declare a interface perto do consumidor") |
| Pointer vs value receiver | Fase 1 `pointers-and-errors` |
| Error handling exaustivo | Fase 1 `pointers-and-errors` |
| Maps (indiretamente, state do model) | Fase 1 `maps` |
| `iota` enums | Não está no livro explicitamente — Go by Example cobre em 5 minutos |
| Funções como valores (`tea.Cmd`) | Fase 1 `iteration` (closures) + Fase 2 reforça |
| Multi-package, visibilidade por capitalização | Fase 1 ensina implicitamente em todo capítulo; Fase 2 DI formaliza |
| JSON tags + `encoding/json` | **Não precisa de capítulo formal**. Aprende-se no próprio projeto. Se travar, Fase 5 `JSON, routing and embedding` é referência pontual |

### Mínimo viável antes do TUI

1. Fase 1 inteira (hello → maps), sem pular.
2. Fase 2: capítulos de DI + mocking. Ensina a desenhar `NoteStore` como interface com `jsonStore` implementação testável.
3. Pula Fase 2 concorrência (goroutines, channels, select) — reservado para o ETL.

### O que fica de fora e não precisa ainda

- Fase 3 (reflection, generics, math) — zero uso no TUI.
- Fase 4 (acceptance tests, mais mocking) — overkill para o TUI.
- Fase 5 (HTTP, WebSockets) — só o JSON vira referência pontual, não pré-requisito.
- Fase 6 (Q&A) — consulta, não sequência.

### Ordem concreta

```
Fase 1 completa → Fase 2 DI → Fase 2 mocking → [começa TUI]
```

Durante o TUI, se faltar algo específico (iota, time.Time, os.UserConfigDir), usar Go by Example como dicionário.
