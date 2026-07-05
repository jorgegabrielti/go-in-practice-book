# Capítulo 03 — Tipos Básicos: O Sistema de Tipos do Go

## Tema Central

Go tem um sistema de tipos **forte, estático e explícito**. Não existe coerção implícita: `int + float64` não compila. Toda conversão é manual e visível no código. Isso elimina bugs silenciosos que assolam linguagens com tipos fracos.

---

## O que o Livro Cobre

- Tipagem forte e estática (sem conversão implícita)
- Overflow silencioso (e.g., `uint8` > 255 "dá a volta")
- `float64` por padrão; nunca usar float para dinheiro (IEEE 754)
- Casting explícito: `TipoDesejado(valor)`
- `byte` = alias de `uint8`; `rune` = alias de `int32` (Unicode)

---

## Aprofundando os Conceitos

### Todos os Tipos Numéricos do Go

**Inteiros com sinal:**
| Tipo | Tamanho | Faixa |
|------|---------|-------|
| `int8` | 8 bits | -128 a 127 |
| `int16` | 16 bits | -32.768 a 32.767 |
| `int32` | 32 bits | -2.147.483.648 a 2.147.483.647 |
| `int64` | 64 bits | -9,2 × 10¹⁸ a 9,2 × 10¹⁸ |
| `int` | 32 ou 64 bits* | Depende da plataforma |

**Inteiros sem sinal:**
| Tipo | Tamanho | Faixa |
|------|---------|-------|
| `uint8` / `byte` | 8 bits | 0 a 255 |
| `uint16` | 16 bits | 0 a 65.535 |
| `uint32` | 32 bits | 0 a 4.294.967.295 |
| `uint64` | 64 bits | 0 a 1,8 × 10¹⁹ |
| `uint` | 32 ou 64 bits* | Depende da plataforma |
| `uintptr` | — | Para armazenar ponteiros |

*Em sistemas 64-bit (praticamente todos hoje), `int` = `int64`.

**Ponto flutuante:**
| Tipo | Bits | Precisão |
|------|------|---------|
| `float32` | 32 | ~7 dígitos decimais |
| `float64` | 64 | ~15-17 dígitos decimais |

**Complexos:**
- `complex64` (float32 real + float32 imaginária)
- `complex128` (float64 real + float64 imaginária)

**Quando usar qual inteiro?**
- `int`: padrão para contadores, índices, aritmética geral
- `int64`: quando precisar de faixa maior ou interoperabilidade explícita
- `uint8`/`byte`: dados binários, pixels, cores
- `int32`/`rune`: code points Unicode

### Overflow: comportamento e detecção

```go
var x uint8 = 255
x++
fmt.Println(x) // 0 — wraps around silenciosamente!

var y int8 = 127
y++
fmt.Println(y) // -128 — two's complement wrap
```

Go não lança exceção em overflow — ele silenciosamente "dá a volta" (wrap-around). Para detectar overflow, você precisa verificar manualmente ou usar a biblioteca `math/big`.

### O problema com float e dinheiro (IEEE 754)

```go
var preco float64 = 0.1 + 0.2
fmt.Println(preco)          // 0.30000000000000004 — NÃO é 0.3!
fmt.Println(preco == 0.3)   // false
```

Isso não é bug do Go — é como números de ponto flutuante binário funcionam em QUALQUER linguagem que use o padrão IEEE 754 (Python, Java, JavaScript — todos têm o mesmo problema).

**Solução para dinheiro:**
```go
// Opção 1: trabalhar com centavos (inteiros)
preco := 1099 // R$ 10,99 representado como 1099 centavos

// Opção 2: usar a biblioteca github.com/shopspring/decimal
import "github.com/shopspring/decimal"
preco := decimal.NewFromString("0.1")
total := preco.Add(decimal.NewFromString("0.2"))
fmt.Println(total) // 0.3 — exato!
```

### `string`: imutável, UTF-8 nativo

```go
s := "Olá, Gopher! 🐹"
fmt.Println(len(s))         // bytes, NÃO runes! "Olá" tem 5 bytes (á = 2 bytes UTF-8)
fmt.Println([]rune(s))      // converter para runes para trabalhar com caracteres Unicode
```

Strings em Go são **sequências de bytes imutáveis**, codificadas em UTF-8 por convenção. O operador `len()` retorna número de bytes, não de caracteres visíveis. Para contar caracteres Unicode corretamente:

```go
import "unicode/utf8"

s := "Olá"
fmt.Println(len(s))                    // 4 bytes (á = 2 bytes)
fmt.Println(utf8.RuneCountInString(s)) // 3 runes (O, l, á)
```

### `rune` vs `byte`

```go
s := "café"

// Iterar por bytes (perigoso para Unicode)
for i := 0; i < len(s); i++ {
    fmt.Printf("%d: %c\n", i, s[i])  // pode quebrar caracteres multibyte
}

// Iterar por runes (correto para Unicode)
for i, r := range s {
    fmt.Printf("%d: %c (U+%04X)\n", i, r, r)
}
```

### Conversão explícita

```go
var i int = 42
var f float64 = float64(i)       // int → float64
var u uint = uint(f)             // float64 → uint (trunca, não arredonda)

var b byte = 65
var r rune = rune(b)             // byte → rune
var s string = string(r)         // rune → string: "A"

// Conversão de string para número
import "strconv"
n, err := strconv.Atoi("42")     // string → int
s2 := strconv.Itoa(42)           // int → string
```

---

## Referências Oficiais

- **Especificação — Tipos numéricos**: https://go.dev/ref/spec#Numeric_types
- **Especificação — String types**: https://go.dev/ref/spec#String_types
- **Especificação — Conversões**: https://go.dev/ref/spec#Conversions
- **Blog: Strings, bytes, runes e characters (Pike)**: https://go.dev/blog/strings
- **Pacote `strconv`**: https://pkg.go.dev/strconv
- **Pacote `unicode/utf8`**: https://pkg.go.dev/unicode/utf8
- **Pacote `math`** (constantes como `math.MaxInt64`): https://pkg.go.dev/math

---

## Exemplos de Código Adicionais

### Verificando limites com o pacote `math`

```go
import "math"

fmt.Println(math.MaxInt8)    // 127
fmt.Println(math.MinInt8)    // -128
fmt.Println(math.MaxUint8)   // 255
fmt.Println(math.MaxFloat64) // 1.7976931348623157e+308
fmt.Println(math.Pi)         // 3.141592653589793
```

### Trabalhando com runes na prática

```go
// Verificar se um caractere é letra maiúscula
import "unicode"

r := 'A'
fmt.Println(unicode.IsUpper(r))   // true
fmt.Println(unicode.IsLetter(r))  // true
fmt.Println(unicode.ToLower(r))   // 'a' (97)
```

### Type alias vs new type

```go
// Alias: são o MESMO tipo (byte = uint8)
type byte = uint8   // definição real na stdlib

// New type: são TIPOS DIFERENTES (não intercambiáveis sem cast)
type Celsius    float64
type Fahrenheit float64

var c Celsius = 100
var f Fahrenheit = Fahrenheit(c * 9/5 + 32)  // conversão explícita necessária
// var f2 Fahrenheit = c   // ERRO: mismatched types
```

---

## Perguntas & Respostas Frequentes

**P: Qual a diferença entre `int` e `int64`?**
R: Em sistemas 64-bit (a maioria atual), `int` tem 64 bits — são iguais em tamanho. A diferença é que `int` e `int64` são **tipos diferentes** para o compilador: você não pode atribuir um ao outro sem conversão explícita.

**P: Por que Go não tem casting implícito como C?**
R: Casting implícito em C é fonte histórica de bugs. Go exige conversões explícitas para deixar claro que você está mudando o tipo intencionalmente. Quando você escreve `float64(i)`, está documentando a intenção no código.

**P: Como comparar dois floats corretamente?**
R: Nunca use `==` com floats para igualdade exata. Use uma margem de tolerância (epsilon):
```go
const epsilon = 1e-9
func iguaisAproximadamente(a, b float64) bool {
    return math.Abs(a-b) < epsilon
}
```

**P: O que é um `rune` literalmente?**
R: Um `rune` é um `int32` que representa um Unicode code point. O caractere 'A' é o rune `65` (U+0041). O emoji '🐹' é o rune `129081` (U+1F439).

**P: Strings são mutáveis em Go?**
R: Não. Strings são imutáveis. Para construir strings dinamicamente use `strings.Builder` ou `fmt.Sprintf`. Concatenação com `+` em loop é O(n²) — evite.

---

## Comparações com Outras Linguagens

### Tipos numéricos

```python
# Python 3: int tem precisão arbitrária (sem overflow!), float é float64
x = 10**100  # 1 googol — funciona!
```

```java
// Java: byte, short, int, long, float, double — sem unsigned nativos
long maximo = Long.MAX_VALUE;  // 9.223.372.036.854.775.807
```

```javascript
// JavaScript: number (IEEE 754 float64 para tudo)
// BigInt para inteiros grandes (ES2020+)
Number.MAX_SAFE_INTEGER  // 9.007.199.254.740.991
```

```go
// Go: tipos explícitos, sized, com e sem sinal
var grande uint64 = 18_446_744_073_709_551_615  // uint64 máximo
import "math/big"
muitoGrande := new(big.Int).Exp(big.NewInt(10), big.NewInt(100), nil)  // 1 googol
```

### Strings e Unicode

```python
# Python 3: strings são sequências de Unicode code points (str)
s = "café"
len(s)  # 4 — conta CARACTERES, não bytes
```

```java
// Java: strings são UTF-16, char é 16 bits
// Problemas com caracteres BMP (acima de U+FFFF)
"café".length()  // 4
```

```go
// Go: strings são bytes UTF-8; rune para Unicode
s := "café"
len(s)                         // 5 bytes (é = 2 bytes em UTF-8)
utf8.RuneCountInString(s)      // 4 runes
```

---

## Armadilhas Comuns

1. **`len()` em strings retorna bytes, não caracteres**: para acentos e emojis, use `utf8.RuneCountInString()`.

2. **Truncamento na conversão float→int**: `int(3.9)` é `3`, não `4`. Go trunca (não arredonda).

3. **Strings e slices de bytes**: `[]byte(s)` e `string(b)` criam **cópias**. Modificar o slice não afeta a string original (e vice-versa).

4. **Literal de rune usa aspas simples, string usa aspas duplas**:
   ```go
   r := 'A'      // rune (int32 = 65)
   s := "A"      // string (1 byte)
   // r == s      // ERRO: tipos incompatíveis
   ```

---

## Quiz de Fixação

1. Por que `0.1 + 0.2` não é `0.3` em Go (nem em Python, Java, ou JavaScript)?
2. Qual é o resultado de `var x uint8 = 255; x++`? Por quê?
3. Qual a diferença entre `byte` e `rune`? Quando usar cada um?
4. Se `s := "café"`, por que `len(s)` não é `4`?
5. Como converter uma `string` para `int` em Go?
6. Por que nunca usar `float64` para representar valores monetários?

---

## 🔬 Dissecando a Sintaxe

### Conversão explícita de tipos — anatomia

```go
var i int     = 42
var f float64 = float64(i)   // [1]
var u uint    = uint(f)      // [2]
```

```
[1]  float64(i)
      ───────  ─
          │    └─ valor a converter: a variável "i" (que é int = 42)
          └─ tipo destino entre parênteses: "transforma i em float64"
              resultado: 42.0 (sem perda de dados int→float64)
              NOTA: Go nunca converte implicitamente — você É OBRIGADO a escrever float64(i)

[2]  uint(f)
      ────  ─
        │   └─ valor a converter: "f" (que é float64 = 42.0)
        └─ tipo destino: uint (inteiro sem sinal)
            resultado: 42 (trunca a parte decimal — 42.9 viraria 42, NÃO 43)
            ATENÇÃO: truncamento, não arredondamento!
```

### Iteração com `range` sobre string — por que o índice pula

```go
s := "café"

for i, r := range s {             // [1] [2] [3]
    fmt.Printf("%d: %c\n", i, r)  // [4]
}
```

```
[1]  for i, r := range s
      ───  ─  ─  ──────  ─
        │   │  │    │    └─ a string "café" sobre a qual iterar
        │   │  │    └─ palavra-chave: itera sobre uma coleção
        │   │  └─ ":=" declara i e r como novas variáveis a cada iteração
        │   └─ "r" = rune (int32): o CARACTERE Unicode atual (não o byte!)
        └─ "i" = índice do BYTE de início da rune (não índice do caractere)

Por que o índice "pula"?
  'c' = 1 byte → i=0
  'a' = 1 byte → i=1
  'f' = 1 byte → i=2
  'é' = 2 bytes em UTF-8 → i=3 (próxima seria i=5, não i=4!)

[4]  fmt.Printf("%d: %c\n", i, r)
                 ──   ──         ─  ─
                  │    │         │  └─ "r" = a rune (o caractere)
                  │    │         └─ "i" = o índice do byte
                  │    └─ %c = formata como CARACTERE (ex: 'é')
                  └─ %d = formata como número decimal (ex: 3)
```

### `byte` vs `rune` — a diferença visual

```go
var b byte = 'A'    // [1]
var r rune = 'é'    // [2]
var s string = "A"  // [3]
```

```
[1]  byte = uint8 = 0 a 255
      'A' entre aspas SIMPLES = literal de caractere = valor 65 (código ASCII de 'A')
      byte guarda 1 caractere ASCII sem problemas

[2]  rune = int32 = pode guardar qualquer Unicode (até 1.114.112 caracteres)
      'é' em UTF-8 precisa de 2 bytes → não cabe em byte, mas cabe em rune
      'é' = U+00E9 = decimal 233

[3]  "A" entre aspas DUPLAS = string (sequência de bytes), não um caractere
      string é diferente de rune — não são intercambiáveis sem conversão
```

---

## 🎙️ Prompts para o NotebookLM

### 🎧 Audio Overview
> "Gere um Audio Overview explicando o sistema de tipos do Go para alguém que vem do Python. Foque em: por que Go tem tantos tipos de inteiro, o problema real do float com dinheiro (com exemplo de bug financeiro), e a diferença entre byte e rune ao trabalhar com texto em português."

### 📋 Briefing Doc
> "Crie um briefing sobre 'Armadilhas de tipos em Go para desenvolvedores iniciantes', cobrindo: (1) overflow silencioso e como detectar, (2) float64 e dinheiro — o bug que ninguém espera, (3) strings em UTF-8 e por que `len()` mente para acentos."

### 📚 Study Guide
> "Crie um guia de estudos do Capítulo 03 com: (1) tabela comparativa de todos os tipos numéricos do Go com faixa de valores, (2) 5 perguntas sobre conversão de tipos com exemplos de código para analisar, (3) exercício: 'O que este código imprime?' para cada armadilha comum."

### 🖼️ Roteiro de Slides
> "Crie um roteiro de 5 slides sobre 'Tipos em Go vs Python': Slide 1: Tipagem estática vs dinâmica. Slide 2: Por que Go tem int8, int16, int32, int64 e Python tem só int. Slide 3: O problema do float com dinheiro (bug universal). Slide 4: UTF-8, byte, rune e por que 'café' tem 5 bytes mas 4 letras. Slide 5: Conversão explícita."

### 💬 Perguntas Profundas para o Chat
- "Me mostre um bug real de software causado por overflow de inteiro. Como o Go poderia ter prevenido?"
- "Por que `0.1 + 0.2 != 0.3` em Go, Python, Java e JavaScript? O que é IEEE 754 em linguagem simples?"
- "Se eu quero guardar o preço de um produto em Go, qual tipo devo usar e por quê? Quais são minhas opções além de float64?"
- "Qual a diferença entre iterar uma string com `for i := 0; i < len(s); i++` e com `for i, r := range s`? Quando cada um dá resultado diferente?"
