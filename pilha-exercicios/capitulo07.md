# 🏋️ Pilha de Exercícios — Capítulo 07: Arrays e Slices

> Exercícios extras sem solução. Resolva por conta própria antes de buscar respostas. Crie suas soluções em `estudos-go/capitulo07/exercicios/practice01/`, `practice02/`, etc.
>
> 🟢 Fácil · 🟡 Médio · 🔴 Desafio
>
> Estes exercícios podem usar qualquer conceito dos capítulos 1 a 7.

---

### 🟢 01. Primeiro e Último

Dado o slice `numeros := []int{10, 25, 3, 87, 42}`, imprima o **primeiro** e o **último** elemento sem usar os índices diretamente (`[0]` e `[4]`). Use `len()` para calcular o índice do último dinamicamente — assim o código funciona para qualquer tamanho de slice.

---

### 🟢 02. Somar Todos

Crie um slice `valores := []int{5, 12, 3, 8, 21, 7}` e use um loop `for range` para somar todos os elementos. Imprima o total ao final.

---

### 🟢 03. Duplicar Todos

Dado `original := []int{1, 2, 3, 4, 5}`, crie um **novo** slice chamado `dobrado` onde cada elemento vale o dobro do original (sem modificar `original`). Imprima os dois.

---

### 🟢 04. Inverter um Slice

Crie um slice de strings com 5 nomes e imprima-o na **ordem inversa** usando um loop. Não use nenhuma função de biblioteca — implemente a lógica de iteração reversa você mesmo.

---

### 🟡 05. Filtro de Pares

Dado `numeros := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}`, crie um novo slice chamado `pares` contendo apenas os números pares. Use `append` para construí-lo. Imprima o resultado e o tamanho.

---

### 🟡 06. Mover para o Fim

Dado `lista := []int{1, 2, 3, 4, 5}`, mova o **primeiro** elemento para o **final** sem alterar a ordem dos demais. O resultado esperado é `[2 3 4 5 1]`. Use fatiamento e `append` — mas cuidado com aliasing.

---

### 🟡 07. Interseção de Dois Slices

Crie dois slices:
```go
a := []int{1, 2, 3, 4, 5}
b := []int{3, 4, 5, 6, 7}
```
Produza um terceiro slice com os elementos que aparecem nos **dois** (interseção). Resultado esperado: `[3 4 5]`. Use loops aninhados (força bruta) — sem Maps por enquanto.

---

### 🟡 08. Rotação

Dado `fila := []string{"Ana", "Beto", "Carla", "Dani"}`, implemente uma **rotação à esquerda**: o primeiro elemento vai para o final, todos os outros avançam uma posição. Resultado esperado: `[Beto Carla Dani Ana]`. Repita a rotação 3 vezes e imprima o resultado após cada rodada.

---

### 🔴 09. Achatar Matriz

Crie uma "matriz" usando um slice de slices:
```go
matriz := [][]int{
    {1, 2, 3},
    {4, 5, 6},
    {7, 8, 9},
}
```
Produza um único slice `achatado` com todos os valores em ordem: `[1 2 3 4 5 6 7 8 9]`. Use `append` com o operador `...` para juntar cada linha ao resultado.

---

### 🔴 10. Chunks

Dado `dados := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}` e um tamanho de bloco `tamanho := 3`, divida `dados` em blocos de tamanho `tamanho`. O resultado deve ser:
```
[[1 2 3] [4 5 6] [7 8 9] [10]]
```
O último bloco pode ter menos elementos que os outros. Use um slice de slices (`[][]int`) para armazenar os blocos.

---

### 🔴 11. Frequência de Elementos

Dado `licao := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}`, conte quantas vezes cada número aparece **sem usar Maps**. Para cada número único no slice, percorra o slice inteiro e conte as ocorrências. Imprima no formato `número: N vez(es)`. Dica: reutilize a lógica de detecção de duplicatas do exercício 03 do livro.

---

### 🔴 12. Interseção Sem Repetição

Evolua o exercício 07: a interseção de `a := []int{1, 2, 2, 3, 4}` e `b := []int{2, 2, 3, 5}` deve ser `[2 3]` — sem duplicatas no resultado, mesmo que o número apareça várias vezes nos dois slices. Combine a lógica de interseção com a lógica do "Matador de Duplicatas" do capítulo.
