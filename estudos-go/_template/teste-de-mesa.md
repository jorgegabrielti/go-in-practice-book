# 🧮 Teste de Mesa — Capítulo XX

> **O que é um teste de mesa?** Uma simulação manual da execução do programa, linha a linha, **sem rodar o código**. Para cada `main.go` do capítulo, a tabela rastreia o estado das variáveis a cada instrução relevante e prevê a saída exata no terminal — incluindo comportamentos inesperados (bugs, efeitos colaterais, armadilhas de sintaxe). Depois de prever, rode `go run` no arquivo real para confirmar.

---

## Formatos de tabela

Use o formato adequado para cada tipo de código:

### Código sequencial (sem loop ou com loop simples)

| Linha | Instrução | Estado das variáveis | Saída produzida |
| :---: | :--- | :--- | :--- |
| `variavel := valor` | declaração | `variavel = valor` | — |
| `fmt.Println(...)` | — | — | `texto impresso` |

### Loop simples (uma única variável de controle)

| Iteração | `i` | Condição | Estado | Saída |
| :---: | :---: | :---: | :--- | :--- |
| 1 | 0 | `true` | `acumulador = X` | `texto` |
| 2 | 1 | `true` | `acumulador = Y` | `texto` |
| — | N | `false` | — | loop termina |

### Loop aninhado (loop dentro de loop)

Para cada iteração do loop **externo**, mostrar uma sub-tabela com as iterações do loop **interno**. Isso expõe exatamente o que cada variável interna vale a cada passo — essencial para rastrear bugs em buscas, filtros e algoritmos de comparação.

**Estrutura:**

**Estado inicial:** `slice = [...]` · `resultado = []`

---

#### Iteração N (externa) — `valor = X`

| Sub-iter. (loop interno) | `variavel_interna` | `condição`? | `bandeirinha` |
| :---: | :---: | :---: | :---: |
| 1 | A | `A == X`? Não | `false` |
| 2 | B | `B == X`? **Sim** → `break` | `true` |

`!bandeirinha` → `false` → **não adiciona** *(ou `true` → **`append(X)`**)*

**`resultado = [...]`** *(sem mudança / novo estado)*

---

*(repita para cada iteração externa)*

---

## `exemplos/ex01/main.go`

**Código:**
```go
// cole aqui o código rastreado
```

*(use o formato de tabela adequado ao tipo do código acima)*

**Saída final no terminal:**
```
(saída prevista aqui)
```

**Observações:** comportamento inesperado, bug ou ponto de atenção encontrado (se houver).

---

## `exercicios/ex01/main.go`

*(repita a estrutura acima para cada exemplo/exercício do capítulo)*

---

## ✅ Conclusão

Resumo do que o teste de mesa revelou: tudo se comportou como esperado? Algum bug de sintaxe (`Println` com verbos de formatação, `Printf` com número errado de argumentos, aliasing de slices, overflow silencioso, etc.) foi encontrado? Vale a pena corrigir o exercício original?
