# Testable Examples

## Contexto

Aparece na seção "Testable Examples" do refactor do cap `integers`. Depois do `Add` ficar verde com doc comment, autor introduz função `ExampleAdd` no `adder_test.go` — não é teste tradicional nem código de produção, é mecanismo Go de **documentação executável**.

```go
import "fmt"

func ExampleAdd() {
    sum := Add(1, 5)
    fmt.Println(sum)
    // Output: 6
}
```

Mora em arquivo `_test.go` junto com `TestX`, mesma pasta, mesmo package.

## Por que

### O que é Testable Example

Função especial que serve dois propósitos ao mesmo tempo:
1. **Vai pra documentação** — renderizada por `pkg.go.dev`, `pkgsite`, `go doc -all` junto com a função documentada.
2. **Roda como teste** — `go test` executa, captura stdout, compara com `// Output: ...`. Falha se output divergir.

### Três regras de detecção

1. **Nome começa com `Example`** (espelha `Test`, `Benchmark`).
2. **Sufixo casa com símbolo:**
   - `ExampleAdd` → exemplo da função `Add`.
   - `Example` (sem sufixo) → exemplo do package.
   - `ExampleUser_Save` → exemplo do método `Save` do tipo `User`.
3. **`// Output: ...` no fim** — define output esperado pra execução.

### Por que existe — problema que resolve

Documentação clássica (README, comentário largo) **apodrece**: alguém renomeia `Add` pra `Sum`, README continua dizendo `Add`. Drift silencioso. Ninguém percebe até usuário tentar copiar e dar erro.

Testable Example resolve por construção:
- É **código compilado**. Renomeou `Add` → `Sum`? Build do test suite quebra. Exemplo precisa atualizar.
- Output é **verificado**. Mudou comportamento de retorno? `// Output: 6` falha. Força revisão.

Documentação nunca diverge da realidade — compilador + test runner garantem.

### Variantes do `// Output:`

- **Sem `// Output:`** → exemplo só compila, não executa. Útil pra código com I/O real (rede, disco) que não cabe em test unitário, mas você quer pelo menos garantir que compila:

  ```go
  func ExampleHTTPGet() {
      resp, _ := http.Get("https://example.com")
      defer resp.Body.Close()
  }
  ```

- **`// Unordered output: ...`** → output não-determinístico (maps, goroutines), compara como conjunto.

### Como aparece no `go test -v`

```
=== RUN   TestAdder
--- PASS: TestAdder (0.00s)
=== RUN   ExampleAdd
--- PASS: ExampleAdd (0.00s)
```

Tratado como entrada normal no test runner.

### Relação com `TestX`

| | `TestX` | `ExampleX` |
|---|---|---|
| Roda no `go test` | sim | sim |
| Vai pra documentação | não | **sim** |
| Forma | `if got != want { t.Errorf(...) }` | `fmt.Println(...)` + `// Output: ...` |
| Foco | corretude exaustiva (edge cases) | demonstração canônica de uso |

Não substituem um ao outro. Coexistem: `TestAdd` cobre casos limite (negativos, zero, overflow), `ExampleAdd` mostra uso típico que usuária vai colar no próprio código.

### Por que aparece no cap 2

Cap 2 é primeiro com função pública genuína (`Add` capitalizado, package biblioteca). É o ponto natural pra apresentar o pacote completo de uma API polida em Go: assinatura clara + named return quando útil + doc comment + Testable Example. Padrão idiomático pra qualquer biblioteca exposta.
