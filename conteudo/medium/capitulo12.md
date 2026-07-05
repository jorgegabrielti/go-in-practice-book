> **Antes de publicar:** conferir checklist no fim do arquivo.

# 🔌 Interfaces em Go: o contrato que liberta

Existe um problema que todo desenvolvedor enfrenta cedo ou tarde: você escreve uma função que funciona perfeitamente para um tipo específico — e de repente precisa que ela funcione para dez tipos diferentes. Em Java, você vai ao cartório: `class Cachorro implements Barulhento`. Em Go, você simplesmente... escreve o método.

Esse é o ponto central do capítulo 12.

---

## A tomada elétrica não sabe o que você vai ligar

Pense em uma tomada na parede. Ela tem um contrato físico: "dois pinos redondos com tal espessura". Não importa se é uma geladeira, um ventilador ou uma furadeira. Se encaixa, recebe energia.

Interfaces em Go funcionam assim. Elas definem **o que** um tipo deve saber fazer, sem se importar com **quem** vai fazer nem **como**. Uma interface é pura definição — sem campos, sem código, sem lógica:

```go
type Barulhento interface {
    FazerBarulho() string
}
```

Qualquer tipo que tiver um método `FazerBarulho() string` automaticamente se torna um `Barulhento`. Sem contrato no cartório. Sem `implements`. Sem herança.

---

## Implementando sem saber que está implementando

Esse é o detalhe que muda a cabeça de quem vem de Java ou C#. Em Go, a satisfação de interface é **implícita**. Você cria o método e pronto — o compilador verifica sozinho.

```go
type Cachorro struct{}

func (c Cachorro) FazerBarulho() string {
    return "Au Au"
}

type Alarme struct{}

func (a Alarme) FazerBarulho() string {
    return "TRIMMMM!"
}
```

Um cachorro e um despertador. Nada mais diferente. Mas ambos agora são `Barulhento` — e podem ser tratados da mesma forma:

```go
func IncomodarVizinho(b Barulhento) {
    fmt.Println("O vizinho está ouvindo:", b.FazerBarulho())
}

IncomodarVizinho(Cachorro{})  // O vizinho está ouvindo: Au Au
IncomodarVizinho(Alarme{})    // O vizinho está ouvindo: TRIMMMM!
```

Isso é polimorfismo — tratar coisas diferentes da mesma maneira.

---

## O coringa: a interface vazia

Se uma interface define requisitos, o que acontece quando você não coloca nenhum?

```go
interface{}  // ou: any (Go 1.18+)
```

Uma interface vazia aceita qualquer coisa, porque todo tipo do Go implementa zero métodos por padrão. É assim que `fmt.Println` funciona — ele recebe `...any` e imprime qualquer coisa que você jogar nele.

Mas há um porém importante: usar `any` em excesso transforma Go em JavaScript. Você perde o compilador como aliado. Use com propósito, não como atalho.

---

## Abrindo a caixa misteriosa: Type Assertion

Quando você recebe um `any` (ou outra interface), às vezes precisa saber o que está dentro para acessar campos específicos. É como receber uma caixa fechada e tentar adivinhar o conteúdo.

Existem dois jeitos de fazer isso. O jeito arriscado:

```go
c := v.(Cachorro)  // panic se v não for Cachorro
```

E o jeito que o Gopher idiomático usa:

```go
dog, ok := v.(Cachorro)  // o comma ok de volta
if ok {
    fmt.Println("Raça:", dog.Raca)
} else {
    fmt.Println("Não era um cachorro.")
}
```

Para testar vários tipos em sequência, existe o Type Switch — um `switch` que distribui por tipo em vez de por valor:

```go
switch valor := v.(type) {
case int:
    fmt.Println("É um número:", valor*2)
case string:
    fmt.Println("É um texto de", len(valor), "chars")
default:
    fmt.Println("Não reconhecido")
}
```

Dentro de cada `case`, `valor` já tem o tipo concreto — sem cast adicional.

---

## O exemplo que une tudo: formas geométricas

O `main.go` do capítulo mostra a beleza das interfaces em ação. Uma única função `ExibirDetalhes` aceita um `Retangulo` e um `Circulo` — tipos completamente diferentes — e calcula área e perímetro de cada um sem saber qual veio:

```go
type Geometrico interface {
    Area() float64
    Perimetro() float64
}

func ExibirDetalhes(g Geometrico) {
    fmt.Printf("Área: %.2f\n", g.Area())
    fmt.Printf("Perímetro: %.2f\n", g.Perimetro())

    // Type Assertion: comportamento específico para Círculo
    if c, ok := g.(Circulo); ok {
        fmt.Printf("Raio: %.2f\n", c.Raio)
    }
}
```

Adicionar um Triângulo amanhã? Basta implementar `Area()` e `Perimetro()`. `ExibirDetalhes` não muda uma linha.

---

## 💡 A dica que muda o design: defina onde usa

Em Java, quem cria a classe `Arquivo` costuma criar também a interface `IArquivo`. Em Go, a lógica é inversa: **a interface pertence a quem consome**.

Se sua função só precisa ler dados, você define a interface `Leitor` localmente:

```go
type Leitor interface {
    Ler() string
}
```

Isso significa que qualquer biblioteca no mundo que tenha um método `Ler() string` pode ser passada para sua função — sem precisar importar nada. **Interfaces pequenas (1 ou 2 métodos) são as mais poderosas.** `io.Reader` e `fmt.Stringer` são os exemplos canônicos da stdlib.

---

## 🦫 Conclusão

Interfaces são o mecanismo que permite que Go seja simples e flexível ao mesmo tempo. Sem herança, sem hierarquia de classes, sem `implements` — apenas contratos implícitos verificados pelo compilador. Escreva o método, satisfaça a interface.

O próximo capítulo vai aprofundar esse poder: **Goroutines e Canais** — onde Go realmente se diferencia com concorrência nativa.

Repositório completo: [Link do repositório]

---

**Checklist antes de publicar:**
- [x] Conteúdo confere com `estudos-go/capitulo12/fonte.txt`
- [x] Texto é original — analogia da tomada reescrita com voz própria
- [x] Todos os blocos de código testados e funcionais
- [ ] Link para o repositório no GitHub incluído
- [ ] Revisão ortográfica feita
- [ ] Capa/imagem de destaque escolhida no Medium
- [ ] Tags do Medium definidas (Golang, Programming, Interfaces, Backend, OOP)
- [ ] Status atualizado em `conteudo/PAINEL.md`
