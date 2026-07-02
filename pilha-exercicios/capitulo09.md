# 🏋️ Pilha de Exercícios — Capítulo 09: Ponteiros

Exercícios extras em dificuldade crescente para praticar Ponteiros em Go. **Sem solução no arquivo** — o objetivo é resolver de verdade.

---

## 🟢 Fáceis

### Ex01 — Dobrador
Crie uma função `dobrar(n *int)` que multiplica o valor original por 2. Na `main`, crie `x := 7`, chame `dobrar(&x)` e imprima `x` — deve mostrar `14`.

### Ex02 — Zerador Condicional
Crie uma função `zerarSeNegativo(n *int)` que zera o valor original somente se ele for negativo. Teste com `a := -5` e `b := 3`.

### Ex03 — Ponteiro Nil Guard
Crie uma função `imprimirValor(p *int)` que verifica se o ponteiro é `nil` antes de desreferenciar. Se for `nil`, imprime `"ponteiro nulo"`. Se não, imprime o valor. Teste nos dois casos.

### Ex04 — Trocar Strings
Repita o exercício do Trocador, mas desta vez com `*string`. Troque os valores de `nome1 := "Ana"` e `nome2 := "Bruno"` usando uma função `trocarStrings(a, b *string)`.

---

## 🟡 Intermediários

### Ex05 — Fábrica de Ponteiros
Crie uma função `novoInt(valor int) *int` que aloca um inteiro com o valor fornecido e retorna o ponteiro. Use `new` internamente. Na `main`, crie três ponteiros com valores 10, 20 e 30 e imprima `*p1 + *p2 + *p3`.

### Ex06 — Modificador de Slice via Ponteiro
Crie uma função `adicionarCem(nums *[]int)` que adiciona o valor `100` ao slice original. Na `main`, crie `lista := []int{1, 2, 3}`, chame a função e imprima `lista` — deve mostrar `[1 2 3 100]`.

### Ex07 — Contador Compartilhado
Crie três funções: `incrementar(c *int)`, `decrementar(c *int)` e `resetar(c *int)`. Use um único `contador := 0` na `main` e chame as três funções em sequências diferentes, imprimindo o valor após cada operação.

### Ex08 — Comparador por Ponteiro
Crie uma função `maior(a, b *int) *int` que retorna o ponteiro para o maior dos dois valores. Na `main`, imprima `*maior(&x, &y)`. Se os dois forem iguais, retorne qualquer um dos dois.

---

## 🔴 Difíceis

### Ex09 — Mini Pilha (Stack)
Implemente uma mini pilha de inteiros usando slice e ponteiros. Crie as funções `push(pilha *[]int, valor int)` e `pop(pilha *[]int) int`. A `pop` deve remover e retornar o último elemento. Teste com push de 3 valores e pop de 2.

### Ex10 — Encadeamento de Modificações
Crie três funções que recebem `*int`: `triplicar`, `subtrairCinco` e `elevarAoQuadrado`. Na `main`, aplique as três em sequência sobre um mesmo valor e imprima o resultado após cada passo. Ex: `x := 3` → `9` → `4` → `16`.

### Ex11 — Ponteiro para Map
Crie uma função `adicionarPalavra(dicionario *map[string]int, palavra string)` que incrementa a contagem da palavra no map original. Na `main`, crie o map com `make`, chame a função 5 vezes (com repetições) e imprima o resultado final.

### Ex12 — Linked List Simples
Implemente um nó de lista encadeada:
```go
type No struct {
    valor    int
    proximo  *No
}
```
Crie uma função `inserirNoFim(cabeca **No, valor int)` que insere um novo nó no fim da lista. Construa uma lista com valores `[1, 2, 3, 4, 5]` e percorra-a imprimindo cada valor.
