# Length vs Capacity em slices

Slice em Go é um **header** com três campos: ponteiro pro array subjacente, length e capacity. Não é o array — é uma "janela" sobre ele.

```
slice header: { ptr, len, cap }
                 │
                 ▼
array:        [_, _, _, _, _]
```

## Os dois números

- **`len(s)`**: quantos elementos você pode acessar agora (`s[0]` até `s[len-1]`)
- **`cap(s)`**: quantos elementos cabem no array subjacente antes de precisar realocar

```go
s := make([]int, 0, 5)
fmt.Println(len(s), cap(s)) // 0 5
```

Length 0 = não tem nada acessível ainda. Capacity 5 = o array por baixo já tem 5 posições reservadas.

## Por que isso importa: `append`

Quando você dá `append`, o Go checa: cabe na capacity atual?

```go
s := make([]int, 0, 5)
s = append(s, 1, 2, 3)  // len=3, cap=5 — mesmo array
s = append(s, 4, 5)     // len=5, cap=5 — mesmo array
s = append(s, 6)        // len=6, cap=10 — REALOCOU
```

Quando estoura a capacity, Go aloca um array novo (geralmente o dobro), copia tudo e te devolve um slice apontando pro novo array. O array antigo vira lixo do GC.

## A pegadinha do compartilhamento

Slices que apontam pro mesmo array compartilham memória:

```go
a := []int{1, 2, 3, 4, 5}
b := a[1:3]              // len=2, cap=4 (vai até o fim de a)
b[0] = 99
fmt.Println(a)           // [1 99 3 4 5] — mudou!
```

`b` tem capacity 4 porque pode crescer até o fim do array original. Se você der `append(b, x)` e couber na cap, **vai sobrescrever `a[3]`**.

## Quando especificar capacity

Se você sabe quantos elementos vai inserir, pré-aloque:

```go
// ruim: várias realocações conforme cresce
nums := []int{}
for i := 0; i < 1000; i++ {
    nums = append(nums, i)
}

// bom: uma alocação só
nums := make([]int, 0, 1000)
for i := 0; i < 1000; i++ {
    nums = append(nums, i)
}
```

Length 0, capacity 1000 — o `append` nunca realoca, só incrementa o len.

## Modelo mental

`len` é a parte "viva" do slice, `cap` é o espaço total já alocado. Os dois existem porque Go separa **o quanto você está usando** de **o quanto está reservado** — isso permite `append` ser amortizado O(1) e permite slicing sem cópia.