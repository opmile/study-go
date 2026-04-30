# Go coverage tool

## Contexto

Aparece no cap `arrays-and-slices` na frase "Go's built-in testing toolkit features a coverage tool". O cap usa coverage para provar que table-driven tests cobrem todos os casos: adiciona linha na tabela, roda `-cover`, vê % subir. Liga TDD a métrica objetiva. Faz parte da toolchain — sem dependência externa, sem instalação.

## Por que

### Comando básico

```bash
go test -cover
```

Output:

```
PASS
coverage: 100.0% of statements
ok  example.com/arrays  0.005s
```

Mostra % de **statements executados** pelos testes do package.

### Como funciona mecanicamente

Quando passa `-cover`, compilador **reescreve teu código** antes de rodar: insere contadores em cada bloco de statements. Cada vez que statement executa, contador incrementa. Ao fim, monta relatório. Não muda binário de produção — instrumentação só existe sob `-cover`.

### Visualização HTML — o ouro

```bash
go test -coverprofile=coverage.out
go tool cover -html=coverage.out
```

Abre browser com source code, **linhas em verde** (cobertas) ou **vermelho** (não cobertas). Vê visualmente o gap.

### Quando coverage é útil

- Detectar **branch não testado** (if/else cujo else nunca rodou).
- Validar refactor — coverage cair = teste não exercita mais a parte tocada.
- Code review — PR com nova função sem teste mostra na cobertura.

### Quando NÃO é

- **100% coverage ≠ código correto.** Cobre linha sem assertar comportamento = falso positivo.
- Não substitui pensar nos casos de borda; mostra só o que **executou**, não o que está **certo**.
- Métrica pode virar teatro: pessoas escrevem teste fraco só pra subir %.

### Statement vs branch coverage

Default Go = **statement coverage**. Cobre "linha rodou ou não". Não diferencia: 

```go
if x > 0 && y > 0 {
```

executou com `x>0,y>0`? Marca coberto. Mas `x<=0` (curto-circuito) não foi testado — branch coverage detectaria, statement não.

Go não tem branch coverage nativo. Comunidade vive bem com statement; é trade-off consciente.

### Flags úteis

```bash
go test -cover ./...                            # cobertura de todos os packages
go test -coverprofile=c.out -covermode=atomic   # seguro pra testes paralelos
go tool cover -func=coverage.out                # texto: cobertura por função
```

`-covermode=atomic` é necessário se testes rodam em paralelo (`t.Parallel()`) — counters viram atômicos.

### Workflow no cap arrays-and-slices

1. Escreve `Sum`, teste passa.
2. `go test -cover` → vê %.
3. Adiciona `SumAllTails`, teste cobre nil case e empty case.
4. `go tool cover -html=coverage.out` → confirma que branch nil foi exercitado.
5. Refactor com `range`, coverage continua igual = não quebrou nada.

Coverage vira sanity check do refactor.
