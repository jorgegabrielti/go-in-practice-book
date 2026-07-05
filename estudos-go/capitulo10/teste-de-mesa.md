# 🧮 Teste de Mesa — Capítulo 10: Structs

Rastreamento manual da execução do `main.go` do capítulo, linha a linha, sem rodar o código.

---

## exemplos/main.go

### Definições (fora do main — sem execução, só compilação)

```go
type Produto struct { Nome string; Preco float64; EmEstoque int; Ativo bool }
type Pedido struct { ID int; Cliente string; Itens []Produto }
```

Sem execução. Apenas registra os tipos no compilador.

---

### Bloco 1 — Criando um Produto

```go
p1 := Produto{
    Nome: "Teclado Mecânico", Preco: 250.00, EmEstoque: 10, Ativo: true,
}
```

| campo | tipo | valor |
|---|---|---|
| `p1.Nome` | `string` | `"Teclado Mecânico"` |
| `p1.Preco` | `float64` | `250.00` |
| `p1.EmEstoque` | `int` | `10` |
| `p1.Ativo` | `bool` | `true` |

---

### Bloco 2 — Imprimindo com `%+v`

```go
fmt.Printf("Produto Detalhado: %+v\n", p1)
```

| expressão | saída |
|---|---|
| `%+v` de `p1` | `{Nome:Teclado Mecânico Preco:250 EmEstoque:10 Ativo:true}` |

**Saída:** `Produto Detalhado: {Nome:Teclado Mecânico Preco:250 EmEstoque:10 Ativo:true}`

---

### Bloco 3 — Ponteiro para Struct

```go
p2 := &p1        // p2 aponta para p1
p2.Preco = 200.00 // Go desreferencia automaticamente: (*p2).Preco = 200.00
fmt.Println("Preço alterado no original:", p1.Preco)
```

| passo | operação | `p1.Preco` |
|---|---|---|
| `p2 := &p1` | `p2` passa a apontar para o endereço de `p1` | `250.00` |
| `p2.Preco = 200.00` | viaja até `p1` e altera o campo `Preco` | `200.00` |

> 💡 `p2.Preco` é açúcar sintático para `(*p2).Preco` — o compilador Go faz isso automaticamente.

**Saída:** `Preço alterado no original: 200`

---

### Bloco 4 — Pedido com Slice de Produtos

```go
pedido := Pedido{
    ID: 1001, Cliente: "Roberto",
    Itens: []Produto{
        p1,                             // p1 copiado (Preco=200.00)
        {Nome: "Mouse", Preco: 50.0},   // struct literal inline
    },
}
```

| campo | valor |
|---|---|
| `pedido.ID` | `1001` |
| `pedido.Cliente` | `"Roberto"` |
| `pedido.Itens[0]` | `{Nome:"Teclado Mecânico", Preco:200, EmEstoque:10, Ativo:true}` |
| `pedido.Itens[1]` | `{Nome:"Mouse", Preco:50, EmEstoque:0, Ativo:false}` |

> ⚠️ `p1` foi passado **por valor** — uma cópia entra no slice. Alterações futuras em `p1` não afetam `pedido.Itens[0]`.

---

### Bloco 5 — Calculando o total (loop)

```go
total := 0.0
for _, item := range pedido.Itens {
    total += item.Preco
}
fmt.Printf("Total do Pedido %d: R$ %.2f\n", pedido.ID, total)
```

| it. | `item.Nome` | `item.Preco` | `total` após |
|---|---|---|---|
| 1 | `"Teclado Mecânico"` | `200.00` | `200.00` |
| 2 | `"Mouse"` | `50.00` | `250.00` |

**Saída:** `Total do Pedido 1001: R$ 250.00`

---

### Saída completa esperada

```
Produto Detalhado: {Nome:Teclado Mecânico Preco:250 EmEstoque:10 Ativo:true}
Preço alterado no original: 200
Total do Pedido 1001: R$ 250.00
```

---

## exercicios/ex01 — O Cadastro de Gamer

```go
type Jogador struct { Nickname string; Nivel int; Vivo bool }
j := Jogador{Nickname: "NoobMaster", Nivel: 1, Vivo: true}
fmt.Println(j)
```

| campo | valor |
|---|---|
| `j.Nickname` | `"NoobMaster"` |
| `j.Nivel` | `1` |
| `j.Vivo` | `true` |

**Saída esperada:** `{NoobMaster 1 true}`

---

## exercicios/ex02 — A Comparação de Retângulos

```go
type Retangulo struct { Largura int; Altura int }
func calcularArea(r Retangulo) int { return r.Largura * r.Altura }

r1 := Retangulo{Largura: 5, Altura: 3}   // área = 15
r2 := Retangulo{Largura: 4, Altura: 4}   // área = 16
```

| variável | `Largura` | `Altura` | área |
|---|---|---|---|
| `r1` | `5` | `3` | `15` |
| `r2` | `4` | `4` | `16` |

`calcularArea(r1) = 15`, `calcularArea(r2) = 16` → `r2` é maior.

**Saída esperada:** `R2 tem a maior área: 16`

---

## exercicios/ex03 — O Sistema de Playlist

```go
type Musica struct { Titulo string; Artista string; DuracaoEmSegundos int }
type Playlist struct { Nome string; Musicas []Musica }

playlist := Playlist{
    Nome: "Rock",
    Musicas: []Musica{
        {Titulo: "Bohemian Rhapsody", Artista: "Queen",          DuracaoEmSegundos: 354},
        {Titulo: "Back in Black",     Artista: "AC/DC",          DuracaoEmSegundos: 255},
        {Titulo: "Stairway to Heaven",Artista: "Led Zeppelin",   DuracaoEmSegundos: 482},
    },
}
```

**Loop + cálculo do total:**

| it. | `Titulo` | `Artista` | `DuracaoEmSegundos` | `totalSegundos` após |
|---|---|---|---|---|
| 1 | `"Bohemian Rhapsody"` | `"Queen"` | `354` | `354` |
| 2 | `"Back in Black"` | `"AC/DC"` | `255` | `609` |
| 3 | `"Stairway to Heaven"` | `"Led Zeppelin"` | `482` | `1091` |

`totalSegundos / 60 = 1091 / 60 = 18.18...` → com `%.2f`: `18.18 min`

**Saída esperada:**
```
Bohemian Rhapsody - Queen
Back in Black - AC/DC
Stairway to Heaven - Led Zeppelin
Duração total da playlist Rock: 18.18 min
```
