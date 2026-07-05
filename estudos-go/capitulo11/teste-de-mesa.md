# 🧮 Teste de Mesa — Capítulo 11: Métodos

Rastreamento manual da execução do `main.go` do capítulo, linha a linha, sem rodar o código.

---

## exemplos/main.go — ContaBancaria

### Definições (fora do main)

| tipo | campos |
|---|---|
| `ContaBancaria` | `Titular string`, `Saldo float64` |

Métodos registrados no compilador:
- `Extrato()` — Value Receiver `(c ContaBancaria)` — leitura, não modifica
- `Depositar(valor float64)` — Pointer Receiver `(c *ContaBancaria)` — modifica `Saldo`
- `Sacar(valor float64) bool` — Pointer Receiver `(c *ContaBancaria)` — modifica `Saldo`, retorna bool

---

### main() — passo a passo

**Estado inicial:**

```go
minhaConta := ContaBancaria{Titular: "César", Saldo: 100.00}
```

| campo | valor |
|---|---|
| `minhaConta.Titular` | `"César"` |
| `minhaConta.Saldo` | `100.00` |

---

**Passo 1 — `minhaConta.Extrato()` (Value Receiver)**

Go cria cópia de `minhaConta` para executar. `c.Titular = "César"`, `c.Saldo = 100.00`.

| expressão | saída |
|---|---|
| `fmt.Printf("Conta de %s \| Saldo: R$ %.2f\n", c.Titular, c.Saldo)` | `Conta de César \| Saldo: R$ 100.00` |

`minhaConta.Saldo` permanece `100.00` — a cópia foi descartada.

---

**Passo 2 — `minhaConta.Depositar(50.00)` (Pointer Receiver)**

Go passa `&minhaConta`. `valor = 50.00 > 0` → entra no `if`.

| operação | `minhaConta.Saldo` após |
|---|---|
| `c.Saldo += 50.00` | `150.00` |

Saída: `Depósito realizado com sucesso.`

**Passo 3 — `minhaConta.Extrato()`**

`c.Saldo = 150.00` → `Conta de César | Saldo: R$ 150.00`

---

**Passo 4 — `sucesso := minhaConta.Sacar(200.00)` (Pointer Receiver)**

`c.Saldo = 150.00`, `valor = 200.00`. Condição: `150.00 >= 200.00` → **false** → vai para o `else`.

| linha | saída | retorno |
|---|---|---|
| `fmt.Println("Saldo insuficiente!")` | `Saldo insuficiente!` | — |
| `return false` | — | `sucesso = false` |

`minhaConta.Saldo` permanece `150.00` — nenhuma alteração feita.

**Passo 5 — `if !sucesso`**

`!false = true` → entra no bloco.

Saída: `Preciso trabalhar mais...`

**Passo 6 — `minhaConta.Extrato()`**

`c.Saldo = 150.00` → `Conta de César | Saldo: R$ 150.00`

---

### Saída completa esperada

```
Conta de César | Saldo: R$ 100.00
Depósito realizado com sucesso.
Conta de César | Saldo: R$ 150.00
Saldo insuficiente!
Preciso trabalhar mais...
Conta de César | Saldo: R$ 150.00
```

---

## exercicios/ex01 — O Carro Acelerado

```go
type Carro struct { VelocidadeAtual int }
// Acelerar(): VelocidadeAtual += 10
// Frear():    VelocidadeAtual -= 10  (mínimo 0)
// main: acelerar 3x, frear 1x
```

| chamada | operação | `VelocidadeAtual` |
|---|---|---|
| inicial | — | `0` |
| `Acelerar()` | `0 + 10` | `10` |
| `Acelerar()` | `10 + 10` | `20` |
| `Acelerar()` | `20 + 10` | `30` |
| `Frear()` | `30 - 10` | `20` |

**Saída esperada:** `Velocidade final: 20`

---

## exercicios/ex02 — O Relógio Digital

```go
type Relogio struct { Horas int; Minutos int }
// AdicionarMinutos(m int): soma minutos, ajusta horas e wrap de 24h
```

**Caso do enunciado: 23:50 + 20min**

| passo | cálculo | resultado |
|---|---|---|
| `totalMin = 50 + 20` | `70` | — |
| `Minutos = 70 % 60` | `10` | `Minutos = 10` |
| `horasExtras = 70 / 60` | `1` | — |
| `Horas = (23 + 1) % 24` | `24 % 24 = 0` | `Horas = 0` |

**Saída esperada:** `00:10`

---

## exercicios/ex03 — A Calculadora Orientada a Métodos

```go
type Calculadora struct { Resultado float64 }
// Somar(v):   Resultado += v
// Subtrair(v): Resultado -= v
// Limpar():   Resultado = 0
```

| chamada | operação | `Resultado` |
|---|---|---|
| inicial | — | `0` |
| `calc.Somar(10)` | `0 + 10` | `10` |
| `calc.Somar(5)` | `10 + 5` | `15` |
| `calc.Subtrair(2)` | `15 - 2` | `13` |

**Saída esperada:** `13`
