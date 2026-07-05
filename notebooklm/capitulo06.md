# Capítulo 06 — Funções: A Arte de Decompor Problemas

## Tema Central

Funções são o bloco fundamental de abstração em Go. Elas transformam código repetitivo em unidades reutilizáveis, testáveis e compreensíveis. Go trata funções como **cidadãos de primeira classe** — podem ser armazenadas em variáveis, passadas como argumentos e retornadas por outras funções.

---

## O que o Livro Cobre

- DRY (Don't Repeat Yourself) — motivação para funções
- Funções com múltiplos retornos (`int, error`)
- Retornos nomeados (named returns) e naked return
- Funções variádicas (`...Tipo`)
- Funções anônimas e closures
- Princípio de responsabilidade única

---

## Aprofundando os Conceitos

### A Anatomia de uma Função em Go

```go
//  palavra-chave  nome         parâmetros         retornos
    func           calcular     (a, b int)          (int, error)  {
        // corpo
        return resultado, nil
    }
```

Variações:
```go
func semRetorno() {}
func umRetorno(x int) int { return x * 2 }
func variosRetornos(x int) (int, int, error) { return x, x*2, nil }
func nomesAbreviados(a, b, c int) int { return a + b + c }  // tipos compartilhados
```

### Múltiplos Retornos: o Padrão `(valor, error)`

O padrão mais importante do Go. Em vez de lançar exceções (Python, Java), funções retornam o erro como um valor normal:

```go
func dividir(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("divisão por zero")
    }
    return a / b, nil
}

// Uso — erro tratado imediatamente
resultado, err := dividir(10, 2)
if err != nil {
    log.Fatal(err)
}
fmt.Println(resultado) // 5
```

Vantagens sobre exceções:
- Erros são **visíveis na assinatura** — você sabe o que pode falhar
- O compilador **força o tratamento** (se você ignora a variável, recebe aviso)
- Sem `try/catch` espalhado — cada falha é tratada no ponto onde ocorre

### Retornos Nomeados — quando usar

```go
// Com retorno nomeado: variáveis já existem no escopo da função
func minMax(slice []int) (min, max int) {
    min, max = slice[0], slice[0]
    for _, v := range slice[1:] {
        if v < min {
            min = v
        }
        if v > max {
            max = v
        }
    }
    return  // "naked return" — retorna min e max
}
```

**Quando usar**: funções curtas (< 10 linhas) onde os nomes documentam claramente o que é retornado. Evite naked returns em funções longas — torna difícil saber o que está sendo retornado sem ler toda a função.

### Funções Variádicas — `...Tipo`

```go
func somar(nums ...int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}

somar(1, 2, 3)      // 6
somar(1)            // 1
somar()             // 0

// Passar um slice com ...
numeros := []int{1, 2, 3, 4, 5}
somar(numeros...)   // 15 — o ... "expande" o slice
```

Dentro da função, `nums` é um `[]int` normal. A diferença está na chamada — o `...` torna a lista de argumentos flexível.

`fmt.Println`, `append`, `fmt.Sprintf` são todos variádicos.

### Funções como Valores — First-Class Functions

```go
// Guardar função em variável
dobrar := func(x int) int { return x * 2 }
fmt.Println(dobrar(5))  // 10

// Passar função como argumento (Higher-Order Function)
func aplicar(nums []int, f func(int) int) []int {
    resultado := make([]int, len(nums))
    for i, n := range nums {
        resultado[i] = f(n)
    }
    return resultado
}

dobrados := aplicar([]int{1, 2, 3}, dobrar)  // [2, 4, 6]

// Tipo de função como parâmetro
type Transformador func(int) int

func aplicarTransformador(n int, t Transformador) int {
    return t(n)
}
```

### Closures — Funções que "Lembram"

Uma closure é uma função anônima que **captura variáveis do escopo externo**:

```go
func criarContador() func() int {
    count := 0  // variável capturada pela closure
    return func() int {
        count++    // modifica a variável capturada
        return count
    }
}

contador := criarContador()
fmt.Println(contador())  // 1
fmt.Println(contador())  // 2
fmt.Println(contador())  // 3

contador2 := criarContador()
fmt.Println(contador2()) // 1 — contador INDEPENDENTE
```

Cada chamada de `criarContador()` cria uma nova `count` independente. Isso é útil para:
- Geradores e sequenciadores
- Configurar callbacks com estado
- Encapsular lógica com contexto

### `defer` — execução postergada

```go
func abrirArquivo(nome string) error {
    f, err := os.Open(nome)
    if err != nil {
        return err
    }
    defer f.Close()  // executado quando a função RETORNAR — qualquer caminho
    
    // usar f...
    return nil
}
```

`defer` garante que o recurso seja liberado independentemente de como a função termina (retorno normal, `return` antecipado em guard clause, ou panic). É a forma idiomática de gerenciar recursos em Go (sem `try/finally`).

**Ordem de execução**: múltiplos `defer` executam em **ordem LIFO** (último a entrar, primeiro a sair):

```go
func main() {
    defer fmt.Println("terceiro")  // executado por último
    defer fmt.Println("segundo")
    defer fmt.Println("primeiro")  // executado primeiro
    fmt.Println("durante")
}
// Saída: durante, primeiro, segundo, terceiro
```

---

## Referências Oficiais

- **Especificação — Function declarations**: https://go.dev/ref/spec#Function_declarations
- **Especificação — Variadic functions**: https://go.dev/ref/spec#Passing_arguments_to_..._parameters
- **Effective Go — Functions**: https://go.dev/doc/effective_go#functions
- **Blog: Defer, Panic, and Recover**: https://go.dev/blog/defer-panic-and-recover
- **Blog: First-class functions**: https://go.dev/blog/functions-in-go
- **Tour de Go — Functions**: https://go.dev/tour/basics/4

---

## Exemplos de Código Adicionais

### Função que retorna função (currying)

```go
// Currying: função que configura e retorna outra função
func multiplicador(fator int) func(int) int {
    return func(x int) int {
        return x * fator
    }
}

duplo := multiplicador(2)
triplo := multiplicador(3)

fmt.Println(duplo(5))   // 10
fmt.Println(triplo(5))  // 15
```

### Middleware pattern com funções

```go
type Handler func(string) string

func comLog(h Handler) Handler {
    return func(entrada string) string {
        fmt.Printf("[LOG] processando: %s\n", entrada)
        resultado := h(entrada)
        fmt.Printf("[LOG] resultado: %s\n", resultado)
        return resultado
    }
}

processar := func(s string) string { return s + " processado" }
comLogProcessar := comLog(processar)
comLogProcessar("dado")
```

### `panic` e `recover`

```go
// panic: similar a uma exceção irrecuperável (reserve para bugs, não erros esperados)
func dividirPanic(a, b int) int {
    if b == 0 {
        panic("divisão por zero")
    }
    return a / b
}

// recover: capturar um panic (usado em frameworks e servidores)
func seguro(f func()) (err error) {
    defer func() {
        if r := recover(); r != nil {
            err = fmt.Errorf("panic capturado: %v", r)
        }
    }()
    f()
    return nil
}
```

---

## Perguntas & Respostas Frequentes

**P: Posso ter uma função com mais de dois retornos?**
R: Sim. Não há limite técnico. Na prática, mais de 3 retornos é sinal de que a função está fazendo coisas demais — considere retornar uma struct.

**P: Qual a diferença entre `panic` e retornar `error`?**
R: `error` é para condições esperadas e recuperáveis (arquivo não encontrado, entrada inválida). `panic` é para bugs — condições que **nunca deveriam acontecer** (índice fora do slice, nil pointer dereference). Nos programas Go, `panic` deve ser raro.

**P: Closures modificam a variável original ou uma cópia?**
R: Modificam a **variável original** (compartilhada por referência). É por isso que o bug de captura de variável de loop existe — todas as closures do loop compartilham o mesmo `i`.

**P: O que é uma função anônima?**
R: Uma função sem nome, definida e usada inline. `func(x int) int { return x * 2 }` é uma função anônima. Pode ser atribuída a uma variável ou chamada imediatamente: `func(x int) int { return x }(42)`.

**P: Por que usar `defer` em vez de chamar a função de cleanup diretamente no fim?**
R: Porque `defer` garante a execução **mesmo se a função retornar cedo** (via guard clause ou erro). Chamada direta no fim da função é esquecida quando há múltiplos `return`.

---

## Comparações com Outras Linguagens

### Múltiplos retornos

```python
# Python: retornar tuple
def dividir(a, b):
    if b == 0:
        return None, "divisão por zero"
    return a / b, None
resultado, erro = dividir(10, 2)
```

```java
// Java: sem múltiplos retornos nativos — usa Exception ou wrapper class
// (throws Exception é o equivalente mais próximo)
double dividir(double a, double b) throws ArithmeticException {
    if (b == 0) throw new ArithmeticException("divisão por zero");
    return a / b;
}
```

```go
// Go: múltiplos retornos nativos, erro como valor
func dividir(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("divisão por zero")
    }
    return a / b, nil
}
```

### Closures e captura de variável

```javascript
// JavaScript tem o mesmo bug de captura de loop
var funcs = [];
for (var i = 0; i < 3; i++) {
    funcs.push(function() { return i; });
}
funcs.map(f => f());  // [3, 3, 3]

// Solução com let (block-scoped)
for (let i = 0; i < 3; i++) {
    funcs.push(function() { return i; });
}
funcs.map(f => f());  // [0, 1, 2]
```

---

## Armadilhas Comuns

1. **Ignorar erros**: `resultado, _ := funcaoQueRetornaErro()` — ignorar o erro com `_` silencia falhas. Sempre trate erros.

2. **Named return + defer**: o defer pode **modificar o retorno nomeado** — comportamento surpreendente:
   ```go
   func comDefer() (resultado int) {
       defer func() { resultado++ }()  // modifica o retorno!
       return 1  // retorna 2, não 1
   }
   ```

3. **Passar slice variádico sem `...`**: `f(slice)` não compila se `f` espera `...T`. Use `f(slice...)`.

4. **Closures em goroutines**: capturar variáveis de loop em goroutines amplifica o bug de captura — e o resultado é race condition.

---

## Quiz de Fixação

1. O que é DRY e como funções implementam esse princípio?
2. Por que Go usa retorno de erro em vez de exceções?
3. O que é uma closure e como ela "lembra" do escopo externo?
4. Qual a diferença entre `f(slice)` e `f(slice...)` para uma função variádica?
5. Em qual ordem os `defer` são executados quando há múltiplos?
6. Quando usar `panic` vs retornar `error`?

---

## 🔬 Dissecando a Sintaxe

### Função com múltiplos retornos — dissecar cada elemento

```go
func dividir(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("divisão por zero")
    }
    return a / b, nil
}
```

```
func  dividir  (a, b float64)  (float64, error)  {
────  ───────  ─────────────   ───────────────   ─
 │      │            │                │          └─ abre o bloco da função
 │      │            │                └─ LISTA DE RETORNOS (entre parênteses quando há mais de um)
 │      │            │                    "float64" = tipo do primeiro retorno (o resultado)
 │      │            │                    "error"   = tipo do segundo retorno (o erro ou nil)
 │      │            └─ PARÂMETROS: "a" e "b" compartilham o tipo "float64"
 │      │                equivale a "(a float64, b float64)"
 │      └─ nome da função (minúscula = privada ao pacote)
 └─ palavra-chave

DENTRO DA FUNÇÃO:
  return 0, errors.New("divisão por zero")
  ──────  ─  ────────────────────────────
     │    │          └─ cria um novo valor do tipo error com a mensagem
     │    └─ vírgula separa os dois valores de retorno
     └─ retorna dois valores: o float (0) e o erro

  return a / b, nil
  ──────  ────  ───
     │      │    └─ nil = "sem erro" (zero value da interface error)
     │      └─ resultado da divisão
     └─ retorna em caso de sucesso: o resultado e nil (sem erro)
```

### Closure — como uma função "lembra" do mundo externo

```go
func criarContador() func() int {   // [1]
    count := 0                       // [2]
    return func() int {              // [3]
        count++                      // [4]
        return count                 // [5]
    }
}
```

```
[1]  func criarContador() func() int
                           ──────────
                               └─ TIPO DO RETORNO: esta função retorna outra função
                                   "func() int" = "uma função que não recebe nada e retorna int"

[2]  count := 0
         └─ variável LOCAL de criarContador
             normalmente morreria quando criarContador retornasse
             MAS: a função anônima abaixo a CAPTURA → ela vive na heap

[3]  return func() int {
            ──────────
                └─ função ANÔNIMA: sem nome, criada e retornada inline
                    essa função CAPTURA "count" do escopo externo

[4]  count++
         └─ incrementa a variável "count" CAPTURADA do escopo de criarContador
             não é uma cópia — é a MESMA variável compartilhada
             a cada chamada, count cresce: 1, 2, 3...

[5]  return count
         └─ retorna o valor atual de count
```

### `defer` — quando é executado?

```go
func abrirArquivo(nome string) error {
    f, err := os.Open(nome)     // [1]
    if err != nil {
        return err              // [2]
    }
    defer f.Close()             // [3]
    // usa f...
    return nil                  // [4]
}
```

```
[1]  f, err := os.Open(nome)
          └─ múltiplo retorno: f é o arquivo, err é o possível erro

[2]  return err
         └─ saída ANTECIPADA — f não foi aberto com sucesso
             defer NÃO foi registrado ainda → f.Close() NÃO será chamado (correto!)

[3]  defer f.Close()
      ─────  ───────
        │       └─ f.Close() = método que fecha e libera o arquivo
        └─ DEFER: "execute isso quando a função retornar — qualquer caminho"
            registra a chamada AGORA mas EXECUTA quando a função terminar
            mesmo se houver return antecipado mais à frente

[4]  return nil
         └─ saída normal — AQUI o defer dispara e f.Close() é chamado
```

---

## 🎙️ Prompts para o NotebookLM

### 🎧 Audio Overview
> "Gere um Audio Overview comparando o tratamento de erros em Go (retornar error) vs Python/Java (try/catch/except). Um apresentador defende exceções, o outro defende o estilo Go. Discutam: qual é mais legível, qual força o programador a ser mais responsável, e o que acontece quando erros são ignorados em cada abordagem."

### 📋 Briefing Doc
> "Crie um briefing sobre closures em Go: o que são, como funcionam internamente (captura por referência), onde são mais úteis (contadores, middlewares, configuração), e a armadilha de captura de variável de loop."

### 📚 Study Guide
> "Crie um guia de estudos do Capítulo 06 com: (1) diagrama textual mostrando o ciclo de vida de um defer (quando é registrado vs quando executa), (2) 5 exercícios progressivos de funções (simples → variádica → closure → HOF → panic/recover), (3) quiz de 8 perguntas sobre múltiplos retornos e closures."

### 🖼️ Roteiro de Slides
> "Crie um roteiro de 6 slides sobre 'Funções em Go'. Slide 1: DRY — o problema que funções resolvem. Slide 2: Múltiplos retornos e o padrão (valor, error). Slide 3: Variádica — como funciona o '...'. Slide 4: Funções como valores — higher-order functions. Slide 5: Closures — funções com memória. Slide 6: defer — o guardião de recursos."

### 💬 Perguntas Profundas para o Chat
- "Por que Go usa retorno de erro em vez de exceções? Quais são as vantagens e desvantagens de cada abordagem na prática?"
- "Closures capturam variáveis por referência em Go. Me dê dois exemplos: um onde isso é exatamente o que queremos, e um onde causa um bug sutil."
- "Qual é a ordem de execução quando há múltiplos `defer` em uma função? Me mostre com um exemplo de código e a saída esperada."
- "O que é 'naked return' (return sem argumentos) e por que deve ser usado só em funções curtas?"
