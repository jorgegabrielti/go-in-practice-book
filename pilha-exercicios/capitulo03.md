# Pilha de Exercícios — Capítulo 03: Tipos Básicos

> Conceitos disponíveis até aqui: tudo dos capítulos 01-02, mais inteiros com/sem sinal, overflow, `float32`/`float64`, casting explícito, strings/UTF-8, `byte`, `rune` e booleanos.

#### Exercício 01 🟢

Declare uma variável `uint8` com o valor `250`, some `10` a ela e imprima o resultado. Explique em um comentário por que o resultado é diferente de `260`.

#### Exercício 02 🟢

Simule uma compra com três itens de preços `float64` diferentes e calcule o total da compra.

#### Exercício 03 🟢

Declare uma variável `float64` com o valor `7.9`, converta-a para `int` e explique em um comentário por que o resultado não é `8`.

#### Exercício 04 🟡

Usando o pacote `unsafe` e a função `unsafe.Sizeof`, imprima o tamanho em bytes de uma variável `int8`, uma `int32`, uma `int64` e uma `float64`.

#### Exercício 05 🟡

Declare uma string contendo acentos e ao menos um emoji. Imprima a diferença entre `len(str)` (contagem de bytes) e `len([]rune(str))` (contagem de caracteres).

#### Exercício 06 🟡

Usando crases (*raw string*), escreva um pequeno trecho de "JSON" como string literal de múltiplas linhas e imprima-o.

#### Exercício 07 🟡

Tente (deixe comentado) escrever `if 1 { ... }` em Go. Documente o erro de compilação e corrija o trecho usando uma comparação booleana válida (ex: `if 1 == 1 { ... }`).

#### Exercício 08 🔴

Calcule a média de três notas inteiras garantindo que o resultado seja `float64` (sem perder a parte decimal por divisão inteira).

#### Exercício 09 🔴

Declare uma variável `uint8` e tente (comentado) atribuir o valor `-1` a ela. Documente o erro de compilação correspondente.

#### Exercício 10 🔴

Declare cinco variáveis com `:=`, cada uma de um tipo diferente (`int`, `float64`, `string`, `bool`, `rune`), e use o verbo `%T` do `fmt.Printf` para descobrir e imprimir o tipo de cada uma.
