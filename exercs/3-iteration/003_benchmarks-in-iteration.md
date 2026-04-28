# Benchmarks no contexto de iteration

## Contexto

Aparece no cap `iteration` da fase-1, na seção de refactor após `Repeat` ficar verde. Livro introduz benchmark como ferramenta de medição para guiar a escolha entre construções alternativas de string (concatenação ingênua com `+=`, `strings.Repeat`, `strings.Builder`).

## Por que

### Forma

```go
func BenchmarkRepeat(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Repeat("a")
    }
}
```

- Mora em arquivo `_test.go` junto com testes.
- Nome começa com `Benchmark` (espelha `Test`, `Example`).
- Recebe `*testing.B` em vez de `*testing.T`.
- Loop interno até `b.N`. Quem decide `N` é o runner — não o autor do benchmark.

### O que `b.N` faz

Runner chama a função várias vezes ajustando `b.N` automaticamente até estabilizar a medição. Começa pequeno (1, 100, 10k...) e aumenta até rodar tempo suficiente para estatística confiável (default ~1s). O autor nunca seta `b.N` — só lê.

### Como rodar

```bash
go test -bench=.            # roda todos os benchmarks
go test -bench=Repeat       # filtra por nome
go test -bench=. -benchmem  # mostra alocações também
```

Output típico:

```
BenchmarkRepeat-8    19531234    61.2 ns/op    16 B/op    1 allocs/op
```

- `BenchmarkRepeat-8` = nome + `GOMAXPROCS` (núcleos).
- `19531234` = `b.N` final (quantas vezes rodou).
- `61.2 ns/op` = tempo médio por operação.
- `B/op` = bytes alocados por chamada (com `-benchmem`).
- `allocs/op` = nº de alocações no heap por chamada.

### Por que aparece em iteration

Cap mostra trade-off de performance entre formas de construir string:

```go
// ingênua: realoca a cada `+=`
var repeated string
for i := 0; i < 5; i++ {
    repeated += character
}

// strings.Repeat: implementação otimizada da stdlib
return strings.Repeat(character, 5)

// strings.Builder: amortiza realocação interna
var sb strings.Builder
for i := 0; i < 5; i++ {
    sb.WriteString(character)
}
return sb.String()
```

Cada versão tem `ns/op` e `allocs/op` diferentes. Benchmark prova com número, não com palpite — é o ponto pedagógico do livro: refactor de performance precisa de medida, não de intuição.

### Diferença vs `TestX`

| | `TestX` | `BenchmarkX` |
|---|---|---|
| Pergunta | "tá certo?" | "quão rápido?" |
| Falha | comparação errada | nunca falha sozinho — mede |
| Roda em | `go test` | `go test -bench=.` (não roda no padrão) |
| Param | `*testing.T` | `*testing.B` |

Default de `go test` **não roda benchmarks** — precisa flag explícita. Razão: medição é cara e não cabe em todo CI run.

### Cuidados

- Benchmark deve fazer **só o trabalho medido** dentro do loop. Setup pesado (criar fixture, abrir conexão) vai antes do `for` ou após `b.ResetTimer()`.
- Compilador pode otimizar e descartar resultado se você não consumir. Padrão: atribuir para var package-level (`var sink string`) ou usar `b.ReportAllocs()`/`runtime.KeepAlive` em casos avançados.
- Benchmark mede a máquina onde roda, não verdade absoluta. Comparar entre versões na mesma máquina, não números absolutos entre máquinas diferentes.

### Conexão com TDD

Test guia corretude. Benchmark guia performance. Workflow do livro: primeiro verde (testes passam), só depois benchmark (mede), só então refactor com ganho mensurável. Sem benchmark, refactor de performance é palpite.
