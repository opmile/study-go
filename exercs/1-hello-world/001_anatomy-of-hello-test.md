# Anatomia de hello_test.go: sintaxe Go linha a linha

## Contexto

Primeiro arquivo de teste em Go, vindo de JUnit em Java/Spring. Aplica-se a `hello_test.go` em `exercs/1-hello-world/`, que testa a função `Hello()` definida em `hello.go`. Serve de gabarito para a sintaxe básica de teste em Go: declaração de função de teste, parâmetro `*testing.T`, short variable declaration, `if`, `t.Errorf` e format verbs.

## Por que

### `package main`

Arquivos no mesmo diretório precisam pertencer ao mesmo pacote. O teste chama `Hello()` diretamente (sem prefixo), então precisa estar no mesmo pacote da função testada — `main`. Análogo a testes em Java que ficam na mesma package da classe pra acessar métodos package-private.

### `import "testing"`

Único import necessário pra escrever testes em Go. Não há JUnit, AssertJ, Mockito no classpath — tudo que o teste usa (`*testing.T`, `t.Errorf`) mora no pacote `testing` da stdlib. Filosofia deliberada: testing é feature de linguagem, não dependência externa.

### `func TestHello(t *testing.T)`

Três regras simultâneas:
- Nome começa com `Test` (maiúsculo). Sem isso, Go não reconhece como teste e ignora. Análogo a `@Test` do JUnit, só que por convenção de nome em vez de anotação.
- Recebe exatamente **um** parâmetro.
- Nome do parâmetro (`t`) vem antes do tipo (`*testing.T`) — ordem oposta a Java.

### `t *testing.T` — o asterisco

`*testing.T` é **ponteiro** para um valor do tipo `testing.T`. Go tem ponteiros explícitos; Java esconde referência a objeto por baixo dos panos.

Tradução prática: `t` não é cópia da struct `testing.T`, é um endereço pra ela. Quando se chama `t.Errorf(...)`, muta-se o estado do objeto de teste real (marcando-o como falho). Se fosse cópia, a falha se perderia.

Regra provisória: `*TipoQualquer` em parâmetro = "referência mutável pra esse tipo".

### `got := Hello()` — short variable declaration

`:=` faz três coisas em uma linha:
1. Declara a variável.
2. Infere o tipo do valor à direita (aqui: `string`).
3. Atribui o valor.

Formas equivalentes:
```go
var got string = Hello()   // tipo explícito
var got = Hello()          // tipo inferido, sintaxe var
got := Hello()             // tipo inferido, short form — só dentro de função
```

Análogo a `var got = hello();` de Java 10+. Diferença crítica: `:=` só vale **dentro de funções** — no nível de pacote, usa-se `var` ou `const`.

### `if got != want { ... }`

If do Go:
- Sem parênteses ao redor da condição (diferente de Java).
- Chaves `{}` são **obrigatórias**, sempre. Não existe `if cond return` em uma linha.
- `!=` e `==` comparam strings por **conteúdo** diretamente. Em Java, `"a" == "a"` pode falhar (precisa `.equals()`); em Go, string é tipo de valor comparável.

### `t.Errorf("got %q want %q", got, want)`

Duas ações:
1. Marca o teste como falho mas **continua** executando o resto da função. Para abortar, usar `t.Fatalf`.
2. Formata e imprime a mensagem.

### Format verbs

- `%q` — envolve o argumento em aspas. Essencial em testes: deixa visíveis espaços, strings vazias, caracteres invisíveis.
- `%s` — string sem aspas.
- `%d` — inteiro em decimal.
- `%v` — valor em formato default, serve pra qualquer tipo, útil em debug.
- `%T` — imprime o tipo do valor (`string`, `int`, etc.).

### Mapa geral JUnit → Go testing

| JUnit                              | Go                             |
|-----------------------------------|-------------------------------|
| `@Test public void testHello()`   | `func TestHello(t *testing.T)` |
| `assertEquals(want, got)`         | `if got != want { t.Errorf(...) }` |
| `@BeforeEach`, `@AfterEach`       | Não há — escreve-se helpers manualmente |
| `fail("msg")`                     | `t.Errorf(...)` (continua) ou `t.Fatalf(...)` (aborta) |
| Biblioteca externa                | Stdlib, sem deps               |

Go testing é deliberadamente minimalista. Não há assertion library oficial — filosofia é "se quer abstrair, escreve uma função helper".
