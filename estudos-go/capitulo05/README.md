# 📖 Capítulo 05: Laços de Repetição (For) — A Pista de Corrida

> **Livro Go na Prática: 30 capítulos para dominar a linguagem**

> 🧮 [Teste de mesa de todo o código deste capítulo](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo05/teste-de-mesa.md)

---

## 🏁 A Pista de Corrida: O Poder da Automação

Computadores são máquinas fantásticas por um motivo principal: eles **adoram** tarefas repetitivas. Eles podem fazer a mesma coisa bilhões de vezes sem ficarem entediados, cansados ou com LER (Lesão por Esforço Repetitivo).

Na maioria das linguagens de programação, você tem um cardápio confuso de opções para repetição: `while`, `do-while`, `for`, `for-each`, `repeat-until`. O Go olha para isso e simplifica: **existe apenas uma palavra-chave para repetição: `for`**. No entanto, esse único construto é flexível o suficiente para se comportar de todas as formas que você precisar.

Neste capítulo, estudamos as seguintes variações e controles de fluxo de repetição:
1. **O `for` Clássico** (O piloto de Fórmula 1)
2. **O `for` como `while`** (O corredor de maratona)
3. **O Loop Infinito** (O servidor eterno)
4. **Controle de Fluxo: `break` e `continue`**
5. **O `range` para varrer coleções** (com o uso de `_` para ignorar dados desnecessários)

---

## 📊 Resumo dos Conceitos

### 1. O `for` Clássico
Herdado da sintaxe clássica da linguagem C, é composto por três partes divididas por ponto e vírgula `;`:
- **Inicialização (Largada):** Onde o loop começa (`i := 0`).
- **Condição (Bandeira Quadriculada):** Enquanto for verdadeira, o loop continua rodando (`i < 5`).
- **Pós-execução (Pit Stop):** O passo que ocorre ao final de cada volta (`i++`).

```go
for i := 0; i < 5; i++ {
    fmt.Println("Volta número:", i)
}
```
*Nota: A variável `i` declarada na inicialização só existe no escopo do loop.*

### 2. O `for` como `while`
Em Go, não existe a palavra-chave `while`. Quando você precisa rodar um loop baseado apenas em uma condição (sem passo ou inicialização definidos de antemão), basta omitir as outras partes do `for`.
```go
energia := 100
for energia > 0 {
    fmt.Println("Correndo... Energia:", energia)
    energia -= 20
}
```

### 3. O Loop Infinito
Usado para serviços em segundo plano (daemons) ou processos que esperam conexões indefinidamente. É simplesmente um `for` sem nenhuma condição ou instrução inicial/final.
```go
for {
    fmt.Println("Rodando para sempre...")
    break // Paramos com o break para evitar travamento
}
```

### 4. Controlando o Fluxo: `break` e `continue`
- **`break` (O freio de emergência):** Aborta a execução do loop imediatamente, pulando para fora dele.
- **`continue` (Pular volta):** Interrompe a iteração atual imediatamente e pula para a próxima volta (volta para o topo do loop e reavalia a condição ou aplica o passo).

### 5. O `range` (Varrer Coleções)
Facilita a varredura de arrays, slices, maps, strings ou canais. Ele retorna dois valores a cada iteração: o índice (ou chave) e o valor do elemento.
```go
nomes := []string{"Ana", "Beto", "Carla"}
for i, nome := range nomes {
    fmt.Printf("Na posição %d tem: %s\n", i, nome)
}
```

---

## 💡 Dica do Gopher: Use `_` (Blank Identifier) para ignorar o que não precisa
Como o compilador do Go proíbe variáveis declaradas e não utilizadas, muitas vezes ao usar `range` você precisará apenas do valor e não do índice. Para isso, use o identificador em branco `_`:
```go
// CERTO: Jogamos o índice no lixo usando o '_'
for _, nome := range nomes {
    fmt.Println(nome)
}
```

---

## 🔬 Exemplos Práticos no Repositório

### 1. [Exemplo 01: O For Clássico](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo05/exemplos/ex01/main.go)
*   **Conceito:** Implementação clássica com inicialização, condição e incremento.

### 2. [Exemplo 02: O For Estilo While](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo05/exemplos/ex02/main.go)
*   **Conceito:** Omitindo a largada e o pit stop para controle puramente condicional.

### 3. [Exemplo 03: Loop Infinito](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo05/exemplos/ex03/main.go)
*   **Conceito:** Um loop `for` sem restrições ou termos de condição.

### 4. [Exemplo 04: O Freio de Emergência com Break](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo05/exemplos/ex04/main.go)
*   **Conceito:** Parando um loop de alta frequência assim que um determinado valor é encontrado.

### 5. [Exemplo 05: Pular Volta com Continue](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo05/exemplos/ex05/main.go)
*   **Conceito:** Pulando o processamento de itens indesejados.

### 6. [Exemplo 06: Iterando com Range](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo05/exemplos/ex06/main.go)
*   **Conceito:** Percorrendo coleções e descartando o índice usando a variável de descarte `_`.

### 7. [Exemplo 07: Demonstração Unificada](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo05/exemplos/ex07/main.go)
*   **Conceito:** Compilação geral de todos os tipos de loops em um único arquivo explicativo (contagem regressiva, acumulador, adivinhação, ímpares).

---

## 🔬 Exercícios Práticos Resolvidos

### 1. [Exercício 01: A Tabuada](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo05/exercicios/ex01/main.go)
*   **Objetivo:** Imprimir a tabuada de um número escolhido de 1 a 10.
*   **Conceito:** Loop `for` clássico e formatação com `Printf`.

### 2. [Exercício 02: O "FizzBuzz"](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo05/exercicios/ex02/main.go)
*   **Objetivo:** Substituir múltiplos de 3 por "Fizz", 5 por "Buzz" e ambos por "FizzBuzz" em um loop de 1 a 50.
*   **Conceito:** Lógica condicional e ordem correta de validação.

### 3. [Exercício 03: A Soma dos Dígitos](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo05/exercicios/ex03/main.go)
*   **Objetivo:** Somar todos os dígitos individuais de um número grande (ex: 12345 -> 15).
*   **Conceito:** Loop condicional (`for` estilo `while`), divisão inteira (`/ 10`) e resto da divisão (`% 10`).
