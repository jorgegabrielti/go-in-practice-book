# 📖 Capítulo 02: Variáveis, Constantes e Tipos de Dados

> **Livro Go na Prática: 30 capítulos para dominar a linguagem**

> 🧮 [Teste de mesa de todo o código deste capítulo](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo02/teste-de-mesa.md)

---

## 📦 Variáveis em Go

Em Go, variáveis são recipientes com tipos fortemente definidos para armazenar dados. A linguagem oferece flexibilidade na declaração, mas mantém a segurança de tipos em tempo de compilação.

### 1. Declaração Padrão (`var`)
Usada principalmente para declarar variáveis sem valor inicial (onde assume o *Zero Value*) ou no escopo global (nível de pacote).
```go
var totalDeProdutos int = 50
var estoqueReservado int // Inicializa com o "Zero Value" (0)
```

### 2. Declaração Curta (`:=`)
A forma mais comum e idiomática de declarar variáveis **dentro de funções**. O Go infere o tipo automaticamente com base no valor atribuído.
```go
produto := "Notebook Gamer" // Tipo inferido: string
preco := 4500.99            // Tipo inferido: float64
```
> [!IMPORTANT]
> O operador `:=` **não pode** ser utilizado fora de funções (no escopo global/do pacote). Fora de funções, utilize sempre `var`.

### 3. Declaração Agrupada
Permite declarar múltiplas variáveis em um único bloco para melhor legibilidade:
```go
var (
    nomeComprador  = "João Silva"
    emailComprador = "joao@email.com"
    ativo          = true
)
```

---

## 🌌 O Mistério do *Zero Value*

Quando você declara uma variável em Go sem atribuir um valor inicial, o Go automaticamente atribui a ela um valor padrão seguro, chamado de **Zero Value**. Diferente de outras linguagens, em Go não existe o perigo de ler lixo de memória ou receber um erro de ponteiro nulo para tipos primitivos.

| Tipo de Dado | Zero Value |
| :--- | :---: |
| `int`, `float64`, etc. | `0` / `0.0` |
| `string` | `""` (vazia) |
| `bool` | `false` |
| Ponteiros, interfaces, maps, slices | `nil` |

---

## 🔒 Constantes (`const`)

Constantes são valores imutáveis definidos em tempo de compilação.
```go
const PontoDeEbulicao = 100
```
*   Não podem ser alteradas depois de declaradas.
*   Não podem usar o operador curto `:=`.

---

## 🛠️ Exercícios Resolvidos

### [Exercício 01: Declaração de Variáveis](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo02/exercicios/ex01/main.go)
*   **Objetivo:** Criar um programa que declare três variáveis para armazenar dados pessoais (`nome`, `idade`, `peso`) e exibi-los formatados.
*   **Conceito Praticado:** Declaração curta de variáveis dentro da função `main` e uso do pacote `fmt`.

### [Exercício 02: Conversor de Temperatura](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo02/exercicios/ex02/main.go)
*   **Objetivo:** Converter o ponto de ebulição da água de Celsius para Fahrenheit utilizando a fórmula $F = (C \times 1.8) + 32$.
*   **Conceito Praticado:** Declaração de constantes com `const`, expressões aritméticas e mistura de tipos em operações básicas.

### [Exercício 03: O Mistério do Zero (Desafio)](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo02/exercicios/ex03/main.go)
*   **Objetivo:** Declarar variáveis sem valor inicial e inspecionar seus *Zero Values*. O desafio extra exigia o uso do verbo `%q` para exibir as aspas da string vazia explicitamente.
*   **Conceito Praticado:** *Zero values*, formatação avançada com `fmt.Printf`.
    *   `%d`: Formata inteiros (base 10).
    *   `%q`: Coloca aspas em strings (útil para debugar strings vazias ou com espaços em branco).
    *   `%t`: Formata booleanos (`true` ou `false`).
