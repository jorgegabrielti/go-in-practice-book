> **Antes de publicar:** conferir checklist no fim do arquivo.

# 🐾 Métodos em Go: quando seus dados aprendem a agir sozinhos

No capítulo anterior, construímos Structs — pastas amarelas organizadas com dados de um paciente, de um produto, de um cachorro. Mas havia algo faltando. Uma struct com `Nome` e `Raca` é como um cachorro de pelúcia: tem forma, ocupa espaço, mas não *faz* nada. Se você quisesse que ele latisse, tinha que criar uma função externa `latir(c Cachorro)` e apertar a barriga do bicho manualmente.

E se o cachorro pudesse latir sozinho? É exatamente isso que os **Métodos** entregam.

---

## 🎯 O que é um Receiver?

Um método em Go é uma função com um parâmetro extra antes do nome — o **Receiver**. Ele "amarra" a função ao tipo:

```go
// Função comum — solta no mundo
func Latir(c Cachorro) {
    fmt.Println(c.Nome, "diz: Au Au!")
}

// Método — pertence ao tipo Cachorro
func (c Cachorro) Latir() {
    fmt.Println(c.Nome, "diz: Au Au!")
}
```

A diferença na chamada é expressiva:

```go
rex := Cachorro{Nome: "Rex"}
Latir(rex)    // função: você carrega o cachorro até a função
rex.Latir()   // método: o cachorro age por conta própria
```

O receiver `(c Cachorro)` declara que `Latir` é propriedade exclusiva do tipo `Cachorro`. `c` é a instância que chamou o método — o equivalente ao `this` de outras linguagens, só que explícito e sem nenhuma mágica escondida.

---

## ⚖️ Value ou Pointer Receiver? A decisão mais importante

Toda vez que você define um método, precisa responder: "esse método muda o estado do objeto?"

**Value Receiver `(c Tipo)`** — Go cria uma cópia para executar o método. Seguro (o original está blindado), mas não consegue alterar nada e é mais lento para structs grandes:

```go
func (c ContaBancaria) Extrato() {
    fmt.Printf("Conta de %s | Saldo: R$ %.2f\n", c.Titular, c.Saldo)
    // c é uma cópia — alterar c.Saldo aqui não afeta a conta real
}
```

**Pointer Receiver `(c *Tipo)`** — Go passa o endereço. Eficiente e capaz de modificar o original:

```go
func (c *ContaBancaria) Depositar(valor float64) {
    if valor > 0 {
        c.Saldo += valor  // modifica a conta real
    }
}
```

**Regra prática**: na dúvida, use Pointer Receiver. É mais barato para a memória e mais flexível. Value Receiver só faz sentido para structs mínimas (um ponto X,Y) que nunca precisam ser alteradas.

---

## 🔧 O Go é gentil: sem `(*c).Campo`

Ao contrário de C, você não precisa desreferenciar manualmente ao usar um Pointer Receiver. O compilador faz isso por você — `c.Saldo` dentro de `(c *ContaBancaria)` já é o saldo real, não uma cópia. Isso remove uma fonte enorme de confusão para quem está migrando de linguagens de nível mais baixo.

---

## 🧱 Métodos além das Structs

Qualquer tipo criado com `type` aceita métodos — não apenas Structs. Isso abre espaço para abstrações elegantes sobre tipos primitivos:

```go
type Dinheiro float64

func (d Dinheiro) String() string {
    return fmt.Sprintf("R$ %.2f", d)
}

var salario Dinheiro = 5000.5
fmt.Println(salario.String())  // R$ 5000.50
```

Um `float64` que sempre sabe se formatar como moeda — sem funções soltas, sem risco de esquecer o formato.

---

## 💡 Dica do Gopher: esqueça Getters e Setters

Quem vem de Java cria reflexivamente `getNome()`, `setNome()`, `getIdade()` para tudo. Em Go, isso é ruído desnecessário. Campo com letra maiúscula já é público — acesse direto:

```go
p.Nome = "Ana"      // acesso direto, Go-way
p.SetNome("Ana")    // desnecessário — não faça isso
```

Só crie um método Setter quando houver lógica de validação real — como `SetIdade` que rejeita valores negativos. Caso contrário, o campo público direto é mais simples, mais legível e mais idiomático.

---

## 🔬 Exemplo do Capítulo: ContaBancaria

O exemplo prático combina os dois tipos de Receiver num único tipo coerente:

```go
type ContaBancaria struct {
    Titular string
    Saldo   float64
}

func (c ContaBancaria) Extrato() { ... }           // Value: só lê
func (c *ContaBancaria) Depositar(v float64) { ... } // Pointer: altera
func (c *ContaBancaria) Sacar(v float64) bool { ... } // Pointer: altera + retorna bool
```

O resultado é uma conta que sabe consultar seu próprio extrato, receber depósitos e rejeitar saques quando o saldo é insuficiente — sem nenhuma função externa e sem expor a lógica de validação para fora do tipo.

---

## 🦫 Conclusão

Métodos transformam Structs de contêineres passivos em entidades ativas. A distinção Value/Pointer Receiver é o coração do capítulo — e ela conecta diretamente com o que aprendemos sobre ponteiros. A partir daqui, o próximo passo natural é **Interfaces**: o mecanismo que permite tratar tipos diferentes de forma uniforme, desde que todos tenham os mesmos métodos.

Repositório completo: [Link do repositório]

---

**Checklist antes de publicar:**
- [x] Conteúdo confere com `estudos-go/capitulo11/fonte.txt`
- [x] Texto é original — analogia do cachorro de pelúcia reescrita com voz própria
- [x] Todos os blocos de código testados e funcionais
- [ ] Link para o repositório no GitHub incluído
- [ ] Revisão ortográfica feita
- [ ] Capa/imagem de destaque escolhida no Medium
- [ ] Tags do Medium definidas (Golang, Programming, OOP, Backend, Methods)
- [ ] Status atualizado em `conteudo/PAINEL.md`
