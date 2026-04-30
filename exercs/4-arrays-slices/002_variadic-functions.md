# Variadic Functions

## Contexto

Aparece no cap `arrays-and-slices` quando o livro introduz `SumAll(numbersToSum ...[]int) []int` — função que recebe N slices e retorna slice com a soma de cada. É o primeiro contato com a sintaxe `...T` em parâmetro, e mostra dois pontos juntos: variadic com tipo composto (`...[]int` = zero ou mais slices de int) e construção de slice de retorno via `append`. Conexão direta com slices porque dentro da função, o parâmetro variadic **é um slice**.

## Por que

### Forma

```go
func Sum(nums ...int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}
```

`nums ...int` = "zero ou mais `int`s". Dentro da função, `nums` é um `[]int` — slice comum. Pode dar `range`, indexar, passar para outra função.

### Chamadas

```go
Sum()              // ok, slice vazio
Sum(1)             // [1]
Sum(1, 2, 3)       // [1, 2, 3]
Sum(1, 2, 3, 4, 5) // [1, 2, 3, 4, 5]
```

Compilador empacota argumentos em slice automaticamente.

### Operador `...` na chamada (spread)

Já tem `[]int` e quer passar para função variadic? Usa `...` no call site:

```go
nums := []int{1, 2, 3}

Sum(nums)      // ERRO: cannot use []int as int
Sum(nums...)   // OK: spread do slice
```

Sem `...`, compilador trata slice como **um argumento** (do tipo errado). Com `...`, **desempacota** elementos como argumentos individuais.

Mesmo operador, dois lados:
- **Definição**: `func F(xs ...int)` — empacota.
- **Chamada**: `F(slice...)` — desempacota.

### Restrições

- Apenas **um** parâmetro variadic por função, e precisa ser o **último**.
- Não pode misturar variadic com argumentos default (Go não tem default args).
- `nums...` precisa ser `[]T` exato, não `[]U` que satisfaça `T`.

### Caso clássico: `fmt.Println`

```go
func Println(a ...any) (n int, err error)
```

Aceita qualquer número de argumentos de qualquer tipo. `any` (alias de `interface{}`) + variadic = aceita tudo.

```go
fmt.Println("oi", 42, true)
```

### Anti-padrão

Variadic é açúcar sintático — não usar por estética. Se sempre passa N fixo, declara N parâmetros. Bom uso é **N genuinamente desconhecido**: `Sum`, `Println`, `path.Join`.

### No cap arrays-and-slices

`SumAll(numbersToSum ...[]int) []int`:

```go
func SumAll(numbersToSum ...[]int) []int {
    var sums []int
    for _, numbers := range numbersToSum {
        sums = append(sums, Sum(numbers...))
    }
    return sums
}
```

Dentro da função, `numbersToSum` é `[][]int`. `range` produz cada slice individual. `Sum(numbers...)` espalha o slice ao chamar a Sum variadic. Mostra os dois lados do operador `...` na mesma função.

### Resumo mental

- `...T` na **definição** = parâmetro vira `[]T`.
- `slice...` na **chamada** = passa elementos individuais.
- Variadic é só açúcar para slice.
- Último parâmetro, único variadic, sem mistura com defaults.
