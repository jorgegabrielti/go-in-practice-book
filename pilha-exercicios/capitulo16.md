# 🏋️ Pilha de Exercícios — Capítulo 16: Channels

> Exercícios extras sem solução. Canais unbuffered são o foco principal — buffered channels chegam nos próximos capítulos.

---

## 🟢 Fácil

### 1. O Canal de Números Pares

Crie uma goroutine que gera os números de 1 a 20 e envia pelo canal apenas os pares. Na main, receba e imprima cada um usando `for range`. Feche o canal ao final da goroutine.

### 2. O Relógio com Canal

Crie uma goroutine que, a cada 500ms, envia a string `"tick"` em um canal de strings. Na main, use `for range` para receber 6 ticks e imprimi-los. Como você vai encerrar o loop? (Dica: feche o canal após os 6 envios na goroutine.)

### 3. O Canal Unidirecional

Crie uma função `gerador(c chan<- int)` que envia os números de 1 a 5. Crie uma função `exibidor(c <-chan int)` que recebe e imprime cada número. Na main, crie o canal bidirecional, passe para ambas as funções e observe que o compilador aceita a conversão implícita.

### 4. O Sinalizador

Crie duas goroutines: `tarefaA` e `tarefaB`, cada uma fazendo um `time.Sleep` de duração diferente (ex: 1s e 2s). Cada uma envia uma string no canal quando termina. Na main, receba duas vezes do canal e imprima qual terminou primeiro.

---

## 🟡 Médio

### 5. O Pipeline de Transformação

Implemente um pipeline de 3 estágios com canais:
- `gerarTexto(c chan<- string)`: envia 5 strings (ex: "olá", "mundo", "go", "channels", "pipeline")
- `converterMaiusculo(entrada <-chan string, saida chan<- string)`: lê de `entrada`, converte com `strings.ToUpper`, envia em `saida`
- `imprimir(c <-chan string)`: lê e imprime cada string

Na main, crie 2 canais, lance as 3 funções como goroutines e feche os canais nos lugares certos.

### 6. O Contador de Palavras Distribuído

Dado um slice de frases, lance uma goroutine por frase. Cada goroutine conta as palavras (`strings.Fields`) e envia a contagem num canal. A main soma todos os resultados. Tente com 5 frases de tamanhos variados.

### 7. O Cache com Canal

Crie uma goroutine que recebe nomes pelo canal, verifica se já foram vistos (usando um `map[string]bool` interno), e envia `true` ou `false` num segundo canal indicando se é duplicata. Teste com: "go", "canal", "go", "goroutine", "canal". Use dois canais separados: um de entrada (`chan string`) e um de resposta (`chan bool`).

### 8. O Fan-In (Junção de Canais)

Crie duas goroutines produtoras: `fonteA` que envia "A1", "A2", "A3" e `fonteB` que envia "B1", "B2", "B3". Crie uma função `merge(a, b <-chan string) <-chan string` que recebe de ambos e retorna um único canal com todos os valores. Na main, consuma o canal mesclado com `for range`.

---

## 🔴 Desafio

### 9. O Semáforo de Goroutines

Implemente um semáforo simples usando um canal buffered (`make(chan struct{}, N)`) para limitar o número de goroutines rodando simultaneamente. Lance 20 goroutines, mas garanta que no máximo 5 rodem ao mesmo tempo. Cada goroutine deve: adquirir o semáforo (enviar no canal), executar (time.Sleep de 300ms), liberar (receber do canal). Use `runtime.NumGoroutine()` para confirmar o limite.

### 10. O Timeout com Canal

Implemente um sistema de timeout sem usar `select` (que vem em capítulos futuros). Crie uma função `comTimeout(trabalho func() string, timeout time.Duration) (string, bool)`. Ela deve: lançar `trabalho` numa goroutine que envia o resultado num canal; lançar um temporizador que envia `""` num segundo canal após `timeout`. Retorne o resultado e `true` se `trabalho` terminou antes, ou `""` e `false` se passou do tempo.

### 11. O Número Primo Distribuído

Gere os números de 2 a 100. Para cada número, lance uma goroutine que verifica se é primo (divisão por todos até √n) e, se for, envia no canal. Na main, colete todos os primos e imprima-os em ordem. Dica: você precisará de uma forma de saber quando todas as goroutines terminaram — use `sync.WaitGroup` em conjunto com o canal.

### 12. O Pipeline com Cancelamento

Implemente um pipeline de geração de números (1, 2, 3, ...) que pode ser cancelado. Crie um canal `done chan struct{}`. A goroutine geradora deve parar de enviar se `done` for fechado. Na main, consuma 10 números, feche `done`, e confirme com `runtime.NumGoroutine()` que a goroutine geradora encerrou. (Este é o padrão de cancelamento que o `context.Context` formaliza no Go.)
