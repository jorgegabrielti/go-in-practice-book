# 🧮 Teste de Mesa — Capítulo 15: Goroutines

> **Atenção:** goroutines são não-determinísticas. A ordem exata de execução varia a cada rodada. O teste de mesa aqui documenta o **comportamento esperado** e os **limites do que podemos prever**.

---

## Código rastreado

`exemplos/exe01/main.go` — 10 goroutines concorrentes com `time.Sleep`.

---

## Rastreamento linha a linha

```
numCpu := runtime.NumCPU()
→ numCpu = N (depende da máquina — ex: 8)

fmt.Printf("Seu computador tem %d núcleos lógicos.\n", numCpu)
```
**Saída:** `Seu computador tem 8 núcleos lógicos.`

```
fmt.Println("\n--- Começando a Linha de Montagem ---")
```
**Saída:**
```

--- Começando a Linha de Montagem ---
```

```
start := time.Now()
→ start = T0 (momento atual)
```

### O loop de goroutines

```
for i := 1; i <= 10; i++ {
    go tarefaPesada(i)
}
```

| Iteração | Ação | Resultado |
|---|---|---|
| i=1 | `go tarefaPesada(1)` | goroutine lançada — retorna **imediatamente** |
| i=2 | `go tarefaPesada(2)` | goroutine lançada — retorna **imediatamente** |
| ... | ... | ... |
| i=10 | `go tarefaPesada(10)` | goroutine lançada — retorna **imediatamente** |

O loop inteiro completa em **microssegundos**. As goroutines começam a rodar em paralelo.

```
fmt.Println("O Chefe (Main) está esperando...")
```
**Saída:** `O Chefe (Main) está esperando...`

> **Importante:** esta linha pode aparecer **antes** dos primeiros "Robô X: Iniciando..." — o scheduler pode não ter dado tempo às goroutines de imprimir ainda.

```
time.Sleep(3 * time.Second)
→ main fica bloqueada por 3 segundos
→ nesse tempo, as 10 goroutines executam tarefaPesada (cada uma leva 2s)
```

### Dentro de cada `tarefaPesada(id)`

```
fmt.Printf("Robô %d: Iniciando trabalho...\n", id)  → imprime
time.Sleep(2 * time.Second)                           → goroutine dorme 2s
fmt.Printf("Robô %d: Terminou!\n", id)               → imprime
```

### Saída esperada das goroutines (ordem **não-determinística**)

Fase 1 — todas as goroutines iniciam quase ao mesmo tempo:
```
Robô 3: Iniciando trabalho...
Robô 7: Iniciando trabalho...
Robô 1: Iniciando trabalho...
Robô 5: Iniciando trabalho...
... (qualquer ordem, todas em ~T0+ε)
```

Após ~2 segundos, todas terminam:
```
Robô 6: Terminou!
Robô 2: Terminou!
Robô 9: Terminou!
... (qualquer ordem, todas em ~T0+2s)
```

```
duracao := time.Since(start)
→ duracao ≈ 3s (o sleep da main, não o das goroutines)

fmt.Printf("\nTempo total: %s (Impressionante, não?)\n", duracao)
```
**Saída:** `\nTempo total: 3.0xxxs (Impressionante, não?)`

---

## Saída total esperada (um exemplo possível)

```
Seu computador tem 8 núcleos lógicos.

--- Começando a Linha de Montagem ---
O Chefe (Main) está esperando...
Robô 4: Iniciando trabalho...
Robô 1: Iniciando trabalho...
Robô 7: Iniciando trabalho...
Robô 2: Iniciando trabalho...
Robô 9: Iniciando trabalho...
Robô 5: Iniciando trabalho...
Robô 3: Iniciando trabalho...
Robô 10: Iniciando trabalho...
Robô 6: Iniciando trabalho...
Robô 8: Iniciando trabalho...
[~2 segundos de silêncio]
Robô 4: Terminou!
Robô 7: Terminou!
Robô 1: Terminou!
Robô 9: Terminou!
Robô 2: Terminou!
Robô 5: Terminou!
Robô 3: Terminou!
Robô 6: Terminou!
Robô 8: Terminou!
Robô 10: Terminou!
[~1 segundo restante do sleep da main]

Tempo total: 3.001234s (Impressionante, não?)
```

---

## O que NÃO podemos prever

| Incógnita | Por quê |
|---|---|
| Ordem dos "Iniciando..." | O scheduler decide qual goroutine roda primeiro |
| Ordem dos "Terminou!" | Depende de quando cada goroutine foi agendada |
| Se "O Chefe..." aparece antes ou depois dos "Iniciando..." | Race condition entre main e goroutines |
| Tempo exato | Depende de carga do sistema |

---

## Exercícios — Trace Conceitual

### exe02 — Loop Traiçoeiro (versão com bug)

```go
for i := 0; i < 5; i++ {
    go func() { fmt.Println(i) }()
}
```

```
Goroutine lançada com closure que captura 'i' por REFERÊNCIA.
Loop completa: i=5 (condição 5 < 5 = false → sai)
Goroutines começam a rodar: todas lêem i=5
Saída provável: "5\n5\n5\n5\n5\n" (mas pode variar)
```

### exe02 — Loop Traiçoeiro (versão corrigida)

```go
for i := 0; i < 5; i++ {
    go func(v int) { fmt.Println(v) }(i)
}
```

```
Goroutine lançada com cópia de i no momento do lançamento.
i=0 → goroutine recebe v=0
i=1 → goroutine recebe v=1
i=2 → goroutine recebe v=2
i=3 → goroutine recebe v=3
i=4 → goroutine recebe v=4
Saída: 0,1,2,3,4 em ordem não-determinística
```

### exe03 — Ping-Pong sem canais

```
go ping()  → goroutine imprime "Ping", dorme 500ms, imprime "Ping", ...
go pong()  → goroutine imprime "Pong", dorme 500ms, imprime "Pong", ...
main dorme 5s

Saída esperada (mas NÃO garantida):
Ping
Pong
Ping
Pong
...

Saída possível (sem sincronia):
Ping
Ping
Pong
Ping
Pong
Pong
...

→ Sem canais, a ordem depende do scheduler. Cap 16 resolve isso.
```
