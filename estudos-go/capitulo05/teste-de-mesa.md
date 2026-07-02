# 🧮 Teste de Mesa — Capítulo 05: Laços de Repetição (for)

> Simulação manual da execução de cada `main.go` do capítulo, linha a linha, sem rodar o código. Depois de prever, confirme com `go run` (⚠️ exceto o `exemplos/ex03`, que é um loop infinito proposital — veja observação).

---

## `exemplos/ex01/main.go`

```go
func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("Volta: nº", i)
	}
}
```

| Iteração | `i` | Condição `i < 5` | Saída |
| :---: | :---: | :---: | :--- |
| 1 | 0 | true | `Volta: nº 0` |
| 2 | 1 | true | `Volta: nº 1` |
| 3 | 2 | true | `Volta: nº 2` |
| 4 | 3 | true | `Volta: nº 3` |
| 5 | 4 | true | `Volta: nº 4` |
| — | 5 | **false** | loop termina |

**Saída final:**
```
Volta: nº 0
Volta: nº 1
Volta: nº 2
Volta: nº 3
Volta: nº 4
```
**Observações:** `for` clássico padrão. Sem armadilhas.

---

## `exemplos/ex02/main.go`

```go
func main() {
	energia := 100
	for energia > 0 {
		fmt.Println("Corendo... Energia: ", energia)
		energia = energia - 20
	}
	fmt.Println("Ufa! Cansei.")
}
```

| Iteração | `energia` (antes) | Condição `energia > 0` | Saída | `energia` (depois) |
| :---: | :---: | :---: | :--- | :---: |
| 1 | 100 | true | `Corendo... Energia:  100` | 80 |
| 2 | 80 | true | `Corendo... Energia:  80` | 60 |
| 3 | 60 | true | `Corendo... Energia:  60` | 40 |
| 4 | 40 | true | `Corendo... Energia:  40` | 20 |
| 5 | 20 | true | `Corendo... Energia:  20` | 0 |
| — | 0 | **false** | loop termina | — |

**Saída final:**
```
Corendo... Energia:  100
Corendo... Energia:  80
Corendo... Energia:  60
Corendo... Energia:  40
Corendo... Energia:  20
Ufa! Cansei.
```
**Observações:** `for` no estilo `while` (só a condição, sem init/pós). Sem armadilhas (apesar do erro de digitação "Corendo" em vez de "Correndo" no texto).

---

## `exemplos/ex03/main.go`

```go
func main() {
	for {
		fmt.Println("Estou rodando para sempre...")
		//break
	}
}
```

| Iteração | Condição | Saída |
| :---: | :---: | :--- |
| 1, 2, 3, ... | sempre `true` (loop infinito, sem `break`) | `Estou rodando para sempre...` repetido infinitamente |

**Saída final:** o programa nunca termina sozinho — imprime `Estou rodando para sempre...` infinitamente até ser interrompido manualmente (Ctrl+C no terminal).

**Observações:** ⚠️ **não execute este arquivo com `go run` sem saber que ele trava o terminal** — é proposital, para demonstrar `for {}` como loop infinito. O `break` está comentado de propósito; se descomentado, o loop imprimiria a linha **uma única vez** e sairia imediatamente.

---

## `exemplos/ex04/main.go`

```go
func main() {
	for i := 0; i < 1000; i++ {
		if i == 5 {
			fmt.Println("Achei o número 5! Parando a busca.")
			break
		}
	}
}
```

| Iteração | `i` | `i == 5`? | Saída |
| :---: | :---: | :---: | :--- |
| 1-5 | 0,1,2,3,4 | false | — |
| 6 | 5 | **true** | `Achei o número 5! Parando a busca.` → `break` |

**Saída final:**
```
Achei o número 5! Parando a busca.
```
**Observações:** o `break` interrompe o loop antes de chegar a `999` — exatamente como o comentário do código documenta. Sem armadilhas.

---

## `exemplos/ex05/main.go`

```go
func main() {
	for i := 0; i < 5; i++ {
		if i == 2 {
			fmt.Println("Pulando o 2 (número do azar)...")
			continue
		}
		fmt.Println("Processando número: ", i)
	}
}
```

| Iteração | `i` | `i == 2`? | Saída |
| :---: | :---: | :---: | :--- |
| 1 | 0 | false | `Processando número:  0` |
| 2 | 1 | false | `Processando número:  1` |
| 3 | 2 | **true** | `Pulando o 2 (número do azar)...` → `continue` (pula o `Println` de baixo) |
| 4 | 3 | false | `Processando número:  3` |
| 5 | 4 | false | `Processando número:  4` |

**Saída final:**
```
Processando número:  0
Processando número:  1
Pulando o 2 (número do azar)...
Processando número:  3
Processando número:  4
```
**Observações:** `continue` pula só a iteração atual, sem interromper o loop. Sem armadilhas.

---

## `exemplos/ex06/main.go`

```go
func main() {
	nomes := []string{"Ana", "Beto", "Carla"}
	for _, nome := range nomes {
		fmt.Println(nome)
	}
}
```

| Iteração | índice (`_`, descartado) | `nome` | Saída |
| :---: | :---: | :---: | :--- |
| 1 | 0 | "Ana" | `Ana` |
| 2 | 1 | "Beto" | `Beto` |
| 3 | 2 | "Carla" | `Carla` |

**Saída final:**
```
Ana
Beto
Carla
```
**Observações:** uso do blank identifier `_` para descartar o índice que o `range` retorna mas o código não usa. Sem armadilhas.

---

## `exemplos/ex07/main.go`

```go
func main() {
	fmt.Println("--- Contagem Regressiva ---")
	for i := 5; i > 0; i-- {
		fmt.Println(i)
		time.Sleep(500 * time.Millisecond)
	}
	fmt.Println("DECOLAR! ->")

	fmt.Println("\n --- Somando até passar de 100 ---")
	soma := 0
	numero := 1
	for soma < 100 {
		soma = soma + numero
		fmt.Printf("Somando %d... Total: %d\n", numero, soma)
		numero++
	}

	fmt.Println("\n--- Tentando acertar a senha ---")
	tentativa := 0
	for {
		tentativa++
		if tentativa == 4 {
			fmt.Println("Senha correta encontrada na tentativa", tentativa)
			break
		}
		fmt.Println("Tentativa", tentativa, "falhou...")
	}

	fmt.Println("\n --- Apenas ímpares ---")
	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Print(i, " ")
	}
	fmt.Println("\nFim!")
}
```

**Bloco 1 — Contagem regressiva:**

| Iteração | `i` | Condição `i > 0` | Saída |
| :---: | :---: | :---: | :--- |
| 1 | 5 | true | `5` (+ pausa de 500ms) |
| 2 | 4 | true | `4` |
| 3 | 3 | true | `3` |
| 4 | 2 | true | `2` |
| 5 | 1 | true | `1` |
| — | 0 | **false** | loop termina |

Depois do loop: `DECOLAR! ->`

**Bloco 2 — Soma até passar de 100:**

| Iteração | `numero` (antes) | `soma` (antes) | `soma` (depois) | Saída |
| :---: | :---: | :---: | :---: | :--- |
| 1 | 1 | 0 | 1 | `Somando 1... Total: 1` |
| 2 | 2 | 1 | 3 | `Somando 2... Total: 3` |
| 3 | 3 | 3 | 6 | `Somando 3... Total: 6` |
| 4 | 4 | 6 | 10 | `Somando 4... Total: 10` |
| 5 | 5 | 10 | 15 | `Somando 5... Total: 15` |
| 6 | 6 | 15 | 21 | `Somando 6... Total: 21` |
| 7 | 7 | 21 | 28 | `Somando 7... Total: 28` |
| 8 | 8 | 28 | 36 | `Somando 8... Total: 36` |
| 9 | 9 | 36 | 45 | `Somando 9... Total: 45` |
| 10 | 10 | 45 | 55 | `Somando 10... Total: 55` |
| 11 | 11 | 55 | 66 | `Somando 11... Total: 66` |
| 12 | 12 | 66 | 78 | `Somando 12... Total: 78` |
| 13 | 13 | 78 | 91 | `Somando 13... Total: 91` |
| 14 | 14 | 91 | 105 | `Somando 14... Total: 105` |
| — | 15 | 105 | — | `105 < 100` é falso → loop termina |

(soma é a soma triangular 1+2+...+14 = 105, primeira vez que passa de 100)

**Bloco 3 — Tentando acertar a senha:**

| Iteração | `tentativa` | `tentativa == 4`? | Saída |
| :---: | :---: | :---: | :--- |
| 1 | 1 | false | `Tentativa 1 falhou...` |
| 2 | 2 | false | `Tentativa 2 falhou...` |
| 3 | 3 | false | `Tentativa 3 falhou...` |
| 4 | 4 | **true** | `Senha correta encontrada na tentativa 4` → `break` |

**Bloco 4 — Apenas ímpares:**

| Iteração | `i` | `i%2==0`? | Saída (`Print`, sem quebra de linha) |
| :---: | :---: | :---: | :--- |
| 0 | 0 | true | (continue) |
| 1 | 1 | false | `1 ` |
| 2 | 2 | true | (continue) |
| 3 | 3 | false | `3 ` |
| 4 | 4 | true | (continue) |
| 5 | 5 | false | `5 ` |
| 6 | 6 | true | (continue) |
| 7 | 7 | false | `7 ` |
| 8 | 8 | true | (continue) |
| 9 | 9 | false | `9 ` |
| 10 | 10 | true | (continue) |

Depois do loop: `\nFim!`

**Saída final no terminal:**
```
--- Contagem Regressiva ---
5
4
3
2
1
DECOLAR! ->

 --- Somando até passar de 100 ---
Somando 1... Total: 1
Somando 2... Total: 3
Somando 3... Total: 6
Somando 4... Total: 10
Somando 5... Total: 15
Somando 6... Total: 21
Somando 7... Total: 28
Somando 8... Total: 36
Somando 9... Total: 45
Somando 10... Total: 55
Somando 11... Total: 66
Somando 12... Total: 78
Somando 13... Total: 91
Somando 14... Total: 105

--- Tentando acertar a senha ---
Tentativa 1 falhou...
Tentativa 2 falhou...
Tentativa 3 falhou...
Senha correta encontrada na tentativa 4

 --- Apenas ímpares ---
1 3 5 7 9 
Fim!
```
**Observações:** a contagem regressiva usa `time.Sleep`, então a execução real demora ~2,5s (5 × 500ms), mas o conteúdo impresso é o mesmo independente do tempo. Sem bugs — código consolida bem os quatro padrões de `for` do capítulo.

---

## `exercicios/ex01/main.go`

```go
func main() {
	numero := 7
	for i := 1; i <= 10; i++ {
		resultado := numero * i
		fmt.Printf("%d x %d = %d\n", numero, i, resultado)
	}
}
```

| Iteração | `i` | `resultado` | Saída |
| :---: | :---: | :---: | :--- |
| 1 | 1 | 7 | `7 x 1 = 7` |
| 2 | 2 | 14 | `7 x 2 = 14` |
| 3 | 3 | 21 | `7 x 3 = 21` |
| 4 | 4 | 28 | `7 x 4 = 28` |
| 5 | 5 | 35 | `7 x 5 = 35` |
| 6 | 6 | 42 | `7 x 6 = 42` |
| 7 | 7 | 49 | `7 x 7 = 49` |
| 8 | 8 | 56 | `7 x 8 = 56` |
| 9 | 9 | 63 | `7 x 9 = 63` |
| 10 | 10 | 70 | `7 x 10 = 70` |

**Saída final:** tabuada do 7 de 1 a 10, exatamente como o enunciado pede. Sem bugs.

---

## `exercicios/ex02/main.go`

```go
func main() {
	for i := 1; i <= 50; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
```

| `i` | `i%3==0 && i%5==0` | `i%3==0` | `i%5==0` | Saída |
| :---: | :---: | :---: | :---: | :--- |
| 1-2 | false | false | false | `1`, `2` |
| 3 | false | true | false | `Fizz` |
| 4 | false | false | false | `4` |
| 5 | false | false | true | `Buzz` |
| 6 | false | true | false | `Fizz` |
| ... | ... | ... | ... | (padrão se repete) |
| 15 | **true** | true | true | `FizzBuzz` |
| ... | ... | ... | ... | ... |
| 50 | false | false | true | `Buzz` |

**Saída final (primeiros e últimos números, padrão clássico do FizzBuzz):**
```
1
2
Fizz
4
Buzz
Fizz
7
8
Fizz
Buzz
11
Fizz
13
14
FizzBuzz
...
49
Buzz
```
**Observações:** a ordem das condições está correta — `FizzBuzz` (múltiplo de 3 **e** 5) é verificado **antes** de `Fizz` e `Buzz` isolados, exatamente como o comentário do exercício avisa que é necessário. Sem bugs.

---

## `exercicios/ex03/main.go`

```go
func main() {
	numero := 12345
	soma := 0
	for numero > 0 {
		soma += numero % 10
		numero /= 10
	}
	fmt.Println("Soma dos digitos: ", soma)
}
```

| Iteração | `numero` (antes) | `numero % 10` | `soma` (depois) | `numero` (depois, `/10`) |
| :---: | :---: | :---: | :---: | :---: |
| 1 | 12345 | 5 | 5 | 1234 |
| 2 | 1234 | 4 | 9 | 123 |
| 3 | 123 | 3 | 12 | 12 |
| 4 | 12 | 2 | 14 | 1 |
| 5 | 1 | 1 | 15 | 0 |
| — | 0 | — | — | `0 > 0` é falso → loop termina |

**Saída final:**
```
Soma dos digitos:  15
```
**Observações:** confirma exatamente o resultado esperado pelo enunciado (`1+2+3+4+5=15`). Sem bugs.

---

## ✅ Conclusão

Nenhum bug encontrado no capítulo 05. Único ponto de atenção prático: `exemplos/ex03` é um loop infinito proposital (`for {}` sem `break`) — não deve ser executado sem saber disso, pois trava o terminal até interrupção manual (Ctrl+C).
