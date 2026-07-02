# 📚 Glossário Acumulativo — Go na Prática

Termos e conceitos-chave de cada capítulo, para revisão rápida sem reabrir todos os READMEs. Atualize este arquivo a cada novo capítulo concluído.

---

## Capítulo 01 — Introdução ao Go

- **Compilado vs. Interpretado**: Go compila para binário nativo antes de rodar; não precisa do código-fonte ou de um interpretador em produção.
- **`package main`**: marca o pacote como executável (gera um binário), não uma biblioteca.
- **`func main()`**: ponto de entrada do programa.
- **Exportação por maiúscula**: identificador iniciado com letra maiúscula é público/exportado (`Println`); minúsculo é privado ao pacote.
- **`go fmt`**: formatador oficial — elimina debates de estilo (tabs, posição de chaves).
- **Filosofia "Less is exponentially more"**: poucas palavras-chave (25), sem "mágica" escondida no código.

## Capítulo 02 — Variáveis, Constantes e Tipos de Dados

- **`var`**: declaração explícita, usada fora de funções ou sem valor inicial.
- **`:=` (short declaration)**: só funciona dentro de funções; tipo inferido pelo valor.
- **Zero Value**: valor padrão automático ao declarar sem inicializar (`0`, `""`, `false`, `nil`).
- **`const`**: valor imutável definido em tempo de compilação; não aceita `:=`.

## Capítulo 03 — Tipos Básicos

- **Tipagem forte e estática**: sem conversão implícita entre tipos numéricos.
- **Overflow**: ao exceder a capacidade de um tipo (ex: `uint8` > 255), o valor "dá a volta" silenciosamente.
- **`float64` por padrão**: mais preciso que `float32`; nunca usar float para dinheiro (problema de representação IEEE 754).
- **Casting explícito**: `TipoDesejado(valor)` — obrigatório entre tipos diferentes.
- **`byte`** = alias de `uint8`; **`rune`** = alias de `int32` (representa um code point Unicode).

## Capítulo 04 — Controle de Fluxo

- **`if`/`else`**: sem parênteses na condição; chave `{` obrigatoriamente na mesma linha.
- **`if` com inicialização curta**: `if x := f(); cond { ... }` — variável só existe no bloco.
- **Operadores lógicos**: `&&` (E), `||` (OU), `!` (NÃO); curto-circuito em `&&`/`||`.
- **`switch`**: `break` implícito por padrão; aceita múltiplos valores por `case` (vírgula).
- **`switch` tagless**: equivalente a `switch true`, alternativa limpa a `if/else if` encadeado.
- **`fallthrough`**: força execução do próximo `case` sem testar sua condição (raro).
- **Guard clauses / early return**: evitar aninhamento profundo tratando casos negativos primeiro.

## Capítulo 05 — Laços de Repetição (`for`)

- **Única palavra-chave de repetição**: `for` cobre os papéis de `while`, `do-while` e loop infinito de outras linguagens.
- **`for` clássico**: `for init; cond; pós { }`.
- **`for` como `while`**: `for cond { }`, omitindo init/pós.
- **Loop infinito**: `for { }`, geralmente com `break` interno.
- **`break`**: aborta o loop. **`continue`**: pula para a próxima iteração.
- **`range`**: percorre arrays, slices, maps, strings e channels; retorna índice/chave e valor.
- **Blank identifier `_`**: descarta valores não utilizados do `range` (exigido pelo compilador, que proíbe variáveis declaradas e não usadas).

## Capítulo 06 — Funções

- **DRY (Don't Repeat Yourself)**: copiar e colar o mesmo bloco mais de duas vezes cria dívida técnica; funções existem para eliminar repetição.
- **Função como mini-fábrica**: parâmetros = entrada, corpo = processamento, retorno = saída; quem chama não precisa saber o "como" (abstração).
- **Múltiplos retornos**: superpoder do Go — uma função pode devolver vários valores (`(int, error)`), permitindo o padrão `resultado, err := funcao()` para tratar erro logo após a chamada.
- **Retornos nomeados (named returns)**: variáveis de retorno já nomeadas na assinatura; permitem o "naked return" (`return` sem argumentos) — usar com moderação, só em funções curtas.
- **Funções variádicas**: parâmetro com `...Tipo` aceita zero ou mais argumentos e se comporta como slice dentro da função (ex: `fmt.Println`).
- **Funções anônimas / closures**: funções são valores em Go, podem ser guardadas em variáveis e usadas como callbacks ou lógica descartável.
- **Princípio de função única**: uma função deve fazer só uma coisa e fazê-la bem (facilita testar, ler e reutilizar).

---

## Capítulo 07 — Arrays e Slices

- **Array**: sequência de tamanho fixo definido em compilação; o tamanho faz parte do tipo (`[3]int ≠ [4]int`). Zero-inicializado automaticamente.
- **Slice**: visão dinâmica sobre um Array; composto por ponteiro, `len` (tamanho visível) e `cap` (capacidade até precisar realocar). Estrutura preferida no dia a dia.
- **`append`**: adiciona elementos a um Slice e **sempre deve ser reatribuído** (`slice = append(slice, valor)`). Quando a capacidade esgota, o Go aloca um Array novo (geralmente o dobro) e copia os dados.
- **Fatiamento** `slice[min:max]`: intervalo semi-aberto `[min, max)` — inclui `min`, exclui `max`. Não copia dados; cria uma nova janela sobre o mesmo Array.
- **Aliasing**: quando dois Slices apontam para o mesmo Array subjacente. Modificar um modifica o outro silenciosamente — armadilha clássica ao usar fatiamento + `append`.
- **`make([]T, len, cap)`**: cria um Slice com Array novo e independente, controlando tamanho e capacidade iniciais. Solução padrão para evitar aliasing ao combinar partes de Slices.
- **`copy(dst, src)`**: copia elementos de `src` para `dst`; útil para criar Slices verdadeiramente independentes e liberar Arrays grandes da memória.

## Capítulo 08 — Maps

- **Map**: estrutura de dados do tipo chave→valor. Cada chave é única; o valor associado pode ser acessado, atualizado ou deletado em tempo O(1) médio.
- **Hash Table**: implementação interna de um Map. Uma função hash transforma a chave em um índice de array — daí a velocidade O(1). O Go usa Hash Table por baixo de todo `map[K]V`.
- **`make(map[K]V)`**: forma idiomática de criar um map inicializado e pronto para uso. Um map declarado sem `make` (apenas `var m map[K]V`) é `nil` e causa panic ao tentar escrever.
- **Zero value em Maps**: ler uma chave inexistente **não gera erro** — o Go retorna o zero value do tipo do valor (`0` para `int`, `""` para `string`, `false` para `bool`). Isso permite o padrão `contador[chave]++` sem verificação prévia.
- **`contador[chave]++`**: idioma de contagem: equivalente a `contador[chave] = contador[chave] + 1`. Na primeira ocorrência, pega `0` (zero value) e guarda `1`; nas seguintes, incrementa o valor existente.
- **Comma ok idiom**: `valor, ok := mapa[chave]` — retorna o valor e um `bool` (`true` se a chave existe, `false` se não). Evita a ambiguidade entre "chave ausente" e "chave com zero value".
- **`delete(mapa, chave)`**: remove uma entrada do map. Não retorna erro se a chave não existir — é uma operação segura mesmo para chaves ausentes.

## Capítulo 09 — Ponteiros

- **Ponteiro**: variável que armazena o *endereço de memória* de outra variável — não o valor em si. Analogia: o papel com o endereço da casa, não a casa.
- **Valor vs. Referência**: passar por valor cria uma cópia independente (seguro, mas caro para dados grandes); passar por referência (ponteiro) compartilha o original (eficiente, mas permite modificação acidental).
- **`&` (operador de endereço)**: obtém o endereço de memória de uma variável. `p := &x` faz `p` apontar para `x`.
- **`*` (operador de desreferência)**: "viaja" até o endereço e acessa o valor armazenado lá. `*p` lê o valor; `*p = 99` modifica o original.
- **`*Tipo` (tipo ponteiro)**: `*int` é "ponteiro para int", `*string` é "ponteiro para string". O tipo do ponteiro e do valor são distintos.
- **`nil` em ponteiros**: zero value de qualquer ponteiro. Desreferenciar um ponteiro `nil` (`*p`) causa **panic** em runtime — o erro mais comum com ponteiros.
- **`new(Tipo)`**: aloca memória para um valor zero do tipo, retorna o ponteiro. Equivalente a `var x Tipo; p := &x`. Usado para tipos básicos e Structs.
- **`make` vs `new`**: `make` inicializa Slices, Maps e Channels e retorna o valor (não ponteiro); `new` aloca qualquer tipo e retorna ponteiro.
- **Regra de Ouro**: use ponteiros quando precisar modificar o original dentro de uma função, ou quando a estrutura for grande demais para copiar. Para tipos pequenos (`int`, `bool`), valores são mais rápidos que ponteiros.

## Capítulo 10

*Conteúdo ainda não transcrito/estudado.*

## Capítulo 11 em diante

Adicione uma nova seção `## Capítulo XX — Título` seguindo o mesmo padrão a cada capítulo concluído.
