> **Antes de publicar:** conferir checklist no fim do arquivo.

# 🔗 A linha de montagem que nunca para: o modelo mental de channels em Go

Há um bug clássico em sistemas concorrentes que todo desenvolvedor descobre da pior forma: dois processos tentando modificar o mesmo dado ao mesmo tempo. Resultado: corrupção silenciosa, race condition, horas de debug.

Go tem uma resposta diferente para esse problema. Não é um lock. Não é um mutex. É uma mudança de paradigma.

---

## O problema que locks não resolvem bem

Imagine dois operários numa fábrica compartilhando um caderno de anotações. Para não escreverem ao mesmo tempo, colocam uma chave na mesa: quem tem a chave, pode escrever. O outro espera.

Funciona. Mas escala mal. Com 10 operários disputando a mesma chave, a maioria do tempo é gasto esperando — não trabalhando. E se alguém esquecer de devolver a chave, o sistema trava para sempre.

Go propõe algo diferente: em vez de compartilhar o caderno, um operário escreve e **passa o caderno** para o próximo. Só um tem o caderno por vez — não por regra imposta, mas pela própria mecânica da troca.

Esse é o channel.

---

## O tubo tipado

Um channel em Go é um tubo por onde dados fluem entre goroutines. O tubo tem um tipo — `chan int` aceita só inteiros, `chan string` só strings. Tentar passar o tipo errado é erro de compilação, não de runtime.

```go
c := make(chan int)  // cria o tubo; canais são referências, precisam de make
```

Para usar: uma seta (`<-`) que indica para onde o dado vai.

```go
c <- 42       // envia: o dado entra no tubo
x := <-c      // recebe: o dado sai do tubo
```

A direção da seta é sempre em relação ao canal. Se aponta para o canal, está entrando. Se aponta para fora, está saindo.

---

## O bloqueio como feature, não bug

Aqui está o detalhe que muda tudo: por padrão, canais em Go são **unbuffered**. O tubo não tem capacidade de armazenamento. Ele é estreito.

Consequência:

Quando uma goroutine envia (`c <- valor`), ela para e espera. Não continua. Fica parada naquela linha até outra goroutine estar pronta para receber do outro lado.

Quando uma goroutine recebe (`x := <-c`), ela também para e espera até alguém enviar.

A troca só acontece quando **as duas goroutines estão prontas ao mesmo tempo**. Isso é sincronização automática — não é efeito colateral, é o design.

Na prática isso significa que você não precisa de `time.Sleep` para garantir que uma goroutine esperou outra. O canal já faz isso.

---

## O que acontece quando não tem ninguém do outro lado

```go
func main() {
    c := make(chan int)
    c <- 10  // trava aqui para sempre
}
```

O Go Runtime detecta quando todas as goroutines estão bloqueadas esperando algo que nunca vai acontecer. Ele encerra o programa com:

```
fatal error: all goroutines are asleep - deadlock!
```

É um mecanismo de segurança. O runtime sabe que se ninguém pode avançar, o programa está preso — e prefere morrer com diagnóstico a ficar rodando em loop infinito silencioso.

---

## Canais têm fim: `close` e `for range`

Quando o produtor termina de enviar dados, ele sinaliza isso fechando o canal:

```go
close(c)  // como lacrar um envelope: "não tem mais nada aqui"
```

O receptor pode usar `for range` no canal — o loop encerra automaticamente quando o canal fecha, sem precisar saber quantos itens virão:

```go
for item := range c {
    processar(item)
}
// sai aqui quando c for fechado
```

Uma regra importante: só quem envia deve fechar. Fechar um canal e continuar enviando gera `panic`. Receber de um canal fechado retorna o zero value do tipo — não trava, não quebra.

---

## Canais direcionais: o compilador como guarda-costas

Você pode passar canais como parâmetros para funções. E pode especificar na assinatura se a função só pode enviar ou só pode receber:

```go
func gerador(c chan<- int) {  // send-only: só pode enviar
    c <- 1
    c <- 2
    close(c)
}

func consumidor(c <-chan int) {  // receive-only: só pode receber
    for v := range c {
        fmt.Println(v)
    }
}
```

Se `gerador` tentar receber (`x := <-c`), o compilador rejeita. Se `consumidor` tentar enviar, idem. Você declara a intenção na assinatura e o compilador garante o contrato — sem runtime error, sem surpresa.

---

## A linha de montagem

O padrão mais poderoso que channels habilitam é o pipeline: goroutines em série, cada uma fazendo uma transformação e passando adiante pelo canal.

```
[gerador] --canal1--> [transformador] --canal2--> [consumidor]
```

Cada estágio roda de forma independente. O bloqueio dos canais garante que ninguém produza mais rápido do que o próximo consegue consumir. É backpressure automático — sem filas explícitas, sem semáforos.

---

## 🦫 O que ficou

Channels não são uma feature opcional do Go — são a forma como Go pensa sobre concorrência. Em vez de proteger dados compartilhados com locks, você transfere a posse dos dados entre goroutines. Quem tem o dado, trabalha com ele. Quem não tem, espera.

No próximo capítulo: `select` — como uma goroutine pode esperar em múltiplos canais ao mesmo tempo.

Repositório completo: [Link do repositório]

---

**Checklist antes de publicar:**
- [x] Conteúdo confere com `estudos-go/capitulo16/fonte.txt`
- [x] Texto é original — sem frases copiadas do fonte.txt
- [x] Todos os blocos de código testados e funcionais
- [ ] Link para o repositório no GitHub incluído
- [ ] Revisão ortográfica feita
- [ ] Capa/imagem de destaque escolhida no Medium
- [ ] Tags do Medium definidas (Golang, Channels, Concurrency, Backend, Programming)
- [ ] Status atualizado em `conteudo/PAINEL.md`
