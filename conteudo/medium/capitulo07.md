> **Antes de publicar:** conferir checklist no fim do arquivo.

# 🗂️ Go na Prática #07: Por Que o Go Tem Duas Formas de Guardar Listas — e Quando Usar Cada Uma

No capítulo anterior, exploramos funções — as mini-fábricas que evitam repetição de código. Agora chegamos num dos tópicos que mais gera confusão em quem está migrando de outras linguagens: as duas estruturas de dados sequenciais do Go, o **Array** e o **Slice**. Eles parecem fazer a mesma coisa, mas têm comportamentos radicalmente diferentes — e entender essa diferença vai poupar horas de depuração no futuro.

---

## 🎟️ O Array: reserva de lugares com quantidade fixa

Pense numa sala de cinema. Antes do filme começar, o sistema de reservas trava o número de assentos: se a sala tem 200 lugares, ela sempre terá 200 — nem mais, nem menos. Mesmo que apenas 50 pessoas apareçam, os outros 150 assentos continuam existindo, ocupando espaço físico na sala.

O Array em Go funciona assim: você define o tamanho na declaração, e ele faz parte do tipo da variável.

```go
var notas [3]int  // três gavetas, todas começando em zero: [0 0 0]
```

Isso tem uma implicação importante que surpreende quem vem de Python ou JavaScript: `[3]int` e `[4]int` são **tipos diferentes**. Você não pode passar um para uma função que espera o outro, nem compará-los diretamente. O compilador trata o tamanho como parte da identidade do tipo — como se `int32` e `int64` fossem tipos separados (porque são).

```go
vogais := [5]string{"A", "E", "I", "O", "U"}
// Tentar fazer vogais[5] = "W" → erro de compilação: índice fora dos limites
```

Arrays são rápidos e previsíveis em memória justamente por isso — o compilador sabe exatamente quanto espaço reservar. Mas no dia a dia, raramente você sabe antecipadamente quantos elementos vai precisar. É aí que entra o Slice.

---

## 🧾 O Slice: a lista que cresce conforme a necessidade

Esqueça a sala de cinema com lugares fixos. Imagine agora uma lista de presença digital que começa vazia e recebe um nome novo a cada pessoa que chega à festa. A lista nunca fica "cheia" — ela simplesmente cresce.

O Slice não armazena dados por conta própria. Por baixo dos panos, ele aponta para um Array comum, mas expõe apenas uma "janela" desse Array. Internamente, um Slice tem três campos:

| Campo | O que guarda |
| :--- | :--- |
| **ponteiro** | onde o Array real começa na memória |
| **len** | quantos elementos estão visíveis agora |
| **cap** | quantos elementos cabem no Array antes de precisar realocar |

```go
var lista []string           // slice nil: ponteiro=nil, len=0, cap=0
numeros := make([]int, 5)    // len=5, cap=5, todos os valores em 0
```

A diferença visual entre Array e Slice é só um par de colchetes: `[5]int` (Array) vs `[]int` (Slice). Mas as consequências são enormes.

---

## ➕ `append`: como a lista cresce

A função `append` é o coração do trabalho com Slices. Você passa o Slice atual e os valores a adicionar, e ela devolve um Slice novo (ou o mesmo, expandido):

```go
numeros := []int{1, 2}
numeros = append(numeros, 3)         // [1 2 3]
numeros = append(numeros, 4, 5, 6)  // múltiplos de uma vez: [1 2 3 4 5 6]
```

> ⚠️ **Nunca esqueça o `=`**. `append` não modifica o Slice original no lugar — ela retorna um novo Slice. Se você escrever apenas `append(numeros, 3)` sem reatribuir, o valor 3 é simplesmente descartado.

Quando a capacidade atual não comporta o novo elemento, o Go executa uma sequência nos bastidores: cria um Array maior (tipicamente o dobro do tamanho atual), copia todos os dados antigos para ele, adiciona o novo elemento e devolve um Slice apontando para esse Array novo. O Array antigo fica disponível para o garbage collector.

```go
numeros := make([]int, 0, 2)  // len=0, cap=2
numeros = append(numeros, 1)  // len=1, cap=2
numeros = append(numeros, 2)  // len=2, cap=2 (cheio!)
numeros = append(numeros, 3)  // cap dobra: len=3, cap=4 (novo array)
```

---

## ✂️ Fatiamento: pegando pedaços de um Slice

A sintaxe `slice[min:max]` cria um Slice que enxerga apenas uma parte do Array subjacente. O intervalo é semi-aberto: inclui `min`, exclui `max`.

```go
times := []string{"Flamengo", "Palmeiras", "Santos", "Grêmio"}
nordeste := times[1:3]  // ["Palmeiras", "Santos"] — índice 1 e 2, para no 3
```

Atalhos úteis:

```go
times[:2]   // do início até o índice 2 (exclusive): ["Flamengo", "Palmeiras"]
times[2:]   // do índice 2 até o fim: ["Santos", "Grêmio"]
times[:]    // cópia da janela, mesmo Array por baixo
```

---

## 🚨 A Armadilha da Memória Compartilhada

Este é o ponto mais importante — e o mais perigoso para quem não conhece.

Quando você fatia um Slice, o resultado **não é uma cópia independente**. Os dois Slices — o original e o fatiado — apontam para o mesmo Array na memória. Modificar um modifica o outro.

```go
herois := []string{"Batman", "Superman", "Mulher Maravilha"}
time1 := herois[:2]  // ["Batman", "Superman"]

time1[0] = "Coringa"

fmt.Println(herois)  // ["Coringa" "Superman" "Mulher Maravilha"]
// Batman sumiu do slice original também!
```

Isso também aparece de forma sutil ao usar `append` sobre uma fatia. Se a capacidade ainda comportar o novo elemento, o `append` vai escrever no Array compartilhado, corrompendo o Slice original:

```go
nums := []int{10, 20, 30, 40, 50}

// ⚠️ JEITO PROBLEMÁTICO
resultado := append(nums[:2], nums[3:]...)
// nums agora está corrompido: [10 20 40 50 50]
```

A forma segura é criar um destino novo e independente com `make`:

```go
// ✅ JEITO SEGURO
resultado := make([]int, 0, len(nums)-1)
resultado = append(resultado, nums[:2]...)
resultado = append(resultado, nums[3:]...)
// nums permanece intacto: [10 20 30 40 50]
// resultado: [10 20 40 50]
```

---

## 💡 Dica do Gopher

Se você carregar um arquivo enorme na memória e quiser trabalhar apenas com os primeiros bytes, não faça:

```go
cabecalho := arquivoGigante[:100]
```

O Go vai manter o arquivo inteiro na memória enquanto `cabecalho` existir, porque `cabecalho` aponta para o mesmo Array. Use `copy()` para criar um Slice verdadeiramente independente:

```go
cabecalho := make([]byte, 100)
copy(cabecalho, arquivoGigante[:100])
// agora arquivoGigante pode ser coletado pelo GC
```

---

## 🔬 Exemplos Práticos

O código do capítulo demonstra os quatro conceitos em sequência:

```go
// 1. Array fixo
var arrayFixo [3]int
arrayFixo[0] = 10
arrayFixo[1] = 20
arrayFixo[2] = 30
fmt.Println("Array fixo:", arrayFixo) // [10 20 30]

// 2. Slice crescendo com append
var sliceDinamico []int
sliceDinamico = append(sliceDinamico, 100)
sliceDinamico = append(sliceDinamico, 200, 300, 400)
fmt.Println("Slice:", sliceDinamico) // [100 200 300 400]

// 3. Observando capacidade dobrar
numeros := make([]int, 0, 2)
numeros = append(numeros, 1, 2)      // cap=2, cheio
numeros = append(numeros, 3)         // cap dobra para 4
fmt.Printf("len=%d, cap=%d\n", len(numeros), cap(numeros)) // len=3, cap=4

// 4. Aliasing em ação
original := []string{"Batman", "Superman", "Mulher Maravilha"}
copia := original[0:3]
copia[0] = "Coringa"
fmt.Println(original) // [Coringa Superman Mulher Maravilha]
```

---

## 🛠️ Exercícios Propostos

1. **A Lista de Convidados (Fácil):** crie um Slice vazio e popule-o com 5 nomes usando `append` dentro de um loop. Imprima a lista e o tamanho.

2. **O Removedor de Itens (Médio):** dado `nums := []int{10, 20, 30, 40, 50}`, remova o elemento de índice 2 (`30`) usando fatiamento e `append`. Atenção à armadilha de aliasing — use `make` para garantir que `nums` não seja corrompido.

3. **O Matador de Duplicatas (Desafio):** dado `licao := []int{2, 5, 2, 8, 5, 9, 2}`, produza `[2 5 8 9]` usando dois loops aninhados. Dica: uma variável booleana "bandeirinha" e `break` são seus aliados aqui.

---

## 🦫 Conclusão

Arrays e Slices parecem simples na superfície, mas o comportamento de memória compartilhada entre Slices fatiados é uma das fontes mais comuns de bugs silenciosos em Go. A regra prática para lembrar: **sempre que precisar de uma cópia independente de parte de um Slice, use `make` + `append` — nunca confie que fatiar é copiar**.

No próximo capítulo, vamos conhecer os **Maps** — a estrutura que vai substituir aquele loop aninhado O(n²) do exercício 3 por uma busca O(1). Até lá, os códigos completos e as notas de estudo estão no repositório:

👉 [github.com/jorgegabrielti/go-in-practice-book](https://github.com/jorgegabrielti/go-in-practice-book)

Se tiver dúvidas ou sugestões, deixa nos comentários. Bora codar! 🐹

---

**Checklist antes de publicar:**
- [x] Conteúdo confere com `estudos-go/capitulo07/fonte.txt` (sem invenções ou lacunas)
- [x] Texto é original — nenhuma frase/analogia copiada ou parafraseada de perto do `fonte.txt`
- [x] Título com analogia clara e chamativa
- [x] Pelo menos um bloco de código testado e funcional
- [x] Link para o repositório no GitHub incluído
- [ ] Revisão ortográfica feita
- [ ] Capa/imagem de destaque escolhida no Medium
- [ ] Tags do Medium definidas (ex: Golang, Programming, Software Development)
- [ ] Status atualizado em `conteudo/PAINEL.md`
