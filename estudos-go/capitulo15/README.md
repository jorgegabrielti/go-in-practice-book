# 📖 Capítulo 15: Goroutines — A Cozinha Frenética e o Mágico Invisível

> **Livro Go na Prática: 30 capítulos para dominar a linguagem**

---

> **Fluxo de trabalho deste capítulo:**
> 1. Transcrever o conteúdo do livro para `fonte.txt` (texto bruto, fiel ao material original).
> 2. Escrever este `README.md` como síntese própria a partir do `fonte.txt`.
> 3. Resolver exemplos/exercícios em `exemplos/` e `exercicios/`.
> 4. Produzir o artigo de Medium e o post de LinkedIn em `conteudo/`, **sempre baseados no `fonte.txt`** — não em memória ou suposição do que o livro disse, e **escritos com palavras próprias** (sem copiar/parafrasear de perto o `fonte.txt`).
> 5. Rastrear a execução de todo o código do capítulo em `teste-de-mesa.md`, prevendo a saída linha a linha antes de rodar `go run`.

> 🧮 [Teste de mesa de todo o código deste capítulo](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo15/teste-de-mesa.md)

---

## 🎯 Tema Central

Goroutines são a forma do Go executar tarefas concorrentes com custo mínimo — 2KB de memória vs. 1MB de uma thread de SO. Com a palavra `go`, qualquer função vira uma tarefa independente. O Go Runtime Scheduler cuida do resto, invisível para você.

---

## 📊 Resumo dos Conceitos

### 1. Goroutine vs. Thread de SO

| | Thread de SO | Goroutine |
|---|---|---|
| Memória inicial | ~1 MB | ~2 KB |
| Criação | Cara (syscall) | Barata (runtime Go) |
| Limite prático | ~10.000 | Milhões |
| Troca de contexto | Microssegundos (caro) | Nanossegundos (barato) |
| Pilha | Tamanho fixo | Cresce/diminui dinamicamente |

### 2. A Sintaxe `go`

Basta prefixar a chamada com `go`:

```go
func cozinhar() { fmt.Println("Cozinhando...") }

cozinhar()     // síncrono — bloqueia até terminar
go cozinhar()  // assíncrono — lança e segue IMEDIATAMENTE
```

### 3. Go Runtime Scheduler — Modelo M:N

O scheduler do Go multiplexar M Goroutines em N Threads de SO (onde N ≈ número de núcleos). Quando uma goroutine bloqueia (I/O, sleep), o scheduler a retira do operário e coloca outra no lugar — sem desperdiçar CPU.

```go
runtime.NumCPU()         // quantos núcleos lógicos
runtime.GOMAXPROCS(1)    // limitar a 1 núcleo (útil para testes)
```

### 4. Goroutines Anônimas

```go
go func() {
    fmt.Println("tarefa rápida em background")
}()  // () executa a função imediatamente como goroutine
```

### 5. O Gotcha do Loop — Captura de Variável

```go
// ❌ Perigoso — todas as goroutines lêem o mesmo 'i'
for i := 0; i < 5; i++ {
    go func() { fmt.Println(i) }()  // pode imprimir "5 5 5 5 5"
}

// ✅ Correto — passa uma cópia de 'i' por parâmetro
for i := 0; i < 5; i++ {
    go func(v int) { fmt.Println(v) }(i)
}
```

> **Nota:** Go 1.22+ corrigiu esse comportamento para variáveis de loop `for`. Em versões anteriores, era necessário sempre passar como parâmetro.

### 6. O Problema "Tchau, Obrigado"

`main` é também uma goroutine. Quando ela termina, **todo o programa morre** — goroutines filhas incluídas. Solução temporária (didática): `time.Sleep`. Solução profissional: `sync.WaitGroup` (Capítulo 18).

---

## 💡 Dica do Gopher — Concorrência ≠ Paralelismo

Rob Pike: *"Concorrência é sobre lidar com muitas coisas ao mesmo tempo. Paralelismo é sobre fazer muitas coisas ao mesmo tempo."*

- **Concorrência** = design do programa (como as tarefas se alternam)
- **Paralelismo** = execução simultânea real (depende de múltiplos núcleos)

Go dá ferramentas de concorrência. Se o hardware permitir, vira paralelismo automaticamente.

---

## 🔬 Exemplos Práticos no Repositório

### 1. [Exemplo 01: A Linha de Montagem de Robôs](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo15/exemplos/exe01/main.go)
*   **Conceito:** 10 goroutines concorrentes, `runtime.NumCPU()`, `time.Sleep` como sincronização temporária.

---

## 🔬 Exercícios Práticos Resolvidos

### 1. [Exercício 01: O Observador de CPUs](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo15/exercicios/exe01/main.go)
*   **Objetivo:** explorar `runtime.NumCPU()` e `runtime.GOMAXPROCS`.
*   **Conceito:** paralelismo, núcleos, scheduler.

### 2. [Exercício 02: O Loop Traiçoeiro](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo15/exercicios/exe02/main.go)
*   **Objetivo:** provocar e corrigir o gotcha de captura de variável em goroutine anônima.
*   **Conceito:** closures, captura por referência vs. por valor.

### 3. [Exercício 03: O Ping-Pong Solitário](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo15/exercicios/exe03/main.go)
*   **Objetivo:** demonstrar que sem canais não há sincronia garantida entre goroutines.
*   **Conceito:** não-determinismo, limitações de `time.Sleep`, motivação para canais (Cap 16).

---

## ✅ Checklist antes de marcar como concluído

- [x] `fonte.txt` com a transcrição completa do capítulo
- [x] Teoria revisada e resumida neste README (com base no `fonte.txt`)
- [x] Todos os exemplos do capítulo têm pasta própria em `exemplos/exe01`
- [x] Todos os exercícios resolvidos têm pasta própria em `exercicios/exe01`, `exe02`, `exe03`
- [x] `go build ./...` e `go vet ./...` passam sem erros
- [x] `go fmt ./...` executado
- [x] Termos novos adicionados ao `estudos-go/GLOSSARIO.md`
- [x] Painel de progresso no `README.md` raiz atualizado
- [x] Artigo de Medium criado em `conteudo/medium/capitulo15.md`
- [x] Post de LinkedIn criado em `conteudo/linkedin/capitulo15.md`
- [x] Pilha de exercícios extras criada em `pilha-exercicios/capitulo15.md`
- [x] Teste de mesa criado em `teste-de-mesa.md`
- [x] Fonte NotebookLM criada em `notebooklm/capitulo15.md`
