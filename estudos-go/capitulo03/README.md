# 📖 Capítulo 03: Tipos Básicos — A Química dos Materiais

> **Livro Go na Prática: 30 capítulos para dominar a linguagem**

> 🧮 [Teste de mesa de todo o código deste capítulo](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo03/teste-de-mesa.md)

---

## 🧪 A Química dos Materiais: Nem Tudo é Igual

Se no capítulo anterior aprendemos a criar caixas (`var`) e colar etiquetas nelas, agora precisamos entender que a **natureza da mercadoria importa**. No Go, que é uma linguagem **fortemente tipada**, entender a "química" de cada tipo é vital. O compilador não faz conversões automáticas ou mágicas. Se você tentar misturar materiais incompatíveis (como somar um inteiro com um decimal), ele impedirá a compilação.

Neste capítulo, estudamos os elementos fundamentais da tabela periódica do Go:
1. **Números Inteiros** (A Solidez dos Blocos)
2. **Números de Ponto Flutuante** (O Fluido Impreciso)
3. **Strings e Runas** (O Texto Imutável)
4. **Booleanos** (O Interruptor)

---

## 📊 Tabela Resumo dos Elementos

### 1. Números Inteiros (`int`)
Representam números sem partes fracionárias. Temos controle total sobre o espaço de memória que cada número ocupa:
*   **Com Sinal (Signed):** `int8` ($-128$ a $127$), `int16`, `int32`, `int64` e `int` (equivalente a 32 ou 64 bits dependendo da arquitetura do computador).
*   **Sem Sinal (Unsigned):** `uint8` ($0$ a $255$), `uint16`, `uint32`, `uint64`.

> [!CAUTION]
> **Overflow (Transbordamento):**
> Se tentarmos guardar um valor maior do que a capacidade do tipo (ex: colocar $256$ em um `uint8`), ocorre o overflow. O número "dá a volta" e retorna a zero, sem nenhum aviso em tempo de execução.

### 2. Números de Ponto Flutuante (`float`)
Representam valores decimais/fracionários:
*   `float32` (Precisão simples, 4 bytes)
*   `float64` (Precisão dupla, 8 bytes). **Sempre use `float64` por padrão.**

> [!WARNING]
> **A Armadilha do IEEE 754:**
> Devido à representação binária de frações decimais, floats podem apresentar pequenas imprecisões matemáticas (como $0.1 + 0.2 = 0.30000000000000004$). Por isso, **nunca use floats para lidar com dinheiro** — use inteiros para contar centavos ou pacotes de precisão arbitrária.

### 3. Strings
Cadeias de caracteres imutáveis codificadas nativamente em UTF-8:
*   **Imutabilidade:** Não podemos alterar caracteres de uma string diretamente na memória. Só podemos criar novas strings.
*   **Raw Strings:** Escritas entre crases (`` ` ``) para preservar quebras de linha e caracteres especiais literalmente.

### 4. Booleanos (`bool`)
Representa apenas `true` ou `false`. O valor padrão (*Zero Value*) é `false`. Go não converte inteiros (como `0` ou `1`) para booleanos automaticamente.

---

## 🔄 Conversão de Tipos (Casting)

Go exige conversões explícitas para operações entre tipos diferentes. A sintaxe é `TipoDesejado(Valor)`.
```go
var inteiro int = 42
var flutuante float64 = float64(inteiro) // Correto

var alto float64 = 99.99
var baixo int = int(alto) // Trunca o decimal (baixo vira 99)
```

---

## 💡 Dica do Gopher: Rune e Byte

*   **`byte`**: Apelido para `uint8`. Usado para dados brutos e binários.
*   **`rune`**: Apelido para `int32`. Representa um caractere Unicode (*code point*). É essencial porque um caractere acentuado ou emoji pode ocupar múltiplos bytes em UTF-8.

---

## 🔬 Exercícios Práticos Resolvidos

### 1. [Exercício 01: O Contador de Quilometragem](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo03/exercicios/ex01/main.go)
*   **Objetivo:** Demonstrar o transbordamento (*overflow*) de um `uint8` ao somar $10$ a uma variável iniciada em $250$.
*   **Conceito:** Overflow numérico.

### 2. [Exercício 02: A Chuva em Milímetros](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo03/exercicios/ex02/main.go)
*   **Objetivo:** Somar a chuva acumulada representada por valores de tipos incompatíveis (`float64` e `int`) utilizando casting.
*   **Conceito:** Casting explícito em Go.

### 3. [Exercício 03: O Inquisidor de Tipos](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo03/exercicios/ex03/main.go)
*   **Objetivo:** Inspecionar o valor e o tipo gerado automaticamente por inferência de tipos para diferentes declarações utilizando o verbo `%T`.
*   **Conceito:** Verbos de formatação e inferência de tipos.
