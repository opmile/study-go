# for: única estrutura de loop em Go

## Contexto

Aparece no cap `iteration` da fase-1. Go intencionalmente não tem `while`, `do-while` nem `foreach` separados — `for` cobre todos os casos via 4 formas sintáticas. Decisão filosófica: uma palavra-chave, várias formas, em vez de várias palavras-chave para casos parecidos.

## Por que

### Forma 1 — Clássica (3-cláusulas)

```go
for i := 0; i < 10; i++ {
    fmt.Println(i)
}
```

Três partes separadas por `;`:
- **init**: `i := 0` — roda uma vez no começo.
- **cond**: `i < 10` — testada antes de cada iteração; falsa = sai.
- **post**: `i++` — roda no fim de cada iteração.

Sem parênteses ao redor (estilo Go). Chaves `{}` obrigatórias mesmo com body de uma linha.

### Forma 2 — Só condição (equivalente a `while`)

```go
for x < 100 {
    x = next(x)
}
```

Omite init e post. `while cond` em outras linguagens vira `for cond`.

### Forma 3 — Infinito

```go
for {
    if done { break }
    work()
}
```

Sem cláusulas. Sai com `break`, `return`, `panic` ou `os.Exit`. Padrão para event loops, servers, goroutines de worker.

### Forma 4 — `for...range`

```go
for i, v := range nums { ... }
```

Iteração sobre estrutura de dados. Detalhada em nota separada (`002_range-deep-dive.md`).

### Controle de fluxo

```go
for i := 0; i < 10; i++ {
    if i == 3 { continue }   // pula resto da iteração
    if i == 7 { break }      // sai do loop
}
```

### Labels para break/continue de loop aninhado

```go
outer:
for i := 0; i < 5; i++ {
    for j := 0; j < 5; j++ {
        if i*j > 6 { break outer }   // sai do loop externo
    }
}
```

Sem label, `break` sai só do loop mais interno. Label evita flag booleana auxiliar para sair de aninhamento.

### Por que Go reduziu para um único loop

- Menos sintaxe pra aprender, menos decisão por linha.
- `for` clássico e `for cond` cobrem todos os casos imperativos sem precisar de duas palavras distintas.
- Filosofia: ortogonalidade — composição de poucas primitivas, não vocabulário grande.

### Caso do cap iteration

`Repeat(character string)` constrói string repetindo N vezes:

```go
func Repeat(character string) string {
    var repeated string
    for i := 0; i < 5; i++ {
        repeated += character
    }
    return repeated
}
```

Cenário canônico de "fazer N vezes" — usa forma clássica. Refactor explora `strings.Repeat`/`strings.Builder` e mede com benchmark.
