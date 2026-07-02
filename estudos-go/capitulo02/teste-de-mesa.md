# 🧮 Teste de Mesa — Capítulo 02: Variáveis, Constantes e Tipos de Dados

> Simulação manual da execução de cada `main.go` do capítulo, linha a linha, sem rodar o código. Depois de prever, confirme com `go run`.

---

## `exemplos/main.go`

**Código:**
```go
package main

import "fmt"

var VersaoDoSistema = "1.0.0"

func main() {
	var totalDeProdutos int = 50
	fmt.Println("Total de Produtos(var):", totalDeProdutos)

	var estoqueReservado int
	fmt.Println("Estoque Reservado (Zero Value:)", estoqueReservado)

	produto := "Notebook Gamer"
	preco := 4500.99

	fmt.Println("Produto: ", produto, "Preço:", preco)

	preco = 4.200
	fmt.Println("Novo Preço:", preco)

	const TaxaDeEntrega = 15.00

	totalFinal := preco + TaxaDeEntrega
	fmt.Println("Total Final: ", totalFinal)

	var (
		nomeComprador  = "João Silva"
		emailComprador = "joao@email.com"

		ativo = true
	)

	fmt.Println("Cliente: ", nomeComprador, emailComprador, "Ativo?", ativo)
}
```

| Linha | Instrução | Estado das variáveis | Saída produzida |
| :---: | :--- | :--- | :--- |
| 5 | `var VersaoDoSistema = "1.0.0"` (pacote) | `VersaoDoSistema = "1.0.0"` (não usada no `main`) | — |
| 8 | `var totalDeProdutos int = 50` | `totalDeProdutos = 50` | — |
| 9 | `fmt.Println(...)` | — | `Total de Produtos(var): 50` |
| 11 | `var estoqueReservado int` | `estoqueReservado = 0` (zero value) | — |
| 12 | `fmt.Println(...)` | — | `Estoque Reservado (Zero Value:) 0` |
| 14 | `produto := "Notebook Gamer"` | `produto = "Notebook Gamer"` | — |
| 15 | `preco := 4500.99` | `preco = 4500.99` | — |
| 17 | `fmt.Println(...)` | — | `Produto:  Notebook Gamer Preço: 4500.99` |
| 19 | `preco = 4.200` | `preco = 4.2` (zero à direita não é armazenado) | — |
| 20 | `fmt.Println(...)` | — | `Novo Preço: 4.2` |
| 22 | `const TaxaDeEntrega = 15.00` | `TaxaDeEntrega = 15` (constante) | — |
| 24 | `totalFinal := preco + TaxaDeEntrega` | `totalFinal = 4.2 + 15 = 19.2` | — |
| 25 | `fmt.Println(...)` | — | `Total Final:  19.2` |
| 27-32 | bloco `var (...)` | `nomeComprador = "João Silva"`, `emailComprador = "joao@email.com"`, `ativo = true` | — |
| 34 | `fmt.Println(...)` | — | `Cliente:  João Silva joao@email.com Ativo? true` |

**Saída final no terminal:**
```
Total de Produtos(var): 50
Estoque Reservado (Zero Value:) 0
Produto:  Notebook Gamer Preço: 4500.99
Novo Preço: 4.2
Total Final:  19.2
Cliente:  João Silva joao@email.com Ativo? true
```

**Observações:** `preco = 4.200` armazena `4.2` — o `.200` é só forma de escrita do literal, não existe "zero à direita" guardado em `float64`. Nenhum bug.

---

## `exercicios/ex01/main.go`

**Código:**
```go
package main

import "fmt"

func main() {
	nome := "Jorge Gabriel"
	idade := "34"
	peso := 77.0

	fmt.Println("Meu nome é", nome, ", tenho", idade, "anos e peso", peso, "kg!")
}
```

| Linha | Instrução | Estado das variáveis | Saída produzida |
| :---: | :--- | :--- | :--- |
| 7 | `nome := "Jorge Gabriel"` | `nome = "Jorge Gabriel"` | — |
| 8 | `idade := "34"` | `idade = "34"` (**string**, não número) | — |
| 9 | `peso := 77.0` | `peso = 77` (`float64`) | — |
| 11 | `fmt.Println(...)` | — | `Meu nome é Jorge Gabriel , tenho 34 anos e peso 77 kg!` |

**Saída final no terminal:**
```
Meu nome é Jorge Gabriel , tenho 34 anos e peso 77 kg!
```

**Observações:** `peso` é `float64` mas vale `77` (inteiro exato) — o Go imprime `77`, não `77.0`, ao usar o formato padrão do `Println`. `idade` foi declarada como `string ("34")` em vez de número — funciona porque só é usada para impressão, mas se o capítulo pedisse uma conta com a idade, essa escolha de tipo causaria erro de compilação (não dá para somar `string + int` sem conversão). Não é um bug de execução, mas vale observar como ponto de atenção de modelagem de dados.

---

## `exercicios/ex02/main.go`

**Código:**
```go
package main

import "fmt"

const PontoDeEbulicao = 100

var temperaturaAtual = 100

func main() {
	fmt.Println("A água ferve a ", PontoDeEbulicao, "Cº ou", (PontoDeEbulicao*1.8 + 32), "Fº")
}
```

| Linha | Instrução | Estado das variáveis | Saída produzida |
| :---: | :--- | :--- | :--- |
| 5 | `const PontoDeEbulicao = 100` | `PontoDeEbulicao = 100` (constante não tipada) | — |
| 7 | `var temperaturaAtual = 100` (pacote, não usada no `main`) | `temperaturaAtual = 100` | — |
| 10 | `fmt.Println(...)` | expressão `PontoDeEbulicao*1.8 + 32` calculada como `float64`: `100*1.8+32 = 212` | `A água ferve a  100 Cº ou 212 Fº` |

**Saída final no terminal:**
```
A água ferve a  100 Cº ou 212 Fº
```

**Observações:** `PontoDeEbulicao` é uma constante **não tipada**; ao ser usada na expressão `PontoDeEbulicao*1.8 + 32` (que contém o literal float `1.8`), o Go promove a conta inteira para `float64` automaticamente — sem precisar de cast explícito, porque é aritmética de constantes em tempo de compilação. O resultado `212` aparece sem `.0` porque é um valor inteiro exato representado em `float64`. `temperaturaAtual` é declarada mas nunca usada dentro de `main` — isso é permitido para variáveis de **pacote** (só dá erro de compilação para variáveis locais não usadas).

---

## `exercicios/ex03/main.go`

**Código:**
```go
package main

import "fmt"

var numero int
var texto string
var boleano bool

func main() {
	fmt.Printf("%d %q %t\n", numero, texto, boleano)
}
```

| Linha | Instrução | Estado das variáveis | Saída produzida |
| :---: | :--- | :--- | :--- |
| 5-7 | declarações de pacote | `numero = 0`, `texto = ""`, `boleano = false` (zero values) | — |
| 10 | `fmt.Printf("%d %q %t\n", numero, texto, boleano)` | — | `0 "" false` |

**Saída final no terminal:**
```
0 "" false
```

**Observações:** o exercício pede para descobrir o que `%q` faz — ele imprime a string entre aspas duplas, tornando visível que o valor é uma string vazia (`""`) em vez de "nada". Sem bugs.

---

## `exercicios/ex04/main.go`

**Código:**
```go
package main

import "fmt"

var total = 10

func main() {
	total := 5
	fmt.Println(total) //Imprime 5 e não 10.

	imprimeTotal()
}

func imprimeTotal() {
	fmt.Println(total) //Imprime 10 e não 5.
}
```

| Linha | Instrução | Estado das variáveis | Saída produzida |
| :---: | :--- | :--- | :--- |
| 5 | `var total = 10` (pacote) | `total (pacote) = 10` | — |
| 8 | `total := 5` (dentro de `main`) | nova variável local `total = 5`, **sombreia** a de pacote dentro do escopo de `main` | — |
| 9 | `fmt.Println(total)` | usa o `total` local | `5` |
| 11 | `imprimeTotal()` | chama a função | — |
| 15 | `fmt.Println(total)` (dentro de `imprimeTotal`) | aqui não existe `total` local, então usa o `total` de pacote | `10` |

**Saída final no terminal:**
```
5
10
```

**Observações:** demonstra *shadowing* de variáveis — o `total := 5` dentro de `main` cria uma variável nova que só existe naquele escopo, sem alterar o `total` de pacote. Os comentários do próprio exercício já confirmam o comportamento esperado. Nenhum bug.

---

## ✅ Conclusão

Nenhum bug de execução encontrado no capítulo 02. Dois pontos de atenção conceituais vale destacar para revisão: (1) `idade` como `string` em `exercicios/ex01` (escolha de tipo que funciona aqui mas limitaria contas futuras) e (2) a promoção automática de constante inteira para `float64` em `exercicios/ex02` por causa do literal `1.8` na expressão.
