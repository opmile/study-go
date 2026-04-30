# Slice operator `[low:high]`

A sintaxe cria uma **nova view** sobre o mesmo array subjacente — não copia dados.

```go
nums := []int{10, 20, 30, 40, 50}
//             0   1   2   3   4   <- índices

nums[1:4] // [20 30 40]
```

Regra: pega de `low` (inclusivo) até `high` (exclusivo). Mesma convenção de Python.

## Omitindo os lados

```go
nums[2:]  // [30 40 50]   — do índice 2 até o fim
nums[:3]  // [10 20 30]   — do início até o índice 3 (exclusivo)
nums[:]   // [10 20 30 40 50] — cópia da view inteira
```

Por isso `numbers[1:]` significa "tudo a partir do índice 1" — útil pra padrão `head, tail`:

```go
head := nums[0]   // 10
tail := nums[1:]  // [20 30 40 50]
```

Esse é o pão-com-manteiga de funções recursivas e processamento de listas em Go.

## Conexão com o que você já viu

Lembrando do header `(ptr, len, cap)`:

```go
nums := []int{10, 20, 30, 40, 50}  // len=5, cap=5
sub := nums[1:3]                    // len=2, cap=4
```

`sub` aponta pro mesmo array, começando no índice 1. A nova `cap` é 4 porque do índice 1 até o fim do array original sobram 4 posições.

E aí volta a pegadinha do compartilhamento:

```go
sub[0] = 999
fmt.Println(nums) // [10 999 30 40 50] — alterou o original
```

## Sintaxe de três índices (avançada)

Existe também `slice[low:high:max]` que controla a capacity da nova view:

```go
sub := nums[1:3:3] // len=2, cap=2
```

Útil quando você quer impedir que `append` numa sub-slice corrompa o array original. Caso de borda, mas vale conhecer.

## Modelo mental

`[low:high]` é uma operação de **janela**, não de cópia. Você está reposicionando o ponteiro e ajustando os contadores do header. Custo O(1), zero alocação. O preço é o compartilhamento de memória — se quiser independência real, precisa copiar explicitamente:

```go
independent := make([]int, len(sub))
copy(independent, sub)
```