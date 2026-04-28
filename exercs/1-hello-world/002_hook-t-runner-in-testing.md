# Hook, *T e Runner em testing

## Contexto

Primeiro contato com `func TestHello(t *testing.T)` em `hello_test.go`. O livro chama `t` de "hook" no framework de testes. A pergunta é: o que esse "hook" é de verdade, por que é ponteiro (`*testing.T` e não `testing.T`), e quem é esse "runner" que o livro menciona.

## Por que

### O que `t` realmente é

"Hook" é linguagem frouxa do livro. O conceito correto é **handle/receiver para estado gerenciado pelo runtime**. Você não criou `t` — o test runner criou, passou pra sua função. Seu acesso ao framework de testes é via métodos em `t`. Sem `t`, sem como marcar fail, logar, pular, registrar subtest.

Padrão generalizado em Go: **função recebe ponteiro ou interface de um objeto controlado por outro sistema; você interage com o sistema chamando métodos nesse objeto.**

Outros objetos com o mesmo padrão em Go:

| Objeto | De quem vem | Pra que serve |
|---|---|---|
| `t *testing.T` | test runner | fail, log, subtest, cleanup |
| `b *testing.B` | benchmark runner | `b.N`, `b.ResetTimer()` |
| `ctx context.Context` | caller da cadeia de chamadas | cancelamento, deadline, values request-scoped |
| `w http.ResponseWriter, r *http.Request` | `net/http` server | escrever resposta, ler request |
| `tx *sql.Tx` | `db.Begin()` | commit/rollback, queries na transação |
| `logger *log.Logger` | `log.New(...)` | escrever logs formatados |

Analogia Java/Spring: mesma forma que `HttpServletRequest req, HttpServletResponse resp` num controller — você não instancia, o container cria e injeta. Você chama métodos pra interagir com a máquina do servlet. Ou `@Autowired` de um bean: framework decide ciclo de vida, você só usa.

Como identificar o padrão quando aparecer: (1) quem criou o objeto? Se não foi você, é handle. (2) O que some se eu não chamar método nele? (3) Ele carrega estado entre chamadas?

### Por que `*T` (ponteiro) e não `T` (valor)

Go passa argumentos **por valor** por default. Função recebe cópia. Exceção: slice, map, channel — são headers que já contêm ponteiro internamente.

Se `TestHello` recebesse `testing.T` (valor), qualquer `t.Fail()` mutaria a cópia. Quando a função retornasse, cópia some, runner leria `failed` do original → sempre `false` → test sempre passa. Catastrófico.

Com `*testing.T` (ponteiro), o método muta o struct original que o runner criou. Runner lê depois, vê `failed = true`.

Regra prática Go: precisa mutar estado visível pro caller → ponteiro. Struct grande (custo de cópia) → ponteiro. Só lê e struct pequeno → valor. Mesma lógica vale pra receiver de método.

Contraste com Java: em Java objeto é sempre referência, não tem escolha. Go força a decisão explícita por chamada — mais verbo, mais controle.

Pegadinha que vai aparecer no capítulo de pointers:

```go
type Wallet struct{ balance int }
func (w Wallet) Deposit(x int) { w.balance += x }  // cópia mutada e descartada

w := Wallet{}
w.Deposit(10)
fmt.Println(w.balance)  // 0
```

Go compila sem avisar. Test vai pegar.

### Quem é o "runner"

Runner = binário `go test`. Quando você digita `go test`:

1. `go test` compila **binário temporário** no diretório, contendo: seu código, seus testes, e uma `main()` **gerada automaticamente** pelo `go test` (você nunca vê).
2. A `main()` gerada usa o package `testing` pra: achar todas as funções `TestXxx(t *testing.T)`; pra cada uma, instanciar `testing.T{}`, passar o ponteiro, chamar a função; após retorno, ler `t.failed`, `t.skipped`, logs e imprimir resultado.
3. Binário sai com exit code 0 (pass) ou 1 (fail). Apaga binário.

Pra ver isso com os olhos: `go test -c` compila sem rodar, aparece `hello.test` no diretório. `./hello.test -test.v` roda manual. `go test -v` mostra cada `TestXxx` correndo em sequência.

"Runner" é então duas coisas juntas: a CLI `go test` (compila + orquestra) e o package `testing` + main gerado (roda dentro do binário, cria `*testing.T`, coleta resultados). Sem mágica — source do `testing` está em `$GOROOT/src/testing/testing.go`.

Diferença importante pra JUnit/Spring: Go não usa anotação. Runner descobre test por **convenção de nome** (`TestXxx` + assinatura `func(t *testing.T)`). Menos config, mais rigor no nome.
