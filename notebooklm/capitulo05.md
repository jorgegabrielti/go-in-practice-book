# Capítulo 05 — Laços de Repetição: O `for` de Mil Formas

## Tema Central

Go tem apenas **uma** palavra-chave de repetição: `for`. Sem `while`, sem `do/while`, sem `foreach`. Isso não é limitação — é simplicidade deliberada. O `for` cobre todos os casos, e o `range` transforma a iteração sobre coleções em algo natural e seguro.

---

## O que o Livro Cobre

- `for` clássico: `for init; cond; pós { }`
- `for` como `while`: `for cond { }`
- Loop infinito: `for { }`
- `break` e `continue`
- `range` para arrays, slices, maps, strings e channels
- Blank identifier `_` para descartar valores não usados

---

## Aprofundando os Conceitos

### As Cinco Formas do `for`

```go
// Forma 1: clássico (init; cond; post)
for i := 0; i < 10; i++ {
    fmt.Println(i)
}

// Forma 2: como while (só condição)
i := 0
for i < 10 {
    fmt.Println(i)
    i++
}

// Forma 3: loop infinito
for {
    entrada := lerTeclado()
    if entrada == "sair" {
        break
    }
}

// Forma 4: range com índice e valor
frutas := []string{"maçã", "banana", "uva"}
for i, fruta := range frutas {
    fmt.Printf("%d: %s\n", i, fruta)
}

// Forma 5: range só com índice (ou só com valor)
for i := range frutas {
    fmt.Println(i)          // 0, 1, 2
}
for _, fruta := range frutas {
    fmt.Println(fruta)      // maçã, banana, uva
}
```

### O `range` em detalhes — o que ele retorna para cada tipo

```go
// Slice/Array: (index int, value T)
for i, v := range []int{10, 20, 30} {}

// Map: (key K, value V) — ORDEM NÃO GARANTIDA
for chave, valor := range map[string]int{"a": 1} {}

// String: (byteIndex int, rune rune) — percorre runes, não bytes!
for i, r := range "café" {
    fmt.Printf("índice %d: rune %c (U+%04X)\n", i, r, r)
}
// índice 0: rune c (U+0063)
// índice 1: rune a (U+0061)
// índice 2: rune f (U+0066)
// índice 3: rune é (U+00E9)  ← note: índice pula de 3 para 5 se houvesse mais

// Channel: só value — bloqueia até receber
for msg := range canal { }

// Integer (Go 1.22+): range sobre um inteiro
for i := range 5 {
    fmt.Println(i) // 0, 1, 2, 3, 4
}
```

**Atenção com strings e range**: o `range` sobre `string` itera por **runes**, não por bytes. O índice `i` é o índice do **byte** de início da rune, não o índice da rune em si. Para caracteres multibyte (como 'é' = 2 bytes), o índice pula mais de 1.

### `break` e `continue` com labels

Labels permitem sair de loops externos quando há loops aninhados:

```go
externo:
for i := 0; i < 5; i++ {
    for j := 0; j < 5; j++ {
        if j == 3 {
            continue externo  // pula para a próxima iteração do loop EXTERNO
        }
        if i == 2 {
            break externo     // sai completamente do loop EXTERNO
        }
        fmt.Printf("(%d,%d) ", i, j)
    }
}
```

Labels são raros mas essenciais em algoritmos com matrizes ou busca em grades.

### `defer` — não é parte do `for`, mas frequentemente aparece em loops

```go
for _, arquivo := range arquivos {
    f, err := os.Open(arquivo)
    if err != nil {
        continue
    }
    defer f.Close()  // CUIDADO! defer acumula até a função retornar, não até o fim do loop
}

// Solução: encapsular em função anônima
for _, arquivo := range arquivos {
    func() {
        f, err := os.Open(arquivo)
        if err != nil {
            return
        }
        defer f.Close()  // agora fecha ao fim da função anônima — correto
        processar(f)
    }()
}
```

### Modificar variável de loop: a armadilha clássica

```go
// ARMADILHA: capturar endereço da variável de loop
funcs := make([]func(), 3)
for i := 0; i < 3; i++ {
    funcs[i] = func() { fmt.Println(i) }  // captura o ENDEREÇO de i!
}
for _, f := range funcs {
    f()  // imprime 3, 3, 3 — não 0, 1, 2!
}

// SOLUÇÃO: criar cópia da variável de loop
for i := 0; i < 3; i++ {
    i := i  // nova variável i no escopo do loop body
    funcs[i] = func() { fmt.Println(i) }
}
// Go 1.22+: o range cria uma nova variável por iteração automaticamente (fix!)
```

---

## Referências Oficiais

- **Especificação — For statements**: https://go.dev/ref/spec#For_statements
- **Especificação — Range clause**: https://go.dev/ref/spec#For_range
- **Effective Go — For**: https://go.dev/doc/effective_go#for
- **Go 1.22 Release Notes (range sobre inteiros)**: https://go.dev/doc/go1.22
- **Tour de Go — For**: https://go.dev/tour/flowcontrol/1

---

## Exemplos de Código Adicionais

### Fibonacci com loop

```go
func fibonacci(n int) []int {
    if n <= 0 {
        return nil
    }
    seq := make([]int, n)
    seq[0] = 0
    if n > 1 {
        seq[1] = 1
    }
    for i := 2; i < n; i++ {
        seq[i] = seq[i-1] + seq[i-2]
    }
    return seq
}
```

### Busca binária — clássico de algoritmo com loop

```go
func buscaBinaria(slice []int, alvo int) int {
    esquerda, direita := 0, len(slice)-1
    for esquerda <= direita {
        meio := (esquerda + direita) / 2
        switch {
        case slice[meio] == alvo:
            return meio
        case slice[meio] < alvo:
            esquerda = meio + 1
        default:
            direita = meio - 1
        }
    }
    return -1  // não encontrado
}
```

### Loop com `range` sobre map em ordem determinística

```go
// Maps em Go têm ordem aleatória — para iterar em ordem:
import "sort"

m := map[string]int{"banana": 3, "maçã": 1, "uva": 2}
chaves := make([]string, 0, len(m))
for k := range m {
    chaves = append(chaves, k)
}
sort.Strings(chaves)

for _, k := range chaves {
    fmt.Printf("%s: %d\n", k, m[k])
}
// banana: 3, maçã: 1, uva: 2 (em ordem alfabética)
```

---

## Perguntas & Respostas Frequentes

**P: Por que Go não tem `while`?**
R: `for cond { }` é o `while`. Uma só palavra-chave reduz o vocabulário da linguagem sem perder expressividade. Veja o FAQ oficial: https://go.dev/doc/faq#no_while

**P: O `range` copia os valores?**
R: Sim — o valor retornado pelo `range` é uma **cópia**. Modificar `v` em `for i, v := range slice` não modifica o slice original. Para modificar, use o índice: `slice[i] = novoValor`.

**P: O `range` em um map tem ordem garantida?**
R: Não. A especificação do Go diz explicitamente que a ordem de iteração sobre maps é aleatória e pode mudar entre execuções. Isso é intencional (evitar dependência em ordem de hash). Para ordem determinística, colete as chaves em um slice e ordene.

**P: Posso usar `break` dentro de `select`?**
R: Sim. `break` dentro de `select` (para channels) encerra o `select`, não o loop externo (se houver). Use labels para sair do loop externo.

**P: O que acontece se o slice for `nil` em um `range`?**
R: Iterar sobre um slice `nil` é seguro — o loop simplesmente não executa nenhuma iteração (zero iterações). Mesmo comportamento para map `nil`.

---

## Comparações com Outras Linguagens

### `for` vs `while` vs `foreach`

```python
# Python: for, while, list comprehension
for i in range(10):       # foreach style
    pass
while condicao:           # while
    pass
[x**2 for x in range(10)]  # comprehension
```

```java
// Java: for, while, do-while, enhanced-for
for (int i = 0; i < 10; i++) {}
while (cond) {}
do {} while (cond);
for (String s : lista) {}  // enhanced for
```

```javascript
// JavaScript: for, while, do-while, for...of, for...in, forEach
for (let i = 0; i < 10; i++) {}
for (const item of array) {}    // iterable
for (const key in obj) {}       // enumerable properties
array.forEach(x => {})
```

```go
// Go: só for — cobre todos os casos
for i := 0; i < 10; i++ {}      // for clássico
for cond {}                      // while
for {}                           // loop infinito
for i, v := range colecao {}    // foreach
```

### Captura de variável de loop

```python
# Python tem o mesmo bug com closures!
funcs = [lambda: i for i in range(3)]
[f() for f in funcs]  # [2, 2, 2] — mesmo problema!
# Solução: lambda i=i: i (default argument)
```

```java
// Java: variável de loop implicitamente final em lambdas
for (int i = 0; i < 3; i++) {
    int j = i;  // necessário — i não é efetivamente final
    funcs.add(() -> System.out.println(j));  // OK
}
```

---

## Armadilhas Comuns

1. **`defer` dentro de loop** acumula até a função retornar — abra e feche recursos na mesma iteração com função anônima.

2. **Modificar slice enquanto itera**: pode causar comportamento indefinido se `append` realocar o slice. Colete modificações e aplique depois.

3. **Range sobre string retorna rune, não byte**: o índice é de byte, não de rune — iteração com `i++` em uma string pode partir runes multibyte ao meio.

4. **Loop infinito acidental**: `for i := 0; i < len(slice); ` sem o incremento é um loop infinito. Go compila mas roda para sempre.

---

## Quiz de Fixação

1. Quais são as cinco formas do `for` em Go?
2. O que `range` retorna para cada tipo (slice, map, string, channel)?
3. Por que iterar sobre um map várias vezes pode retornar ordens diferentes?
4. O que acontece se você modificar o valor retornado por `range`? O slice original muda?
5. Como usar `break` para sair de um loop externo quando há loops aninhados?
6. Por que `defer f.Close()` dentro de um loop `for` pode ser problemático?

---

## 🔬 Dissecando a Sintaxe

### O `for` com `range` — a forma mais usada no dia a dia

```go
frutas := []string{"maçã", "banana", "uva"}

for i, fruta := range frutas {
    fmt.Printf("%d: %s\n", i, fruta)
}
```

```
for  i, fruta  :=  range  frutas  {
───  ─  ─────   ─  ─────  ──────  ─
 │   │    │     │    │       │    └─ abre o bloco do loop
 │   │    │     │    │       └─ o slice/map/string/channel a percorrer
 │   │    │     │    └─ palavra-chave que "abre" a coleção para iteração
 │   │    │     └─ ":=" declara as variáveis i e fruta como NOVAS a cada iteração
 │   │    └─ "fruta" = o VALOR do elemento atual (cópia do elemento do slice)
 │   └─ "i" = o ÍNDICE (posição) do elemento atual: 0, 1, 2...
 └─ palavra-chave do loop

O QUE ACONTECE A CADA ITERAÇÃO:
  Iteração 1: i=0, fruta="maçã"
  Iteração 2: i=1, fruta="banana"
  Iteração 3: i=2, fruta="uva"
  → loop termina automaticamente ao esgotar a coleção

DESCARTAR COM _:
  for _, fruta := range frutas  → ignora o índice (compilador exige usar ou descartar)
  for i := range frutas         → ignora o valor (só dois-pontos-igual sem vírgula)
```

### O `for` clássico — dissecar cada parte

```go
for i := 0; i < 10; i++ {
    fmt.Println(i)
}
```

```
for  i := 0  ;  i < 10  ;  i++  {
───  ──────   ─  ──────  ─  ───  ─
 │      │     │     │    │   │   └─ abre o bloco
 │      │     │     │    │   └─ PÓS-ITERAÇÃO: executado no FIM de cada iteração
 │      │     │     │    │       "i++" equivale a "i = i + 1"
 │      │     │     │    └─ segundo separador ";"
 │      │     │     └─ CONDIÇÃO: avaliada ANTES de cada iteração
 │      │     │         se false → loop termina
 │      │     └─ primeiro separador ";"
 │      └─ INICIALIZAÇÃO: executada UMA vez antes do loop começar
 │          ":=" declara i como nova variável (só existe dentro do for)
 └─ palavra-chave

ORDEM DE EXECUÇÃO:
  1. i := 0           (inicialização — uma vez)
  2. i < 10?          (condição — verifica)
  3. fmt.Println(i)   (bloco — executa)
  4. i++              (pós — incrementa)
  5. Volta para 2... repete até condição ser false
```

### `break` com label — saindo do loop externo

```go
externo:                         // [1]
for i := 0; i < 5; i++ {
    for j := 0; j < 5; j++ {
        if j == 3 {
            break externo        // [2]
        }
    }
}
```

```
[1]  externo:
      ────────
         └─ LABEL: um rótulo que nomeia o bloco de código a seguir
             nome seguido de ":" — por convenção em minúsculo
             pode ser qualquer nome válido de identificador

[2]  break externo
      ─────  ───────
        │       └─ nome do label: "sai do bloco marcado com 'externo:'"
        └─ palavra-chave: interrompe a execução
            SEM o label: só sairia do for interno (j)
            COM o label: sai do for externo (i) — ambos encerram
```

---

## 🎙️ Prompts para o NotebookLM

### 🎧 Audio Overview
> "Gere um Audio Overview explicando o `for` do Go para um desenvolvedor JavaScript. O apresentador explica: por que Go tem só uma palavra de loop, o que `range` faz por baixo dos panos, e a armadilha de captura de variável de loop em closures — com o exemplo do bug que imprime '3, 3, 3' em vez de '0, 1, 2'."

### 📋 Briefing Doc
> "Crie um briefing sobre 'As cinco formas do for em Go', com um exemplo de código para cada forma e a situação ideal de uso de cada uma."

### 📚 Study Guide
> "Crie um guia de estudos do Capítulo 05 com: (1) tabela de o que `range` retorna para cada tipo (slice, map, string, channel), (2) 5 exercícios de loop progressivos (Fibonacci, busca linear, contagem de vogais, inversão de string, busca binária), (3) quiz de 8 perguntas sobre `range`, `break` e `continue`."

### 🖼️ Roteiro de Slides
> "Crie um roteiro de 6 slides sobre 'Laços de Repetição em Go'. Slide 1: Por que só existe `for`. Slide 2: As 5 formas do for (código lado a lado). Slide 3: O que range retorna para cada tipo. Slide 4: break e continue — diferença. Slide 5: Labels — quando usar. Slide 6: A armadilha de defer dentro de loop."

### 💬 Perguntas Profundas para o Chat
- "Me explique o bug de captura de variável de loop em closures: por que `for i := 0; i < 3; i++ { go func() { fmt.Println(i) }() }` imprime '3, 3, 3' e não '0, 1, 2'?"
- "Por que o `range` de um map retorna ordens diferentes a cada execução? Isso é um bug ou uma feature?"
- "O que acontece com a memória quando um slice enorme é fatiado e a janela pequena é mantida? Como `copy` resolve isso?"
- "Qual a diferença entre `for range s` e `for i := 0; i < len(s); i++` ao iterar sobre uma string com acentos?"
