# Capítulo 11 — Métodos: Quando Seus Dados Aprendem a Agir

## Tema Central

Um Método é uma função associada a um tipo. Em vez de `latir(cachorro)`, você escreve `cachorro.Latir()`. Métodos transformam Structs de contêineres passivos em entidades com comportamento — é o passo do Go em direção à orientação a objetos, feito com a simplicidade característica da linguagem.

---

## O que o Livro Cobre

- Analogia: Cachorro de Pelúcia vs Cachorro Real
- Sintaxe do Receiver: `func (c Cachorro) Latir()`
- Value Receiver `(c Tipo)`: cópia — leitura, sem modificação
- Pointer Receiver `(c *Tipo)`: endereço — pode modificar o original
- `MudarNomeFalho` (value) vs `MudarNomeReal` (pointer) — demonstração
- Métodos em tipos não-struct: `type Dinheiro float64` com `String()`
- Dica: sem Getters/Setters desnecessários
- `main.go`: ContaBancaria com Extrato (value), Depositar (pointer), Sacar (pointer + bool)
- 3 exercícios: Carro Acelerado, Relógio Digital, Calculadora Orientada a Métodos

---

## Aprofundando os Conceitos

### Receiver: o parâmetro especial

```go
// Função comum — solta no mundo
func latirFuncao(c Cachorro) {
    fmt.Println(c.Nome, "diz: Au!")
}

// Método — pertence ao tipo Cachorro
func (c Cachorro) Latir() {
    fmt.Println(c.Nome, "diz: Au!")
}
```

O receiver `(c Cachorro)` antes do nome da função é o que transforma uma função em método. `c` é a instância que invocou o método — equivalente ao `this` de Java/Python, mas **sempre explícito**.

```go
rex := Cachorro{Nome: "Rex"}
latirFuncao(rex)  // função: "você carrega o cachorro até a fábrica"
rex.Latir()       // método: "o cachorro age por conta própria"
```

### Value Receiver — cópia, somente leitura

```go
type Temperatura struct {
    Celsius float64
}

// Value Receiver — Go faz uma cópia de Temperatura
func (t Temperatura) Fahrenheit() float64 {
    return t.Celsius*9/5 + 32  // t é cópia — original intacto
}

func (t Temperatura) MudarFalho(nova float64) {
    t.Celsius = nova  // muda a CÓPIA, não o original!
}

temp := Temperatura{Celsius: 100}
fmt.Println(temp.Fahrenheit())  // 212 — só lê, funciona perfeitamente
temp.MudarFalho(25)
fmt.Println(temp.Celsius)       // 100 — não mudou! bug silencioso
```

**Quando usar Value Receiver:**
- O método só precisa ler dados (não modificar)
- A struct é pequena (Ponto, Temperatura, Cor)
- O tipo é imutável por design (similar a string)

### Pointer Receiver — modifica o original

```go
func (t *Temperatura) MudarReal(nova float64) {
    t.Celsius = nova  // modifica o original via ponteiro
}

temp := Temperatura{Celsius: 100}
temp.MudarReal(25)
fmt.Println(temp.Celsius)  // 25 — modificado!
```

**Quando usar Pointer Receiver:**
- O método precisa modificar o estado da struct
- A struct é grande (evitar cópia cara)
- Consistência: se algum método usa pointer receiver, todos devem usar

**Regra de ouro**: na dúvida, use Pointer Receiver. É mais barato (copia apenas o ponteiro) e mais flexível (pode modificar ou não).

### Auto-desreferência — Go é gentil

Ao contrário de C, Go automaticamente insere `*` e `&` quando necessário:

```go
temp := Temperatura{Celsius: 100}  // temp é Temperatura, não *Temperatura

// Para chamar Pointer Receiver em valor, Go insere & automaticamente
temp.MudarReal(25)  // Go compila como (&temp).MudarReal(25)

// Para chamar Value Receiver em ponteiro, Go insere * automaticamente
p := &temp
fmt.Println(p.Fahrenheit())  // Go compila como (*p).Fahrenheit()
```

**Limitação**: a auto-inserção de `&` só funciona se a variável é **endereçável** (variável nomeada). Não funciona em retornos de função temporários:
```go
NovaConta("Jorge", 1000).Depositar(50)  // Se NovaConta retorna Conta (não *Conta), ERRO
```

### Métodos em Tipos Não-Struct

Qualquer tipo definido com `type` no mesmo pacote aceita métodos:

```go
type Celsius float64
type Fahrenheit float64

func (c Celsius) ParaFahrenheit() Fahrenheit {
    return Fahrenheit(c*9/5 + 32)
}

func (f Fahrenheit) ParaCelsius() Celsius {
    return Celsius((f - 32) * 5 / 9)
}

var fervura Celsius = 100
fmt.Println(fervura.ParaFahrenheit())  // 212

// Slice com métodos (type alias de slice)
type Numeros []int

func (n Numeros) Soma() int {
    total := 0
    for _, v := range n {
        total += v
    }
    return total
}

nums := Numeros{1, 2, 3, 4, 5}
fmt.Println(nums.Soma())  // 15
```

### A Interface `Stringer` — o método mais importante para debug

```go
// Definida em fmt:
type Stringer interface {
    String() string
}

// Quando um tipo implementa String(), fmt.Println o usa automaticamente
type Dinheiro float64

func (d Dinheiro) String() string {
    return fmt.Sprintf("R$ %.2f", d)
}

salario := Dinheiro(5000.5)
fmt.Println(salario)       // R$ 5000.50 — usa String() automaticamente
fmt.Printf("%v\n", salario) // R$ 5000.50
```

Implementar `String()` é o padrão Go para customizar como um tipo é exibido.

### Method Sets — quais métodos uma variável pode chamar

| Tipo da variável | Pode chamar |
|---|---|
| `T` (valor) | Value Receivers |
| `*T` (ponteiro) | Value Receivers + Pointer Receivers |

```go
type Conta struct{ Saldo float64 }
func (c Conta) Extrato() {}      // Value Receiver
func (c *Conta) Depositar() {}   // Pointer Receiver

var c Conta
c.Extrato()     // OK
c.Depositar()   // OK (Go insere & automaticamente)

var p *Conta
p.Extrato()     // OK (Go insere * automaticamente)
p.Depositar()   // OK
```

Isso importa para **Interfaces**: uma variável do tipo `T` (não ponteiro) não satisfaz uma interface que exige Pointer Receiver.

### Getters/Setters: o que Go pensa

```go
// Java-way (não faça em Go)
func (p *Pessoa) GetNome() string { return p.Nome }
func (p *Pessoa) SetNome(nome string) { p.Nome = nome }

// Go-way: acesso direto se não há lógica
p.Nome = "Ana"  // simples e idiomático

// Go-way com validação: Setter tem sentido quando há lógica real
func (p *Pessoa) SetIdade(idade int) error {
    if idade < 0 || idade > 150 {
        return errors.New("idade inválida")
    }
    p.Idade = idade
    return nil
}
```

Effective Go diz: se você tem um Getter para um campo `Owner`, nomeie o método `Owner()` (não `GetOwner()`). Setter seria `SetOwner()`.

---

## Referências Oficiais

- **Especificação — Method declarations**: https://go.dev/ref/spec#Method_declarations
- **Especificação — Method sets**: https://go.dev/ref/spec#Method_sets
- **Effective Go — Methods**: https://go.dev/doc/effective_go#methods
- **Effective Go — Getters**: https://go.dev/doc/effective_go#Getters
- **Tour de Go — Methods**: https://go.dev/tour/methods/1
- **Blog: Go's Declaration Syntax (contexto de receivers)**: https://go.dev/blog/declaration-syntax
- **Pacote `fmt` — interface Stringer**: https://pkg.go.dev/fmt#Stringer

---

## Exemplos de Código Adicionais

### Method Chaining (Builder Pattern)

```go
type QueryBuilder struct {
    tabela    string
    condicoes []string
    limite    int
}

func (q *QueryBuilder) Da(tabela string) *QueryBuilder {
    q.tabela = tabela
    return q  // retorna o ponteiro para encadeamento
}

func (q *QueryBuilder) Onde(cond string) *QueryBuilder {
    q.condicoes = append(q.condicoes, cond)
    return q
}

func (q *QueryBuilder) Limite(n int) *QueryBuilder {
    q.limite = n
    return q
}

func (q *QueryBuilder) SQL() string {
    where := strings.Join(q.condicoes, " AND ")
    return fmt.Sprintf("SELECT * FROM %s WHERE %s LIMIT %d", q.tabela, where, q.limite)
}

// Uso encadeado
sql := (&QueryBuilder{}).
    Da("usuarios").
    Onde("ativo = true").
    Onde("idade > 18").
    Limite(10).
    SQL()
```

### Implementando `sort.Interface`

```go
type Produto struct {
    Nome  string
    Preco float64
}

type ByPreco []Produto

func (b ByPreco) Len() int           { return len(b) }
func (b ByPreco) Less(i, j int) bool { return b[i].Preco < b[j].Preco }
func (b ByPreco) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }

produtos := []Produto{{"B", 30}, {"A", 10}, {"C", 20}}
sort.Sort(ByPreco(produtos))
// [{A 10} {C 20} {B 30}]
```

### Máquina de estados com métodos

```go
type Estado string

const (
    Pendente   Estado = "pendente"
    Processando Estado = "processando"
    Concluido  Estado = "concluído"
    Cancelado  Estado = "cancelado"
)

type Pedido struct {
    ID     int
    Estado Estado
}

func (p *Pedido) Iniciar() error {
    if p.Estado != Pendente {
        return fmt.Errorf("pedido não está pendente: %s", p.Estado)
    }
    p.Estado = Processando
    return nil
}

func (p *Pedido) Concluir() error {
    if p.Estado != Processando {
        return fmt.Errorf("pedido não está em processamento: %s", p.Estado)
    }
    p.Estado = Concluido
    return nil
}
```

---

## Perguntas & Respostas Frequentes

**P: Posso ter métodos com o mesmo nome em tipos diferentes?**
R: Sim! Métodos são escopados ao tipo. `func (c Cachorro) Falar()` e `func (g Gato) Falar()` são métodos completamente independentes. Isso é a base das Interfaces.

**P: O receiver `c` em `(c Cachorro)` é como um parâmetro normal?**
R: Sim — é exatamente um parâmetro, com as mesmas regras de cópia. A diferença é que ele aparece antes do nome do método e é chamado com notação de ponto.

**P: Posso definir método em um tipo de outro pacote?**
R: Não. Métodos só podem ser definidos no mesmo pacote do tipo. Para "estender" um tipo externo, use wrapping:
```go
type MeuSlice []int  // novo tipo baseado em []int do runtime
func (s MeuSlice) Soma() int { ... }  // método válido — MeuSlice é deste pacote
```

**P: Qual a diferença entre método e função que recebe o tipo como parâmetro?**
R: Funcionalmente equivalentes. A diferença é semântica e ergonômica:
- `c.Latir()` — a ação pertence ao tipo (orientado a objetos)
- `Latir(c)` — função utilitária (procedural)
Go não obriga a orientação a objetos — escolha o que faz mais sentido para o domínio.

**P: O que é `method expression`?**
R: Você pode tratar um método como uma função passando o tipo como primeiro argumento:
```go
latir := Cachorro.Latir     // func(Cachorro)
latir(rex)                  // equivale a rex.Latir()

// Com pointer receiver
depositar := (*Conta).Depositar  // func(*Conta, float64)
depositar(&conta, 100)
```

---

## Comparações com Outras Linguagens

### Métodos e `this`

```python
class Cachorro:
    def __init__(self, nome):
        self.nome = nome
    
    def latir(self):  # self é o receiver implícito
        print(f"{self.nome}: Au!")

rex = Cachorro("Rex")
rex.latir()
```

```java
class Cachorro {
    String nome;
    void latir() {  // this é implícito
        System.out.println(this.nome + ": Au!");
    }
}
```

```go
type Cachorro struct{ Nome string }
func (c Cachorro) Latir() {  // receiver EXPLÍCITO — sem this implícito
    fmt.Println(c.Nome, ": Au!")
}
```

### Getters/Setters

```java
// Java: sempre getters e setters (encapsulamento clássico)
public String getNome() { return nome; }
public void setNome(String nome) { this.nome = nome; }
```

```python
# Python: property decorator para simular getters/setters
@property
def nome(self): return self._nome
@nome.setter
def nome(self, valor): self._nome = valor
```

```go
// Go: acesso direto para campos simples
p.Nome = "Ana"  // sem setter — Go-way
// Setter só quando há lógica de validação
func (p *Pessoa) SetNome(nome string) error { ... }
```

---

## Armadilhas Comuns

1. **Value Receiver que tenta modificar**: `func (c Conta) Depositar(v float64) { c.Saldo += v }` — compila mas não faz nada de útil. Use Pointer Receiver.

2. **Misturar receivers**: ter alguns métodos com Value e outros com Pointer Receiver no mesmo tipo causa confusão sobre qual pode satisfazer uma interface. Seja consistente.

3. **Nil pointer receiver com métodos**: um Pointer Receiver pode ser chamado com receptor nil (se o método verificar):
   ```go
   func (c *Conta) Extrato() {
       if c == nil {
           fmt.Println("conta inválida")
           return
       }
       fmt.Printf("Saldo: %.2f\n", c.Saldo)
   }
   var c *Conta
   c.Extrato()  // funciona! imprime "conta inválida"
   ```

4. **Método retornando cópia de si mesmo**: value receiver que retorna o receiver retorna uma cópia modificada — o original não muda:
   ```go
   func (c Conta) ComSaldo(v float64) Conta {
       c.Saldo = v
       return c  // retorna cópia modificada
   }
   nova := conta.ComSaldo(500)  // conta original intacta, nova tem 500
   ```

---

## Quiz de Fixação

1. Qual é a diferença entre Value Receiver e Pointer Receiver?
2. Por que `func (c Conta) Depositar(v float64) { c.Saldo += v }` não funciona como esperado?
3. Como Go lida com a chamada `temp.MudarReal(25)` se `MudarReal` tem Pointer Receiver mas `temp` é um valor?
4. O que é o Method Set de um tipo `T` vs `*T`?
5. Por que métodos não podem ser definidos em tipos de outros pacotes?
6. Quando faz sentido criar Getters e Setters em Go?

---

## 🔬 Dissecando a Sintaxe

### Value Receiver vs Pointer Receiver — anatomia comparada

```go
// Value Receiver
func (c ContaBancaria) Extrato() {
    fmt.Printf("Saldo: R$ %.2f\n", c.Saldo)
}

// Pointer Receiver
func (c *ContaBancaria) Depositar(valor float64) {
    if valor > 0 {
        c.Saldo += valor
    }
}
```

```
Value Receiver:
func  (c ContaBancaria)  Extrato  ()  {
────   ─  ─────────────   ──────   ─  ─
  │    │        │            │     │  └─ abre o bloco
  │    │        │            │     └─ parâmetros (vazio: não recebe nada além do receiver)
  │    │        │            └─ nome do método (maiúscula = exportado)
  │    │        └─ tipo do receiver: "ContaBancaria" (sem asterisco)
  │    │            Go cria uma CÓPIA de ContaBancaria para executar este método
  │    │            modificar c.Saldo aqui não afeta a conta original
  │    └─ nome local do receiver: "c" — como "this" em Java, mas explícito
  └─ palavra-chave

Pointer Receiver:
func  (c *ContaBancaria)  Depositar  (valor float64)  {
            ─
            └─ "*" antes do tipo = PONTEIRO receiver
                Go passa o ENDEREÇO da ContaBancaria, não uma cópia
                modificar c.Saldo aqui AFETA a conta original
                c.Saldo é atalho para (*c).Saldo — o compilador insere (*) automaticamente
```

### Chamada de método — a auto-desreferência

```go
conta := ContaBancaria{Titular: "Jorge", Saldo: 100}  // [1]

conta.Extrato()    // [2] — Value Receiver: OK
conta.Depositar(50) // [3] — Pointer Receiver em valor: Go resolve
```

```
[1]  conta é do tipo ContaBancaria (valor, não ponteiro)

[2]  conta.Extrato()
         └─ Extrato tem Value Receiver (c ContaBancaria)
             Go faz: ContaBancaria.Extrato(conta)  ← passa cópia de conta
             conta original: intacta

[3]  conta.Depositar(50)
         └─ Depositar tem Pointer Receiver (c *ContaBancaria)
             conta é um VALOR, não ponteiro — mas Go resolve automaticamente:
             Go faz: (&conta).Depositar(50)  ← obtém endereço e passa
             conta original: MODIFICADA (Saldo passa de 100 para 150)

             ATENÇÃO: isso só funciona porque "conta" é endereçável (variável nomeada)
             Se fosse retorno de função temporário: NovaConta().Depositar(50) → ERRO
```

### Método em tipo não-struct — `type Dinheiro float64`

```go
type Dinheiro float64                    // [1]

func (d Dinheiro) String() string {      // [2]
    return fmt.Sprintf("R$ %.2f", d)    // [3]
}

var salario Dinheiro = 5000.5
fmt.Println(salario)                    // [4]
```

```
[1]  type Dinheiro float64
      ────  ────────  ──────
        │       │         └─ tipo BASE: Dinheiro é um float64 por baixo
        │       └─ nome do novo tipo (diferente de float64 — tipo distinto)
        └─ palavra-chave: cria novo tipo

[2]  func (d Dinheiro) String() string
              ────────   ──────   ──────
                 │          │         └─ tipo de retorno: uma string
                 │          └─ nome: "String" é um nome ESPECIAL
                 │              implementa a interface fmt.Stringer automaticamente
                 └─ receiver é Dinheiro (não float64 — tipo base não aceita métodos)

[3]  fmt.Sprintf("R$ %.2f", d)
                  ─────────  ─
                       │     └─ d = o valor Dinheiro (usado como float64 aqui)
                       └─ "%.2f" = formatar como float com 2 casas decimais

[4]  fmt.Println(salario)
         └─ fmt detecta que Dinheiro implementa fmt.Stringer (tem método String())
             chama salario.String() automaticamente
             imprime: "R$ 5000.50" em vez de "5000.5"
```

---

## 🎙️ Prompts para o NotebookLM

### 🎧 Audio Overview
> "Gere um Audio Overview explicando Métodos em Go comparando com Java e Python. Foque em: o que é o receiver (o 'this' explícito do Go), por que Value Receiver é como tirar foto e Pointer Receiver é como ter a chave da casa, a regra prática de quando usar cada um, e por que Go não precisa de Getters e Setters para tudo."

### 📋 Briefing Doc
> "Crie um briefing sobre 'A decisão mais importante ao definir métodos em Go: Value Receiver ou Pointer Receiver'. Inclua: a regra de ouro (na dúvida, use Pointer Receiver), os 3 sinais de que Value Receiver é apropriado, e o bug silencioso de usar Value Receiver quando Pointer era necessário."

### 📚 Study Guide
> "Crie um guia de estudos do Capítulo 11 com: (1) tabela comparativa Value Receiver vs Pointer Receiver (quando usar, o que pode e não pode fazer, custo de memória), (2) exercício: dada uma lista de métodos, classifique quais devem ser Value e quais devem ser Pointer Receiver, (3) quiz de 8 perguntas sobre Métodos."

### 🖼️ Roteiro de Slides
> "Crie um roteiro de 6 slides sobre 'Métodos em Go'. Slide 1: O que é um Receiver — o this explícito. Slide 2: Value Receiver — quando usar. Slide 3: Pointer Receiver — quando usar. Slide 4: Auto-desreferência — Go resolve conta.Depositar(). Slide 5: Métodos em tipos não-struct. Slide 6: Getters e Setters — quando fazer e quando não fazer."

### 💬 Perguntas Profundas para o Chat
- "Por que usar Value Receiver em um método que modifica um campo da struct é um bug silencioso? Mostre com um exemplo concreto de ContaBancaria."
- "O que é o Method Set de um tipo T versus *T? Por que isso importa para Interfaces?"
- "Implementar a interface `fmt.Stringer` com o método `String() string` — como o Go detecta isso automaticamente? O que é duck typing?"
- "Por que Effective Go diz que um método Getter chamado `Owner()` é preferível a `GetOwner()`? Qual a filosofia por trás disso?"
