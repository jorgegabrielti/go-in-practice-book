# 🧮 Teste de Mesa — Capítulo 12: Interfaces

Rastreamento manual da execução do `main.go` e dos exercícios, linha a linha, **antes de rodar `go run`**.

---

## main.go — Bloco 1: Implementação da interface `Geometrico`

### Tipos e métodos definidos (nenhuma execução ainda)

```go
type Geometrico interface { Area() float64; Perimetro() float64 }
type Retangulo struct { Largura, Altura float64 }
type Circulo struct { Raio float64 }
```

| Tipo | Satisfaz `Geometrico`? | Motivo |
|---|---|---|
| `Retangulo` | ✅ Sim | tem `Area() float64` e `Perimetro() float64` |
| `Circulo` | ✅ Sim | tem `Area() float64` e `Perimetro() float64` |

> Nenhuma linha de código precisa declarar `implements` — o compilador verifica o Method Set automaticamente.

---

## main.go — Bloco 2: `func main()` — ExibirDetalhes(r)

```go
r := Retangulo{Largura: 10, Altura: 5}
```

| Variável | Tipo | Valor |
|---|---|---|
| `r` | `Retangulo` | `{Largura:10, Altura:5}` |

```go
ExibirDetalhes(r)
```

Dentro de `ExibirDetalhes(g Geometrico)`, `g` recebe uma cópia de `r` como `Geometrico`:

| Passo | Expressão | Cálculo | Resultado |
|---|---|---|---|
| 1 | `fmt.Println("--- Forma Geométrica ---")` | — | imprime `--- Forma Geométrica ---` |
| 2 | `g.Area()` | `r.Largura * r.Altura` = `10 * 5` | `50.00` |
| 3 | `fmt.Printf("Área: %.2f\n", 50.0)` | — | imprime `Área: 50.00` |
| 4 | `g.Perimetro()` | `2*10 + 2*5` = `20 + 10` | `30.00` |
| 5 | `fmt.Printf("Perímetro: %.2f\n", 30.0)` | — | imprime `Perímetro: 30.00` |
| 6 | `c, ok := g.(Circulo)` | `g` guarda um `Retangulo`, não `Circulo` | `ok = false` |
| 7 | `if ok` | `false` | bloco NÃO executa |

**Saída parcial:**
```
--- Forma Geométrica ---
Área: 50.00
Perímetro: 30.00
```

---

## main.go — Bloco 3: `ExibirDetalhes(c)`

```go
c := Circulo{Raio: 3}
ExibirDetalhes(c)
```

Dentro de `ExibirDetalhes`, `g` recebe `c` (Circulo):

| Passo | Expressão | Cálculo | Resultado |
|---|---|---|---|
| 1 | `fmt.Println("--- Forma Geométrica ---")` | — | imprime separador |
| 2 | `g.Area()` | `math.Pi * 3 * 3` = `3.14159... * 9` | `≈ 28.27` |
| 3 | `fmt.Printf("Área: %.2f\n", 28.27...)` | — | imprime `Área: 28.27` |
| 4 | `g.Perimetro()` | `2 * math.Pi * 3` = `6.28318...` | `≈ 18.85` |
| 5 | `fmt.Printf("Perímetro: %.2f\n", 18.85...)` | — | imprime `Perímetro: 18.85` |
| 6 | `c2, ok := g.(Circulo)` | `g` guarda um `Circulo` | `ok = true`, `c2 = {Raio:3}` |
| 7 | `if ok` | `true` | bloco EXECUTA |
| 8 | `fmt.Printf("Raio: %.2f ...\n", c2.Raio)` | `c2.Raio = 3` | imprime `Raio: 3.00 (Específico de Círculo)` |

**Saída parcial:**
```
--- Forma Geométrica ---
Área: 28.27
Perímetro: 18.85
Raio: 3.00 (Específico de Círculo)
```

---

## main.go — Bloco 4: Interface Vazia (`[]any`)

```go
var listaGenerica []any
listaGenerica = append(listaGenerica, 10)
listaGenerica = append(listaGenerica, "Olá")
listaGenerica = append(listaGenerica, true)
fmt.Println("\nLista Genérica:", listaGenerica)
```

| Passo | Operação | Estado de `listaGenerica` |
|---|---|---|
| 1 | `var listaGenerica []any` | `nil` (slice vazio) |
| 2 | `append(..., 10)` | `[10]` — int guardado como `any` |
| 3 | `append(..., "Olá")` | `[10, "Olá"]` — string guardada como `any` |
| 4 | `append(..., true)` | `[10, "Olá", true]` — bool guardado como `any` |
| 5 | `fmt.Println(...)` | imprime `\nLista Genérica: [10 Olá true]` |

---

## Saída Total Esperada do main.go

```
--- Forma Geométrica ---
Área: 50.00
Perímetro: 30.00
--- Forma Geométrica ---
Área: 28.27
Perímetro: 18.85
Raio: 3.00 (Específico de Círculo)

Lista Genérica: [10 Olá true]
```

---

## Exercício 1 — A Impressora

### Tipos e métodos

```go
type Imprimivel interface { Imprimir() string }

type Livro struct { Titulo, Autor string }
func (l Livro) Imprimir() string { return "Livro: " + l.Titulo + " — " + l.Autor }

type Carro struct { Modelo, Marca string }
func (c Carro) Imprimir() string { return "Carro: " + c.Modelo + " da " + c.Marca }
```

### Trace do main

```go
itens := []Imprimivel{
    Livro{Titulo: "Go na Prática", Autor: "Angie"},
    Carro{Modelo: "Civic", Marca: "Honda"},
}
for _, item := range itens {
    fmt.Println(item.Imprimir())
}
```

| Iteração | `item` (tipo concreto) | `item.Imprimir()` | Saída |
|---|---|---|---|
| 1 | `Livro{...}` | `"Livro: Go na Prática — Angie"` | `Livro: Go na Prática — Angie` |
| 2 | `Carro{...}` | `"Carro: Civic da Honda"` | `Carro: Civic da Honda` |

---

## Exercício 2 — O Processador de Pagamentos

```go
type Pagamento interface { Pagar(valor float64) }

type Boleto struct{}
func (b Boleto) Pagar(v float64) { fmt.Printf("Pagando R$ %.2f via Boleto\n", v) }

type Cartao struct{}
func (c Cartao) Pagar(v float64) { fmt.Printf("Pagando R$ %.2f via Cartão\n", v) }

func FinalizarCompra(p Pagamento, valor float64) { p.Pagar(valor) }
```

| Chamada | Tipo concreto | Saída |
|---|---|---|
| `FinalizarCompra(Boleto{}, 150.00)` | `Boleto` | `Pagando R$ 150.00 via Boleto` |
| `FinalizarCompra(Cartao{}, 89.90)` | `Cartao` | `Pagando R$ 89.90 via Cartão` |

---

## Exercício 3 — O Type Switch

```go
func Classificar(v any) {
    switch valor := v.(type) {
    case string:
        fmt.Println("É um texto de tamanho", len(valor))
    case int:
        fmt.Println("É um número: dobro =", valor*2)
    case bool:
        fmt.Println("É booleano: valor invertido =", !valor)
    default:
        fmt.Println("Tipo desconhecido")
    }
}
```

| Chamada | `v.(type)` | `case` executado | Saída |
|---|---|---|---|
| `Classificar("Go")` | `string` | string | `É um texto de tamanho 2` |
| `Classificar(21)` | `int` | int | `É um número: dobro = 42` |
| `Classificar(true)` | `bool` | bool | `É booleano: valor invertido = false` |
| `Classificar(3.14)` | `float64` | default | `Tipo desconhecido` |

> **Nota sobre Type Switch**: a variável `valor` dentro de cada `case` já tem o tipo concreto, sem necessidade de cast adicional. Em `case string:`, `valor` é `string`; em `case int:`, `valor` é `int`.
