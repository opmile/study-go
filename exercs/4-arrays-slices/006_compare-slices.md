# `reflect.DeepEqual` em testes

O problema: **slices não são comparáveis com `==`** em Go.

```go
a := []int{2, 9}
b := []int{2, 9}
a == b // erro de compilação: slice can only be compared to nil
```

Por quê? `==` em Go faz comparação rasa. Para structs e arrays compara campo a campo, mas slices são headers `(ptr, len, cap)` — comparar headers compararia ponteiros, não conteúdo. Go preferiu proibir a comparação a deixar você cair nessa armadilha.

## O que `reflect.DeepEqual` faz

Compara **recursivamente** o conteúdo, não a identidade:

```go
reflect.DeepEqual([]int{2, 9}, []int{2, 9}) // true
reflect.DeepEqual([]int{2, 9}, []int{9, 2}) // false (ordem importa)
```

Funciona com slices, maps, structs aninhadas, ponteiros — desce na estrutura e compara valor por valor.

## Por que no teste

```go
if !reflect.DeepEqual(got, want) {
    t.Errorf(...)
}
```

`got` e `want` são `[]int`. Sem `reflect.DeepEqual`, você teria que escrever um loop manual comparando `len` e cada índice. O package `reflect` resolve isso numa linha.

## Caveat importante

`reflect.DeepEqual` é **lento** (usa reflection em runtime) e **permissivo demais** em alguns casos:

```go
reflect.DeepEqual([]int{}, []int(nil)) // false
```

Slice vazio e slice nil são "diferentes" pra ele, embora sejam equivalentes na maioria dos contextos. Por isso, em código de produção, evita-se. Em testes é aceitável porque a clareza compensa o custo.

## Alternativas modernas

- **`slices.Equal`** (Go 1.21+): mais rápido e específico para slices, faz o que você quer aqui.
  ```go
  if !slices.Equal(got, want) { ... }
  ```
- **`go-cmp`** (Google): biblioteca de teste com mensagens de diff melhores, padrão em projetos sérios.

## Modelo mental

`reflect` é a porta de entrada para inspecionar tipos em runtime. `DeepEqual` é seu uso mais comum em testes porque resolve a limitação do `==` para tipos compostos. Mas é uma ferramenta de conveniência — quando você tem uma alternativa tipada (`slices.Equal`, `maps.Equal`), prefira ela.

---

# Benchmarking 

## Estrutura básica

```go
func BenchmarkReflectDeepEqual(b *testing.B) {
    got := []int{1, 2, 3, 4, 5}
    want := []int{1, 2, 3, 4, 5}
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        reflect.DeepEqual(got, want)
    }
}

func BenchmarkSlicesEqual(b *testing.B) {
    got := []int{1, 2, 3, 4, 5}
    want := []int{1, 2, 3, 4, 5}
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        slices.Equal(got, want)
    }
}
```

Roda com:
```
go test -bench=. -benchmem
```

`b.N` é ajustado automaticamente pelo runtime até obter uma medição estatisticamente estável. Você não escolhe o número de iterações — o Go escolhe.

## O que observar

Saída típica:
```
BenchmarkReflectDeepEqual-8    10000000    150 ns/op    48 B/op    2 allocs/op
BenchmarkSlicesEqual-8        500000000      3 ns/op     0 B/op    0 allocs/op
```

Três métricas:
- **ns/op**: nanossegundos por operação
- **B/op**: bytes alocados (com `-benchmem`)
- **allocs/op**: número de alocações no heap

A diferença entre os dois aqui é gigante — `reflect` aloca porque empacota valores em `reflect.Value`, enquanto `slices.Equal` é um loop direto sem alocação.

## Cuidados que costumam morder

**`b.ResetTimer()`** depois do setup, para não contar a inicialização.

**Resultado precisa ser usado**, senão o compilador pode otimizar a chamada inteira:
```go
var result bool
for i := 0; i < b.N; i++ {
    result = slices.Equal(got, want)
}
_ = result
```
Para casos extremos, existe `runtime.KeepAlive` ou a variável global `var sink`.

**Variar o tamanho da entrada** com sub-benchmarks para ver como escala:
```go
for _, n := range []int{10, 100, 1000, 10000} {
    b.Run(fmt.Sprintf("n=%d", n), func(b *testing.B) {
        s1 := make([]int, n)
        s2 := make([]int, n)
        b.ResetTimer()
        for i := 0; i < b.N; i++ {
            slices.Equal(s1, s2)
        }
    })
}
```

Aí você vê se a relação entre os dois é constante ou se uma escala pior.

## Comparando duas implementações de forma rigorosa

A ferramenta padrão é **`benchstat`** (do golang.org/x/perf):

```
go test -bench=. -count=10 > old.txt
# muda o código
go test -bench=. -count=10 > new.txt
benchstat old.txt new.txt
```

`-count=10` roda cada benchmark 10 vezes, e o `benchstat` calcula média, desvio e p-value entre as duas séries. Sem isso, você está olhando uma medição única e ruído pode te enganar.

## Modelo mental

Benchmark em Go responde "qual é mais rápido em condições controladas". Mas duas coisas valem lembrar: medir microbenchmarks fora do contexto real pode ser enganoso (cache quente, branch prediction, inlining se comportam diferente em produção), e a diferença só importa se a função estiver no caminho crítico. Para `reflect.DeepEqual` vs `slices.Equal`, o veredito já é conhecido — `slices.Equal` ganha por ordens de grandeza — mas rodar você mesma uma vez é um excelente exercício para internalizar a metodologia.