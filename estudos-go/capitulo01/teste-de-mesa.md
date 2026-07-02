# 🧮 Teste de Mesa — Capítulo 01: Introdução & Origens do Go

> Simulação manual da execução de cada `main.go` do capítulo, linha a linha, sem rodar o código. Depois de prever, confirme com `go run`.

---

## `exemplos/ola-mundo/main.go`

**Código:**
```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, World! I'm a Gopher := !")
}
```

| Linha | Instrução | Estado das variáveis | Saída produzida |
| :---: | :--- | :--- | :--- |
| 6 | `fmt.Println("Hello, World! I'm a Gopher := !")` | — (nenhuma variável) | `Hello, World! I'm a Gopher := !` |

**Saída final no terminal:**
```
Hello, World! I'm a Gopher := !
```

**Observações:** o `:=` dentro da string é só texto, não tem efeito de sintaxe ali — é uma "piada" do autor mostrando o operador de declaração curta dentro da própria saudação. Nenhum comportamento inesperado.

---

## `exercicios/ex01_identificador/main.go`

**Código:**
```go
package main

import "fmt"

func main() {
	fmt.Println("My name is: Jorge Gabriel")
	fmt.Println("I'm from: Belém, PA")
	fmt.Println("And, my favorite food is: Pizza!")
}
```

| Linha | Instrução | Estado das variáveis | Saída produzida |
| :---: | :--- | :--- | :--- |
| 8 | `fmt.Println("My name is: Jorge Gabriel")` | — | `My name is: Jorge Gabriel` |
| 9 | `fmt.Println("I'm from: Belém, PA")` | — | `I'm from: Belém, PA` |
| 10 | `fmt.Println("And, my favorite food is: Pizza!")` | — | `And, my favorite food is: Pizza!` |

**Saída final no terminal:**
```
My name is: Jorge Gabriel
I'm from: Belém, PA
And, my favorite food is: Pizza!
```

**Observações:** três chamadas independentes de `Println`, cada uma fecha com `\n` automático. Sem armadilhas.

---

## `exercicios/ex02_artista/main.go`

**Código:**
```go
package main

import "fmt"

func main() {
	fmt.Println("* * * * *")
	fmt.Println("*       *")
	fmt.Println("*       *")
	fmt.Println("*       *")
	fmt.Println("* * * * *")
}
```

| Linha | Instrução | Estado das variáveis | Saída produzida |
| :---: | :--- | :--- | :--- |
| 7 | `fmt.Println("* * * * *")` | — | `* * * * *` |
| 8 | `fmt.Println("*       *")` | — | `*       *` |
| 9 | `fmt.Println("*       *")` | — | `*       *` |
| 10 | `fmt.Println("*       *")` | — | `*       *` |
| 11 | `fmt.Println("* * * * *")` | — | `* * * * *` |

**Saída final no terminal:**
```
* * * * *
*       *
*       *
*       *
* * * * *
```

**Observações:** desenha um quadrado vazado de asteriscos. Sem armadilhas.

---

## `exercicios/ex03_quebrador/main.go`

**Código:**
```go
package main

import "fmt"

func main() {
	fmt.Println("Olá mundo, eu sou gopher")
}
```

| Linha | Instrução | Estado das variáveis | Saída produzida |
| :---: | :--- | :--- | :--- |
| 6 | `fmt.Println("Olá mundo, eu sou gopher")` | — | `Olá mundo, eu sou gopher` |

**Saída final no terminal:**
```
Olá mundo, eu sou gopher
```

**Observações:** este exercício pede para o estudante **provocar erros de propósito** (trocar `package main` por `package batata`, remover o `import "fmt"`, trocar `Println` por `println` minúsculo) e observar as mensagens do compilador. O arquivo no repositório está na versão "correta" (sem os erros), já que os erros são para serem testados manualmente no terminal e depois revertidos — não fazem sentido versionados. Resultado esperado de cada experimento, caso quisesse testar:
- `package batata`: erro de compilação porque o Go exige `package main` para gerar um binário executável (`go run` espera encontrar `func main()` dentro de `package main`).
- Remover `import "fmt"`: erro `undefined: fmt`.
- `fmt.println` (minúsculo): erro `undefined: fmt.println` — Go é case-sensitive e `println` (com p minúsculo) não é exportado pelo pacote `fmt` (existe uma função builtin `println` sem pacote, mas não é o que está sendo chamado aqui com o prefixo `fmt.`).

---

## ✅ Conclusão

Todo o capítulo 01 é composto apenas por chamadas diretas a `fmt.Println` com strings literais — não há variáveis, lógica condicional nem armadilhas de tipo. Nenhum bug encontrado. O único exercício "especial" (`ex03_quebrador`) é proposital para gerar erros de compilação e não tem saída de execução normal a ser rastreada além da versão corrigida.
