# 🧮 Teste de Mesa — Capítulo 08: Maps

Rastreamento manual da execução de cada `main.go` do capítulo, linha a linha, sem rodar o código.

---

## exercicios/exe01 — O Contador de Palavras

```go
palavras := []string{"banana", "maçã", "laranja", "maçã", "banana"}
contador := make(map[string]int)
for _, palavra := range palavras {
    contador[palavra]++
}
fmt.Println("Contagem de palavras:", contador)
```

**Estado inicial:**

| variável | valor |
|---|---|
| `palavras` | `["banana", "maçã", "laranja", "maçã", "banana"]` |
| `contador` | `{}` (map vazio) |

**Iteração do loop `range palavras`:**

| it. | `palavra` | `contador[palavra]` antes | operação | `contador[palavra]` depois | estado do map |
|---|---|---|---|---|---|
| 1 | `"banana"` | `0` ← zero value (chave ausente) | `0 + 1` | `1` | `{"banana": 1}` |
| 2 | `"maçã"` | `0` ← zero value (chave ausente) | `0 + 1` | `1` | `{"banana": 1, "maçã": 1}` |
| 3 | `"laranja"` | `0` ← zero value (chave ausente) | `0 + 1` | `1` | `{"banana": 1, "maçã": 1, "laranja": 1}` |
| 4 | `"maçã"` | `1` ← chave existe | `1 + 1` | `2` | `{"banana": 1, "maçã": 2, "laranja": 1}` |
| 5 | `"banana"` | `1` ← chave existe | `1 + 1` | `2` | `{"banana": 2, "maçã": 2, "laranja": 1}` |

**Saída esperada:**
```
Contagem de palavras: map[banana:2 laranja:1 maçã:1]
```

> ⚠️ A ordem das chaves no output **não é garantida** — Maps em Go não preservam ordem de inserção. O Go pode imprimir as chaves em qualquer sequência a cada execução.

> ⚠️ O comentário do enunciado diz `"banana": 3`, mas o slice só contém 2 ocorrências de `"banana"` — o correto é `banana: 2`. Provável typo no exercício.

---

## exercicios/exe02 — O Dicionário de Cores (Comma Ok Idiom)

```go
cores := map[string]string{
    "Vermelho" : "Red",
    "Verde"    : "Green",
    "Amarelo"  : "Yellow",
    "Laranja"  : "Orange",
}
corBuscada := "Vermelho"
traducao, ok := cores[corBuscada]
if ok {
    fmt.Printf("A cor \"%s\" em Inglês é %s\n", corBuscada, traducao)
} else {
    fmt.Printf("A cor \"%s\" não foi encontrada!\n", corBuscada)
}
```

**Estado inicial:**

| variável | valor |
|---|---|
| `cores` | `{"Vermelho":"Red", "Verde":"Green", "Amarelo":"Yellow", "Laranja":"Orange"}` |
| `corBuscada` | `"Vermelho"` |

**Linha por linha:**

| linha | expressão | resultado |
|---|---|---|
| `traducao, ok := cores["Vermelho"]` | chave `"Vermelho"` existe no map | `traducao = "Red"`, `ok = true` |
| `if ok` | `true` | entra no bloco `if` |
| `fmt.Printf(...)` | `corBuscada="Vermelho"`, `traducao="Red"` | imprime linha abaixo |

**Saída esperada:**
```
A cor "Vermelho" em Inglês é Red
```

**Cenário alternativo — chave ausente (`corBuscada := "Azul"`):**

| linha | expressão | resultado |
|---|---|---|
| `traducao, ok := cores["Azul"]` | chave `"Azul"` não existe | `traducao = ""`, `ok = false` |
| `if ok` | `false` | vai para o `else` |
| `fmt.Printf(...)` | `corBuscada="Azul"` | imprime linha abaixo |

```
A cor "Azul" não foi encontrada!
```

> 💡 Sem o comma ok, `traducao` seria `""` nos dois casos (chave ausente e chave com valor `""`). O `bool ok` é a única forma confiável de distinguir os dois cenários.
