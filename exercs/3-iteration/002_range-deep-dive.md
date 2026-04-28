# range: iteração sobre estruturas de dados

## Contexto

Aparece no cap `iteration` ao lado das outras formas de `for`. `range` é palavra-chave que `for` aceita como cláusula alternativa às 3-cláusulas clássicas. Abstrai iteração sobre tipos iteráveis: array, slice, string, map, channel, inteiro (1.22+) e função iterator (1.23+).

## Por que

### Forma geral

```go
for <vars> := range <expr> {
    // body
}
```

`<expr>` deve ser tipo iterável. Compilador rejeita outros.

### Comportamento por tipo

#### Slice / array

```go
xs := []string{"a", "b", "c"}
for i, v := range xs {
    fmt.Println(i, v)
}
```

- `i` = índice (`int`, 0-based).
- `v` = **cópia** do elemento.

Mutar `v` não afeta o slice. Para mutar, usa índice:

```go
for i := range xs {
    xs[i] = strings.ToUpper(xs[i])
}
```

#### String

```go
for i, r := range "café" {
    fmt.Printf("%d: %c\n", i, r)
}
// 0 c, 1 a, 2 f, 3 é
```

- `i` = índice em **bytes** (não em runes).
- `r` = **rune** (codepoint Unicode, `int32`).

`range` em string decodifica UTF-8 automaticamente. Índice pula entre caracteres multibyte (`é` ocupa 2 bytes, próximo `i` salta).

Diferença crítica vs. `for i := 0; i < len(s); i++`: indexar string com `s[i]` retorna `byte`, não `rune`. Default certo é `range`.

#### Map

```go
for k, v := range m { ... }
```

- `k` = chave, `v` = valor (cópia).
- **Ordem é não-determinística por design.** Go aleatoriza intencionalmente para evitar código que dependa de ordem que a especificação não garante.

Para ordem estável: extrai chaves, ordena, itera:

```go
keys := make([]string, 0, len(m))
for k := range m { keys = append(keys, k) }
sort.Strings(keys)
for _, k := range keys { ... }
```

#### Channel

```go
for v := range ch { ... }
```

- Apenas uma variável: `v` = valor recebido.
- Loop bloqueia esperando próximo valor.
- Sai quando channel fecha (`close(ch)`).

Sem `close`, loop trava para sempre. Aprofundado na fase 2 (concurrency).

#### Inteiro (Go 1.22+)

```go
for i := range 5 { ... }   // 0..4
```

Substitui `for i := 0; i < N; i++` quando só interessa contar.

#### Função iterator (Go 1.23+)

Range-over-func — permite tipo customizado ser iterável. Avançado, fora do escopo da fase 1.

### As 4 formas de declarar vars

```go
for k, v := range m { }   // ambos
for k := range m    { }   // só primeira
for _, v := range m { }   // descarta primeira (blank identifier)
for range m         { }   // descarta tudo, só itera (1.22+)
```

`_` é o **blank identifier**: descarta valor que não vai usar. Compilador exige consumir toda variável declarada — `_` é exceção.

### Pegadinha histórica: closures e variável de loop

```go
funcs := []func(){}
for _, v := range []int{1, 2, 3} {
    funcs = append(funcs, func() { fmt.Println(v) })
}
for _, f := range funcs { f() }
// Pré-1.22: imprime 3 3 3
// Go 1.22+: imprime 1 2 3
```

Antes de 1.22, `v` era a mesma variável reusada entre iterações; closures capturavam referência, viam o último valor. Go 1.22 mudou: cada iteração cria `v` novo.

Workaround pré-1.22 (ainda aparece em código legado):

```go
for _, v := range items {
    v := v   // shadow: nova variável por iteração
    go func() { fmt.Println(v) }()
}
```

### Resumo mental

- `range` = "me dá os elementos disso, um por iteração".
- 1ª var = índice/chave; 2ª var = valor (sempre cópia).
- String = bytes pelo índice, runes pelo valor.
- Map = ordem aleatória por design.
- Channel = bloqueia, sai no `close`.
- `_` descarta o que não usa.
- Para mutar, usa índice (`xs[i] = ...`), não a cópia.
