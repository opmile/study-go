# Subtests com t.Run

## Contexto

Introduzido na v4 do capítulo `hello-world` (seção "Hello, world... again"). O test `TestHello` começa a precisar validar mais de um cenário — o caso normal (`Hello("Chris")`) e o caso de fallback (`Hello("")`). Em vez de criar `TestHelloNormal` e `TestHelloEmpty` como funções separadas, o livro usa `t.Run` para aninhar subtests dentro de `TestHello`.

## Por que

### O que `t.Run` faz mecanicamente

```go
func TestHello(t *testing.T) {
    t.Run("caso A", func(t *testing.T) { /* ... */ })
    t.Run("caso B", func(t *testing.T) { /* ... */ })
}
```

Cada `t.Run` cria um **novo `*testing.T` filho**, com escopo isolado. O `t` dentro do closure é um ponteiro diferente do `t` externo — apesar do mesmo nome, a variável local sombreia a de fora. O runner executa cada subtest sequencialmente (ou em paralelo com `t.Parallel()`). Fail em um subtest **não interrompe** os outros. O test pai `TestHello` passa apenas se todos os subtests passam.

### Output estruturado

```
--- FAIL: TestHello (0.00s)
    --- PASS: TestHello/caso_A (0.00s)
    --- FAIL: TestHello/caso_B (0.00s)
        hello_test.go:15: got "X" want "Y"
```

O nome do subtest vira sufixo no path do test pai. Espaços viram underscores.

### Rodar subtest específico

```bash
go test -run TestHello/caso_A     # só A
go test -run TestHello/caso       # prefixo: A e B
```

### Por que usar em vez de N funções `TestX` separadas

- **Agrupa por comportamento** — todos os casos de `Hello` ficam embaixo de `TestHello`, o que documenta a API visualmente.
- **Setup compartilhado** — variáveis declaradas em `TestHello` são visíveis nos closures dos subtests, reduzindo duplicação.
- **Isolamento de fail** — se `caso_A` falha, `caso_B` ainda roda. Funções separadas também rodam separadas, mas perdem o agrupamento semântico.

### Semântica de fail dentro de subtest

- `t.Fail()` marca o subtest como falho e **continua** executando o resto do closure.
- `t.FailNow()` ou `t.Fatal()` aborta apenas o subtest corrente e roda o próximo `t.Run`.
- O test pai é marcado como fail automaticamente assim que qualquer subtest falha.

### Uso canônico: table-driven test

Padrão Go idiomático (aparece formalmente no capítulo `structs-methods-and-interfaces`):

```go
cases := []struct{
    name, input, want string
}{
    {"nome normal", "Chris", "Hello, Chris"},
    {"empty", "", "Hello, World"},
}
for _, c := range cases {
    t.Run(c.name, func(t *testing.T) {
        got := Hello(c.input)
        if got != c.want {
            t.Errorf("got %q want %q", got, c.want)
        }
    })
}
```

Cada linha da tabela vira um subtest nomeado. Adicionar caso = adicionar linha, sem duplicar estrutura de teste.

### Analogia Java/JUnit 5

- `t.Run("nome", ...)` corresponde conceitualmente a `@Nested` + `@DisplayName`.
- Table-driven tests correspondem a `@ParameterizedTest` + `@MethodSource`.

Diferença: Go faz tudo com closures dentro de uma função comum, sem anotações nem descoberta via reflection. Menos mágica, mais código direto.
