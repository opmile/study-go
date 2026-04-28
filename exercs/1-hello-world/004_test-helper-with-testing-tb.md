# Helper de teste com testing.TB e t.Helper()

## Contexto

Aparece no refactor da v4 do capítulo `hello-world`. Depois de passar os testes do caso normal (`Hello("Chris")`) e do caso de fallback (`Hello("")`), o livro extrai a assertion repetida em uma função helper `assertCorrectMessage`. O trecho final do refactor fica:

```go
func TestHello(t *testing.T) {
    t.Run("saying hello to people", func(t *testing.T) {
        got := Hello("Chris")
        want := "Hello, Chris"
        assertCorrectMessage(t, got, want)
    })

    t.Run("empty string defaults to 'world'", func(t *testing.T) {
        got := Hello("")
        want := "Hello, World"
        assertCorrectMessage(t, got, want)
    })
}

func assertCorrectMessage(t testing.TB, got, want string) {
    t.Helper()
    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}
```

## Por que

### Por que extrair o helper

Cada subtest repetia o mesmo bloco de comparação e `t.Errorf`. Test code é código: TDD trata o refactor pós-green como parte do ciclo, aplicável tanto à produção quanto aos testes. Com o helper, adicionar um novo cenário de teste passa a ser uma linha em vez de três, e a intenção (`assertCorrectMessage`) fica explícita na leitura.

### Por que `testing.TB` como tipo do parâmetro em vez de `*testing.T`

`testing.TB` é uma interface da stdlib que tanto `*testing.T` (testes) quanto `*testing.B` (benchmarks) satisfazem implicitamente. Assinatura simplificada:

```go
type TB interface {
    Error(...)
    Errorf(...)
    Fatal(...)
    Fatalf(...)
    Helper()
    Log(...)
    // ... outros
}
```

Usar `testing.TB` no parâmetro permite que o mesmo helper seja chamado a partir de `TestXxx` e `BenchmarkXxx` sem duplicação:

```go
func TestHello(t *testing.T)      { assertCorrectMessage(t, ...) }
func BenchmarkHello(b *testing.B) { assertCorrectMessage(b, ...) }
```

Se o parâmetro fosse `*testing.T`, apenas testes poderiam chamar — perdia o reuso em benchmarks. Esta é a primeira aparição prática de **interface satisfaction implícita** em Go: `*testing.T` e `*testing.B` nunca declaram `implements TB`; eles simplesmente possuem os métodos certos, e o compilador aceita.

### Por que interface em Go é passada por valor (não ponteiro)

Um valor de interface em Go é um header de duas palavras: `{ tipo-concreto, ponteiro-pro-dado }`. Quando `t *testing.T` é passado para um parâmetro `testing.TB`, o runtime monta esse header — o ponteiro interno continua apontando para o mesmo `testing.T{}` que o runner criou. Mutar via `tb.Errorf(...)` ainda alcança o struct original.

Consequência: `*testing.TB` (ponteiro para interface) é quase sempre errado — introduz indireção dupla sem ganho. Regra Go: struct mutável → `*T`; interface → passa por valor; nunca `*Interface`.

### Por que `t.Helper()`

Sem `t.Helper()`, quando `t.Errorf` dispara dentro do helper, o runner reporta a linha do `Errorf` **dentro do próprio helper**:

```
hello_test.go:29: got "X" want "Y"    ← linha do t.Errorf no helper
```

Todo fail aponta para a mesma linha do helper, independentemente de qual subtest falhou. Diagnóstico inútil.

Com `t.Helper()`, o runner marca aquela função como "não contar no stack de reporte". O output aponta para o call site:

```
hello_test.go:11: got "X" want "Y"    ← linha do assertCorrectMessage() dentro do subtest
```

Agora a linha reportada é exatamente onde o test real está — dá para navegar até a falha.

`Helper()` não muda pass/fail — só muda a linha reportada. Se os testes passam, `Helper()` é no-op invisível; o valor aparece somente na hora de uma falha.

Regra prática: toda função helper de teste chama `t.Helper()` na primeira linha. Custo zero, valor quando precisar.

### Shorthand de parâmetros do mesmo tipo

```go
func assertCorrectMessage(t testing.TB, got, want string) {
```

Quando dois ou mais parâmetros consecutivos têm o mesmo tipo, pode-se declarar o tipo apenas no último: `got, want string` em vez de `got string, want string`. Mesma regra vale para retornos nomeados, e é o estilo idiomático em todo código Go.
