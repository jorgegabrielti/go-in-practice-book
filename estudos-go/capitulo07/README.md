# 📖 Capítulo 07: Arrays e Slices — A Lista Fixa e a Lista Mágica

> **Livro Go na Prática: 30 capítulos para dominar a linguagem**

---

> **Fluxo de trabalho deste capítulo:**
> 1. Transcrever o conteúdo do livro para `fonte.txt` (texto bruto, fiel ao material original).
> 2. Escrever este `README.md` como síntese própria a partir do `fonte.txt`.
> 3. Resolver exemplos/exercícios em `exemplos/` e `exercicios/`.
> 4. Produzir o artigo de Medium e o post de LinkedIn em `conteudo/`, **sempre baseados no `fonte.txt`** — não em memória ou suposição do que o livro disse, e **escritos com palavras próprias** (sem copiar/parafrasear de perto o `fonte.txt`).
> 5. Rastrear a execução de todo o código do capítulo em `teste-de-mesa.md`, prevendo a saída linha a linha antes de rodar `go run`.

> 🧮 [Teste de mesa de todo o código deste capítulo](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo07/teste-de-mesa.md)

---

## 🎯 Tema Central

Em Go, **Array** e **Slice** resolvem o mesmo problema — guardar uma coleção de valores do mesmo tipo — mas com filosofias opostas. O Array é rígido: seu tamanho é definido em compilação e faz parte do tipo (um `[3]int` e um `[4]int` são tipos incompatíveis). O Slice é dinâmico: cresce conforme necessário, encolhe com fatiamento e é a estrutura preferida do dia a dia. Entender os dois é essencial porque **todo Slice vive sobre um Array escondido** — e ignorar isso gera bugs sutis de memória compartilhada.

---

## 📊 Resumo dos Conceitos

### 1. Arrays — tamanho fixo, tipo imutável

A declaração exige tamanho e tipo; valores são zero-inicializados automaticamente:

```go
var notas [3]int         // [0 0 0]
vogais := [5]string{"A", "E", "I", "O", "U"}
```

O tamanho é parte do tipo: `[3]int ≠ [4]int`. Isso torna Arrays raramente úteis no código de aplicação, mas muito eficientes em memória quando o tamanho é realmente fixo.

---

### 2. Slices — janela sobre um Array

Um Slice tem três componentes internos:

| Campo | Significado |
| :--- | :--- |
| **Ponteiro** | onde o Array subjacente começa |
| **len** | quantos elementos estão visíveis agora |
| **cap** | quantos cabem antes de precisar realocar |

Formas de criar:

```go
var lista []string           // nil slice (len=0, cap=0)
numeros := make([]int, 5)    // len=5, cap=5, todos zeros
```

---

### 3. `append` — crescimento automático

```go
numeros := []int{1, 2}
numeros = append(numeros, 3)  // [1 2 3]
```

Quando a capacidade esgota, o Go aloca um Array novo (geralmente o dobro do tamanho), copia os dados e descarta o antigo. Por isso `append` **sempre deve ser reatribuído**: `slice = append(slice, valor)`.

---

### 4. Fatiamento — `slice[min:max]`

Intervalo semi-aberto `[min, max)`: inclui `min`, exclui `max`.

```go
frutas := []string{"Maçã", "Banana", "Uva", "Pera"}
lanche := frutas[1:3]  // ["Banana", "Uva"]
```

Atalhos: `frutas[:2]` (do início até 2), `frutas[2:]` (do 2 até o fim), `frutas[:]` (tudo).

---

### 5. ⚠️ Armadilha: memória compartilhada

Fatiar **não copia** dados. O novo Slice aponta para o mesmo Array do original:

```go
original := []string{"Batman", "Superman", "Mulher Maravilha"}
copia := original[:2]
copia[0] = "Coringa"
// original agora é ["Coringa" "Superman" "Mulher Maravilha"] — inesperado!
```

Para evitar isso ao combinar pedaços de um Slice, use `make` + `append`:

```go
resultado := make([]int, 0, len(nums)-1)
resultado = append(resultado, nums[:2]...)
resultado = append(resultado, nums[3:]...)
```

---

### 6. Iterando com `range`

```go
for indice, valor := range frutas {
    fmt.Println(indice, valor)
}
```

Use `_` para descartar o índice quando não precisar dele.

---

## 💡 Dica do Gopher

> **Cuidado com memory leaks ao fatiar arquivos grandes.** Se você carregar 1 GB de dados e fazer `pequeno := gigante[:5]`, o Go mantém o 1 GB inteiro na memória enquanto `pequeno` existir, porque `pequeno` aponta para o mesmo Array. Use `copy()` para criar um Slice verdadeiramente independente e liberar o original para o garbage collector.

---

## 🔬 Exemplos Práticos no Repositório

### 1. [Exemplo 01: Arrays, Slices, Capacidade e Aliasing](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo07/exemplos/ex01/main.go)
- **Conceito:** demonstra Arrays fixos, Slices dinâmicos com `append`, crescimento de capacidade com `make`, e o efeito de aliasing ao fatiar e modificar um Slice.

---

## 🔬 Exercícios Práticos Resolvidos

### 1. [Exercício 01: A Lista de Convidados](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo07/exercicios/ex01/main.go)
- **Objetivo:** criar um Slice vazio e populá-lo com 5 nomes usando `append` dentro de um loop.
- **Conceito:** `append`, `fmt.Sprintf`, `len`.

### 2. [Exercício 02: O Removedor de Itens](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo07/exercicios/ex02/main.go)
- **Objetivo:** remover o elemento de índice 2 de `[10 20 30 40 50]` sem usar função nativa de remoção.
- **Conceito:** fatiamento + `append` + armadilha de aliasing (e como evitá-la com `make`).

### 3. [Exercício 03: O Matador de Duplicatas](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo07/exercicios/ex03/main.go)
- **Objetivo:** dado `[2 5 2 8 5 9 2]`, produzir `[2 5 8 9]` — apenas valores únicos.
- **Conceito:** loops aninhados, variável "bandeirinha" (`jaExiste`), `break` para otimização.

---

## ✅ Checklist antes de marcar como concluído

- [x] `fonte.txt` com a transcrição completa do capítulo
- [x] Teoria revisada e resumida neste README (com base no `fonte.txt`)
- [x] Todos os exemplos do capítulo têm pasta própria em `exemplos/ex01`, `ex02`... (dois dígitos)
- [x] Todos os exercícios resolvidos têm pasta própria em `exercicios/ex01`, `ex02`... (dois dígitos)
- [ ] `go build ./...` e `go vet ./...` passam sem erros
- [ ] `go fmt ./...` executado
- [x] Termos novos adicionados ao `estudos-go/GLOSSARIO.md`
- [x] Painel de progresso no `README.md` raiz atualizado
- [x] Artigo de Medium criado em `conteudo/medium/capitulo07.md` (a partir do `_template.md`, baseado no `fonte.txt`, com texto original)
- [x] Post de LinkedIn criado em `conteudo/linkedin/capitulo07.md` (a partir do `_template.md`, baseado no `fonte.txt`, com texto original)
- [ ] Status atualizado em `conteudo/PAINEL.md`
- [ ] Pilha de exercícios extras (mín. 10, dificuldade crescente) criada em `pilha-exercicios/capitulo07.md`
- [x] Teste de mesa de todo o código do capítulo criado em `teste-de-mesa.md`
