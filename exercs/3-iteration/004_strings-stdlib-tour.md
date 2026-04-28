# Tour da stdlib `strings`

## Contexto

No fim do cap `iteration`, autor sugere: "Have a look through the strings package. Find functions you think could be useful and experiment with them by writing tests like we have here." Investir tempo na stdlib de `strings` é um dos atalhos de maior retorno em Go — código idiomático assume domínio dela; reinventar `Contains` à mão é red flag em code review.

## Por que

### Núcleo (usado quase sempre)

| Função | O que faz | Exemplo |
|---|---|---|
| `Contains(s, substr)` | substring presente? | `strings.Contains("hello", "ell")` → `true` |
| `HasPrefix(s, p)` | começa com? | `strings.HasPrefix("file.go", "file")` → `true` |
| `HasSuffix(s, sfx)` | termina com? | `strings.HasSuffix("file.go", ".go")` → `true` |
| `Index(s, substr)` | posição da 1ª ocorrência ou -1 | `strings.Index("hello", "l")` → `2` |
| `Count(s, substr)` | nº de ocorrências | `strings.Count("banana", "a")` → `3` |

### Transformação

| Função | O que faz |
|---|---|
| `ToUpper`/`ToLower` | case |
| `TrimSpace` | tira whitespace das pontas |
| `Trim(s, "abc")` | tira chars do conjunto das pontas |
| `TrimPrefix(s, p)`/`TrimSuffix(s, sfx)` | remove só se presente |
| `Replace(s, old, new, n)` | substitui n ocorrências (-1 = todas) |
| `ReplaceAll(s, old, new)` | açúcar para `Replace(..., -1)` |

`Title` está deprecated — usar `cases.Title` de `golang.org/x/text/cases`.

### Split / Join

| Função | O que faz |
|---|---|
| `Split(s, sep)` | string → `[]string` |
| `SplitN(s, sep, n)` | split com limite de partes |
| `Fields(s)` | split por whitespace, ignora múltiplos espaços |
| `Join(parts, sep)` | `[]string` → string |

### Construção performática

| Tipo/função | Quando usar |
|---|---|
| `strings.Builder` | concat em loop (motivo do refactor no cap iteration) |
| `strings.Repeat(s, n)` | repetir N vezes |
| `strings.NewReader(s)` | string como `io.Reader` |

### Comparação

| Função | O que faz |
|---|---|
| `EqualFold(a, b)` | igualdade case-insensitive |
| `Compare(a, b)` | retorna -1/0/1 (geralmente prefere `==`/`<` direto) |

### Por que investir tempo aqui

- **Gratuito**: vem com toolchain, sem dependência externa.
- **Idiomático**: código Go assume domínio da `strings`.
- **Composição**: `Split` + `TrimSpace` + `Join` resolve ~80% de manipulação de texto sem regex.

### Estratégia mínima de exploração

1. Abrir `pkg.go.dev/strings`, ler índice de funções (não doc inteira de cada).
2. Para cada que parecer útil, escrever teste pequeno:
   ```go
   func TestContains(t *testing.T) {
       if !strings.Contains("hello", "ell") {
           t.Error("expected true")
       }
   }
   ```
3. Rodar `go test`, observar comportamento, passar para a próxima.
4. Marcar mental as 10–15 do núcleo (tabelas acima). Resto consulta sob demanda.

### O que NÃO precisa decorar

- `IndexByte`, `IndexRune`, `LastIndex*`, `IndexFunc`, `Map`, `EqualFold` em loops quentes — quando precisar, ler doc.
- Funções com `Func` no nome (`TrimFunc`, `IndexFunc`) — caso avançado, predicado custom.
- `Builder` com tuning fino (`Grow`) — só quando perfilamento mostrar gargalo.

### Pacotes adjacentes úteis

- `strconv` — `Atoi`, `Itoa`, parse de números.
- `unicode` — checagens (`IsLetter`, `IsDigit`).
- `regexp` — quando texto fica complexo demais para `strings`.
- `bytes` — versão de `strings` para `[]byte`, mesma API.

Cap iteration só pede `strings`. Outros entram conforme próximos capítulos.
