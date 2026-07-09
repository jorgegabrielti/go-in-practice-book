# 📚 Fonte NotebookLM — Capítulo 15: Goroutines

> Material de referência aprofundado para o NotebookLM. Use junto com `fonte.txt` e `README.md` do capítulo.

---

## 1. Conceitos Fundamentais

### Goroutine vs. Thread de SO

Uma goroutine **não é** uma thread de sistema operacional. É uma "thread verde" — uma abstração gerenciada pelo runtime do Go.

| Característica | Thread de SO | Goroutine |
|---|---|---|
| Memória inicial | ~1 MB | ~2 KB (500x menor) |
| Criação | Cara (syscall ao SO) | Barata (runtime Go) |
| Limite prático | ~10.000 | Milhões |
| Troca de contexto | Microssegundos | Nanossegundos |
| Pilha | Tamanho fixo (pré-alocado) | Dinâmica (cresce/encolhe) |
| Gerenciamento | Sistema Operacional | Go Runtime Scheduler |

**Por que a pilha dinâmica importa:** uma thread de SO precisa reservar memória suficiente para o pior caso antes de saber qual será o pior caso. Uma goroutine começa pequena e cresce sob demanda — sem desperdício.

---

### Go Runtime Scheduler — Modelo M:N

O scheduler opera em três camadas:

```
[Goroutines — milhões]
        ↓  (multiplexação M:N)
[Threads de SO — dezenas]
        ↓
[Núcleos de CPU — unidades]
```

**Como ele evita desperdício:**
- Goroutine bloqueia em I/O → scheduler a remove do operário imediatamente
- Coloca outra goroutine pronta no lugar
- O núcleo nunca fica ocioso aguardando I/O

**Variáveis de controle:**

```go
runtime.NumCPU()          // núcleos lógicos disponíveis
runtime.GOMAXPROCS(N)     // quantos núcleos o Go pode usar (padrão: todos)
runtime.NumGoroutine()    // goroutines ativas neste momento
```

---

## 2. Dissecando a Sintaxe

### 2.1 A palavra `go`

```go
go tarefaPesada(42)
```

| Parte | Significado |
|---|---|
| `go` | Palavra-chave que lança uma nova goroutine |
| `tarefaPesada` | Qualquer função previamente declarada |
| `(42)` | Argumentos — avaliados imediatamente no momento do lançamento |
| Retorno | Ignorado — goroutines não retornam valores para o chamador |

> ⚠️ **Retorno ignorado:** se `tarefaPesada` retornar um erro, esse retorno é perdido. Para comunicar resultados, use canais (Capítulo 16).

---

### 2.2 Goroutine anônima com closure

```go
go func() {
    fmt.Println("trabalho em background")
}()
```

| Parte | Significado |
|---|---|
| `func()` | Define uma função sem nome (literal) |
| `{ ... }` | Corpo da função |
| `()` final | Chama a função imediatamente — sem esses parênteses ela não seria executada |
| `go` antes | Lança essa chamada como goroutine |

---

### 2.3 O gotcha do loop com closure

```go
// ❌ VERSÃO COM BUG
for i := 0; i < 5; i++ {
    go func() {
        fmt.Println(i)  // 'i' capturado por referência
    }()
}
// Saída provável: 5 5 5 5 5
```

**Por que falha:**
1. Loop executa rapidamente — termina em microssegundos
2. Goroutines são agendadas para depois — levam alguns nanossegundos para arrancar
3. Quando as goroutines lêem `i`, o loop já terminou: `i == 5` (valor que fez a condição `i < 5` ser falsa)
4. Todas as goroutines compartilham a mesma variável `i` (referência, não cópia)

```go
// ✅ VERSÃO CORRIGIDA
for i := 0; i < 5; i++ {
    go func(v int) {
        fmt.Println(v)  // 'v' é uma cópia local
    }(i)               // 'i' é avaliado AGORA e copiado para 'v'
}
// Saída: 0 1 2 3 4 (em ordem não-determinística)
```

| Parte | Significado |
|---|---|
| `(v int)` | Parâmetro `v` — cada goroutine recebe sua própria cópia |
| `}(i)` | Passa o valor atual de `i` como argumento — avaliado imediatamente |

> **Nota Go 1.22+:** a partir do Go 1.22, variáveis de loop `for` têm escopo por iteração por padrão, eliminando esse bug. Em versões anteriores (<=1.21), a correção manual com parâmetro ainda é necessária.

---

### 2.4 `runtime.GOMAXPROCS`

```go
runtime.GOMAXPROCS(1)  // força uso de apenas 1 núcleo de CPU
```

| Argumento | Efeito |
|---|---|
| `0` | Retorna o valor atual sem alterar |
| `1` | Simula ambiente single-core — útil para reproduzir bugs de concorrência |
| `N` | Define N núcleos — valor ≥ `runtime.NumCPU()` não traz ganho extra |
| `-1` | Restaura o padrão (todos os núcleos) |

---

## 3. Conceitos Avançados

### 3.1 O problema "Tchau, Obrigado"

`main` é ela mesma uma goroutine — a goroutine raiz. Quando retorna, o runtime encerra todo o programa imediatamente, sem esperar goroutines filhas.

```go
func main() {
    go trabalhoLento()
    // sem nada aqui: 'trabalhoLento' nunca imprime nada
}
```

**Soluções:**

| Abordagem | Quando usar |
|---|---|
| `time.Sleep` | Apenas para fins didáticos — impreciso e frágil |
| `sync.WaitGroup` | Solução idiomática — capítulo 18 |
| `channel` de sinal | Quando goroutines precisam reportar conclusão com dados |
| `context.Context` | Para cancelamento e timeout (capítulo 20+) |

---

### 3.2 Goroutine leak

Uma goroutine leak ocorre quando uma goroutine é lançada mas nunca termina e nunca é coletada pelo GC.

```go
func vazar() {
    go func() {
        for {
            time.Sleep(time.Second) // nunca termina
        }
    }()
}
```

**Consequências:**
- Memória cresce indefinidamente
- CPU consumida desnecessariamente
- Detectável com `runtime.NumGoroutine()` ou pprof

**Causa comum:** goroutine esperando em canal que nunca recebe/envia.

---

### 3.3 Concorrência vs. Paralelismo

Rob Pike: *"Concorrência é sobre estrutura. Paralelismo é sobre execução."*

| | Concorrência | Paralelismo |
|---|---|---|
| Definição | Lidar com múltiplas tarefas | Executar múltiplas tarefas ao mesmo instante |
| Requer | Design correto | Hardware multicore |
| Exemplo | 1 chef, 3 panelas (alternando) | 3 chefs, 3 panelas (simultâneo) |
| Go provê | ✅ (goroutines + canais) | ✅ (automático com GOMAXPROCS > 1) |

Go é fundamentalmente uma linguagem de concorrência. O paralelismo emerge automaticamente quando GOMAXPROCS > 1 e há trabalho suficiente.

---

### 3.4 Comparação com outras linguagens

| | Go | JavaScript | Java | Python (asyncio) |
|---|---|---|---|---|
| Modelo | Goroutines (M:N) | Event loop (1:1) | Threads SO / Virtual threads (Java 21) | Coroutines (1:1) |
| Sintaxe | `go func()` | `async/await` / callbacks | `new Thread()` / `ExecutorService` | `async def` / `await` |
| Memória/tarefa | ~2KB | ~heap JS | ~1MB (threads tradicionais) | Varia |
| Bloqueio em I/O | Scheduler substitui | Event loop não bloqueia | Thread bloqueia (ou usa NIO) | Yield implícito |
| Compartilhamento de estado | Via canais (idiomático) | Monothread (evita o problema) | Locks / synchronized | GIL (limitação) |

---

## 4. Exemplos Extras

### 4.1 Fan-out simples

```go
func processarItem(id int) {
    time.Sleep(100 * time.Millisecond) // simula trabalho
    fmt.Printf("Item %d processado\n", id)
}

func main() {
    for i := 1; i <= 50; i++ {
        go processarItem(i)
    }
    time.Sleep(2 * time.Second) // aguarda (hack didático)
}
```

50 itens processados em ~100ms em vez de 5 segundos.

### 4.2 Usando `runtime.NumGoroutine` para diagnóstico

```go
fmt.Printf("Goroutines antes: %d\n", runtime.NumGoroutine())
for i := 0; i < 10; i++ {
    go func() { time.Sleep(5 * time.Second) }()
}
fmt.Printf("Goroutines depois: %d\n", runtime.NumGoroutine())
// Output: antes: 1, depois: 11 (main + 10 filhas)
```

---

## 5. Prompts para NotebookLM

### 🎙️ Audio Overview
> "Crie um diálogo de podcast entre dois especialistas em Go: um explicando goroutines para alguém vindo de Java com threads tradicionais, e o outro fazendo perguntas céticas sobre onde isso pode dar errado. Inclua o problema do 'Tchau, Obrigado', goroutine leaks, e o gotcha do loop."

### 🎥 Video Overview
> "Crie um roteiro de vídeo explicativo mostrando, com animação, o que acontece internamente quando 10 goroutines são lançadas: como o Go Runtime Scheduler as distribui nos núcleos de CPU, o que acontece quando uma goroutine bloqueia em I/O, e como isso difere de 10 threads de SO."

### 📊 Slides
> "Crie 8 slides sobre goroutines: (1) problema das threads de SO, (2) goroutine como solução, (3) tabela comparativa, (4) sintaxe go, (5) Go Runtime Scheduler e Modelo M:N, (6) o gotcha do loop com código lado a lado, (7) Tchau-Obrigado e soluções, (8) concorrência vs. paralelismo."

### ❓ Perguntas e Respostas
> "Com base no material do capítulo 15, crie 10 perguntas de revisão de múltipla escolha sobre goroutines, cobrindo: memória, scheduler, sintaxe go, gotcha de closure, Tchau-Obrigado, concorrência vs. paralelismo. Inclua as respostas corretas e explicações."

### 🔍 Estudo Aprofundado
> "Explique o funcionamento interno do Go Runtime Scheduler com o Modelo M:N. Como ele decide qual goroutine rodar? O que é 'work stealing'? Como ele lida com goroutines bloqueadas em syscall vs. I/O de rede? Qual a relação com GOMAXPROCS?"

### 🔄 Comparativo
> "Compare goroutines em Go com: (1) async/await em JavaScript/TypeScript, (2) threads em Java com ExecutorService, (3) coroutines em Kotlin. Quais são as vantagens e desvantagens de cada abordagem para construir um servidor HTTP com alta concorrência?"
