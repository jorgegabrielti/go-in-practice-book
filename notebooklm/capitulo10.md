# Capítulo 10 — Structs: A Planta Baixa dos Seus Dados

## Tema Central

Uma Struct é um tipo composto que agrupa campos de tipos diferentes sob um único nome. É a forma do Go de modelar entidades do mundo real: uma Pessoa tem Nome, Idade e Email; um Produto tem Código, Descrição e Preço. Sem Structs, você estaria gerenciando 300 variáveis soltas.

---

## O que o Livro Cobre

- Analogia: "Pasta Amarela" com fichas organizadas
- `type Pessoa struct { ... }` — definição de tipo
- 3 formas de instanciação: literal com nomes, posicional, zero value
- Acesso a campos com `.`
- `%+v` para debug (imprime nomes dos campos)
- Structs aninhadas (Boneca Russa)
- Structs anônimas (Guardanapo — uso único)
- Ponteiro para Struct e auto-desreferência
- Embedding (Incorporação) e Promoted Fields
- Composição sobre Herança
- `main.go`: Produto + Pedido + slice de Produto, loop calculando total
- 3 exercícios: Cadastro de Gamer, Comparação de Retângulos, Sistema de Playlist

---

## Aprofundando os Conceitos

### Definição de Struct — o novo tipo

```go
// Definição (planta baixa — sem memória alocada)
type Pessoa struct {
    Nome  string
    Idade int
    Email string
}

// O tipo Pessoa é tão válido quanto int ou string
// O compilador verifica tipos em tempo de compilação
```

Structs são definidas com `type` no nível do pacote (normalmente) ou dentro de funções (raras vezes). O nome começa com maiúscula para exportar o tipo.

### As Três Formas de Instanciação

```go
// Forma 1: Literal com nomes de campos (PREFERIDA)
p1 := Pessoa{
    Nome:  "Jorge",
    Idade: 30,
    Email: "jorge@email.com",
}

// Forma 2: Posicional (EVITAR — frágil a mudanças na ordem)
p2 := Pessoa{"Jorge", 30, "jorge@email.com"}

// Forma 3: Zero value (todos os campos recebem zero value)
var p3 Pessoa  // Nome: "", Idade: 0, Email: ""
p3.Nome = "Maria"  // preencher depois
```

**Por que evitar posicional?** Se você adicionar um campo no meio da struct, o código posicional silenciosamente usa valores errados nos campos subsequentes.

### Structs são passadas por valor

```go
type Ponto struct{ X, Y int }

func mover(p Ponto) {
    p.X += 10  // modifica a CÓPIA
}

original := Ponto{1, 2}
mover(original)
fmt.Println(original)  // {1, 2} — intacto
```

Para modificar via função, passe ponteiro:

```go
func moverPonteiro(p *Ponto) {
    p.X += 10  // modifica o original
}
moverPonteiro(&original)
fmt.Println(original)  // {11, 2}
```

### Ponteiro para Struct — auto-desreferência

Go faz um favor ao programador: `p.Campo` funciona tanto para `Ponto` quanto para `*Ponto`:

```go
p := &Pessoa{Nome: "Jorge"}
// Sem auto-deref (como em C): (*p).Nome
// Com auto-deref do Go: p.Nome — o compilador insere (*) automaticamente
fmt.Println(p.Nome)  // "Jorge" — funciona!
p.Idade = 31         // modifica o original
```

### Structs Aninhadas

```go
type Endereco struct {
    Rua    string
    Cidade string
    Estado string
}

type Usuario struct {
    Nome     string
    Idade    int
    Endereco Endereco  // campo do tipo Endereco
}

u := Usuario{
    Nome:  "Jorge",
    Idade: 30,
    Endereco: Endereco{
        Rua:    "Rua das Flores, 123",
        Cidade: "São Paulo",
        Estado: "SP",
    },
}

fmt.Println(u.Endereco.Cidade)  // São Paulo — acesso encadeado
```

### Struct Anônima — descartável

```go
// Para um único uso — sem criar tipo nomeado
config := struct {
    Host string
    Port int
}{
    Host: "localhost",
    Port: 8080,
}

// Útil para tabelas de testes (table-driven tests)
tests := []struct {
    entrada  string
    esperado int
}{
    {"hello", 5},
    {"go", 2},
    {"gopher", 6},
}
```

### Embedding — Composição sobre Herança

Em vez de herança (`Cachorro extends Animal`), Go usa **Embedding**:

```go
type Animal struct {
    Nome  string
    Vida  int
}

func (a Animal) Respirar() {
    fmt.Println(a.Nome, "respira")
}

type Cachorro struct {
    Animal        // Embedding — sem nome de campo!
    Raca string
}

rex := Cachorro{
    Animal: Animal{Nome: "Rex", Vida: 100},
    Raca:   "Labrador",
}

// Promoted Fields — acessar campos/métodos do Animal diretamente
fmt.Println(rex.Nome)   // "Rex" — promovido de Animal.Nome
rex.Respirar()          // chamada ao método de Animal
rex.Raca = "Labrador"   // campo próprio do Cachorro
```

**O que é Promoted Field?** Os campos e métodos do tipo embedado são "promovidos" ao tipo que o contém. É açúcar sintático para `rex.Animal.Nome` — o compilador gera o acesso indireto automaticamente.

**Por que não herança?**
- Herança cria acoplamento: mudar a classe pai quebra todas as subclasses
- Composição é explícita: você sabe exatamente o que está reutilizando
- Go pode embutir múltiplos tipos (sem diamond problem do C++)
- Mais fácil de testar — você pode trocar o tipo embedado

### Igualdade de Structs

Structs são comparáveis com `==` se todos os seus campos forem comparáveis:

```go
type Ponto struct{ X, Y int }

p1 := Ponto{1, 2}
p2 := Ponto{1, 2}
fmt.Println(p1 == p2)  // true — comparação campo a campo

type ComSlice struct {
    Dados []int  // slice não é comparável
}
// c1 == c2  // ERRO DE COMPILAÇÃO
```

### Tags de Struct — metadados para serialização

```go
type Usuario struct {
    Nome  string `json:"nome"`
    Idade int    `json:"idade,omitempty"`  // omitir se zero
    Senha string `json:"-"`               // nunca serializar
}

import "encoding/json"
u := Usuario{Nome: "Jorge", Idade: 30}
dados, _ := json.Marshal(u)
fmt.Println(string(dados))  // {"nome":"Jorge","idade":30}
```

Tags de struct são strings de metadados lidas em runtime por pacotes como `encoding/json`, `gorm`, `validate`, etc. O formato é `` `chave:"valor"` ``.

---

## Referências Oficiais

- **Especificação — Struct types**: https://go.dev/ref/spec#Struct_types
- **Especificação — Promoted methods**: https://go.dev/ref/spec#Method_sets
- **Effective Go — Structs**: https://go.dev/doc/effective_go#composite_literals
- **Blog: JSON and Go**: https://go.dev/blog/json
- **Tour de Go — Structs**: https://go.dev/tour/moretypes/2
- **Pacote `encoding/json`**: https://pkg.go.dev/encoding/json
- **Pacote `reflect`** (para inspecionar tags): https://pkg.go.dev/reflect

---

## Exemplos de Código Adicionais

### Método constructor (factory function)

```go
type Conta struct {
    Titular string
    Saldo   float64
}

// Constructor: garante que a struct seja criada em estado válido
func NovaConta(titular string, depositoInicial float64) (*Conta, error) {
    if titular == "" {
        return nil, errors.New("titular obrigatório")
    }
    if depositoInicial < 0 {
        return nil, errors.New("depósito inicial deve ser não-negativo")
    }
    return &Conta{Titular: titular, Saldo: depositoInicial}, nil
}

c, err := NovaConta("Jorge", 1000)
```

### Struct embedding múltiplo

```go
type Nadador struct {
    VelocidadeNatacao float64
}

type Corredor struct {
    VelocidadeCorrida float64
}

type Triatleta struct {
    Nadador
    Corredor
    Nome string
}

t := Triatleta{
    Nadador:  Nadador{2.5},
    Corredor: Corredor{12.0},
    Nome:     "Atlas",
}
fmt.Println(t.VelocidadeNatacao)  // 2.5 — promovido de Nadador
fmt.Println(t.VelocidadeCorrida)  // 12.0 — promovido de Corredor
```

### Comparação por valor em testes

```go
func TestNovaConta(t *testing.T) {
    c1 := Conta{Titular: "A", Saldo: 100}
    c2 := Conta{Titular: "A", Saldo: 100}
    
    if c1 != c2 {
        t.Errorf("esperado %v, obtido %v", c1, c2)
    }
}
```

---

## Perguntas & Respostas Frequentes

**P: Qual a diferença entre `Struct{}` e `new(Struct)`?**
R: `&Struct{}` e `new(Struct)` produzem o mesmo resultado: um `*Struct` com campos zero. A forma literal `&Struct{}` é preferida por ser mais explícita e permitir inicialização no mesmo lugar.

**P: Posso ter métodos em uma struct definida em outro pacote?**
R: Não. Métodos só podem ser definidos no mesmo pacote do tipo. Para "estender" um tipo de outro pacote, use Embedding ou uma função utilitária.

**P: O que acontece se dois tipos embedados têm um campo com o mesmo nome?**
R: Conflito de nome. O compilador exige que você use o nome do tipo embedado para desambiguar: `t.Nadador.Velocidade` vs `t.Corredor.Velocidade`.

**P: Structs são thread-safe?**
R: Não. Ler e escrever no mesmo campo de struct de múltiplas goroutines sem sincronização é race condition. Use `sync.Mutex` ou canais.

**P: Por que usar `%+v` em vez de `%v` no fmt?**
R: `%v` imprime só os valores: `{Jorge 30}`. `%+v` imprime com nomes dos campos: `{Nome:Jorge Idade:30}`. Muito mais útil para debug.

**P: Posso ter uma struct com zero campos?**
R: Sim. `type Vazia struct{}` tem tamanho zero. É usado como "token" em sets (`map[string]struct{}`) e como tipo para sinais em channels.

---

## Comparações com Outras Linguagens

### Struct vs Class

```python
# Python: class com __init__
class Pessoa:
    def __init__(self, nome, idade):
        self.nome = nome
        self.idade = idade
    
    def apresentar(self):
        return f"{self.nome}, {self.idade} anos"

p = Pessoa("Jorge", 30)
```

```java
// Java: class com construtor
public class Pessoa {
    private String nome;
    private int idade;
    
    public Pessoa(String nome, int idade) {
        this.nome = nome;
        this.idade = idade;
    }
    // getters, setters...
}
Pessoa p = new Pessoa("Jorge", 30);
```

```go
// Go: struct + métodos separados; sem class, sem construtor obrigatório
type Pessoa struct {
    Nome  string
    Idade int
}

func (p Pessoa) Apresentar() string {
    return fmt.Sprintf("%s, %d anos", p.Nome, p.Idade)
}

p := Pessoa{Nome: "Jorge", Idade: 30}
```

### Herança vs Composição

```java
// Java: herança de classe
class Animal { void respirar() {} }
class Cachorro extends Animal { void latir() {} }
// Cachorro HERDA de Animal — acoplamento forte
```

```python
# Python: herança múltipla (MRO complexo)
class Animal:
    def respirar(self): pass
class Cachorro(Animal):
    def latir(self): pass
```

```go
// Go: composição via embedding
type Animal struct{}
func (a Animal) Respirar() {}

type Cachorro struct {
    Animal  // composição — sem herança
    Raca string
}
// Cachorro TEM um Animal, não É um Animal (semântica diferente mas resultado similar)
```

---

## Armadilhas Comuns

1. **Instanciação posicional com ordem errada**: `Pessoa{"30", "Jorge"}` compila (se os tipos baterem) mas os valores estão trocados — bug silencioso.

2. **Cópia não intencional**: passar struct grande por valor em funções hot path é lento. Use ponteiro.

3. **Embedded type com nome ambíguo**: dois tipos embedados com campos de mesmo nome causam compilação com erro. Acesse pelo nome do tipo: `t.Nadador.Campo`.

4. **Tags de struct ignoradas**: tags não têm efeito sem um pacote que as leia via `reflect`. `json:"nome"` só funciona com `encoding/json`.

5. **Struct com campos não exportados e JSON**: campos não exportados (minúscula) são ignorados pelo `encoding/json`. Sempre use maiúscula para campos que precisam ser serializados.

---

## Quiz de Fixação

1. Qual das 3 formas de instanciação é preferida e por quê?
2. O que são Promoted Fields em Embedding?
3. Por que a Composição é preferida sobre a Herança em Go?
4. O que significa `%+v` no contexto de `fmt.Printf`?
5. Por que `type ComSlice struct { Dados []int }` não é comparável com `==`?
6. O que é uma struct anônima e quando você a usaria?

---

## 🔬 Dissecando a Sintaxe

### Definição e instanciação de Struct — anatomia completa

```go
type Pessoa struct {        // [1]
    Nome  string            // [2]
    Idade int               // [3]
    Email string
}

p := Pessoa{               // [4]
    Nome:  "Jorge",         // [5]
    Idade: 30,
    Email: "jorge@email.com",
}
```

```
[1]  type  Pessoa  struct  {
      ────  ──────  ──────
        │      │       └─ palavra-chave: "o que se segue é uma estrutura de campos"
        │      └─ nome do novo tipo (maiúscula = exportado para outros pacotes)
        └─ palavra-chave: "estou criando um novo tipo chamado..."

[2]  Nome  string
      ────  ──────
        │      └─ tipo do campo: qualquer tipo válido do Go
        └─ nome do campo (maiúscula = campo exportado/público)
            minúscula = campo privado ao pacote

[3]  Idade int
         └─ dois campos na mesma struct: Pessoa TEM um Nome E TEM uma Idade

[4]  p := Pessoa{
               └─ instanciação com literal: "cria um valor do tipo Pessoa"
                  ≠ definição (que ficou no type): aqui MEMÓRIA É ALOCADA

[5]  Nome:  "Jorge",
      ────   ───────  ─
        │       │     └─ vírgula obrigatória após cada campo (inclusive o último)
        │       └─ valor do campo: uma string literal
        └─ nome do campo seguido de ":"
            FORMA PREFERIDA: nomear o campo torna o código resistente a mudanças
```

### Embedding — dissecar a composição

```go
type Animal struct {               // [1]
    Nome string
    Vida int
}

func (a Animal) Respirar() {       // [2]
    fmt.Println(a.Nome, "respira")
}

type Cachorro struct {             // [3]
    Animal                         // [4]
    Raca string
}

rex := Cachorro{                   // [5]
    Animal: Animal{Nome: "Rex", Vida: 100},
    Raca:   "Labrador",
}

fmt.Println(rex.Nome)  // [6]
rex.Respirar()         // [7]
```

```
[1]  type Animal struct { ... }
         └─ tipo base — define os dados e comportamentos compartilhados

[2]  func (a Animal) Respirar()
         └─ método de Animal — pertence ao tipo Animal

[3]  type Cachorro struct { ... }
         └─ tipo que VAI USAR Animal via composição

[4]  Animal   ← SEM nome de campo e SEM tipo explícito de campo
      ──────
         └─ EMBEDDING: Cachorro incorpora Animal
             NÃO é "var Animal Animal" (seria um campo nomeado)
             a ausência de nome é o sinal de embedding para o compilador
             efeito: todos os campos e métodos de Animal são PROMOVIDOS para Cachorro

[5]  Cachorro{Animal: Animal{...}}
                ──────
                   └─ ao inicializar, o campo embedded se chama pelo tipo: "Animal"

[6]  rex.Nome
          └─ acessa Animal.Nome DIRETAMENTE — sem escrever rex.Animal.Nome
              o compilador gera rex.Animal.Nome automaticamente (Promoted Field)

[7]  rex.Respirar()
          └─ chama Animal.Respirar() DIRETAMENTE — sem escrever rex.Animal.Respirar()
              o compilador gera a chamada correta automaticamente
```

### Tag de Struct — metadados legíveis por pacotes externos

```go
type Usuario struct {
    Nome  string `json:"nome"`              // [1]
    Idade int    `json:"idade,omitempty"`   // [2]
    Senha string `json:"-"`                 // [3]
}
```

```
[1]  `json:"nome"`
      ────  ────
        │     └─ "nome" = nome do campo no JSON (diferente do nome em Go)
        └─ "json" = qual pacote lê essa tag (encoding/json neste caso)
        backticks ` ` = string raw literal (permite aspas duplas dentro)

[2]  `json:"idade,omitempty"`
                  ─────────
                       └─ opção: omite o campo do JSON se o valor for zero value
                           Idade=0 → campo ausente no JSON; Idade=30 → presente

[3]  `json:"-"`
            ─
            └─ "-" = NUNCA incluir este campo no JSON (mesmo que tenha valor)
                útil para campos sensíveis (senhas, tokens)
```

---

## 🎙️ Prompts para o NotebookLM

### 🎧 Audio Overview
> "Gere um Audio Overview sobre Structs em Go para um desenvolvedor Java. Explique: por que Go não tem classes, como Structs + métodos substituem classes, a diferença entre Embedding e herança (composição vs extensão), e por que 'Cachorro TEM um Animal' em vez de 'Cachorro É um Animal' é melhor design."

### 📋 Briefing Doc
> "Crie um briefing sobre 'Composição sobre Herança em Go: por que Embedding é melhor que extends'. Inclua: o problema real da herança profunda, como Promoted Fields funcionam, quando o nome ambíguo de dois embeddings causa erro, e um exemplo de refatoração de hierarquia de herança para composição."

### 📚 Study Guide
> "Crie um guia de estudos do Capítulo 10 com: (1) comparação das 3 formas de instanciação com análise de prós e contras de cada uma, (2) exercício de modelagem: dado um sistema de e-commerce, identifique quais entidades viram Structs e quais campos cada uma tem, (3) quiz de 8 perguntas sobre Structs e Embedding."

### 🖼️ Roteiro de Slides
> "Crie um roteiro de 6 slides sobre 'Structs em Go'. Slide 1: O problema das 300 variáveis soltas. Slide 2: type + struct — a planta baixa. Slide 3: As 3 formas de instanciação — qual usar. Slide 4: Structs aninhadas vs Embedding. Slide 5: Promoted Fields — como funciona. Slide 6: Tags de Struct — o que são e para que servem."

### 💬 Perguntas Profundas para o Chat
- "Qual é a diferença entre um campo nomeado do tipo Animal (`tipo Animal`) e um embedding de Animal (só `Animal`) em uma struct? O que muda no acesso e na semântica?"
- "Por que Structs são passadas por valor em Go? Em Java, objetos são sempre referências — qual a implicação prática dessa diferença?"
- "O que são tags de struct? Me dê 3 exemplos reais de pacotes que usam tags e o que cada um faz com elas."
- "Se tenho `type Cachorro struct { Animal; Felino }` e tanto Animal quanto Felino têm um campo `Nome`, o que acontece quando acesso `cachorro.Nome`?"
