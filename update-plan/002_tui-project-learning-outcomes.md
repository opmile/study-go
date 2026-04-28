# O que o projeto Notas TUI vai te ensinar de Go na prática

## Contexto

Projeto Notas TUI (bloco de notas em terminal com abas, ecossistema Charm) foi escolhido como primeiro hands-on em Go. Objetivo: solidificar fundamentos — sintaxe, organização de pacotes, interfaces básicas, error handling, I/O — em código real motivador, sem encarar ainda os conceitos avançados (concorrência crua, pipelines, context) que ficam reservados para outro projeto.

## Por que

A implementação do TUI força colisão prática com os seguintes pontos de Go:

### Structs e tipos

Uso extensivo de `Note`, `model`, tags JSON, `time.Time`. Aprende-se **zero values** (struct com campos zerados é construção segura), **struct literals** e serialização via `encoding/json` com tags.

### Enums idiomáticos com `iota`

O padrão `type mode int` seguido de `const (modeNormal mode = iota; modeEditing; modeRenaming)` é como Go faz enum — não há keyword `enum`. Aprende-se tipo nomeado a partir de int e convenção `iota` para geração de valores sequenciais.

### Organização em múltiplos pacotes

Estrutura `storage/`, `keys/`, `styles/` força imports entre pacotes do mesmo módulo. Capitalização como visibility vira prática diária: `NoteStore` exportado, `jsonStore` privado ao pacote. Aprende-se a desenhar superfície pública mínima por pacote.

### Interfaces implícitas no ponto certo

`NoteStore` é interface pequena definida perto do consumo, com implementação `jsonStore` separada. Aprende-se o padrão Go idiomático: interface declarada onde é usada, não onde o tipo concreto vive. Intuição base para interfaces mais sofisticadas no ETL.

### Error handling exaustivo

Todo I/O de arquivo e todo `json.Marshal`/`json.Unmarshal` retorna `error`. Internaliza-se o ritmo `if err != nil { return ..., err }` por saturação — único jeito real de engolir o padrão vindo de exceptions de Java.

### Métodos com receivers (value vs pointer)

O contrato do Bubbletea — `func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd)` — exige decisão explícita sobre value vs pointer receiver. Aprende-se implicação prática: value receiver retorna cópia modificada (padrão Bubbletea/Elm), pointer receiver muta estado in-place (mais comum em Go de produção).

### Funções como valores

`tea.Cmd` é `func() tea.Msg` — função que retorna mensagem. Primitivo pouco destacado em Java pré-8, aqui é base cotidiana para efeitos assíncronos. Aprende-se a passar função como argumento e retorno sem cerimônia.

### Stdlib essencial

- `os.UserConfigDir()` para descobrir `~/.config/...` de forma cross-platform.
- `encoding/json` com tags e `time.Time`.
- `os.ReadFile` / `os.WriteFile` para persistência de arquivo.

Trio que aparece em qualquer código Go de produção.

### Dos 7 modelos mentais Go-específicos

O TUI cobre **4 dos 7**: interfaces implícitas, zero values, ponteiros/receivers, erros como valores. Ficam reservados para o ETL: **goroutines e channels cruas, pipeline pattern, composição via embedding, `context.Context`, slices em cenários problemáticos**. O TUI entrega fundamentos sólidos sem ocupar mentalmente os três slots mais difíceis, que serão o foco do projeto seguinte.
