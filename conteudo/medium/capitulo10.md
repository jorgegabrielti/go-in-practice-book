> **Antes de publicar:** conferir checklist no fim do arquivo.

# 🗂️ Structs em Go: pare de voar com 300 variáveis soltas

No capítulo anterior, aprendemos que ponteiros são papéis com endereços — eficientes para compartilhar e modificar dados. Hoje esse conceito ganha um parceiro natural: a **Struct**, a estrutura que agrupa dados relacionados em uma entidade coerente.

A situação sem Struct é conhecida. Você está modelando pacientes numa clínica: `nome1`, `idade1`, `peso1`, `nome2`, `idade2`, `peso2`... Cem pacientes depois, são 300 variáveis soltas voando pelo código. Passar os dados do João para um especialista exige passar 3 variáveis separadas — e esquecer uma delas pode custar um diagnóstico errado.

A Struct resolve isso com uma metáfora simples: a **Pasta Amarela**. Em vez de papéis soltos, você grampeia Nome, Idade e Peso numa pasta e entrega a pasta inteira.

---

## 📐 A planta baixa do seu dado

Definir uma Struct é criar um molde — a forma que os dados vão ter. Ela não ocupa memória sozinha, apenas diz ao compilador o que esperar:

```go
type Pessoa struct {
    Nome  string
    Idade int
    Peso  float64
}
```

A partir daí, `Pessoa` é um tipo de primeira classe no seu programa — tão válido quanto `int` ou `string`. Você pode criar variáveis do tipo `Pessoa`, passar para funções, guardar em slices, usar como chave de map.

---

## 🖊️ Três formas de preencher o formulário

**Forma 1 — literal com nomes (a certa):**

```go
cliente := Pessoa{
    Nome:  "João",
    Idade: 30,
    Peso:  80.5,
}
```

Resistente a mudanças: se você reordenar os campos na definição da struct, esse código continua funcionando.

**Forma 2 — posicional (a perigosa):**

```go
cliente := Pessoa{"César", 30, 80.5}
```

Omite os nomes, depende da ordem exata. Se você inverter `Idade` e `Peso` na definição amanhã, esse código compila mas produz dados errados — silenciosamente. Evite.

**Forma 3 — zero value:**

```go
var cliente Pessoa
// cliente.Nome == "", cliente.Idade == 0, cliente.Peso == 0.0
```

Todos os campos recebem seus zero values automaticamente. Útil quando você vai preencher campo a campo depois.

---

## 🔗 Structs dentro de Structs (A Boneca Russa)

Um campo de uma struct pode ser de outro tipo struct. O endereço é mais complexo do que uma simples string? Modele como struct também:

```go
type Endereco struct {
    Rua    string
    Cidade string
}

type Usuario struct {
    Nome     string
    Endereco Endereco
}
```

O acesso encadeia pontos:

```go
fmt.Println(u.Endereco.Cidade)  // "São Paulo"
```

---

## 📌 Ponteiro para Struct: o atalho que o Go te dá

Combinando o capítulo anterior com este: quando você cria um ponteiro para uma struct, o Go dispensa a sintaxe `(*p).Campo` e deixa usar `p.Campo` diretamente. O compilador desreferencia por você:

```go
p2 := &p1
p2.Preco = 200.00  // equivale a (*p2).Preco = 200.00 — altera p1 original
```

Isso é extremamente comum em Go. Funções que modificam structs quase sempre recebem `*MinhaStruct` por parâmetro.

---

## 🧩 Composição no lugar de Herança

Se você vem de Java ou C#, vai procurar a palavra `extends`. Ela não existe em Go. Em vez de "Cachorro **é um** Animal" (herança), Go usa "Cachorro **tem um** Animal" (composição via Embedding):

```go
type Animal struct {
    Vida int
}

type Cachorro struct {
    Animal       // Embedding — sem nome de campo
    Raca  string
}
```

O Go promove automaticamente os campos de `Animal` para `Cachorro` — é o que chamam de **Promoted Fields**. Você acessa `cachorro.Vida` diretamente, como se o campo fosse do próprio `Cachorro`. Sem hierarquia de classes, sem `super`, sem complicação.

---

## 💡 Dica do Gopher: não seja o Maníaco do Ponteiro (de volta)

No capítulo anterior, a regra era: não use ponteiro para tudo. Ela continua válida aqui. Structs pequenas (2-3 campos de tipos básicos) são mais rápidas por valor — a CPU as mantém em cache. Use `*MinhaStruct` quando a struct for grande, ou quando precisar que a função modifique o original.

---

## 🔬 Exemplo do Capítulo

O exemplo prático modela um sistema de e-commerce simples — `Produto` e `Pedido` com `[]Produto` — e demonstra os três pontos do capítulo juntos: instância literal, ponteiro para struct alterando o original, e loop em slice de structs calculando total:

```go
p2 := &p1
p2.Preco = 200.00  // p1.Preco agora é 200

pedido := Pedido{
    ID: 1001, Cliente: "Roberto",
    Itens: []Produto{p1, {Nome: "Mouse", Preco: 50.0}},
}

total := 0.0
for _, item := range pedido.Itens {
    total += item.Preco
}
fmt.Printf("Total do Pedido %d: R$ %.2f\n", pedido.ID, total)
// Total do Pedido 1001: R$ 250.00
```

---

## 🦫 Conclusão

Structs são a base de tudo que vem pela frente em Go: métodos, interfaces, JSON, bancos de dados. Dominar a criação, o acesso a campos, o aninhamento e a combinação com ponteiros é o que separa o código com variáveis soltas do código que realmente modela o domínio do problema.

No próximo capítulo, vamos adicionar comportamento às Structs — os **Métodos**.

Repositório completo: [Link do repositório]

---

**Checklist antes de publicar:**
- [x] Conteúdo confere com `estudos-go/capitulo10/fonte.txt`
- [x] Texto é original — analogia da Pasta Amarela reescrita com voz própria
- [x] Todos os blocos de código testados e funcionais
- [ ] Link para o repositório no GitHub incluído
- [ ] Revisão ortográfica feita
- [ ] Capa/imagem de destaque escolhida no Medium
- [ ] Tags do Medium definidas (Golang, Programming, Software Development, Backend, Structs)
- [ ] Status atualizado em `conteudo/PAINEL.md`
