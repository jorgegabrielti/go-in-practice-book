# 🧮 Teste de Mesa — Capítulo 03: Tipos Básicos

> Simulação manual da execução de cada `main.go` do capítulo, linha a linha, sem rodar o código. Depois de prever, confirme com `go run`. **Este capítulo tem dois bugs reais de sintaxe** — destacados abaixo.

---

## `exemplos/exemplo1/main.go`

**Código:**
```go
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var pequeno uint8 = 255
	var grande int64 = 9223372036854775807

	fmt.Println("--- Inteiros ---")
	fmt.Println("Pequeno: %d (Tipo: %T)\n", pequeno, pequeno)

	fmt.Println("Grande: ", grande)

	var numero int = 10
	fmt.Println("Tamanho de 'int' neste computador: %d bytes\n", unsafe.Sizeof(numero))

	fmt.Println("\n--- Floats ---")

	var a float64 = 0.1
	var b float64 = 0.2
	var resultado = a + b

	fmt.Println("0.1 + 0.2 =", resultado)
	fmt.Println("É exatamente 0.3?", resultado == 0.3)

	fmt.Println("\n --- Strings ---")
	str := "Go é []"
	fmt.Println("String: ", str)
	fmt.Println("Tamanho em Bytes(len): ", len(str))

	runes := []rune(str)
	fmt.Println("Quantidade de caracteres (Runes):", len(runes))

	fmt.Println("\n --- Conversões ---")
	var nota1 int = 85
	var nota2 int = 90

	mediaErrada := (nota1 + nota2) / 2
	fmt.Println("Média inteira (Errada):", mediaErrada)

	mediaCert := float64(nota1+nota2) / 2.0
	fmt.Println("Média correta:", mediaCert)
}
```

| Linha | Instrução | Estado das variáveis | Saída produzida |
| :---: | :--- | :--- | :--- |
| 9 | `var pequeno uint8 = 255` | `pequeno = 255` | — |
| 10 | `var grande int64 = 9223372036854775807` | `grande = 9223372036854775807` (limite de `int64`) | — |
| 12 | `fmt.Println("--- Inteiros ---")` | — | `--- Inteiros ---` |
| 13 | `fmt.Println("Pequeno: %d (Tipo: %T)\n", pequeno, pequeno)` | **🐞 BUG**: usa `Println`, não `Printf` — `%d`/`%T` **não são interpretados**, ficam literais na string | `Pequeno: %d (Tipo: %T)` *(quebra de linha pelo `\n` literal dentro da string)* seguido de ` 255 255` na linha seguinte |
| 15 | `fmt.Println("Grande: ", grande)` | — | `Grande:  9223372036854775807` |
| 17 | `var numero int = 10` | `numero = 10` | — |
| 18 | `fmt.Println("Tamanho de 'int'...%d bytes\n", unsafe.Sizeof(numero))` | **🐞 BUG**: mesmo problema — `Println` em vez de `Printf` | `Tamanho de 'int' neste computador: %d bytes` seguido de ` 8` na linha seguinte (em arquitetura de 64 bits) |
| 20 | `fmt.Println("\n--- Floats ---")` | — | linha vazia, depois `--- Floats ---` |
| 22-23 | `var a, b float64 = 0.1, 0.2` | `a = 0.1`, `b = 0.2` | — |
| 24 | `var resultado = a + b` | `resultado = 0.30000000000000004` (erro de representação binária IEEE 754) | — |
| 26 | `fmt.Println("0.1 + 0.2 =", resultado)` | — | `0.1 + 0.2 = 0.30000000000000004` |
| 27 | `fmt.Println("É exatamente 0.3?", resultado == 0.3)` | comparação `0.30000000000000004 == 0.3` → `false` | `É exatamente 0.3? false` |
| 29 | `fmt.Println("\n --- Strings ---")` | — | linha vazia, depois ` --- Strings ---` |
| 30 | `str := "Go é []"` | `str = "Go é []"` | — |
| 31 | `fmt.Println("String: ", str)` | — | `String:  Go é []` |
| 32 | `fmt.Println("Tamanho em Bytes(len): ", len(str))` | `é` ocupa 2 bytes em UTF-8 → `len(str)` conta bytes, não caracteres | `Tamanho em Bytes(len):  8` |
| 34 | `runes := []rune(str)` | `runes` tem 7 elementos (G,o,espaço,é,espaço,[,]) | — |
| 35 | `fmt.Println("Quantidade de caracteres (Runes):", len(runes))` | — | `Quantidade de caracteres (Runes): 7` |
| 37 | `fmt.Println("\n --- Conversões ---")` | — | linha vazia, depois ` --- Conversões ---` |
| 38-39 | `var nota1, nota2 int = 85, 90` | `nota1 = 85`, `nota2 = 90` | — |
| 41 | `mediaErrada := (nota1 + nota2) / 2` | divisão inteira: `175 / 2 = 87` (trunca, perde `.5`) | — |
| 42 | `fmt.Println("Média inteira (Errada):", mediaErrada)` | — | `Média inteira (Errada): 87` |
| 44 | `mediaCert := float64(nota1+nota2) / 2.0` | `float64(175) / 2.0 = 87.5` | — |
| 45 | `fmt.Println("Média correta:", mediaCert)` | — | `Média correta: 87.5` |

**Saída final no terminal (prevista, incluindo os bugs):**
```
--- Inteiros ---
Pequeno: %d (Tipo: %T)
 255 255
Grande:  9223372036854775807
Tamanho de 'int' neste computador: %d bytes
 8

--- Floats ---
0.1 + 0.2 = 0.30000000000000004
É exatamente 0.3? false

 --- Strings ---
String:  Go é []
Tamanho em Bytes(len):  8
Quantidade de caracteres (Runes): 7

 --- Conversões ---
Média inteira (Errada): 87
Média correta: 87.5
```

**Observações — 🐞 dois bugs reais encontrados:** nas linhas 13 e 18, o código usa `fmt.Println` com uma string que contém verbos de formatação (`%d`, `%T`) como se fosse `fmt.Printf`. O `Println` **não interpreta verbos** — ele só junta os argumentos com espaço e imprime, então `%d` e `%T` aparecem literalmente no texto, e os valores (`pequeno`, `pequeno` / `unsafe.Sizeof(numero)`) acabam impressos depois, fora de contexto, em vez de substituírem os verbos. A correção é trocar `Println` por `Printf` nessas duas linhas. O restante do arquivo (floats, strings/runes, conversões) está correto e ilustra bem os conceitos do capítulo.

---

## `exercicios/ex01/main.go`

**Código:**
```go
package main

import "fmt"

var quilometragem uint8 = 250

func main() {
	fmt.Println("Kilometragem =", quilometragem+10)
}
```

| Linha | Instrução | Estado das variáveis | Saída produzida |
| :---: | :--- | :--- | :--- |
| 5 | `var quilometragem uint8 = 250` (pacote) | `quilometragem = 250` | — |
| 8 | `fmt.Println("Kilometragem =", quilometragem+10)` | `250 + 10 = 260`, mas `uint8` só vai até `255` → overflow silencioso: `260 - 256 = 4` | `Kilometragem = 4` |

**Saída final no terminal:**
```
Kilometragem = 4
```

**Observações:** comportamento esperado e já documentado no próprio comentário do exercício — `uint8` "dá a volta" (wrap-around) silenciosamente ao passar de 255, sem erro de compilação ou de execução. Não é um bug do código, é o conceito sendo demonstrado de propósito.

---

## `exercicios/ex02/main.go`

**Código:**
```go
package main

import "fmt"

var chuvaOntem float64 = 10.5
var chuvaHoje int = 10
var total float64 = chuvaOntem + float64(chuvaHoje)

func main() {
	fmt.Println("Resultados das chuvas de ontem + hoje: ", total)
}
```

| Linha | Instrução | Estado das variáveis | Saída produzida |
| :---: | :--- | :--- | :--- |
| 5 | `var chuvaOntem float64 = 10.5` | `chuvaOntem = 10.5` | — |
| 6 | `var chuvaHoje int = 10` | `chuvaHoje = 10` | — |
| 7 | `var total float64 = chuvaOntem + float64(chuvaHoje)` | cast explícito `float64(10) = 10.0`; `total = 10.5 + 10.0 = 20.5` | — |
| 10 | `fmt.Println(...)` | — | `Resultados das chuvas de ontem + hoje:  20.5` |

**Saída final no terminal:**
```
Resultados das chuvas de ontem + hoje:  20.5
```

**Observações:** cast explícito (`float64(chuvaHoje)`) feito corretamente antes da soma — exatamente o que o exercício pedia para corrigir o erro original de "somar tipos diferentes sem conversão" (que daria erro de compilação `mismatched types float64 and int`). Nenhum bug.

---

## `exercicios/ex03/main.go`

**Código:**
```go
package main

import "fmt"

func main() {
	x := 10
	y := 10.0
	z := "10"

	fmt.Printf("Valor: %v, Tipo: %T\n", x, y, z)
}
```

| Linha | Instrução | Estado das variáveis | Saída produzida |
| :---: | :--- | :--- | :--- |
| 6 | `x := 10` | `x = 10` (`int`) | — |
| 7 | `y := 10.0` | `y = 10` (`float64`) | — |
| 8 | `z := "10"` | `z = "10"` (`string`) | — |
| 10 | `fmt.Printf("Valor: %v, Tipo: %T\n", x, y, z)` | **🐞 BUG**: a string de formato tem só **2 verbos** (`%v`, `%T`) mas são passados **3 argumentos** (`x`, `y`, `z`) | `Valor: 10, Tipo: float64` + `%!(EXTRA string=10)` |

**Saída final no terminal (prevista, incluindo o bug):**
```
Valor: 10, Tipo: float64
%!(EXTRA string=10)
```

**Observações — 🐞 bug real encontrado:** o `Printf` consome os argumentos na ordem em que aparecem para preencher os verbos: `%v` recebe `x` (`10`) e `%T` recebe `y` (revelando o tipo de `y`, que é `float64`, não o de `x`). Como `z` não tem verbo correspondente, o Go anexa a marca de erro `%!(EXTRA string=10)` ao final da saída, sinalizando "argumento extra não usado". O exercício queria imprimir valor e tipo dos **três** (`x`, `y`, `z`), mas a string de formato precisaria de 3 pares de verbos, por exemplo: `"x: %v (%T)\ny: %v (%T)\nz: %v (%T)\n"`, passando os 6 argumentos correspondentes (ou três chamadas de `Printf` separadas, uma por variável).

---

## ✅ Conclusão

O capítulo 03 tem **dois bugs reais de sintaxe**, ambos no uso de `fmt`: `exemplos/exemplo1/main.go` usa `Println` em vez de `Printf` com texto contendo verbos de formatação (linhas 13 e 18), e `exercicios/ex03/main.go` passa 3 argumentos para uma string `Printf` com apenas 2 verbos, gerando a marca `%!(EXTRA ...)`. Vale a pena corrigir esses dois arquivos para que a saída realmente demonstre o que o capítulo pretende ensinar (tipos e formatação), em vez de mostrar o efeito colateral do uso incorreto de `fmt`.
