# 🧮 Teste de Mesa — Capítulo 06: Funções

> Simulação manual da execução de cada `main.go` do capítulo, linha a linha, sem rodar o código. Depois de prever, confirme com `go run`.

---

## `exemplos/ex01/main.go`

```go
func boasVindas(nome string) {
	fmt.Printf("Olá, %s! Bem-vindo ao sistema.\n", nome)
}

func geometria(largura, altura float64) (float64, float64) {
	area := largura * altura
	perimetro := 2 * (largura + altura)
	return area, perimetro
}

func dividir(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("Proibido dividor por zero")
	}
	return a / b, nil
}

func calcularMedia(notas ...float64) float64 {
	total := 0.0
	if len(notas) == 0 {
		return 0
	}
	for _, nota := range notas {
		total += nota
	}
	return total / float64(len(notas))
}

func main() {
	boasVindas("Jorge Gabriel")

	a, p := geometria(5.0, 3.0)
	fmt.Printf("Retângulo 5x3 -> ÁREA: %.2f, Perímetro: %.2f\n", a, p)

	resultado, err := dividir(10, 2)
	if err != nil {
		fmt.Println("Erro: ", err)
	} else {
		fmt.Println("10/2 =", resultado)
	}

	_, err2 := dividir(10, 0)
	if err2 != nil {
		fmt.Println("Tentativa de dividir por 10/0 falhou com sucesso: ", err2)
	}

	mediaTurma := calcularMedia(8.5, 9.0, 5.5, 10.0)
	fmt.Printf("Media da turma: %.2f\n", mediaTurma)
}
```

| Linha | Instrução | Estado das variáveis | Saída produzida |
| :---: | :--- | :--- | :--- |
| `boasVindas("Jorge Gabriel")` | chama função simples | `nome = "Jorge Gabriel"` (parâmetro local) | `Olá, Jorge Gabriel! Bem-vindo ao sistema.` |
| `a, p := geometria(5.0, 3.0)` | `area = 5*3 = 15`, `perimetro = 2*(5+3) = 16` | `a = 15`, `p = 16` | — |
| `fmt.Printf(...)` | — | — | `Retângulo 5x3 -> ÁREA: 15.00, Perímetro: 16.00` |
| `resultado, err := dividir(10, 2)` | `b != 0` → retorna `(5, nil)` | `resultado = 5`, `err = nil` | — |
| `if err != nil ... else` | `err == nil` → ramo `else` | — | `10/2 = 5` |
| `_, err2 := dividir(10, 0)` | `b == 0` → retorna `(0, erro)` | `err2 = "Proibido dividor por zero"` | — |
| `if err2 != nil` | `true` | — | `Tentativa de dividir por 10/0 falhou com sucesso:  Proibido dividor por zero` |
| `mediaTurma := calcularMedia(8.5, 9.0, 5.5, 10.0)` | variádica recebe `notas = [8.5, 9.0, 5.5, 10.0]`; `total = 33.0`; `33.0/4 = 8.25` | `mediaTurma = 8.25` | — |
| `fmt.Printf(...)` | — | — | `Media da turma: 8.25` |

**Saída final no terminal:**
```
Olá, Jorge Gabriel! Bem-vindo ao sistema.
Retângulo 5x3 -> ÁREA: 15.00, Perímetro: 16.00
10/2 = 5
Tentativa de dividir por 10/0 falhou com sucesso:  Proibido dividor por zero
Media da turma: 8.25
```
**Observações:** demonstra função simples, múltiplos retornos, tratamento de erro idiomático (`valor, err :=`) e função variádica. A mensagem de erro `"Proibido dividor por zero"` tem um pequeno erro de digitação ("dividor" em vez de "dividir"), mas não afeta a execução — só a qualidade do texto. Sem bugs de lógica.

---

## `exercicios/ex01/main.go`

```go
func soma(a, b int) (int) {
	return a + b
}

func subtracao(a, b int) (int) {
	return a - b
}

func multiplicacao(a, b int) (int) {
	return a * b
}

func main() {
	resultadoSoma := soma(10, 5)
	fmt.Printf("Soma %d + %d = %d\n", 10, 5, resultadoSoma)

	resultadoSubtracao := subtracao(10, 5)
	fmt.Printf("Subtracao %d - %d = %d\n", 10, 5, resultadoSubtracao)

	resultadoMultiplicacao := multiplicacao(10, 5)
	fmt.Printf("Multiplicacao %d * %d = %d\n", 10, 5, resultadoMultiplicacao)
}
```

| Linha | Instrução | Estado das variáveis | Saída produzida |
| :---: | :--- | :--- | :--- |
| `resultadoSoma := soma(10, 5)` | `10 + 5 = 15` | `resultadoSoma = 15` | — |
| `fmt.Printf(...)` | — | — | `Soma 10 + 5 = 15` |
| `resultadoSubtracao := subtracao(10, 5)` | `10 - 5 = 5` | `resultadoSubtracao = 5` | — |
| `fmt.Printf(...)` | — | — | `Subtracao 10 - 5 = 5` |
| `resultadoMultiplicacao := multiplicacao(10, 5)` | `10 * 5 = 50` | `resultadoMultiplicacao = 50` | — |
| `fmt.Printf(...)` | — | — | `Multiplicacao 10 * 5 = 50` |

**Saída final no terminal:**
```
Soma 10 + 5 = 15
Subtracao 10 - 5 = 5
Multiplicacao 10 * 5 = 50
```
**Observações:** sem bugs. A sintaxe `func soma(a, b int) (int)` usa parênteses redundantes no tipo de retorno (poderia ser só `int` sem parênteses) — não é erro, apenas estilo menos idiomático; `go fmt` não o altera, mas a convenção comum omite os parênteses para retorno único.

---

## `exercicios/ex02/main.go`

```go
func analisarPreco(preco float64) (bool, string) {
	if preco >= 100.0 {
		return true, "Caro"
	}
	return false, "Barato"
}

func main() {
	fmt.Println(analisarPreco(50.0))
	fmt.Println(analisarPreco(100.0))
}
```

| Linha | Instrução | Estado das variáveis | Saída produzida |
| :---: | :--- | :--- | :--- |
| `analisarPreco(50.0)` | `50.0 >= 100.0` → `false` → retorna `(false, "Barato")` | — | `false Barato` |
| `analisarPreco(100.0)` | `100.0 >= 100.0` → `true` → retorna `(true, "Caro")` | — | `true Caro` |

**Saída final no terminal:**
```
false Barato
true Caro
```
**Observações:** `Println` recebendo diretamente os dois valores de retorno da função (sem armazenar em variáveis antes) — sintaxe válida em Go porque `Println` aceita argumentos variádicos e os múltiplos retornos se espalham automaticamente quando a chamada da função é o único argumento. Sem bugs.

---

## `exercicios/ex03/main.go`

```go
func soma(a, b int) int { return a + b }
func subtracao(a, b int) int { return a - b }
func multiplicacao(a, b int) int { return a * b }

func main() {
	fmt.Println(soma(10, 5), subtracao(10, 5), multiplicacao(10, 5))
}
```

| Linha | Instrução | Estado das variáveis | Saída produzida |
| :---: | :--- | :--- | :--- |
| `soma(10, 5)` | `10+5=15` | — | — |
| `subtracao(10, 5)` | `10-5=5` | — | — |
| `multiplicacao(10, 5)` | `10*5=50` | — | — |
| `fmt.Println(15, 5, 50)` | três argumentos separados por espaço | — | `15 5 50` |

**Saída final no terminal:**
```
15 5 50
```
**Observações:** versão mais compacta (funções de uma linha) do mesmo exercício de `ex01`/conversão de tipos, agora chamando as três funções diretamente dentro do `Println`. Sem bugs.

---

## ✅ Conclusão

Nenhum bug de execução encontrado no capítulo 06. Ponto de estilo (não bug): `exercicios/ex01` usa `(int)` com parênteses no tipo de retorno, redundante mas válido. Mensagem de erro em `exemplos/ex01` tem um typo ("dividor") sem impacto funcional.
