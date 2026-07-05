# Capítulo 12 — Interfaces: O Plugue Universal

## Tema Central

Interfaces são o mecanismo de abstração mais poderoso do Go. Elas definem **contratos implícitos**: qualquer tipo que implementar os métodos exigidos satisfaz a interface automaticamente, sem declaração explícita. Isso é o que torna Go flexível sem herança de classes.

---

## O que o Livro Cobre

- Analogia da Tomada Elétrica — interface como contrato de pinos
- `type Barulhento interface { FazerBarulho() string }` — definição pura
- Duck Typing implícito — sem `implements`
- Polimorfismo: `IncomodarVizinho(b Barulhento)` aceita Cachorro e Alarme
- Interface vazia `interface{}` / `any` — aceita tudo, use com moderação
- Type Assertion: `v.(Tipo)` arriscado vs `v, ok := v.(Tipo)` seguro
- Type Switch: `switch v := x.(type)`
- Dica: defina interfaces onde são usadas, não onde são criadas
- Interfaces pequenas são as melhores: `io.Reader`, `fmt.Stringer`
- `main.go`: `Geometrico` com `Retangulo` e `Circulo`, `ExibirDetalhes`, `[]any`
- 3 exercícios: A Impressora, O Processador de Pagamentos, O Type Switch

---

## Aprofundando os Conceitos

### O que exatamente o compilador verifica

Quando você escreve `func f(b Barulhento)` e passa um `Cachorro`, o compilador verifica o **Method Set** de `Cachorro`:

- Tem `FazerBarulho() string`? ✅ → `Cachorro` satisfaz `Barulhento`
- A assinatura precisa ser exata: nome, parâmetros e tipo de retorno devem bater

```go
// Barulhento exige: FazerBarulho() string

// ✅ Satisfaz
func (c Cachorro) FazerBarulho() string { return "Au Au" }

// ❌ NÃO satisfaz — tipo de retorno diferente
func (c Gato) FazerBarulho() int { return 42 }

// ❌ NÃO satisfaz — nome diferente
func (c Gato) Barulho() string { return "Miau" }
```

### Como interfaces são representadas internamente

Internamente, um valor de interface é uma **dupla**: `(tipo concreto, ponteiro para valor)`:

```
interface Barulhento
┌──────────────────────────────┐
│  tipo: *Cachorro             │  ← qual tipo está guardado
│  valor: 0xC0001040B0 → {Rex} │  ← ponteiro para o dado
└──────────────────────────────┘
```

Isso explica duas coisas importantes:
1. Uma interface `nil` tem **ambos** tipo e valor nulos — `(nil, nil)`
2. Um ponteiro `nil` guardado em interface **não é nil interface**: `(*Cachorro)(nil)` é `(tipo=*Cachorro, valor=nil)` — diferente de `nil`

```go
var c *Cachorro = nil
var b Barulhento = c   // NÃO é nil interface!
fmt.Println(b == nil)  // false — tipo está preenchido
```

### Diferença entre `any` / `interface{}` e generics

`any` = "qualquer tipo, verificado em runtime via type assertion"
Generics (`T any`) = "qualquer tipo, verificado em tempo de compilação"

```go
// any: perde type safety, usa type assertion em runtime
func somar(a, b any) any {
    return a.(int) + b.(int)  // panic se não for int!
}

// Generics (Go 1.18+): seguro em compilação
func somarGenerico[T int | float64](a, b T) T {
    return a + b
}
```

Use `any` quando o tipo realmente não importa (ex: log, cache genérico). Use generics quando precisa de type safety com múltiplos tipos.

### Nil interface vs interface com valor nil

O bug mais sutil com interfaces em Go:

```go
func retornarErro(falhou bool) error {
    var err *MeuErro = nil
    if falhou {
        err = &MeuErro{"algo deu errado"}
    }
    return err  // SEMPRE retorna interface não-nil! (tipo=*MeuErro, valor=nil)
}

err := retornarErro(false)
if err != nil {  // true — mesmo que err.(*MeuErro) seja nil!
    fmt.Println("erro:", err)  // executa!
}
```

**Solução correta:**
```go
func retornarErro(falhou bool) error {
    if falhou {
        return &MeuErro{"algo deu errado"}
    }
    return nil  // retorna nil interface (tipo=nil, valor=nil)
}
```

### Interfaces da stdlib que você vai usar sempre

| Interface | Pacote | Método(s) | Para que serve |
|---|---|---|---|
| `io.Reader` | `io` | `Read(p []byte) (n int, err error)` | Ler de qualquer fonte (arquivo, rede, buffer) |
| `io.Writer` | `io` | `Write(p []byte) (n int, err error)` | Escrever em qualquer destino |
| `fmt.Stringer` | `fmt` | `String() string` | Customizar impressão com Println |
| `error` | builtin | `Error() string` | Tipos de erro customizados |
| `sort.Interface` | `sort` | `Len/Less/Swap` | Ordenar qualquer slice |
| `http.Handler` | `net/http` | `ServeHTTP(w, r)` | Rotas HTTP |

### Compondo interfaces

Interfaces podem ser compostas de outras interfaces:

```go
type Leitor interface {
    Ler() string
}

type Escritor interface {
    Escrever(s string)
}

// Composição: LeitorEscritor exige ambos
type LeitorEscritor interface {
    Leitor
    Escritor
}
```

`io.ReadWriter` da stdlib é exatamente isso — composição de `io.Reader` e `io.Writer`.

---

## Referências Oficiais

- **Especificação — Interface types**: https://go.dev/ref/spec#Interface_types
- **Especificação — Type assertions**: https://go.dev/ref/spec#Type_assertions
- **Especificação — Type switches**: https://go.dev/ref/spec#Type_switches
- **Effective Go — Interfaces**: https://go.dev/doc/effective_go#interfaces
- **Effective Go — Interface checks**: https://go.dev/doc/effective_go#interface_methods
- **Tour de Go — Interfaces**: https://go.dev/tour/methods/9
- **Blog: The Go Blog — Interfaces and other types**: https://go.dev/doc/effective_go#interfaces_and_types
- **Blog: Laws of Reflection (explica a dupla tipo+valor)**: https://go.dev/blog/laws-of-reflection
- **Pacote `io`** (io.Reader, io.Writer): https://pkg.go.dev/io
- **Pacote `fmt`** (fmt.Stringer): https://pkg.go.dev/fmt#Stringer

---

## Exemplos de Código Adicionais

### Implementando `error` customizado

```go
type ErrSaldoInsuficiente struct {
    Saldo     float64
    Tentativa float64
}

func (e *ErrSaldoInsuficiente) Error() string {
    return fmt.Sprintf("saldo %.2f insuficiente para saque de %.2f", e.Saldo, e.Tentativa)
}

// Usar
func sacar(saldo, valor float64) error {
    if valor > saldo {
        return &ErrSaldoInsuficiente{Saldo: saldo, Tentativa: valor}
    }
    return nil
}

err := sacar(100, 200)
if err != nil {
    var e *ErrSaldoInsuficiente
    if errors.As(err, &e) {  // type assertion segura para errors
        fmt.Printf("Tentou sacar R$ %.2f mas tinha só R$ %.2f\n", e.Tentativa, e.Saldo)
    }
}
```

### Interface check em tempo de compilação

```go
// Garante em compilação que MinhaStruct implementa MinhaInterface
var _ MinhaInterface = (*MinhaStruct)(nil)
// Se MinhaStruct não implementar MinhaInterface, erro de compilação aqui
```

Esse padrão é usado por bibliotecas para documentar intenção e capturar falha de implementação cedo.

### Usando `io.Writer` para escrever em qualquer destino

```go
func escreverRelatorio(w io.Writer, dados []string) {
    for _, linha := range dados {
        fmt.Fprintln(w, linha)
    }
}

// Mesma função, destinos diferentes
escreverRelatorio(os.Stdout, dados)           // terminal
escreverRelatorio(arquivo, dados)             // arquivo
escreverRelatorio(&strings.Builder{}, dados)  // buffer em memória
```

---

## Perguntas & Respostas Frequentes

**P: Por que Go não tem `implements` explícito como Java?**
R: Implementação implícita permite retroativamente satisfazer interfaces de bibliotecas que você não controla. Se `os.File` tem `Read([]byte) (int, error)`, ele automaticamente satisfaz `io.Reader` — sem modificar o pacote `os`. Em Java, você precisaria que `os.File` declarasse `implements io.Reader`, o que é impossível se você não controla o código.

**P: Posso implementar uma interface com Pointer Receiver e passar um valor (não ponteiro)?**
R: Não. Se o método que satisfaz a interface tem Pointer Receiver (`func (c *Cachorro) FazerBarulho()`), só `*Cachorro` satisfaz a interface, não `Cachorro`. Valores não são automaticamente endereçáveis em contexto de interface.

**P: Interface nil é igual a interface com valor nil?**
R: Não — esse é o bug mais sutil de Go. `var b Barulhento = nil` é `nil`. Mas `var c *Cachorro = nil; var b Barulhento = c` NÃO é nil interface — tem tipo `*Cachorro`, só o valor é nil.

**P: Qual a diferença entre Type Assertion e Type Switch?**
R: Type Assertion testa um único tipo: `v.(Cachorro)`. Type Switch testa múltiplos: `switch v.(type)`. Use type assertion quando você tem certeza (ou quer o comma ok para um tipo). Use type switch quando precisa tratar múltiplas possibilidades.

**P: Posso usar interface como chave de map?**
R: Sim, mas com cuidado. O tipo concreto guardado na interface precisa ser comparável. Se em runtime o tipo concreto não for comparável (ex: slice), ocorre panic. Em geral, evite interfaces como chave de map.

---

## Comparações com Outras Linguagens

### Implementação implícita vs explícita

```java
// Java: EXPLÍCITO — deve declarar implements
interface Barulhento { String fazerBarulho(); }
class Cachorro implements Barulhento {
    public String fazerBarulho() { return "Au Au"; }
}
```

```python
# Python: duck typing dinâmico — sem verificação em compilação
class Cachorro:
    def fazer_barulho(self): return "Au Au"

def incomodar(b):
    print(b.fazer_barulho())  # erro só em runtime se não tiver o método
```

```typescript
// TypeScript: structural typing — similar ao Go, mas em runtime JS
interface Barulhento { fazerBarulho(): string; }
// Qualquer objeto com fazerBarulho(): string satisfaz — verificado em compilação
```

```go
// Go: implícito, verificado em COMPILAÇÃO — melhor dos dois mundos
type Barulhento interface { FazerBarulho() string }
// Cachorro satisfaz se e somente se tem exatamente esse método
// Verificado em compilação — sem surpresas em runtime
```

---

## 🔬 Dissecando a Sintaxe

### Definição de interface — anatomia

```go
type  Geometrico  interface  {
────  ──────────  ─────────  ─
  │       │           │      └─ abre o bloco da interface
  │       │           └─ palavra-chave: "o que se segue é uma interface"
  │       └─ nome da interface (maiúscula = exportada)
  └─ palavra-chave: "crio um novo tipo chamado..."

    Area()     float64     // [1]
    Perimetro() float64    // [2]
}
```

```
[1]  Area()  float64
      ────    ───────
        │        └─ tipo de retorno: o método deve retornar float64
        └─ assinatura: nome + parâmetros (vazio) + retorno
            NÃO tem corpo — apenas a assinatura
            qualquer tipo que tiver "Area() float64" satisfaz este requisito

[2]  Perimetro() float64
         └─ segundo requisito: O tipo precisa ter AMBOS para satisfazer Geometrico
             se tiver só Area mas não Perimetro → NÃO satisfaz → erro de compilação
```

### Type Assertion segura — dissecar o comma ok

```go
dog, ok := v.(Cachorro)
───   ──  ─  ─────────
  │    │  │      └─ tipo concreto que queremos extrair da interface
  │    │  └─ ":=" declara dog e ok como novas variáveis
  │    └─ segundo retorno: bool
  │        true  = v guarda um Cachorro
  │        false = v guarda outro tipo (dog será zero value de Cachorro)
  └─ primeiro retorno: o valor concreto do tipo Cachorro
      se ok=false: dog = Cachorro{} (zero value — sem campos preenchidos)

SEM ok: v.(Cachorro) — se v não for Cachorro → PANIC em runtime
COM ok: nunca panic — trate ok=false com guard clause
```

### Type Switch — dissecar a sintaxe especial

```go
switch  valor  :=  v.(type)  {
──────  ─────   ─  ─────────
  │       │     │       └─ "(type)" é sintaxe ESPECIAL do switch — não é uma expressão normal
  │       │     │           nenhuma outra construção Go usa .(type)
  │       │     └─ ":=" declara "valor" como nova variável
  │       └─ nome da variável que recebe o valor TIPADO em cada case
  └─ palavra-chave

case  int:
────  ───
  │     └─ tipo concreto a testar (sem parênteses, sem aspas)
  └─ dentro deste case: "valor" É do tipo int — pode usar valor*2 sem cast

case  string:
         └─ dentro deste case: "valor" É do tipo string — pode usar len(valor)
```

---

## 🎙️ Prompts para o NotebookLM

### 🎧 Audio Overview
> "Gere um Audio Overview explicando Interfaces em Go para um desenvolvedor Java. Um apresentador fica surpreso que Go não tem `implements`. O outro explica Duck Typing, por que implementação implícita é mais flexível, o bug sutil de nil interface vs interface com valor nil, e por que interfaces pequenas são preferidas."

### 📋 Briefing Doc
> "Crie um briefing sobre 'As 3 regras de ouro de Interfaces em Go': (1) defina onde usa, não onde cria, (2) prefira interfaces pequenas (1-2 métodos), (3) evite `any` — use tipos concretos ou generics."

### 📚 Study Guide
> "Crie um guia de estudos do Capítulo 12 com: (1) tabela das interfaces mais importantes da stdlib (io.Reader, fmt.Stringer, error, sort.Interface, http.Handler) com seus métodos e usos, (2) exercício de diagnóstico: dado um código com interface nil bug, identifique e corrija, (3) quiz de 8 perguntas."

### 🖼️ Roteiro de Slides
> "Crie um roteiro de 6 slides sobre 'Interfaces em Go'. Slide 1: A tomada elétrica — o contrato. Slide 2: Duck Typing — sem implements. Slide 3: Polimorfismo em ação. Slide 4: Interface vazia (any) — poder e perigo. Slide 5: Type Assertion e Type Switch. Slide 6: Defina onde usa — a regra de ouro."

### 💬 Perguntas Profundas para o Chat
- "O que é o bug de 'nil interface vs interface com valor nil' em Go? Mostre com código o que acontece e como corrigir."
- "Por que `io.Reader` (1 método) é mais útil que uma interface com 10 métodos? Qual o princípio de design por trás disso?"
- "Qual a diferença entre usar `any` e usar Generics (`[T any]`) em Go? Me dê um exemplo onde cada abordagem é a correta."
- "Como o compilador do Go verifica se um tipo satisfaz uma interface? Mostra um exemplo onde o erro de compilação aponta exatamente o método que falta."
