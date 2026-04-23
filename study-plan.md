# Study Plan — learn-go-with-tests

**Critério de progresso:** modelos mentais consolidados por semana, não capítulos concluídos.
**Critério de conclusão por capítulo:** conseguir explicar o conceito central em termos de _como difere do Java_ sem consultar o livro.

---

## Fase 1 — Fundamentos e desaprendizado de Java (Semanas 1–2)

Foco: estabelecer o modelo mental de Go como linguagem com filosofia diferente — não Java com sintaxe diferente.

| Capítulo | Modelo mental central | Mapeamento Java |
|---|---|---|
| Hello, World | Ciclo TDD em Go; `go test`; estrutura de pacotes | `@Test` do JUnit → arquivo `_test.go` |
| Integers | Funções, tipos primitivos, documentação via `Example` | Métodos estáticos, Javadoc |
| Iteration | `for` é o único loop em Go | `for`, `while`, `do-while` colapsados em um |
| Arrays and slices | Slice como view sobre array; `append`; nil slice | `int[]` vs `ArrayList` — diferença fundamental de mutação |
| Structs, methods & interfaces | Interfaces satisfeitas implicitamente | **Ruptura central:** Java exige `implements`; Go infere pela assinatura |
| Pointers & errors | `error` como valor de retorno; `*T` explícito | Exceptions vs. `if err != nil`; sem NullPointerException — zero values |
| Maps | `map[K]V`; comma-ok idiom | `HashMap<K, V>`; acesso seguro sem `getOrDefault` |

**Checkpoint da fase:** escrever um programa que lê um mapa, itera sobre ele e retorna erro sem usar nenhuma exception.

---

## Fase 2 — Idiomas Go: inversão de dependência e concorrência (Semanas 3–4)

Foco: como Go resolve os problemas que Java resolve com frameworks (Spring, ExecutorService).

| Capítulo | Modelo mental central | Mapeamento Java |
|---|---|---|
| Dependency Injection | DI sem container; interface implícita como contrato | Spring `@Autowired` vs. passar dependência no construtor |
| Mocking | Mocks manuais via interface; sem Mockito | Mockito/`@Mock` — Go força você a entender o que está mockando |
| Concurrency | Goroutines leves; channels como pipe de comunicação | `Thread`, `ExecutorService`, `CompletableFuture` |
| Select | Multiplexação de channels; timeout pattern | `CompletableFuture.anyOf` + timeout explícito |
| Sync | `sync.Mutex`, `sync.WaitGroup`; quando *não* usar channel | `synchronized`, `CountDownLatch`, `Semaphore` |
| Context | Propagação de cancelamento pela call chain | `ThreadLocal` + `Future.cancel()` — mas sem leak de goroutine |

**Checkpoint da fase:** implementar um worker pool simples com goroutines e channel sem usar `sync.Mutex`.

---

## Fase 3 — Recursos avançados da linguagem (Semana 5)

Foco: ferramentas que aparecem em código de produção — vale entender para ler, não necessariamente dominar agora.

| Capítulo | Modelo mental central | Mapeamento Java |
|---|---|---|
| Reflection | `reflect` package; uso raro, custo alto | `java.lang.reflect` — mesmo tradeoff de legibilidade |
| Intro to property-based tests | Testes por propriedades, não exemplos | JUnit Theories / QuickCheck — verifica invariantes |
| Reading files | `io.Reader` como abstração universal de I/O | `InputStream`, `BufferedReader` — mesmo conceito, interface menor |
| Generics | Type parameters `[T any]`; type constraints | Generics Java desde 1.5 — Go adicionou em 1.18, filosofia mais restrita |
| Revisiting arrays with generics | Aplicar generics em código já escrito | Refatorar `ArrayList<Object>` para `ArrayList<T>` |

---

## Fase 4 — Fundamentos de teste (Semana 6)

Foco: como testar Go profissionalmente — além do ciclo TDD básico.

| Capítulo | Modelo mental central |
|---|---|
| Intro to acceptance tests | Testes de fora para dentro; black-box vs. white-box |
| Scaling acceptance tests | Estrutura de projeto quando testes crescem |
| Working without mocks | Quando mock é o problema, não a solução |
| Refactoring Checklist | Checklist concreto; leitura rápida, referência recorrente |

---

## Fase 5 — Construir uma aplicação (Semanas 7–8)

Foco: aplicação do mundo real — **esta é a fase mais relevante para o trabalho na Arco.** HTTP server em Go é o core do que times com essa stack fazem.

| Capítulo | Modelo mental central | Relevância Arco |
|---|---|---|
| HTTP server | `net/http`; handlers; routing sem framework | Alta — padrão de servidor Go |
| JSON, routing & embedding | `encoding/json`; struct tags; embedding vs. herança | Alta — serialização de API |
| IO and sorting | Interfaces `io.Reader`/`io.Writer`; `sort.Interface` | Média — I/O de arquivos e ordenação |
| Command line & package structure | `flag` package; estrutura de pacotes em projeto real | Média — CLIs e múltiplos binários |
| Time | `time.Time`; mock de tempo em testes | Média — testes determinísticos |
| WebSockets | `gorilla/websocket`; upgrade HTTP | Baixa no início, alta em features realtime |

**Checkpoint da fase:** ter um HTTP server funcionando que aceita JSON, persiste em memória e responde com os dados corretos — tudo com testes passando.

---

## Fase 6 — Q&A e consolidação (Semana 9)

Leitura opcional — consultar conforme aparecer necessidade real.

| Capítulo | Quando ler |
|---|---|
| OS Exec | Se precisar rodar subprocessos |
| Error types | Quando `if err != nil` não for suficiente para comunicar contexto |
| Context-aware Reader | Se implementar I/O com cancelamento |
| Revisiting HTTP Handlers | Após terminar a fase 5 — refatora o que foi feito |

---

## Referência rápida — conceitos que mais travam quem vem do Java

| Conceito Go | O que confunde | Regra prática |
|---|---|---|
| Interface implícita | Parece que qualquer struct satisfaz qualquer interface por acidente | Se compilou, satisfaz. Confie no compilador. |
| `nil` vs zero value | `nil` existe, mas zero values previnem muitos NPE | Structs têm zero value útil; ponteiros têm `nil` |
| `error` como retorno | Tentação de criar hierarquia de exceptions | Sempre checar `err != nil` imediatamente após a chamada |
| Goroutines vs threads | Goroutines são muito mais baratas — muda a intuição de custo | Criar goroutine é barato; vazar goroutine é caro |
| `defer` | Executa na saída da função, não do bloco | Pense em `finally` que se registra no ponto de uso |
| Capitalização como visibilidade | `Exported` vs `unexported` — sem `public`/`private` keyword | Letra maiúscula = público para outros pacotes |

---

## Como usar este plano

- **Não pule fases.** A fase 2 depende da 1; a fase 5 depende da 2.
- **Checkpoint antes de avançar fase.** Se não passar no checkpoint, revise os capítulos mais fracos da fase.
- **Fase 6 é consulta, não sequência.** Leia quando o problema aparecer.
- **Tempo por capítulo:** estimativa de 2–4h por capítulo (leitura + TDD + reflexão). Fases 1–2 podem ser mais rápidas pelo mapeamento Java.
