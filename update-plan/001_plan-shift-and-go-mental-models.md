# Virada do plano: de livro linear para projeto ETL, e os modelos mentais Go-específicos

## Contexto

Com estágio iminente na Arco Educação e provável atuação em data pipeline (ETL para mapeamento de escolas com convênio), o caminho ótimo é trocar a trilha linear do livro por um projeto real que force colisão com os conceitos Go-específicos. 

## Por que

### Os 7 modelos mentais que tornam Go ≠ Java

Esses são os pontos onde o modelo mental do Java *atrapalha* e onde se perde tempo se não houver colisão prática:

1. **Goroutines e channels.** Não é `ExecutorService` disfarçado. A filosofia é "compartilhar memória comunicando" em vez de "comunicar compartilhando memória". Channels são primitivo de linguagem, não biblioteca. Pipeline pattern (stages conectados por channels) é *o* idioma canônico de Go — se não for internalizado, escreve-se Go-que-parece-Java usando mutex desnecessariamente.

2. **Interface satisfaction implícita.** Não existe `implements`. Qualquer tipo que tenha os métodos da interface satisfaz ela automaticamente, sem declarar intenção. Isso muda desenho de API: você define interfaces pequenas no ponto de *consumo*, não no ponto de *declaração do tipo*. Java força acoplamento explícito entre tipo e contrato; Go deixa o contrato emergir do uso.

3. **Composição via embedding** vs herança. Go não tem `extends`. Você embute um tipo dentro de outro e os métodos ficam promovidos — parecido com delegation, mas sintático. Força pensar em composição antes de hierarquia.

4. **Erros como valores** (`if err != nil`). Não há exceptions, não há `try/catch`. Toda função que pode falhar retorna `(resultado, error)`. A verbosidade é deliberada: torna visível onde o erro é tratado. Não é só sintaxe — muda como você desenha fluxos, pipelines e recuperação.

5. **Slices vs arrays.** `[]T` parece `List<T>` de Java, mas morde em casos não-óbvios: slices compartilham backing array por baixo. Mutar um slice pode mutar outro que você não sabia estar conectado. Gera bugs silenciosos se o modelo mental estiver em "ArrayList".

6. **Ponteiros e receivers** (value vs pointer receiver). Go tem ponteiros explícitos (`*T`). Cada método é declarado com receiver value (`func (t T)`) ou pointer (`func (t *T)`) — decisão de design em cada método, com implicações em mutação, performance e interface satisfaction. Em Java todo método opera sobre `this` (referência implícita); em Go é explícito e escolhido.

7. **Zero values.** Go não tem `null` como default. Cada tipo tem seu "zero": `0` pra int, `""` pra string, `nil` pra ponteiros/slices/maps, struct com campos zerados pra struct. Isso muda padrões de construção — você não precisa de constructor pra "inicializar direito", pode contar que o zero value é seguro (quando o tipo é bem desenhado).

Cerimônia de sintaxe (declaração de variável, if, switch, loops) é absorvida em leitura passiva. Modelos mentais acima exigem colisão prática — de preferência num projeto que *obrigue* a encará-los.

### Por que ETL é o projeto certo

ETL de escolas bate nos 7 pontos quase sem esforço:

- Leitura de fontes heterogêneas → **interfaces implícitas** (`io.Reader`, `Source` próprio).
- Parsing de CSV/JSON → **structs, zero values, slices**.
- Transformação paralela → **goroutines e channels**.
- Múltiplas fontes/sinks → **composição e embedding**.
- Falhas parciais (registro ruim não mata job) → **erros como valores**.
- Cancelamento e timeout → **`context.Context`** (primitivo onipresente em Go de produção).

Além disso, é o provável tipo de trabalho real no time da Arco, então o esforço de aprender o domínio junto é recuperado.

### Rejeitado: "construir um CRUD"

CRUD em Spring é zona de conforto — replica-se a arquitetura Java sem ser forçada a pensar em Go. Resultado: Go-que-parece-Java, sem aprendizado real. ETL pipeline empurra para estilos idiomáticos de Go desde o v0.

### Formato de trabalho a partir daqui

Aluna codifica. Ao travar, traz: (1) o que tentou, (2) hipótese sendo testada, (3) erro ou comportamento estranho. Claude devolve contra-pergunta ou explicação cirúrgica — não entrega código pronto antes da tentativa. Fim de cada fase: nota em `doc/` amarrada ao conceito consolidado. Livro `learn-go-with-tests` vira referência por capítulo quando um conceito exige aprofundamento.
