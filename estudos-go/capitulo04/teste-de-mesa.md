# 🧮 Teste de Mesa — Capítulo 04: Controle de Fluxo

> Simulação manual da execução de cada `main.go` do capítulo, linha a linha, sem rodar o código. Depois de prever, confirme com `go run`.

---

## `exemplos/ex01/main.go`

```go
func main() {
	idade := 16
	if idade >= 18 {
		fmt.Println("Pode entrar. Divirta-se!")
	} else {
		fmt.Println("Desculpe, volte ano que vem")
	}
}
```

| Linha | Instrução | Estado das variáveis | Saída produzida |
| :---: | :--- | :--- | :--- |
| `idade := 16` | declaração | `idade = 16` | — |
| `if idade >= 18` | `16 >= 18` → `false` | — | — |
| `else` | ramo executado | — | `Desculpe, volte ano que vem` |

**Saída final:**
```
Desculpe, volte ano que vem
```
**Observações:** sem armadilhas.

---

## `exemplos/ex02/main.go`

```go
func main() {
	temperatura := 25
	if temperatura > 30 {
		fmt.Println("Ligando resfriamento máximo!")
	} else if temperatura > 20 {
		fmt.Println("Temperatura agradável")
	} else {
		fmt.Println("Ligando aquecedor")
	}
}
```

| Linha | Instrução | Estado das variáveis | Saída produzida |
| :---: | :--- | :--- | :--- |
| `temperatura := 25` | declaração | `temperatura = 25` | — |
| `if temperatura > 30` | `25 > 30` → `false` | — | — |
| `else if temperatura > 20` | `25 > 20` → `true` | — | `Temperatura agradável` |

**Saída final:**
```
Temperatura agradável
```
**Observações:** sem armadilhas.

---

## `exemplos/ex03/main.go`

```go
func main() {
	carteira := true
	sobrio := false

	if carteira && sobrio {
		fmt.Println("Pode dirigir")
	} else if carteira && !sobrio {
		fmt.Println("Você tem carteira, mas bebeu. Chame um Uber!")
	} else {
		fmt.Println("Você nem deveria estar perto do volante.")
	}
}
```

| Linha | Instrução | Estado das variáveis | Saída produzida |
| :---: | :--- | :--- | :--- |
| declarações | — | `carteira = true`, `sobrio = false` | — |
| `if carteira && sobrio` | `true && false` → `false` | — | — |
| `else if carteira && !sobrio` | `true && !false` = `true && true` → `true` | — | `Você tem carteira, mas bebeu. Chame um Uber!` |

**Saída final:**
```
Você tem carteira, mas bebeu. Chame um Uber!
```
**Observações:** boa demonstração de `&&` combinado com negação `!`.

---

## `exemplos/ex04/main.go`

```go
func main() {
	dia := "Sabado"
	switch dia {
	case "Segunda":
		fmt.Println("Dia de começar a semana!")
	case "Sexta":
		fmt.Println("Sextou!")
	case "Sabado", "Domingo":
		fmt.Println("Fim de semana!")
	default:
		fmt.Println("Bora trabalhar!")
	}
}
```

| Linha | Instrução | Estado das variáveis | Saída produzida |
| :---: | :--- | :--- | :--- |
| `dia := "Sabado"` | declaração | `dia = "Sabado"` | — |
| `switch dia` | compara `dia` contra cada `case` em ordem | — | — |
| `case "Sabado", "Domingo"` | `"Sabado"` bate (case com múltiplos valores) | — | `Fim de semana!` |

**Saída final:**
```
Fim de semana!
```
**Observações:** demonstra `case` com múltiplos valores separados por vírgula. Sem armadilhas.

---

## `exemplos/ex05/main.go`

```go
func main() {
	hora := 20
	switch {
	case hora < 12:
		fmt.Println("Bom dia!")
	case hora < 18:
		fmt.Println("Boa tarde!")
	default:
		fmt.Println("Boa noite!")
	}
}
```

| Linha | Instrução | Estado das variáveis | Saída produzida |
| :---: | :--- | :--- | :--- |
| `hora := 20` | declaração | `hora = 20` | — |
| `switch {` (sem expressão = `switch true`) | avalia cada `case` como condição booleana | — | — |
| `case hora < 12` | `20 < 12` → `false` | — | — |
| `case hora < 18` | `20 < 18` → `false` | — | — |
| `default` | nenhum `case` bateu | — | `Boa noite!` |

**Saída final:**
```
Boa noite!
```
**Observações:** `switch` tagless (sem expressão), equivalente a uma cadeia de `if/else if`. Sem armadilhas.

---

## `exemplos/ex06/main.go`

```go
func main() {
	numero := 42
	if numero%2 == 0 {
		fmt.Println(numero, "é Par")
	} else {
		fmt.Println(numero, "é ímpar")
	}

	if x := 10 * 2; x > 15 {
		fmt.Println("X é grande: ", x)
	}

	sabor := "Chocolate"
	switch sabor {
	case "chocolate":
		fmt.Println("Meu favorito!")
	case "baunilha":
		fmt.Println("Clássico!")
	case "morango", "flocos":
		fmt.Println("Frutado!")
	default:
		fmt.Println("Hmm, interessante.")
	}

	hoje := time.Now().Weekday()
	fmt.Println("Hoje é: ", hoje)
	switch {
	case hoje == time.Saturday || hoje == time.Sunday:
		fmt.Println("É fim de semana!")
	case hoje == time.Monday:
		fmt.Println("Força, guerreiro. É segunda!")
	default:
		fmt.Println("Bora trabalhar!")
	}

	nivel := 1
	fmt.Println("\nNíveis de Acesso.")
	switch nivel {
	case 1:
		fmt.Println("- Acesso básico")
		fallthrough
	case 2:
		fmt.Println("- Acesso moderado")
		fallthrough
	case 3:
		fmt.Println("- Acesso total")
	default:
		fmt.Println("Nível desconhecido")
	}
}
```

| Linha | Instrução | Estado das variáveis | Saída produzida |
| :---: | :--- | :--- | :--- |
| `numero := 42` | declaração | `numero = 42` | — |
| `if numero%2 == 0` | `42 % 2 = 0` → `true` | — | `42 é Par` |
| `if x := 10 * 2; x > 15` | inicialização curta: `x = 20` (escopo só do `if`); `20 > 15` → `true` | `x = 20` (local ao bloco) | `X é grande:  20` |
| `sabor := "Chocolate"` | declaração (note o **C** maiúsculo) | `sabor = "Chocolate"` | — |
| `switch sabor { case "chocolate": ... }` | `"Chocolate" != "chocolate"` (case-sensitive!) — nenhum case bate | — | `Hmm, interessante.` (cai no `default`) |
| `hoje := time.Now().Weekday()` | depende da data real de execução | `hoje` = dia da semana atual | — |
| `fmt.Println("Hoje é: ", hoje)` | — | — | `Hoje é:  <dia da semana>` (varia conforme a data de execução) |
| `switch { case hoje == Saturday \|\| Sunday ... }` | depende do dia da semana real | — | varia: `É fim de semana!`, `Força, guerreiro. É segunda!` ou `Bora trabalhar!` |
| `nivel := 1` | declaração | `nivel = 1` | — |
| `fmt.Println("\nNíveis de Acesso.")` | — | — | linha vazia, depois `Níveis de Acesso.` |
| `switch nivel { case 1: ...fallthrough }` | `case 1` bate, imprime e cai (fallthrough) para `case 2` **sem testar a condição** | — | `- Acesso básico` |
| `case 2: ...fallthrough` | executado por fallthrough, não por match | — | `- Acesso moderado` |
| `case 3:` | executado por fallthrough, não por match (`nivel` continua `1`) | — | `- Acesso total` |
| (sem fallthrough no `case 3`) | switch termina, `default` **não** executa | — | — |

**Saída final no terminal (a parte de data/hora varia conforme o dia em que o `go run` for executado):**
```
42 é Par
X é grande:  20
Hmm, interessante.
Hoje é:  <dia da semana atual>
<mensagem correspondente ao dia>

Níveis de Acesso.
- Acesso básico
- Acesso moderado
- Acesso total
```

**Observações:** ponto interessante (não é bug, é demonstração proposital): `sabor := "Chocolate"` (C maiúsculo) nunca bate com `case "chocolate"` (minúsculo) porque comparação de string em Go é case-sensitive — cai no `default` mostrando "Hmm, interessante." em vez de "Meu favorito!". Isso pode confundir quem só lê o código rapidamente sem notar a diferença de capitalização. O `fallthrough` funciona exatamente como documentado: força a entrada no próximo `case` sem testar sua condição, e por isso os três `Println` de nível são todos impressos mesmo `nivel` sendo sempre `1`.

---

## `exercicios/ex01/main.go`

```go
var login string = "admin"
var senha string = "12345"

func main() {
	if login == "admin" && senha == "12345" {
		fmt.Println("Acesso liberado!")
	} else {
		fmt.Println("Login ou senha inválidos.")
	}
}
```

| Linha | Instrução | Estado das variáveis | Saída produzida |
| :---: | :--- | :--- | :--- |
| declarações (pacote) | — | `login = "admin"`, `senha = "12345"` | — |
| `if login == "admin" && senha == "12345"` | `true && true` → `true` | — | `Acesso liberado!` |

**Saída final:**
```
Acesso liberado!
```
**Observações:** sem armadilhas.

---

## `exercicios/ex02/main.go`

```go
var sigla string = "He"

func main() {
	switch sigla {
	case "He":
		fmt.Println("He => Hélio")
	case "O":
		fmt.Println("O => Oxigênio")
	case "Li":
		fmt.Println("Li => Lítio")
	default:
		fmt.Println("Elemento desconhecido")
	}
}
```

| Linha | Instrução | Estado das variáveis | Saída produzida |
| :---: | :--- | :--- | :--- |
| declaração (pacote) | — | `sigla = "He"` | — |
| `switch sigla { case "He": ... }` | `"He" == "He"` → bate no primeiro case | — | `He => Hélio` |

**Saída final:**
```
He => Hélio
```
**Observações:** sem armadilhas.

---

## `exercicios/ex03/main.go`

```go
func main() {
	ano := 2026
	switch {
	case ano%4 == 0:
		fmt.Println("Ano bissexto")
	default:
		fmt.Println("Ano não bissexto")
	}
}
```

| Linha | Instrução | Estado das variáveis | Saída produzida |
| :---: | :--- | :--- | :--- |
| `ano := 2026` | declaração | `ano = 2026` | — |
| `switch { case ano%4 == 0 ... }` | `2026 % 4 = 2` → `false` | — | `Ano não bissexto` |

**Saída final:**
```
Ano não bissexto
```

**Observações — ⚠️ ponto de atenção conceitual (não é bug de sintaxe, mas a regra está incompleta):** a regra real do calendário gregoriano para ano bissexto é "divisível por 4, **exceto** se divisível por 100, a menos que também seja divisível por 400" (ex: 2000 foi bissexto, mas 1900 não foi). O código só verifica `ano%4 == 0`, o que dá o resultado correto para a grande maioria dos anos (incluindo 2026, que realmente não é bissexto), mas erraria para anos como 1900 (diria "bissexto" incorretamente). Vale mencionar isso como uma melhoria possível do exercício, embora não afete o resultado para o ano testado (2026).

---

## ✅ Conclusão

Nenhum bug de sintaxe encontrado no capítulo 04. Dois pontos de atenção conceituais vale destacar: em `exemplos/ex06`, a comparação `"Chocolate" != "chocolate"` no switch é case-sensitive e cai no `default` — comportamento correto do Go, mas fácil de não notar lendo rápido; e em `exercicios/ex03`, a regra de ano bissexto está simplificada (só checa `%4`), o que funciona para o ano testado mas não cobre a exceção dos séculos não-bissextos (ex: 1900).
