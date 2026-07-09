# 🏋️ Pilha de Exercícios — Capítulo 15: Goroutines

> Exercícios extras sem solução. Use `time.Sleep` como sincronização temporária quando necessário — WaitGroups vêm no Capítulo 18.

---

## 🟢 Fácil

### 1. O Relógio Paralelo

Crie uma função `relogio(nome string, intervalo time.Duration)` que, em loop infinito, imprime `"{nome}: tick"` e dorme pelo `intervalo`. Lance dois relógios como goroutines com intervalos diferentes (ex: 500ms e 800ms). Faça a main dormir 5 segundos e observe a intercalação das impressões.

### 2. A Contagem Regressiva Paralela

Crie 3 funções: `contarAte(n int, nome string)` que imprime `"{nome}: {i}"` de 1 até n, dormindo 100ms entre cada número. Lance as 3 como goroutines com valores diferentes (5, 8, 3). Observe que as contagens se misturam — isso é concorrência em ação.

### 3. O Diagnóstico de Hardware

Crie um programa que imprime: quantos núcleos lógicos a máquina tem (`runtime.NumCPU()`), quantos está usando atualmente (`runtime.GOMAXPROCS(0)` retorna o valor atual), e quantas goroutines estão rodando nesse momento (`runtime.NumGoroutine()`). Rode antes e depois de lançar algumas goroutines.

### 4. O Lançador Sequencial vs. Concorrente

Implemente uma função `trabalho(id, duracaoMs int)` que simula uma tarefa com `time.Sleep`. Primeiro, chame 5 tarefas **sequencialmente** e meça o tempo. Depois, chame as mesmas 5 tarefas com `go` e meça o tempo. Compare os resultados com `time.Since`.

---

## 🟡 Médio

### 5. O Gerador de Thumbnails Simulado

Simule um pipeline de processamento de imagens: crie uma função `processarImagem(id int)` que imprime `"Processando imagem {id}..."`, dorme um tempo aleatório entre 100ms e 500ms (use `rand.Intn`), e imprime `"Imagem {id} pronta!"`. Lance 20 goroutines simultaneamente e meça o tempo total.

### 6. O Closure Correto

Escreva um loop `for i := 0; i < 10; i++` que lança goroutines anônimas. Primeiro, escreva a versão com o bug (sem passar `i` como parâmetro) e documente em comentário qual saída incorreta você observou. Depois escreva a versão corrigida. Compare as saídas.

### 7. A Fábrica de Workers

Crie uma função `worker(id int, tarefas []string)` que processa cada tarefa da lista imprimindo `"Worker {id}: processando '{tarefa}'"` com 200ms de delay entre cada uma. Divida uma lista de 9 tarefas entre 3 workers (3 cada) e lance cada worker como goroutine. Observe a intercalação.

### 8. O Timeout Forçado

Crie uma função `tarefaLenta()` que dorme 10 segundos. Lance-a como goroutine. Na main, use `time.Sleep(2 * time.Second)` e então imprima "Timeout! Continuando sem esperar." Observe que a goroutine foi "abandonada" — o programa termina sem ela completar. O que isso ensina sobre o problema "Tchau, Obrigado"?

---

## 🔴 Desafio

### 9. O Benchmark de Goroutines

Meça quanto tempo leva para criar N goroutines sem fazer nada (`go func() {}()`). Teste com N = 100, 1.000, 10.000, 100.000 e 1.000.000. Use `time.Since` para medir. O que você observa? Isso comprova a leveza das goroutines comparado a threads de SO.

### 10. O Fan-Out Manual

Crie uma lista de 20 URLs fictícias (strings como `"https://api.exemplo.com/item/1"`). Sem usar canais, simule um "download" de cada URL com uma goroutine anônima que dorme um tempo aleatório (50–300ms) e imprime `"Baixado: {url}"`. Controle o fim com `time.Sleep`. Observe que downloads mais rápidos terminam primeiro, independentemente da ordem de lançamento.

### 11. O Goroutine Leak

Crie uma função `vazamento()` que lança uma goroutine com `go func() { for { time.Sleep(time.Second) } }()` — um loop infinito. Chame `vazamento()` 10 vezes. Use `runtime.NumGoroutine()` para mostrar quantas goroutines ficaram rodando. O que acontece quando o programa termina? Isso ilustra o problema de goroutine leak — goroutines que nunca terminam e consomem memória.

### 12. O Pipeline Sem Canal

Implemente 3 estágios de processamento como funções: `estagio1(id int)` (lê dado), `estagio2(id int)` (processa), `estagio3(id int)` (salva). Para cada item de 1 a 10, lance uma goroutine que chama os 3 estágios em sequência dentro da mesma goroutine (sem canais). Compare com a versão sequencial. Reflita: o que seria diferente se cada estágio fosse uma goroutine separada? (Isso é o que canais resolvem — Cap 16.)
