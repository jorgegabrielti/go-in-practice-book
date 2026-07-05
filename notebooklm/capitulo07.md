# Capítulo 07 — Arrays e Slices: A Diferença que Muda Tudo

## Tema Central

Arrays e Slices são superficialmente parecidos, mas fundamentalmente diferentes. Um Array é uma estrutura rígida de tamanho fixo; um Slice é uma janela dinâmica e flexível. No dia a dia do Go, você usa Slices quase que exclusivamente — mas entender o Array por baixo é essencial para evitar o bug de aliasing.

---

## O que o Livro Cobre

- Array: tamanho fixo em compilação; tamanho faz parte do tipo
- Slice: ponteiro + len + cap sobre um Array subjacente
- `append`: sempre reatribuir; dobra a capacidade ao esgotar
- Fatiamento `slice[min:max]`: intervalo `[min, max)`
- Aliasing: duas janelas no mesmo Array
- `make([]T, len, cap)`: cria Slice independente
- `copy(dst, src)`: cópia real entre Slices

---

## Aprofundando os Conceitos

### Array: o que o compilador sabe sobre você

```go
// Tamanho faz parte do tipo
var a [3]int         // [3]int — diferente de [4]int
var b [4]int

a = b  // ERRO: cannot use b (type [4]int) as type [3]int

// Inicialização
c := [3]int{10, 20, 30}   // com valores
d := [...]int{10, 20, 30} // tamanho inferido pelo compilador: [3]int
e := [5]int{1: 10, 3: 30} // inicialização por índice: [0, 10, 0, 30, 0]
```

Arrays em Go são **passados por valor** — a função recebe uma cópia completa. Para arrays grandes, isso é ineficiente. Slices resolvem isso.

### A Estrutura Interna de um Slice

Um Slice não armazena dados — ele é um **descritor** de 3 campos:

```
┌─────────────────────────────────┐
│  Slice                          │
│  ┌──────────┐                   │
│  │ ponteiro ├──────────────────►│ Array subjacente: [10, 20, 30, 40, 50]
│  ├──────────┤                   │
│  │ len: 3   │   (elementos visíveis)
│  ├──────────┤
│  │ cap: 5   │   (elementos até o fim do array)
│  └──────────┘
└─────────────────────────────────┘
```

```go
arr := [5]int{10, 20, 30, 40, 50}
s := arr[1:4]        // Slice sobre arr
fmt.Println(s)       // [20, 30, 40]
fmt.Println(len(s))  // 3
fmt.Println(cap(s))  // 4 (do índice 1 até o fim do array = 4 elementos)
```

### `append` — o segredo dos bastidores

```go
s := []int{1, 2, 3}  // cap = 3

// Adicionar dentro da capacidade: mesmo array, só atualiza len
s = append(s, 4)     // cap = 3, mas len = 4... espera, cap insuficiente!

// Quando cap esgota: Go aloca um NOVO array (geralmente 2x), copia tudo
// e retorna novo Slice apontando para o novo array
```

A estratégia de crescimento de capacidade (doubling) garante que `append` seja amortizado O(1). Se você sabe a quantidade final, use `make` com capacidade pré-alocada:

```go
// Ineficiente: muitas realocações
s := []int{}
for i := 0; i < 1000; i++ {
    s = append(s, i)  // pode realocar muitas vezes
}

// Eficiente: uma única alocação
s := make([]int, 0, 1000)
for i := 0; i < 1000; i++ {
    s = append(s, i)  // nunca realoca
}
```

### Fatiamento: a janela sobre o array

```go
s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

// s[low:high] — inclui low, exclui high
a := s[2:5]   // [2, 3, 4]
b := s[:3]    // [0, 1, 2] — equivale a s[0:3]
c := s[7:]    // [7, 8, 9] — equivale a s[7:len(s)]
d := s[:]     // cópia do slice (mesmo array!)

// Três índices: s[low:high:max] — controla a capacidade
e := s[2:5:7] // [2, 3, 4] com cap = 5 (7-2), não 8
```

### Aliasing: a armadilha mais sutil

```go
original := []int{1, 2, 3, 4, 5}
janela := original[1:3]  // [2, 3] — mesma memória!

janela[0] = 99           // modifica o ORIGINAL também!
fmt.Println(original)    // [1, 99, 3, 4, 5]

// Append dentro da capacidade: também modifica o original!
janela = append(janela, 100)
fmt.Println(original)    // [1, 99, 100, 4, 5] — surpresa!

// Quando append excede a capacidade: alocação nova, sem mais aliasing
janela = append(janela, 200, 300, 400)  // agora janela tem seu próprio array
```

**Regra de ouro**: se você fatiar um Slice e usar `append` no resultado, pode ou não modificar o original — depende da capacidade. Para slices independentes, use `copy`.

### `make` e `copy`: criando Slices verdadeiramente independentes

```go
original := []int{1, 2, 3, 4, 5}

// make: cria array novo
independente := make([]int, len(original))
copy(independente, original)
independente[0] = 99
fmt.Println(original)     // [1, 2, 3, 4, 5] — intacto

// Atalho com append
independente2 := append([]int{}, original...)
```

`copy(dst, src)` retorna o número de elementos copiados: `min(len(dst), len(src))`.

### Slice de Slices (matriz 2D)

```go
// Matriz 3x3
matriz := make([][]int, 3)
for i := range matriz {
    matriz[i] = make([]int, 3)
}
matriz[1][2] = 42
fmt.Println(matriz)  // [[0 0 0] [0 0 42] [0 0 0]]
```

---

## Referências Oficiais

- **Especificação — Slice types**: https://go.dev/ref/spec#Slice_types
- **Blog: Go Slices — usage and internals**: https://go.dev/blog/slices-intro
- **Blog: Arrays, slices e strings (Pike, 2013)**: https://go.dev/blog/slices
- **Effective Go — Slices**: https://go.dev/doc/effective_go#slices
- **Pacote `slices` (Go 1.21+)**: https://pkg.go.dev/slices
- **Pacote `sort`**: https://pkg.go.dev/sort

---

## Exemplos de Código Adicionais

### Remover elemento de um Slice por índice

```go
func remover(s []int, i int) []int {
    // Sem preservar ordem: troca com o último e trunca
    s[i] = s[len(s)-1]
    return s[:len(s)-1]
}

// Preservando ordem: O(n) mas mantém a sequência
func removerOrdem(s []int, i int) []int {
    return append(s[:i], s[i+1:]...)
}
```

### Stack (pilha) com Slice

```go
type Stack[T any] struct {
    itens []T
}

func (s *Stack[T]) Push(v T) {
    s.itens = append(s.itens, v)
}

func (s *Stack[T]) Pop() (T, bool) {
    if len(s.itens) == 0 {
        var zero T
        return zero, false
    }
    v := s.itens[len(s.itens)-1]
    s.itens = s.itens[:len(s.itens)-1]
    return v, true
}
```

### `slices` package (Go 1.21+)

```go
import "slices"

s := []int{3, 1, 4, 1, 5, 9, 2, 6}
slices.Sort(s)                    // ordena in-place
fmt.Println(slices.Contains(s, 5)) // true
idx, found := slices.BinarySearch(s, 5)
fmt.Println(slices.Max(s))        // 9
```

---

## Perguntas & Respostas Frequentes

**P: Qual a diferença entre `nil` slice e slice vazio?**
R: `var s []int` é `nil` (ponteiro nulo). `s := []int{}` é um slice não-nil com comprimento 0. Para a maioria dos fins práticos, ambos se comportam igual: `len(nil_slice) == 0`, `append` funciona nos dois, `range` funciona nos dois. A diferença aparece em serialização JSON (`nil` → `null`, `[]int{}` → `[]`).

**P: `append` é sempre seguro?**
R: Desde que você sempre reatribua o resultado (`s = append(s, v)`). Se você ignorar o retorno (`append(s, v)` sem atribuição), o elemento é descartado. O compilador não avisa.

**P: Quando a capacidade dobra, ela sempre dobra exatamente?**
R: Aproximadamente. O Go usa uma fórmula complexa que considera o tamanho atual, o tipo dos elementos e limites de página de memória. Em versões recentes (Go 1.18+) a estratégia foi ajustada para crescer mais suavemente para slices grandes. Mas para finalidades práticas, "dobra" é uma boa aproximação.

**P: Como saber se dois Slices apontam para o mesmo Array?**
R: Usando o pacote `reflect` ou comparando os ponteiros: `&a[0] == &b[0]`. Mas em código normal você não precisa fazer isso — o design correto evita aliasing.

**P: Arrays são passados por valor ou referência?**
R: **Por valor** — sempre copiados. Slices são passados por valor também (cópia do descritor: ponteiro + len + cap), mas o ponteiro aponta para o mesmo array subjacente. Por isso modificações no slice dentro de uma função afetam o original; mas `append` na função não afeta o len/cap do slice externo.

---

## Comparações com Outras Linguagens

### Arrays e listas

```python
# Python: list é dinâmica, sem tipo fixo
lista = [1, 2, 3]
lista.append(4)   # dinâmico
lista[0] = "texto"  # tipagem dinâmica permite
```

```java
// Java: array é tamanho fixo; ArrayList é dinâmica
int[] array = new int[3];
ArrayList<Integer> lista = new ArrayList<>();
lista.add(4);
```

```javascript
// JavaScript: Array é sempre dinâmico
const arr = [1, 2, 3];
arr.push(4);
arr[10] = 99;  // gap preenchido com undefined
```

```go
// Go: Array (fixo) vs Slice (dinâmico)
arr := [3]int{1, 2, 3}  // array: fixo, parte do tipo
s := []int{1, 2, 3}     // slice: dinâmico
s = append(s, 4)
```

### Cópia de arrays

```python
# Python: cópia rasa com slice
copia = lista[:]        # ou list(lista)
```

```go
// Go: copy() ou append([]int{}, s...)
copia := make([]int, len(s))
copy(copia, s)
```

---

## Armadilhas Comuns

1. **Não reatribuir `append`**: `append(s, v)` sem `s =` descarta o resultado silenciosamente.

2. **Fatiamento cria aliasing**: `janela := s[1:3]` não copia dados — modificações se propagam.

3. **`len` vs `cap` em comparações**: checar `len(s) == 0` é seguro para slice nil ou vazio. Nunca comparar `s == nil` para "slice vazio" — um slice vazio não-nil `!= nil`.

4. **Crescimento de Slice invalida ponteiros**: após `append` que realoca, qualquer ponteiro para elementos antigos do slice fica dangling. Nunca guarde ponteiros para elementos de slices que sofrerão append.

---

## Quiz de Fixação

1. Por que `[3]int` e `[4]int` são tipos diferentes em Go?
2. O que são os três campos internos de um Slice?
3. O que acontece com a capacidade quando `append` excede o limit?
4. Por que modificar `janela[0]` pode modificar o slice original?
5. Qual a diferença entre `copy(dst, src)` e `dst = append(dst, src...)`?
6. Como criar uma cópia verdadeiramente independente de um Slice?

---

## 🔬 Dissecando a Sintaxe

### `make` para criar Slice — anatomia completa

```go
s := make([]int, 5, 10)
```

```
make  (  []int  ,  5  ,  10  )
────     ─────     ─     ──
  │        │       │      └─ capacidade (cap): quantos elementos o array subjacente comporta
  │        │       │          OPCIONAL: se omitido, cap = len
  │        └─ comprimento (len): quantos elementos estão visíveis e acessíveis agora
  │          (todos inicializados com zero value: [0, 0, 0, 0, 0])
  └─ tipo do slice: "[]int" = slice de inteiros
      "[" "]" sem número = SLICE (com número = Array: "[5]int")

resultado: s = [0, 0, 0, 0, 0] com len=5 e cap=10
→ posso fazer append mais 5 vezes sem realocar o array subjacente
```

### Fatiamento `s[low:high]` — a mecânica da janela

```go
s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
a := s[2:5]
```

```
s  [  2  :  5  ]
─     ─     ─
│     │     └─ HIGH (excluído): o slice vai ATÉ o índice 5, mas NÃO inclui o 5
│     │         → último elemento incluído = índice 4
│     └─ LOW (incluído): o slice começa NO índice 2 (elemento valor=2)
└─ o slice original

resultado: a = [2, 3, 4]   ← elementos nos índices 2, 3, 4
           len(a) = 3       ← high - low = 5 - 2 = 3
           cap(a) = 8       ← do índice 2 até o fim do array original = 10 - 2 = 8

REGRA: s[low:high]  → inclui low, EXCLUI high
       "intervalo semi-aberto [low, high)"
```

### `append` — por que SEMPRE reatribuir

```go
s = append(s, 42)
─    ──────  ─  ──
│      │     │   └─ valor(es) a adicionar (variádico: pode ser um ou muitos)
│      │     └─ o slice original
│      └─ função builtin: adiciona elementos ao fim do slice
└─ REATRIBUIÇÃO OBRIGATÓRIA ← essa é a parte que confunde

POR QUE REATRIBUIR?
  append pode retornar:
  (a) o MESMO slice com len incrementado (se cap era suficiente)
  (b) um NOVO slice apontando para um NOVO array maior (se cap foi excedido)

  Se você não reatribuir → o retorno é descartado → parece que nada aconteceu
  O compilador NÃO avisa sobre isso — é um bug silencioso comum
```

### `copy` — a cópia real (sem aliasing)

```go
dst := make([]int, len(src))
n := copy(dst, src)
```

```
copy  (  dst  ,  src  )
────     ───     ───
  │       │       └─ slice ORIGEM (de onde copiar)
  │       └─ slice DESTINO (para onde copiar)
  └─ função builtin

retorno "n": número de elementos copiados = min(len(dst), len(src))

DIFERENÇA vs fatiamento:
  janela := src[1:3]   → NÃO copia dados — cria outra janela no MESMO array
  copy(dst, src)        → copia os DADOS — arrays independentes
  modificar dst NÃO afeta src (e vice-versa)
```

---

## 🎙️ Prompts para o NotebookLM

### 🎧 Audio Overview
> "Gere um Audio Overview explicando a diferença entre Array e Slice em Go como uma metáfora física. Explique: por que Slice não armazena dados (só o endereço), o que são len e cap, por que fatiamento cria aliasing, e o bug que ninguém espera quando dois slices compartilham o mesmo array e um faz append."

### 📋 Briefing Doc
> "Crie um briefing sobre 'Aliasing em Slices Go — o bug invisível', explicando: (1) quando aliasing acontece, (2) como detectar que dois slices compartilham o mesmo array, (3) três formas de criar um slice verdadeiramente independente."

### 📚 Study Guide
> "Crie um guia de estudos do Capítulo 07 com: (1) diagrama textual da estrutura interna de um Slice (ponteiro + len + cap), (2) exercício passo a passo de fatiamento (trace o estado de len/cap após cada operação), (3) 8 perguntas sobre append, copy e aliasing."

### 🖼️ Roteiro de Slides
> "Crie um roteiro de 6 slides sobre 'Arrays e Slices em Go'. Slide 1: Array vs Slice — diferença fundamental. Slide 2: Os 3 campos internos de um Slice. Slide 3: make — len vs cap — quando usar cada. Slide 4: Fatiamento — o intervalo semi-aberto. Slide 5: O problema do aliasing com exemplos. Slide 6: copy vs append como solução."

### 💬 Perguntas Profundas para o Chat
- "Por que `append(s, v)` sem atribuição é um bug silencioso? O compilador não deveria avisar?"
- "O que acontece com a memória quando faço `pequeno := grande[0:3]` e descarto `grande`? O array grande ainda existe na memória?"
- "Qual a diferença entre `s := []int{}`, `var s []int` e `s := make([]int, 0)`? São equivalentes em tudo?"
- "Como a estratégia de duplicar a capacidade em append garante O(1) amortizado? Me explique com um cálculo simples."
