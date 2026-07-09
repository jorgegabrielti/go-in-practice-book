> **Antes de publicar:** conferir checklist no fim do arquivo.

# 🍳 O segredo do Go que deixa outros servidores com inveja

Existe uma pergunta que todo desenvolvedor faz quando descobre Go: "como um programa consegue lidar com milhares de requisições simultâneas consumindo tão pouca memória?"

A resposta tem um nome: goroutine.

---

## O problema que todo mundo ignora até quebrar em produção

Imagine que você tem um servidor que recebe pedidos. Cada pedido precisa buscar dados em um banco, esperar uma resposta de uma API externa, processar o resultado. Com threads tradicionais — o jeito de Java, C++, Python — cada pedido ocupa uma thread. Uma thread ocupa cerca de 1MB de memória só para existir. Dez mil pedidos simultâneos? Dez gigabytes só em threads paradas esperando I/O.

Esse é o modelo que Go jogou no lixo.

---

## A thread que não é uma thread

Uma goroutine começa com cerca de 2KB de memória. Quando precisa de mais, a pilha cresce. Quando libera, encolhe. Você pode ter um milhão de goroutines rodando num servidor modesto sem transpirar.

A diferença não é só de tamanho. É de filosofia.

Thread de SO é um cidadão do sistema operacional — pesado, caro de criar, caro de trocar. Goroutine é um cidadão do runtime do Go — leve, gerenciado internamente, transparente para você.

Para lançar qualquer função como goroutine, existe exatamente uma palavra:

```go
go minhaFuncao()
```

É isso. Não tem pool de threads para configurar, não tem executor service, não tem ThreadPoolExecutor com número mágico de workers. Uma palavra, e a função começa a rodar de forma independente.

---

## O gerente que você nunca vê

Se você tem 10.000 goroutines e 8 núcleos de CPU, alguém precisa decidir qual goroutine roda em qual núcleo. Esse alguém é o Go Runtime Scheduler.

Ele opera em background, invisível, usando um modelo chamado M:N — muitas goroutines multiplexadas em poucas threads de SO. O que torna ele especialmente eficiente é o que acontece quando uma goroutine bloqueia: em vez de deixar um núcleo ocioso esperando um arquivo ou uma resposta de rede, o scheduler imediatamente substitui a goroutine bloqueada por outra que tem trabalho a fazer.

Do seu lado como programador: você não configura nada disso. Você escreve `go` e segue em frente.

---

## O que "síncrono" e "assíncrono" realmente significam aqui

```go
buscarDados()    // síncrono: próxima linha só executa quando isso terminar
go buscarDados() // assíncrono: próxima linha executa imediatamente
```

Quando você escreve `go buscarDados()`, o programa não espera. Ele lança a tarefa e continua. A goroutine roda em background, no seu próprio ritmo.

Isso tem uma consequência que pega muita gente de surpresa: a função `main` também é uma goroutine. Quando ela termina, o programa inteiro morre — incluindo todas as goroutines filhas que ainda estavam trabalhando.

```go
func main() {
    go fazerAlgoImportante()
    // programa termina aqui — fazerAlgoImportante pode nunca completar
}
```

É como o dono do restaurante apagar a luz sem avisar os cozinheiros. Para esperar os filhos terminarem, existe o `sync.WaitGroup` — que o livro aborda no capítulo 18. Por ora, o truque didático é um `time.Sleep`.

---

## O armadilha que todo iniciante cai

Goroutines anônimas dentro de loops têm um comportamento contraintuitivo:

```go
for i := 0; i < 5; i++ {
    go func() {
        fmt.Println(i)  // qual valor de 'i' isso vai imprimir?
    }()
}
```

A resposta não é 0, 1, 2, 3, 4. É provavelmente 5, 5, 5, 5, 5.

Por quê? A closure captura a variável `i` por referência — não uma cópia. Quando as goroutines finalmente rodam, o loop já terminou e `i` vale 5. A correção é simples e idiomática:

```go
for i := 0; i < 5; i++ {
    go func(v int) {
        fmt.Println(v)  // v é uma cópia de i no momento do lançamento
    }(i)
}
```

---

## Concorrência não é paralelismo

Rob Pike, um dos criadores do Go, tem uma distinção que ficou famosa: concorrência é sobre lidar com muitas coisas ao mesmo tempo; paralelismo é sobre fazer muitas coisas ao mesmo tempo.

Um único chef alternando entre três panelas é concorrente. Três chefs, cada um na sua panela, são paralelos.

Go te dá ferramentas de concorrência — goroutines e canais. Se a máquina tiver múltiplos núcleos disponíveis, o scheduler automaticamente distribui o trabalho em paralelo. Você escreve concorrente; o hardware decide se vira paralelo.

---

## Na prática: 10 robôs em 2 segundos

O exemplo do capítulo é revelador. Dez funções que cada uma demora 2 segundos. Sequencial: 20 segundos. Com goroutines:

```go
for i := 1; i <= 10; i++ {
    go tarefaPesada(i)
}
time.Sleep(3 * time.Second)
```

Tempo total: ~2 segundos. Os 10 robôs trabalharam ao mesmo tempo.

---

## 🦫 O que ficou

Goroutine é a aposta do Go. Uma abstração leve sobre concorrência que permite escrever código direto — sem callbacks, sem promises, sem async/await — mas com execução naturalmente concorrente. O scheduler cuida da distribuição. Você cuida da lógica.

No próximo capítulo: canais — a forma como goroutines conversam entre si sem guerra de dados.

Repositório completo: [Link do repositório]

---

**Checklist antes de publicar:**
- [x] Conteúdo confere com `estudos-go/capitulo15/fonte.txt`
- [x] Texto é original — sem frases copiadas do fonte.txt
- [x] Todos os blocos de código testados e funcionais
- [ ] Link para o repositório no GitHub incluído
- [ ] Revisão ortográfica feita
- [ ] Capa/imagem de destaque escolhida no Medium
- [ ] Tags do Medium definidas (Golang, Concurrency, Goroutines, Backend, Programming)
- [ ] Status atualizado em `conteudo/PAINEL.md`
