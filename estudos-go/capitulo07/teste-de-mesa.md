# 🧮 Teste de Mesa — Capítulo 07: Arrays e Slices

> Simulação manual da execução de cada `main.go` do capítulo, linha a linha, sem rodar o código. Depois de prever, confirme com `go run`. **Este capítulo é o motivo da criação do teste de mesa** — o exercício 02 expõe na prática o efeito de slices compartilharem o mesmo array por trás dos panos.

---

## `exemplos/ex01/main.go`

```go
func main() {
	// --- 1. Arrays (fixos) ---
	var arrayFixo [3]int
	arrayFixo[0] = 10
	arrayFixo[1] = 20
	arrayFixo[2] = 30
	fmt.Println("Array fixo: ", arrayFixo)

	vogais := [5]string{"A", "E", "I", "O", "U"}
	fmt.Printf("Vogais (Tamanho %d): %v\n", len(vogais), vogais)

	// --- 2. Slices (Dinâmicos) ---
	var sliceDinamico []int
	fmt.Println("\nSlice Inicial: ", sliceDinamico)
	fmt.Println("Vazio?", sliceDinamico == nil)

	sliceDinamico = append(sliceDinamico, 100)
	fmt.Println("Slice Crescido: ", sliceDinamico)
	sliceDinamico = append(sliceDinamico, 200, 300, 400)
	fmt.Println("Slice Crescido: ", sliceDinamico)

	// --- 3. Capacidade vs Tamanho ---
	numeros := make([]int, 0, 2)
	fmt.Printf("\nLen: %d, Cap: %d, Array: %v\n", len(numeros), cap(numeros), numeros)

	numeros = append(numeros, 1)
	numeros = append(numeros, 2)
	fmt.Printf("Len: %d, Cap: %d, Array: %v(Cheio!)\n", len(numeros), cap(numeros), numeros)

	numeros = append(numeros, 3)
	fmt.Printf("Len: %d, Cap: %d, Array: %v (Cresceu!)\n", len(numeros), cap(numeros), numeros)

	// --- 4. Slicing e Referência ---
	original := []string{"Batman", "Superman", "Mulher Maravilha"}
	copiaCopia := original[0:3]
	fmt.Println("\nOriginal: ", original)
	fmt.Println("Cópia (Slice): ", copiaCopia)

	copiaCopia[0] = "Coringa"
	fmt.Println("--- Depois de mudar a cópia ---")
	fmt.Println("Cópia: ", copiaCopia)
	fmt.Println("Original: ", original)
}
```

| Linha | Instrução | Estado das variáveis | Saída produzida |
| :---: | :--- | :--- | :--- |
| `var arrayFixo [3]int` | zero values | `arrayFixo = [0 0 0]` | — |
| `arrayFixo[0]=10; [1]=20; [2]=30` | preenchimento manual | `arrayFixo = [10 20 30]` | — |
| `fmt.Println("Array fixo: ", arrayFixo)` | — | — | `Array fixo:  [10 20 30]` |
| `vogais := [5]string{...}` | array literal | `vogais = [A E I O U]` | — |
| `fmt.Printf(...)` | `len(vogais)=5` | — | `Vogais (Tamanho 5): [A E I O U]` |
| `var sliceDinamico []int` | slice nil | `sliceDinamico = nil` (len=0, cap=0) | — |
| `fmt.Println("\nSlice Inicial: ", sliceDinamico)` | — | — | linha vazia, depois `Slice Inicial:  []` |
| `fmt.Println("Vazio?", sliceDinamico == nil)` | `nil == nil` → `true` | — | `Vazio? true` |
| `append(sliceDinamico, 100)` | array novo alocado | `sliceDinamico = [100]` | — |
| `fmt.Println(...)` | — | — | `Slice Crescido:  [100]` |
| `append(sliceDinamico, 200, 300, 400)` | append variádico com 3 valores | `sliceDinamico = [100 200 300 400]` | — |
| `fmt.Println(...)` | — | — | `Slice Crescido:  [100 200 300 400]` |
| `numeros := make([]int, 0, 2)` | len=0, cap=2 | `numeros = []` | — |
| `fmt.Printf(...)` | — | — | `Len: 0, Cap: 2, Array: []` (precedido de linha vazia) |
| `append(numeros, 1)` | cabe na capacidade (1≤2) | `numeros = [1]`, len=1, cap=2 | — |
| `append(numeros, 2)` | cabe (2≤2) | `numeros = [1 2]`, len=2, cap=2 (**cheio**) | — |
| `fmt.Printf(...)` | — | — | `Len: 2, Cap: 2, Array: [1 2](Cheio!)` |
| `append(numeros, 3)` | excede capacidade → Go realoca, geralmente **dobra** a capacidade | `numeros = [1 2 3]`, len=3, cap=4 (comportamento típico do runtime atual; valor exato de cap pode variar entre versões do Go) | — |
| `fmt.Printf(...)` | — | — | `Len: 3, Cap: 4, Array: [1 2 3] (Cresceu!)` |
| `original := []string{...}` | slice literal | `original = [Batman Superman Mulher Maravilha]` | — |
| `copiaCopia := original[0:3]` | fatia **toda** a faixa — **aponta para o mesmo array** de `original` | `copiaCopia = [Batman Superman Mulher Maravilha]` (mesmo array subjacente) | — |
| `fmt.Println(...)` (Original, Cópia) | — | — | `Original:  [Batman Superman Mulher Maravilha]` e `Cópia (Slice):  [Batman Superman Mulher Maravilha]` |
| `copiaCopia[0] = "Coringa"` | escreve na posição 0 do **array compartilhado** | `copiaCopia = [Coringa Superman Mulher Maravilha]`, **e `original` também muda**, pois é o mesmo array por trás | — |
| `fmt.Println(...)` (Depois) | — | — | `--- Depois de mudar a cópia ---`, `Cópia:  [Coringa Superman Mulher Maravilha]`, `Original:  [Coringa Superman Mulher Maravilha]` |

**Saída final no terminal:**
```
Array fixo:  [10 20 30]
Vogais (Tamanho 5): [A E I O U]

Slice Inicial:  []
Vazio? true
Slice Crescido:  [100]
Slice Crescido:  [100 200 300 400]

Len: 0, Cap: 2, Array: []
Len: 2, Cap: 2, Array: [1 2](Cheio!)
Len: 3, Cap: 4, Array: [1 2 3] (Cresceu!)

Original:  [Batman Superman Mulher Maravilha]
Cópia (Slice):  [Batman Superman Mulher Maravilha]
--- Depois de mudar a cópia ---
Cópia:  [Coringa Superman Mulher Maravilha]
Original:  [Coringa Superman Mulher Maravilha]
```
**Observações:** comportamento esperado e documentado pelo próprio comentário do código ("O Batman virou Coringa no original também!"). É a demonstração intencional do "ALERTA DE PERIGO" descrito no `fonte.txt`: fatiar (`original[0:3]`) não copia dados, apenas cria uma nova "janela" sobre o mesmo array. Sem bugs — o comportamento é o conceito sendo ensinado.

---

## `exercicios/ex01/main.go`

```go
func main() {
	convidados := []string{}
	for i := 0; i < 5; i++ {
		convidados = append(convidados, fmt.Sprintf("Convidado %d", i+1))
	}
	fmt.Println("Lista de convidados: ", convidados)
	fmt.Println("Tamanho da lista: ", len(convidados))
}
```

| Iteração | `i` | `fmt.Sprintf("Convidado %d", i+1)` | `convidados` (depois do append) |
| :---: | :---: | :--- | :--- |
| 1 | 0 | "Convidado 1" | `[Convidado 1]` |
| 2 | 1 | "Convidado 2" | `[Convidado 1 Convidado 2]` |
| 3 | 2 | "Convidado 3" | `[Convidado 1 Convidado 2 Convidado 3]` |
| 4 | 3 | "Convidado 4" | `[... Convidado 4]` |
| 5 | 4 | "Convidado 5" | `[... Convidado 5]` |

**Saída final no terminal:**
```
Lista de convidados:  [Convidado 1 Convidado 2 Convidado 3 Convidado 4 Convidado 5]
Tamanho da lista:  5
```
**Observações:** `fmt.Sprintf` funciona como `Printf`, mas em vez de imprimir no terminal **retorna a string formatada**, que pode ser guardada em variável (aqui, passada direto para `append`). Sem bugs.

---

## `exercicios/ex02/main.go`

```go
func main() {
	// Jeito problemático (comentado)
	//nums := []int{10, 20, 30, 40, 50}
	//fmt.Println("Teste: ", nums[:2])
	//fmt.Println("Teste: ", nums[3:])
	//resultado := append(nums[:2], nums[3:]...)
	//fmt.Println("Resulado: \n", resultado)
	//fmt.Println("Original", nums)

	// Jeito certo
	nums := []int{10, 20, 30, 40, 50}
	resultado := make([]int, 0, len(nums)-1)
	resultado = append(resultado, nums[:2]...)
	resultado = append(resultado, nums[3:]...)
	fmt.Println("Resultado: ", resultado)
}
```

| Linha | Instrução | Estado das variáveis | Saída produzida |
| :---: | :--- | :--- | :--- |
| `nums := []int{10, 20, 30, 40, 50}` | declaração | `nums = [10 20 30 40 50]` | — |
| `resultado := make([]int, 0, len(nums)-1)` | cria slice **novo e independente**, len=0, cap=4 | `resultado = []` | — |
| `append(resultado, nums[:2]...)` | espalha `[10 20]` para dentro do array novo de `resultado` (não reaproveita o array de `nums`) | `resultado = [10 20]` | — |
| `append(resultado, nums[3:]...)` | espalha `[40 50]` para dentro do mesmo array novo | `resultado = [10 20 40 50]` | — |
| `fmt.Println("Resultado: ", resultado)` | — | — | `Resultado:  [10 20 40 50]` |

**Saída final no terminal:**
```
Resultado:  [10 20 40 50]
```

**Observações — comparação com a versão comentada ("jeito problemático"):** este é o ponto que motivou a criação do teste de mesa. A versão ativa do código (linhas 19-26) usa `make` para criar um array **novo**, e por isso `nums` permanece intacto — `[10 20 30 40 50]` — mesmo depois dos `append`s em `resultado`. Já a versão comentada (`append(nums[:2], nums[3:]...)`) reaproveitaria o array de `nums[:2]` como destino do primeiro `append`, sobrescrevendo a posição do índice 2 (que continha `30`) com o primeiro valor de `nums[3:]` (`40`) — o que corrompe `nums` por baixo dos panos, fazendo `nums` virar `[10 20 40 40 50]` e `resultado` aparecer como `[10 20 40 50]` (mas compartilhando array com `nums`, então uma alteração posterior em um afetaria o outro). Foi exatamente esse comportamento (`Original [10 20 40 50 50]` ao testar manualmente uma variação com `Println` extra) que o usuário observou e que gerou a explicação sobre aliasing de slices e o operador `...`. A versão final no arquivo já está corrigida com `make` + `append` em slice novo — comportamento correto e seguro.

---

## `exercicios/ex03/main.go`

```go
func main() {
	licao := []int{2, 5, 2, 8, 5, 9, 2}
	unicos := []int{}

	for _, valor := range licao {
		jaExiste := false

		for _, existente := range unicos {
			if existente == valor {
				jaExiste = true
				break
			}
		}

		if !jaExiste {
			unicos = append(unicos, valor)
		}
	}
	fmt.Println("Lista de unicos:", unicos)
}
```

**Estado inicial:** `licao = [2 5 2 8 5 9 2]` · `unicos = []`

---

### Iteração 1 — `valor = 2`

| Sub-iter. (loop interno) | `existente` | `existente == valor`? | `jaExiste` |
| :---: | :---: | :---: | :---: |
| `unicos` vazio | — | loop não executa | `false` |

`!jaExiste` → `true` → **`append(2)`**

**`unicos = [2]`**

---

### Iteração 2 — `valor = 5`

| Sub-iter. | `existente` | `existente == valor`? | `jaExiste` |
| :---: | :---: | :---: | :---: |
| 1 | 2 | `2 == 5`? Não | `false` |
| (fim do loop) | — | — | `false` |

`!jaExiste` → `true` → **`append(5)`**

**`unicos = [2 5]`**

---

### Iteração 3 — `valor = 2` *(duplicata)*

| Sub-iter. | `existente` | `existente == valor`? | `jaExiste` |
| :---: | :---: | :---: | :---: |
| 1 | 2 | `2 == 2`? **Sim** → `break` | `true` |

`!jaExiste` → `false` → **não adiciona**

**`unicos = [2 5]`** *(sem mudança)*

---

### Iteração 4 — `valor = 8`

| Sub-iter. | `existente` | `existente == valor`? | `jaExiste` |
| :---: | :---: | :---: | :---: |
| 1 | 2 | `2 == 8`? Não | `false` |
| 2 | 5 | `5 == 8`? Não | `false` |
| (fim do loop) | — | — | `false` |

`!jaExiste` → `true` → **`append(8)`**

**`unicos = [2 5 8]`**

---

### Iteração 5 — `valor = 5` *(duplicata)*

| Sub-iter. | `existente` | `existente == valor`? | `jaExiste` |
| :---: | :---: | :---: | :---: |
| 1 | 2 | `2 == 5`? Não | `false` |
| 2 | 5 | `5 == 5`? **Sim** → `break` | `true` |

`!jaExiste` → `false` → **não adiciona**

**`unicos = [2 5 8]`** *(sem mudança)*

---

### Iteração 6 — `valor = 9`

| Sub-iter. | `existente` | `existente == valor`? | `jaExiste` |
| :---: | :---: | :---: | :---: |
| 1 | 2 | `2 == 9`? Não | `false` |
| 2 | 5 | `5 == 9`? Não | `false` |
| 3 | 8 | `8 == 9`? Não | `false` |
| (fim do loop) | — | — | `false` |

`!jaExiste` → `true` → **`append(9)`**

**`unicos = [2 5 8 9]`**

---

### Iteração 7 — `valor = 2` *(duplicata)*

| Sub-iter. | `existente` | `existente == valor`? | `jaExiste` |
| :---: | :---: | :---: | :---: |
| 1 | 2 | `2 == 2`? **Sim** → `break` | `true` |

`!jaExiste` → `false` → **não adiciona**

**`unicos = [2 5 8 9]`** *(sem mudança)*

---

**Saída final no terminal:**
```
Lista de unicos: [2 5 8 9]
```

**Observações:** o `break` no loop interno é uma otimização importante — nas iterações 3 e 7 (`valor = 2`), o número é encontrado já no índice 0 de `unicos`, e o loop para imediatamente sem varrer o resto. Na iteração 5 (`valor = 5`), varre apenas 2 elementos antes de parar. Sem `break`, o loop continuaria até o fim de `unicos` mesmo depois de encontrar o valor — ineficiente. A variável `jaExiste` sendo redeclarada como `false` no início de cada iteração externa é essencial: sem esse reset, um `true` de uma iteração anterior contaminaria as seguintes e nenhum número seria adicionado após o primeiro duplicado encontrado.

---

## ✅ Conclusão

Nenhum bug no código ativo do capítulo 07. O exercício 03 demonstra claramente o padrão de "bandeirinha booleana" (`jaExiste`) combinado com loop aninhado para busca linear — a solução de força bruta O(n²) que o capítulo pede antes de apresentar Maps (capítulo 08) como solução O(n). O grande aprendizado do capítulo é conceitual: fatiar um slice (`slice[a:b]`) não copia dados, apenas cria uma nova visão sobre o mesmo array subjacente, e usar `append` sobre essa fatia pode sobrescrever o array original silenciosamente — a forma segura é alocar um destino novo com `make` e usar `append(destino, origem...)`.
