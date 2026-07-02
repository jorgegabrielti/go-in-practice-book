# 📖 Capítulo 04: Controle de Fluxo — O Guarda de Trânsito

> **Livro Go na Prática: 30 capítulos para dominar a linguagem**

> 🧮 [Teste de mesa de todo o código deste capítulo](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo04/teste-de-mesa.md)

---

## 🚦 O Guarda de Trânsito: Tomando Decisões

Se até agora aprendemos a declarar variáveis e tipos (o que permite ao nosso programa guardar dados), o **Controle de Fluxo** é o cérebro que permite ao programa tomar caminhos diferentes com base em condições. Sem controle de fluxo, um programa só sabe executar linha após linha de cima para baixo. No Go, a simplicidade impera: as estruturas de decisão são diretas e sem redundâncias.

Neste capítulo, estudamos as seguintes estruturas de decisão:
1. **O Condicional `if` / `else`** (A bifurcação clássica)
2. **`if` com Inicialização Curta** (*Short Statement*, a joia de escopo do Go)
3. **Operadores Lógicos** (`&&`, `||`, `!`) e Curto-Circuito
4. **`switch`** (O triador de trilhos, com break automático e múltiplos casos)
5. **Switch Sem Expressão** (Alternativa limpa ao `if/else if`)
6. **O Polêmico `fallthrough`** (Continuidade de casos)

---

## 📊 Resumo dos Conceitos

### 1. O Condicional `if` e `else`
Em Go, não usamos parênteses `()` ao redor da condição. A chave de abertura `{` **obrigatoriamente** deve ficar na mesma linha da condição.
```go
idade := 16
if idade >= 18 {
    fmt.Println("Acesso liberado")
} else {
    fmt.Println("Acesso negado")
}
```

### 2. A Joia do Go: `if` com Inicialização Curta
Você pode declarar uma variável temporária e testá-la na mesma linha do `if`. O escopo dessa variável fica restrito ao bloco condicional, sendo destruída logo em seguida.
```go
if nota := calcularNota(); nota < 0 {
    fmt.Println("Erro: Nota inválida")
}
// 'nota' não existe aqui fora!
```

### 3. Operadores Lógicos
*   `&&` (AND): Ambas as condições devem ser verdadeiras.
*   `||` (OR): Pelo menos uma condição deve ser verdadeira.
*   `!` (NOT): Inverte o valor lógico.
> **Curto-Circuito:** Se o Go avaliar `A && B` e `A` for falso, ele nem chega a avaliar `B`, otimizando a performance.

### 4. Switch: Sem `break` manual
No Go, o `break` é implícito no fim de cada `case`. Você também pode agrupar múltiplos valores no mesmo caso usando vírgulas.
```go
dia := "Sabado"
switch dia {
case "Segunda":
    fmt.Println("Trabalho")
case "Sabado", "Domingo":
    fmt.Println("Fim de semana!")
default:
    fmt.Println("Dia de semana normal")
}
```

### 5. Switch Sem Expressão (Tagless)
Funciona como um `switch true`. O compilador executa o primeiro `case` que avaliar como verdadeiro.
```go
switch {
case hora < 12:
    fmt.Println("Bom dia")
case hora < 18:
    fmt.Println("Boa tarde")
default:
    fmt.Println("Boa noite")
}
```

### 6. O Polêmico `fallthrough`
Usado para forçar a execução do próximo `case` sem validar sua condição (comportamento herdado do C/C++). É raramente utilizado na prática.

---

## 💡 Dica do Gopher: Evite Ninhos Profundos (*Early Return*)
Em vez de aninhar múltiplos `if`s (formato "flecha"), utilize cláusulas de guarda (*Guard Clauses*) e retorne o fluxo o mais cedo possível. Mantenha o "caminho feliz" (*happy path*) com o menor recuo horizontal possível.

---

## 🔬 Exemplos Práticos no Repositório

### 1. [Exemplo 01: O Porteiro de Balada](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo04/exemplos/ex01/main.go)
*   **Conceito:** Condicional `if` / `else` clássico sem parênteses.

### 2. [Exemplo 02: O Termostato](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo04/exemplos/ex02/main.go)
*   **Conceito:** Uso do `else if` encadeado para múltiplas faixas de temperatura.

### 3. [Exemplo 03: Habilitação e Sobriedade](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo04/exemplos/ex03/main.go)
*   **Conceito:** Lógica combinada com os operadores `&&` e `!`.

### 4. [Exemplo 04: Dias da Semana](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo04/exemplos/ex04/main.go)
*   **Conceito:** `switch` clássico comparando strings com agrupamento de casos.

### 5. [Exemplo 05: Fases do Dia](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo04/exemplos/ex05/main.go)
*   **Conceito:** `switch` sem expressão (tagless) atuando como uma cadeia de `if-else`.

### 6. [Exemplo 06: Resumo Geral e Fallthrough](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo04/exemplos/ex06/main.go)
*   **Conceito:** Arquivo unificado exercitando o `if` curto, o `switch` sem expressão e a palavra-chave `fallthrough`.

---

## 🔬 Exercícios Práticos Resolvidos

### 1. [Exercício 01: O Verificador de Senha](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo04/exercicios/ex01/main.go)
*   **Objetivo:** Validar login e senha com operadores lógicos.
*   **Conceito:** Operadores lógicos (`&&`) e desvio condicional simples.

### 2. [Exercício 02: A Tabela Periódica](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo04/exercicios/ex02/main.go)
*   **Objetivo:** Traduzir siglas de elementos químicos em seus nomes reais.
*   **Conceito:** `switch` clássico baseado em strings e comportamento `default`.

### 3. [Exercício 03: O Detetive de Ano Bissexto](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo04/exercicios/ex03/main.go)
*   **Objetivo:** Determinar se um determinado ano é bissexto.
*   **Conceito:** Expressão lógica usando divisibilidade (`ano % 4 == 0`) e desvio em `switch`.
