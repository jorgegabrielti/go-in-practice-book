# 📚 Fonte NotebookLM — Capítulo 16: Channels

> Material de referência aprofundado para o NotebookLM. Use junto com `fonte.txt` e `README.md` do capítulo.

---

## 1. Conceitos Fundamentais

### O Mantra do Go

> *"Não compartilhe memória para se comunicar. Comunique-se para compartilhar memória."*

Em linguagens tradicionais, múltiplas threads compartilham variáveis e usam locks para evitar corrupção. Em Go, o idioma preferido é transferir a posse do dado entre goroutines pelo canal — quem tem o dado é quem trabalha com ele.

| Abordagem | Mecanismo | Risco |
|---|---|---|
| Memória compartilhada | Mutex / Lock | Race condition, deadlock por lock esquecido |
| Comunicação por canal | `chan` + `<-` | Deadlock (detectável pelo runtime) |

---

### Anatomia de um Channel

```
Goroutine A                     Goroutine B
    │                               │
    │  c <- valor (ENVIAR)          │
    │──────────────────────────────►│  msg := <-c (RECEBER)
    │                               │
    │    ◄── sincronização: ambas desbloqueiam ao mesmo tempo ──►
    │
```

O canal não armazena — ele **sincroniza**. Os dois lados precisam estar prontos ao mesmo tempo (para canais unbuffered).

---

### Estados possíveis de um canal

| Estado do canal | Operação de envio | Operação de recebimento |
|---|---|---|
| Aberto, sem receptor | **Bloqueia** | — |
| Aberto, sem produtor | — | **Bloqueia** |
| Aberto, ambos prontos | Transfere e desbloqueia | Transfere e desbloqueia |
| Fechado | **Panic** | Retorna zero value imediatamente |
| `nil` | **Bloqueia para sempre** | **Bloqueia para sempre** |

> Canal `nil` é um canal declarado mas não inicializado: `var c chan int`. Operações nele bloqueiam para sempre — outro tipo de deadlock silencioso.

---

## 2. Dissecando a Sintaxe

### 2.1 Criação de canal

```go
c := make(chan int)
```

| Parte | Significado |
|---|---|
| `make` | Channels são tipos referência — precisam ser inicializados |
| `chan` | Palavra-chave que declara o tipo canal |
| `int` | Tipo dos dados que o canal transporta |
| `c` | Variável que referencia o canal criado |

---

### 2.2 Operador `<-`

```go
canal <- valor    // ENVIAR
valor := <-canal  // RECEBER
```

| Expressão | Leitura em voz alta | Direção |
|---|---|---|
| `c <- 42` | "envie 42 para c" | dado entra no canal |
| `x := <-c` | "receba de c e atribua a x" | dado sai do canal |
| `<-c` (sem atribuição) | "descarte o valor de c" | dado sai e é ignorado |

---

### 2.3 Canais direcionais

```go
func produtor(c chan<- int) { ... }   // Send-Only
func consumidor(c <-chan int) { ... } // Receive-Only
```

| Tipo | Símbolo | Pode enviar | Pode receber | Uso típico |
|---|---|---|---|---|
| `chan int` | Bidirecional | ✅ | ✅ | Criação do canal na main |
| `chan<- int` | Send-Only | ✅ | ❌ | Parâmetro de função produtora |
| `<-chan int` | Receive-Only | ❌ | ✅ | Parâmetro de função consumidora |

> Go converte implicitamente `chan int` para `chan<- int` ou `<-chan int` quando necessário — sem cast explícito.

---

### 2.4 `close` e `for range`

```go
close(c)  // fecha o canal

for v := range c {
    fmt.Println(v)
}
// loop termina automaticamente quando c for fechado
```

| Parte | Significado |
|---|---|
| `close(c)` | Sinaliza que não haverá mais envios. Só o produtor deve chamar. |
| `for v := range c` | Recebe valores do canal sequencialmente; sai quando fechado |
| Receber de canal fechado | Retorna zero value do tipo + `ok = false` |

**Forma com ok:**
```go
v, ok := <-c
if !ok {
    // canal fechado
}
```

---

## 3. Conceitos Avançados

### 3.1 Buffered Channels (mencionado como próximo passo)

```go
c := make(chan int, 3)  // buffer de capacidade 3
```

Canal com buffer **não bloqueia imediatamente** ao enviar — só bloqueia quando o buffer está cheio. Permite desacoplar temporalmente produtor e consumidor.

> Capítulo atual foca em **unbuffered**. Buffered channels são introduzidos no próximo capítulo.

---

### 3.2 O Padrão Pipeline

```
[Estágio 1] --chan1--> [Estágio 2] --chan2--> [Estágio 3]
```

Cada estágio é uma goroutine que lê de um canal, transforma, e escreve em outro. O bloqueio dos canais garante **backpressure automático**: nenhum estágio produz mais rápido do que o próximo consome.

```go
func gerar(nums ...int) <-chan int {
    c := make(chan int)
    go func() {
        for _, n := range nums { c <- n }
        close(c)
    }()
    return c
}

func quadrado(in <-chan int) <-chan int {
    c := make(chan int)
    go func() {
        for v := range in { c <- v * v }
        close(c)
    }()
    return c
}

func main() {
    for v := range quadrado(gerar(1, 2, 3, 4, 5)) {
        fmt.Println(v) // 1 4 9 16 25
    }
}
```

---

### 3.3 Deadlock — Análise Completa

**Deadlock ocorre quando:**
1. Uma goroutine envia num canal sem receptor
2. Uma goroutine recebe de um canal sem produtor
3. Duas goroutines esperam uma pela outra (circular)

**Go Runtime detecta e encerra com:**
```
fatal error: all goroutines are asleep - deadlock!
```

> Isso só funciona quando **todas** as goroutines estão bloqueadas. Se houver goroutine rodando em loop infinito, o runtime não detecta o deadlock das outras.

---

### 3.4 Comparação com outras linguagens

| | Go (channels) | Java (BlockingQueue) | Python (queue.Queue) | JavaScript |
|---|---|---|---|---|
| Sintaxe | `c <- v` / `v := <-c` | `.put(v)` / `.take()` | `.put(v)` / `.get()` | N/A (event loop) |
| Bloqueio | Automático (unbuffered) | Automático | Automático | N/A |
| Deadlock detection | Runtime | Não (hung threads) | Não | N/A |
| Tipagem | Compile-time | Runtime (generics) | Runtime | N/A |
| Canais direcionais | ✅ (`chan<-`, `<-chan`) | ❌ | ❌ | N/A |

---

## 4. Exemplos Extras

### 4.1 Canal como semáforo (padrão avançado)

```go
// Canal buffered como semáforo — limita goroutines simultâneas
sem := make(chan struct{}, 5) // máximo 5 simultâneas

for i := 0; i < 20; i++ {
    sem <- struct{}{} // adquire vaga (bloqueia se cheio)
    go func(id int) {
        defer func() { <-sem }() // libera vaga ao terminar
        trabalhar(id)
    }(i)
}
```

### 4.2 Canal de notificação (`chan struct{}`)

```go
done := make(chan struct{})

go func() {
    trabalhoLongo()
    close(done) // sinaliza conclusão
}()

<-done // espera conclusão
fmt.Println("Terminado!")
```

`chan struct{}` é idiomático para sinais — `struct{}` ocupa 0 bytes de memória.

---

## 5. Prompts para NotebookLM

### 🎙️ Audio Overview
> "Crie um diálogo de podcast entre dois engenheiros Go: um explicando por que o Go prefere channels a mutexes, e o outro desafiando com casos onde mutex ainda é necessário. Inclua deadlock, o padrão pipeline, canais direcionais e o mantra do Rob Pike."

### 🎥 Video Overview
> "Crie um roteiro de vídeo animado mostrando visualmente como um channel unbuffered funciona: duas goroutines se aproximando do canal, o bloqueio de cada lado, a transferência atômica do dado, e o desbloqueio simultâneo. Compare com a versão mutex (disputando um lock)."

### 📊 Slides
> "Crie 9 slides sobre channels: (1) o mantra do Go, (2) o operador `<-` com exemplos visuais, (3) canal unbuffered e o bloqueio, (4) deadlock e detecção pelo runtime, (5) `close` e `for range`, (6) regras de fechamento (3 regras), (7) canais direcionais com tabela, (8) padrão pipeline, (9) canal vs mutex — quando usar cada um."

### ❓ Perguntas e Respostas
> "Com base no material do capítulo 16, crie 10 perguntas de revisão sobre channels em Go: criação com make, operador `<-`, bloqueio, deadlock, close, for range, canais direcionais, zero value de canal fechado, canal nil. Inclua respostas e explicações."

### 🔍 Estudo Aprofundado
> "Explique o funcionamento interno de um channel unbuffered em Go: como o runtime implementa a sincronização? O que acontece na memória quando `c <- v` bloqueia? Como o scheduler trata goroutines bloqueadas em canais versus goroutines em time.Sleep? Qual a diferença de implementação entre unbuffered e buffered?"

### 🔄 Comparativo
> "Compare o modelo de concorrência por channels do Go com: (1) o modelo de Atores (Erlang/Akka), (2) Communicating Sequential Processes (CSP) — a teoria matemática que Go implementa, (3) async/await com filas em JavaScript/Node.js. Qual deles é mais próximo de channels? Quais são as diferenças práticas?"
