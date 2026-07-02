# 🧮 Teste de Mesa — Capítulo 09: Ponteiros

Rastreamento manual da execução de cada `main.go` do capítulo, linha a linha, sem rodar o código.

---

## exemplos/main.go — Demonstração completa de ponteiros

```go
numero := 42
var ponteiro *int = &numero
```

**Estado inicial:**

| variável | tipo | valor |
|---|---|---|
| `numero` | `int` | `42` |
| `ponteiro` | `*int` | endereço de `numero` (ex: `0xc000014090`) |

---

### Bloco 1 — Entendendo Endereços

| linha | expressão | saída no terminal |
|---|---|---|
| `fmt.Println("Valor de numero:", numero)` | `numero = 42` | `Valor de numero: 42` |
| `fmt.Println("Endereço de numero:", ponteiro)` | `ponteiro` = endereço | `Endereço de numero: 0xc000014090` (varia a cada execução) |
| `fmt.Println("O que tem no endereço:", *ponteiro)` | `*ponteiro` → viaja até `numero` → `42` | `O que tem no endereço: 42` |

---

### Bloco 2 — Mudando valor via ponteiro

| linha | operação | estado após |
|---|---|---|
| `*ponteiro = 100` | viaja até o endereço de `numero`, escreve `100` lá | `numero = 100` |
| `fmt.Println("Novo valor de numero:", numero)` | `numero` já é `100` | `Novo valor de numero: 100` |

> 💡 `numero` foi alterado sem tocar nele diretamente — só via `*ponteiro`. Isso é o efeito colateral de um ponteiro.

---

### Bloco 3 — Teste Valor vs Referência

**Estado inicial:** `a = 10`, `b = 10`

| chamada | o que acontece internamente | `a` após | `b` após |
|---|---|---|---|
| `zerarValor(a)` | Go copia `a` → parâmetro local `x = 10`; `x = 0` altera só a cópia | `10` (intacto) | `10` |
| `zerarPonteiro(&b)` | Go passa o endereço de `b`; `*x = 0` viaja até `b` e zera o original | `10` | `0` |

| linha | saída |
|---|---|
| `fmt.Println("A (Valor):", a)` | `A (Valor): 10` |
| `fmt.Println("B (Ponteiro):", b)` | `B (Ponteiro): 0` |

---

### Bloco 4 — Ponteiro de Ponteiro (Inception)

**Estado:** `numero = 100`, `ponteiro` aponta para `numero`

| variável | tipo | valor |
|---|---|---|
| `numero` | `int` | `100` |
| `ponteiro` | `*int` | endereço de `numero` |
| `pp` | `**int` | endereço de `ponteiro` |

| linha | expressão | saída |
|---|---|---|
| `fmt.Println("Endereço do ponteiro:", pp)` | `pp` = endereço de `ponteiro` | `Endereço do ponteiro: 0xc000...` (varia) |
| `fmt.Println("Valor lá no fundo:", **pp)` | `**pp` → viaja até `ponteiro` → viaja até `numero` → `100` | `Valor lá no fundo: 100` |

**Saída completa esperada do programa:**
```
Valor de numero: 42
Endereço de numero: 0xc000014090
O que tem no endereço: 42
Novo valor de numero: 100
A (Valor): 10
B (Ponteiro): 0
Endereço do ponteiro: 0xc000012028
Valor lá no fundo: 100
```
> ⚠️ Os endereços hexadecimais variam a cada execução — apenas os valores `42`, `100`, `10`, `0` são previsíveis.

---

## exercicios/ex01 — O Trocador

```go
func trocar(a, b *int) {
    *a, *b = *b, *a
}
// main: x := 5, y := 10 → trocar(&x, &y)
```

| passo | operação | `x` | `y` |
|---|---|---|---|
| inicial | — | `5` | `10` |
| `trocar(&x, &y)` chamado | `a = &x`, `b = &y` | — | — |
| `*a, *b = *b, *a` | avalia direita: `*b=10`, `*a=5` → atribui: `*a=10`, `*b=5` | `10` | `5` |

**Saída esperada:** `10 5`

---

## exercicios/ex02 — O Incrementador

```go
func incrementar(c *int) {
    *c++
}
// main: contador := 0, loop 10x chamando incrementar(&contador)
```

| iteração | `*c++` | `contador` após |
|---|---|---|
| 1 | `0 + 1` | `1` |
| 2 | `1 + 1` | `2` |
| ... | ... | ... |
| 10 | `9 + 1` | `10` |

**Saída esperada:** `10`

---

## exercicios/ex03 — O Detetive de Endereços

```go
texto := "Go"
fmt.Println("Endereço de texto:", &texto)   // ex: 0xc000014098
verEndereco(texto)                           // passa por VALOR

func verEndereco(t string) {
    fmt.Println("Endereço dentro da função:", &t)  // endereço DIFERENTE
}
```

| variável | endereço |
|---|---|
| `texto` (na `main`) | `0xc000014098` (exemplo) |
| `t` (cópia na função) | `0xc0000140b0` (diferente!) |

**Conclusão:** os endereços são **diferentes** — a função recebeu uma cópia de `texto`, alocada em outro lugar da memória. Isso confirma a analogia da Foto: a função trabalha com a foto (cópia), não com a casa original.

**Saída esperada (valores exatos variam):**
```
Endereço de texto: 0xc000014098
Endereço dentro da função: 0xc0000140b0
```
