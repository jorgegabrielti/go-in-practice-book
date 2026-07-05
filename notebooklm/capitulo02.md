# Capítulo 02 — Variáveis, Constantes e Tipos de Dados

## Tema Central

Go é uma linguagem de tipagem **estática e forte**: toda variável tem um tipo definido em tempo de compilação, e não há conversão implícita entre tipos diferentes. Ao mesmo tempo, Go tem **inferência de tipo** com `:=`, o que a torna expressiva sem a verbosidade de Java.

---

## O que o Livro Cobre

- `var` para declaração explícita (dentro ou fora de funções)
- `:=` para declaração com inferência de tipo (só dentro de funções)
- Zero Value: valor padrão de cada tipo quando declarado sem inicialização
- `const` para valores imutáveis em tempo de compilação

---

## Aprofundando os Conceitos

### As Quatro Formas de Declarar uma Variável

```go
// Forma 1: var com tipo e valor
var nome string = "Jorge"

// Forma 2: var com inferência de tipo
var nome = "Jorge"

// Forma 3: := (short variable declaration) — SÓ dentro de funções
nome := "Jorge"

// Forma 4: var block (para múltiplas variáveis)
var (
    nome  string  = "Jorge"
    idade int     = 30
    ativo bool    = true
)
```

**Quando usar cada forma:**
- `:=` é o padrão dentro de funções (mais conciso)
- `var` é obrigatório em nível de pacote (fora de funções)
- `var` também é preferido quando o valor inicial é o zero value (deixa a intenção clara)

### Zero Value: a decisão de design mais elegante do Go

Em C, variáveis não inicializadas contêm lixo de memória — comportamento indefinido e bugs devastadores. Em Java, campos de objeto têm zero value mas variáveis locais não são inicializadas (erro de compilação se usadas). Em Go, **toda variável tem zero value sempre**, sem exceção.

| Tipo | Zero Value |
|------|-----------|
| `int`, `int8`, `int16`, `int32`, `int64` | `0` |
| `uint`, `uint8`, `uint16`, `uint32`, `uint64` | `0` |
| `float32`, `float64` | `0.0` |
| `complex64`, `complex128` | `(0+0i)` |
| `bool` | `false` |
| `string` | `""` (string vazia) |
| `pointer` | `nil` |
| `slice` | `nil` |
| `map` | `nil` |
| `channel` | `nil` |
| `interface` | `nil` |
| `struct` | todos os campos com seus zero values |

Isso torna o código mais previsível e elimina toda uma classe de bugs.

### `const` vs `var`: a diferença real

```go
const Pi = 3.14159265358979323846  // definido em compilação, imutável
var   pi = 3.14159265358979323846  // definido em runtime, mutável

Pi = 3.0  // ERRO DE COMPILAÇÃO
pi = 3.0  // OK
```

**iota**: o gerador automático de constantes:

```go
type DiaDaSemana int

const (
    Segunda DiaDaSemana = iota  // 0
    Terca                       // 1
    Quarta                      // 2
    Quinta                      // 3
    Sexta                       // 4
    Sabado                      // 5
    Domingo                     // 6
)
```

`iota` é reiniciado a 0 em cada bloco `const`. É o mecanismo Go para simular enumerações.

```go
// Exemplo avançado com iota: potências de 2
const (
    _  = iota             // ignora o 0
    KB = 1 << (10 * iota) // 1 << 10 = 1024
    MB                    // 1 << 20 = 1.048.576
    GB                    // 1 << 30 = 1.073.741.824
)
```

### Escopo de variáveis

```go
var global = "visível em todo o pacote"

func main() {
    local := "visível só em main"
    
    if condicao := true; condicao {
        dentroDoIf := "visível só aqui"
        _ = local      // OK
        _ = dentroDoIf // OK
    }
    // dentroDoIf não existe mais aqui
    // condicao não existe mais aqui
}
```

### Blank identifier `_`

O compilador Go proíbe variáveis declaradas e não usadas. O `_` é a válvula de escape:

```go
_, err := funcaoQueRetornaDoisValores()
// Estamos dizendo: "o primeiro valor não me interessa"
```

---

## Referências Oficiais

- **Especificação — Declarações de variáveis**: https://go.dev/ref/spec#Variable_declarations
- **Especificação — Short variable declarations**: https://go.dev/ref/spec#Short_variable_declarations
- **Especificação — Constantes**: https://go.dev/ref/spec#Constant_declarations
- **Especificação — iota**: https://go.dev/ref/spec#Iota
- **Especificação — Zero values**: https://go.dev/ref/spec#The_zero_value
- **Effective Go — Declarations and scope**: https://go.dev/doc/effective_go#declarations_and_scope
- **Blog: Constants (Pike)**: https://go.dev/blog/constants

---

## Exemplos de Código Adicionais

### Multiple assignment (swap sem variável temporária)

```go
a, b := 10, 20
fmt.Println(a, b)  // 10 20

a, b = b, a        // swap!
fmt.Println(a, b)  // 20 10
```

Em Python isso também funciona (`a, b = b, a`). Em Java você precisa de uma variável temporária.

### Constantes untyped (sem tipo explícito)

```go
const N = 500         // untyped integer constant
var x int = N         // OK
var y float64 = N     // OK também — N é "adaptável"
var z int32 = N       // OK

const M int = 500     // typed integer constant
var w float64 = M     // ERRO: cannot use M (type int) as type float64
```

Constantes untyped têm um "tipo padrão" (para `500` é `int`) mas se adaptam ao contexto. Constantes tipadas são rígidas como variáveis.

---

## Perguntas & Respostas Frequentes

**P: Posso reatribuir uma variável declarada com `:=`?**
R: Sim. `:=` só proíbe re-declarar no mesmo escopo, não reatribuir. `x := 1; x = 2` é válido. O que não é válido é `x := 1; x := 2` (dois := para o mesmo nome no mesmo escopo).

**P: Qual a diferença entre `nil` e zero value?**
R: `nil` é o zero value de tipos que são ponteiros por natureza: slices, maps, channels, funções, interfaces, e ponteiros. Para `int`, o zero value é `0`; para `string`, é `""`. `nil` não existe para esses tipos básicos.

**P: Por que o Go não tem `null` como Java?**
R: `null` em Java causa `NullPointerException` — o erro mais comum em Java. Go usa zero value por padrão (previsível) e `nil` só onde faz sentido semântico. Além disso, o Go encourage verificação explícita de `nil` antes de usar, tornando o código mais robusto.

**P: Posso declarar uma constante dentro de uma função?**
R: Sim. Constantes podem ser declaradas tanto em nível de pacote quanto dentro de funções.

**P: `var x int` é diferente de `var x int = 0`?**
R: São equivalentes. O Go garante que `x` será `0` nos dois casos. A primeira forma deixa mais claro que a inicialização com zero é **intencional** (você quer o zero value), enquanto a segunda é redundante.

---

## Comparações com Outras Linguagens

### Declaração de variáveis

```python
# Python: tipagem dinâmica, sem declaração
nome = "Jorge"
nome = 42  # válido! tipo mudou silenciosamente
```

```java
// Java: tipagem estática explícita, verbosa
String nome = "Jorge";
int idade = 30;
final double PI = 3.14;  // constante
```

```javascript
// JavaScript: var (function-scoped), let (block-scoped), const
let nome = "Jorge";
const PI = 3.14;
// typeof dinâmico — pode mudar de tipo
```

```go
// Go: tipagem estática com inferência
nome := "Jorge"     // string, imutável de tipo
const PI = 3.14     // constante
nome = 42           // ERRO DE COMPILAÇÃO
```

### Zero Value vs comportamento de outras linguagens

```java
// Java — variável local não inicializada: erro de compilação
// Campo de classe: null/0/false (como Go)
int x;
System.out.println(x);  // ERRO: variable x might not have been initialized
```

```python
# Python — sem zero value; usar antes de atribuir é NameError
print(x)  # NameError: name 'x' is not defined
```

```go
// Go — zero value SEMPRE garantido
var x int
fmt.Println(x)  // 0 — sem erro, sem lixo de memória
```

---

## Armadilhas Comuns

1. **`:=` fora de função**: `x := 10` no nível do pacote causa erro. Use `var x = 10`.

2. **Shadowing**: uma variável interna com o mesmo nome "esconde" a externa:
   ```go
   x := 10
   if true {
       x := 20  // NOVA variável, não reatribuição!
       fmt.Println(x)  // 20
   }
   fmt.Println(x)  // 10 — a de fora não mudou
   ```

3. **Constantes não podem ser endereçadas**: `&Pi` não compila. Constantes não têm endereço de memória.

4. **iota em múltiplos blocos const**: iota é reiniciado para 0 em cada novo bloco `const`, não é global.

---

## Quiz de Fixação

1. Qual a diferença entre `var x int` e `x := 0`?
2. Por que `:=` não funciona fora de funções?
3. O que o `iota` faz em um bloco `const`?
4. Se eu declarar `var m map[string]int` sem inicializar, qual é o valor de `m`?
5. O que é shadowing e por que é perigoso?
6. Por que o Go não tem `null`?

---

## 🔬 Dissecando a Sintaxe

### As quatro formas de declaração — anatomia comparada

```go
var nome string = "Jorge"   // Forma 1
var nome = "Jorge"          // Forma 2
nome := "Jorge"             // Forma 3
```

```
Forma 1:  var  nome  string  =  "Jorge"
           ───  ────  ──────  ─  ───────
            │    │      │     │     └─ valor inicial (literal string — sempre entre aspas duplas)
            │    │      │     └─ operador de atribuição
            │    │      └─ tipo explícito: "esta variável só pode guardar texto"
            │    └─ nome da variável (minúscula = privada ao pacote)
            └─ palavra-chave de declaração (obrigatória fora de funções)

Forma 2:  var  nome  =  "Jorge"
           ───  ────  ─  ───────
            │    │    │     └─ Go INFERE o tipo: como "Jorge" é string, nome vira string
            │    │    └─ operador de atribuição
            │    └─ nome da variável
            └─ palavra-chave var (ainda necessária — infere o tipo, não a palavra-chave)

Forma 3:  nome  :=  "Jorge"
           ────  ──  ───────
            │     │     └─ valor inicial (Go infere o tipo)
            │     └─ ":=" = declaração + atribuição em uma operação só
            │         (":=" SÓ funciona DENTRO de funções — nunca no nível do pacote)
            └─ nome da variável
```

### Anatomia do `iota` — contador automático de constantes

```go
type DiaDaSemana int       // [1]

const (                    // [2]
    Segunda DiaDaSemana = iota  // [3] [4] [5]
    Terca                       // [6]
    Quarta                      // [7]
)
```

```
[1]  type DiaDaSemana int
      ────  ────────────  ───
        │         │         └─ tipo base: DiaDaSemana é um int por baixo
        │         └─ nome do novo tipo (maiúscula = exportado)
        └─ palavra-chave: cria um novo tipo distinto (não é alias — é tipo novo)

[2]  const (
      ─────
         └─ abre um bloco de constantes (todas compartilham o mesmo iota)

[3]  Segunda
      ───────
         └─ nome da constante (maiúscula = exportada para outros pacotes)

[4]  DiaDaSemana
      ───────────
         └─ tipo da constante (poderia ser omitido nas seguintes — Go herda o tipo)

[5]  = iota
      ──────
         └─ iota é um "contador mágico" do compilador
             começa em 0 no primeiro item de cada bloco const
             incrementa automaticamente a cada nova linha do bloco
             Segunda = 0, Terca = 1, Quarta = 2...

[6]  Terca
      ─────
         └─ tipo e expressão OMITIDOS: Go repete "DiaDaSemana = iota"
             iota agora vale 1 → Terca = 1

[7]  Quarta
         └─ iota agora vale 2 → Quarta = 2
```

---

## 🎙️ Prompts para o NotebookLM

### 🎧 Audio Overview
> "Gere um Audio Overview onde dois programadores debatem: 'Por que o Go tem Zero Value em vez de null?' Um deles vem de Java e sofreu com NullPointerException. O outro explica como o Zero Value do Go resolve o problema de forma elegante. Incluam exemplos de bugs reais que Zero Value previne."

### 📋 Briefing Doc
> "Crie um briefing sobre declaração de variáveis em Go comparando `:=` vs `var`, explicando em quais situações cada um deve ser usado, com exemplos de código para cada caso."

### 📚 Study Guide
> "Crie um guia de estudos do Capítulo 02 com: (1) tabela completa de zero values para cada tipo do Go, (2) 6 perguntas de múltipla escolha sobre `:=` vs `var`, (3) 3 exercícios de fixação progressivos sobre constantes e iota."

### 🖼️ Roteiro de Slides
> "Crie um roteiro de 5 slides sobre 'Variáveis e Constantes em Go' para um desenvolvedor JavaScript. Mostre lado a lado: `let/const` do JS vs `var/:=` e `const` do Go. Destaque o que é igual e o que é diferente."

### 💬 Perguntas Profundas para o Chat
- "Qual é a diferença prática entre `var x int` e `x := 0`? Em qual situação eu escolheria um ou outro?"
- "O que é 'shadowing' de variável em Go? Me dê um exemplo concreto onde isso causa um bug difícil de encontrar."
- "Por que constantes untyped em Go são mais flexíveis que constantes typed? Me explique com um exemplo que mostra a diferença."
- "Em Python, posso fazer `x = 'texto'` e depois `x = 42` na mesma variável. Por que o Go não permite isso e qual o benefício?"
