# 🧮 Teste de Mesa — Capítulo 16: Channels

> **Atenção:** canais unbuffered envolvem goroutines — a ordem de saída pode variar. O trace abaixo documenta o comportamento esperado e os pontos de sincronização.

---

## Código rastreado

`exemplos/exe01/main.go` — Ping Pong + Gerador e Consumidor.

---

## Exemplo 1: Ping Pong

### Setup

```go
mesa := make(chan string)  // canal unbuffered de strings
go jogarPingPong("Ping", mesa)
go jogarPingPong("Pong", mesa)
mesa <- "bola"  // main envia a primeira bola
```

| Passo | Quem | Ação | Estado do canal | Bloqueia? |
|---|---|---|---|---|
| 1 | main | `mesa <- "bola"` | tentando enviar | Sim — espera alguém receber |
| 2 | goroutine "Ping" | `bola := <-mesa` | dado transferido | Desbloqueados ambos |
| 3 | goroutine "Ping" | `Printf("Ping: Recebi bola")` | livre | Não |
| 4 | goroutine "Ping" | `time.Sleep(500ms)` | livre | Não (apenas espera) |
| 5 | goroutine "Ping" | `Printf("Ping: Devolvendo...")` | livre | Não |
| 6 | goroutine "Ping" | `mesa <- bola` | tentando enviar | Sim — espera "Pong" receber |
| 7 | goroutine "Pong" | `bola := <-mesa` | dado transferido | Desbloqueados ambos |
| 8 | goroutine "Pong" | `Printf("Pong: Recebi bola")` | livre | Não |
| ... | ... | ciclo se repete | ... | ... |

### Saída esperada (primeiras trocas)

```
=== Exemplo 1: Ping Pong ===
Ping: Recebi bola
Ping: Devolvendo...
Pong: Recebi bola
Pong: Devolvendo...
Ping: Recebi bola
Ping: Devolvendo...
...
Jogo encerrado!
```

> **Por que é determinístico aqui?** O canal unbuffered força uma sincronização estrita: Ping recebe → Ping envia → Pong recebe → Pong envia. Não há como ambos terem a bola ao mesmo tempo.

---

## Exemplo 2: Gerador e Consumidor

### Código rastreado

```go
esteira := make(chan int)
go gerarNumeros(esteira)  // goroutine em background
calcularQuadrado(esteira)  // roda na main — usa for range
```

### `gerarNumeros` (goroutine)

```
i=1: esteira <- 1  → BLOQUEIA até calcularQuadrado ler
     calcularQuadrado lê 1 → ambos desbloqueiam
     Printf("Gerando: 1")

i=2: esteira <- 2  → BLOQUEIA
     calcularQuadrado lê 2 → desbloqueiam
     ...

i=5: esteira <- 5  → BLOQUEIA
     calcularQuadrado lê 5 → desbloqueiam

close(esteira)  → sinaliza fim
```

### `calcularQuadrado` (main)

```
for numero := range esteira:
  numero=1 → Printf("Quadrado de 1 = 1")
  numero=2 → Printf("Quadrado de 2 = 4")
  numero=3 → Printf("Quadrado de 3 = 9")
  numero=4 → Printf("Quadrado de 4 = 16")
  numero=5 → Printf("Quadrado de 5 = 25")
  canal fechado → loop termina
```

### Saída esperada (determinística — canal força ordem)

```
=== Exemplo 2: Gerador e Consumidor ===
Gerando: 1
Quadrado de 1 = 1
Gerando: 2
Quadrado de 2 = 4
Gerando: 3
Quadrado de 3 = 9
Gerando: 4
Quadrado de 4 = 16
Gerando: 5
Quadrado de 5 = 25
Processamento concluído.
```

> **Nota:** a ordem "Gerando X" / "Quadrado de X" pode se intercalar dependendo do scheduler, mas **cada número é processado antes do próximo ser enviado** — o canal garante isso.

> **Importante:** nenhum `time.Sleep` na main. O `for range` bloqueia até o canal fechar, segurando a main naturalmente.

---

## Exercícios — Trace Conceitual

### exe01 — O Correio Elegante

```
goroutine enviarMensagens:
  c <- "Olá"   → BLOQUEIA até main receber
main: fmt.Println(<-c) → imprime "Olá", desbloqueia goroutine

goroutine:
  c <- "Mundo" → BLOQUEIA
main: fmt.Println(<-c) → imprime "Mundo"

goroutine:
  c <- "Do"    → BLOQUEIA
main: fmt.Println(<-c) → imprime "Do"

goroutine:
  c <- "Go"    → BLOQUEIA
main: fmt.Println(<-c) → imprime "Go"

goroutine: retorna (função termina)
main: termina normalmente

// Se tentar <-c uma 5ª vez:
// main bloqueia esperando envio que nunca acontece
// Go Runtime detecta: todos dormindo
// fatal error: all goroutines are asleep - deadlock!
```

### exe02 — A Soma Distribuída

```
c := make(chan int)

go somar([1,2,3], c)  → goroutine A calcula 6, c <- 6 (BLOQUEIA)
go somar([4,5,6], c)  → goroutine B calcula 15, c <- 15 (BLOQUEIA)

x := <-c  → recebe 6 ou 15 (ordem não garantida), desbloqueia goroutine A ou B
y := <-c  → recebe o outro valor

Printf("Total: %d\n", x+y)
→ "Total: 21"
```

### exe03 — O Temporizador Manual

```
c := make(chan bool)
go temporizador(c)    → goroutine dorme 2s, depois c <- true (BLOQUEIA)
fmt.Println("Esperando...")
<-c                   → main BLOQUEIA aqui por ~2 segundos
                         quando goroutine enviar true → main desbloqueia
fmt.Println("Pronto!")
```

Saída:
```
Esperando...
[~2 segundos de silêncio]
Pronto!
```
