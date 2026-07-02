# 🏋️ Pilha de Exercícios — Capítulo 08: Maps

Exercícios extras em dificuldade crescente para praticar Maps em Go. **Sem solução no arquivo** — o objetivo é resolver de verdade.

---

## 🟢 Fáceis

### Ex01 — Inventário de Frutas
Dado um slice `[]string{"maçã", "banana", "maçã", "laranja", "banana", "maçã"}`, crie um map que conte quantas vezes cada fruta aparece. Imprima o resultado.

### Ex02 — Tradução Reversa
Crie um map `pt→en` com pelo menos 5 palavras. Em seguida, crie um segundo map `en→pt` invertendo as chaves e valores do primeiro. Imprima o map invertido.

### Ex03 — Verificador de Presença
Crie um map `map[string]bool` representando uma lista de convidados confirmados. Dada uma lista de nomes a verificar, imprima quem está confirmado e quem não está (use o comma ok idiom).

### Ex04 — Contagem de Caracteres
Dada uma string qualquer (ex: `"gopher"`) percorra caractere por caractere e construa um `map[rune]int` com a frequência de cada letra. Imprima o resultado.

---

## 🟡 Intermediários

### Ex05 — Agrupador por Inicial
Dado um slice de nomes `[]string{"Ana", "Bruno", "Alice", "Carlos", "Beatriz", "Andre"}`, crie um `map[string][]string` onde a chave é a letra inicial e o valor é o slice de nomes que começam com aquela letra.

### Ex06 — Placar de Partidas
Simule um placar de campeonato. Dado um slice de resultados no formato `[]string{"Brasil 3 Argentina 1", "Argentina 2 Brasil 2", "Brasil 1 Argentina 0"}`, construa um map com o total de gols marcados por cada seleção.

### Ex07 — Frequência de Palavras (com ordenação)
Repita o exercício exe01 do capítulo (contador de palavras), mas desta vez exiba o resultado **ordenado alfabeticamente por chave** usando `sort.Strings()` sobre um slice das chaves do map.

### Ex08 — Cache de Resultados
Crie uma função `quadrado(n int) int` que calcula `n * n`. Use um `map[int]int` como cache: antes de calcular, verifique se o resultado já existe no map; se sim, retorne do cache; se não, calcule, armazene no map e retorne. Teste com chamadas repetidas para o mesmo número.

---

## 🔴 Difíceis

### Ex09 — Anagrama
Dadas duas strings (ex: `"listen"` e `"silent"`), determine se elas são anagramas usando Maps. Dica: construa um `map[rune]int` para cada string contando a frequência de cada letra e compare os dois maps.

### Ex10 — Índice Invertido
Dado um slice de frases `[]string{"o gato subiu", "o rato roeu", "o gato roeu"}`, construa um índice invertido: um `map[string][]int` onde a chave é cada palavra única e o valor é o slice de índices das frases onde essa palavra aparece.

### Ex11 — Top N Palavras
Dado um texto longo (use uma string com pelo menos 20 palavras repetidas), construa um `map[string]int` de frequência e depois encontre as **3 palavras mais frequentes**. Imprima-as em ordem decrescente de frequência.

### Ex12 — Grafo de Adjacência
Represente um grafo simples usando `map[string][]string` onde cada chave é um nó e o valor é o slice de vizinhos. Implemente uma função que, dado um nó de origem, retorne todos os nós alcançáveis via busca em largura (BFS).
