# Capítulo 08 — Maps: O Guarda-Volumes de Alta Performance

## Tema Central

Maps são a estrutura de dados chave→valor do Go. São implementados como Hash Tables — a mesma estrutura usada por dicionários do Python, HashMap do Java e objetos do JavaScript. O resultado é acesso médio em O(1): independentemente de ter 10 ou 10 milhões de entradas, encontrar um valor leva o mesmo tempo médio.

---

## O que o Livro Cobre

- Analogia: guarda-volumes de aeroporto (chave = número, valor = mochila)
- `make(map[K]V)` como forma idiomática de criar Map
- Zero value em Maps: chave ausente retorna zero value do tipo valor
- Padrão `contador[chave]++` para contagem idiomática
- Comma ok idiom: `valor, ok := mapa[chave]`
- `delete(mapa, chave)` para remover entradas
- Ordem de iteração de Maps é aleatória

---

## Aprofundando os Conceitos

### Como uma Hash Table funciona internamente

Uma Hash Table transforma uma chave em um índice de array usando uma **função hash**:

```
chave "Jorge" → função hash → índice 42 → array[42] = "valor"
```

Propriedades:
- A mesma chave sempre produz o mesmo índice (determinístico)
- Chaves diferentes podem produzir o mesmo índice (**colisão**)
- O Go resolve colisões com encadeamento (linked list ou bucket)
- Em média: O(1) para leitura, escrita e deleção
- No pior caso (muitas colisões): O(n) — mas improvável com bom hash

O Go usa uma função hash baseada em AES quando disponível no hardware (a maioria dos processadores modernos), o que torna o hashing extremamente rápido.

### Tipos válidos como chave de Map

Qualquer tipo **comparável** com `==` pode ser chave:
- `int`, `float64`, `bool`, `string`
- `pointer`, `channel`, `interface`
- Arrays (`[3]int` é comparável, `[]int` NÃO)
- Structs cujos campos são todos comparáveis

```go
// Válido
map[string]int{}
map[int]string{}
map[[3]int]string{}        // array como chave — OK
map[struct{ x, y int }]string{}  // struct como chave — OK

// INVÁLIDO — compilação falha
map[[]int]string{}         // slice NÃO é comparável
map[map[string]int]string{} // map NÃO é comparável
```

### Criação: `make` vs literal

```go
// make — criação idiomática (vazio mas pronto para uso)
m := make(map[string]int)

// Literal — inicialização com valores
m := map[string]int{
    "maçã":   5,
    "banana": 3,
    "uva":    8,
}

// ARMADILHA: declaração sem make = nil map
var m map[string]int  // nil!
m["chave"] = 1        // panic: assignment to entry in nil map
```

Com `make`, você pode fornecer uma capacidade inicial (hint):
```go
m := make(map[string]int, 1000)  // pré-aloca para ~1000 elementos (evita re-hash)
```

### Zero Value em Maps — a mecânica exata

```go
m := make(map[string]int)

// Leitura de chave inexistente: retorna ZERO VALUE, sem panic, sem erro
valor := m["inexistente"]
fmt.Println(valor)  // 0 (zero value de int)

// Isso torna possível o padrão de contagem
palavras := []string{"go", "go", "gopher", "go"}
contador := make(map[string]int)
for _, p := range palavras {
    contador[p]++  // primeira ocorrência: 0++ = 1; próximas: n++
}
// contador = {"go": 3, "gopher": 1}
```

### Comma Ok Idiom — distinguindo ausente de zero

O problema: se o zero value é um valor legítimo, como distinguir "não existe" de "existe com valor zero"?

```go
notas := map[string]int{"Ana": 10, "Bob": 0}

// Sem comma ok: impossível distinguir
nota := notas["Carlos"]  // 0 — Carlos não existe ou tirou zero?

// Com comma ok: distinção clara
nota, existe := notas["Carlos"]
if !existe {
    fmt.Println("Carlos não está no sistema")
} else {
    fmt.Printf("Carlos tirou: %d\n", nota)  // mesmo se nota == 0
}
```

O `ok` é o segundo valor de retorno — uma convenção do Go para operações que podem ou não produzir um resultado. O mesmo padrão aparece em type assertions e channels.

### `delete` — seguro mesmo para chaves ausentes

```go
m := make(map[string]int)
m["chave"] = 1

delete(m, "chave")     // remove — OK
delete(m, "ausente")   // chave não existe — NENHUM ERRO, operação silenciosa
```

### Iteração sobre Maps — ordem aleatória intencional

```go
m := map[string]int{"b": 2, "a": 1, "c": 3}

// Ordem aleatória — DIFERENTE a cada execução
for k, v := range m {
    fmt.Printf("%s: %d\n", k, v)
}

// Para ordem determinística:
import "sort"
chaves := make([]string, 0, len(m))
for k := range m {
    chaves = append(chaves, k)
}
sort.Strings(chaves)
for _, k := range chaves {
    fmt.Printf("%s: %d\n", k, m[k])
}
```

A aleatoriedade é **proposital**: evitar que programas dependam de uma ordem de iteração que pode mudar entre versões do Go. Também é uma proteção contra ataques de hash collision DoS.

### Operações avançadas com Maps

```go
// Verificar tamanho
tamanho := len(m)

// Map de Slices (agrupar elementos)
grupos := make(map[string][]string)
grupos["frutas"] = append(grupos["frutas"], "maçã")
grupos["frutas"] = append(grupos["frutas"], "banana")
grupos["legumes"] = append(grupos["legumes"], "cenoura")

// Map de Maps (estrutura hierárquica)
escola := make(map[string]map[string]int)
escola["turma-A"] = map[string]int{"Ana": 9, "Bob": 7}
escola["turma-B"] = map[string]int{"Carlos": 8}

// Verificar antes de inserir em map aninhado
if escola["turma-C"] == nil {
    escola["turma-C"] = make(map[string]int)
}
escola["turma-C"]["Diana"] = 10
```

---

## Referências Oficiais

- **Especificação — Map types**: https://go.dev/ref/spec#Map_types
- **Especificação — Index expressions (comma ok)**: https://go.dev/ref/spec#Index_expressions
- **Effective Go — Maps**: https://go.dev/doc/effective_go#maps
- **Blog: Go maps in action**: https://go.dev/blog/maps
- **Tour de Go — Maps**: https://go.dev/tour/moretypes/19
- **Pacote `maps` (Go 1.21+)**: https://pkg.go.dev/maps

---

## Exemplos de Código Adicionais

### Frequência de palavras (exemplo clássico)

```go
func frequencia(texto string) map[string]int {
    contador := make(map[string]int)
    palavras := strings.Fields(texto)  // split por whitespace
    for _, palavra := range palavras {
        palavra = strings.ToLower(palavra)
        contador[palavra]++
    }
    return contador
}
```

### Set (conjunto) com Map

Go não tem tipo Set nativo — mas `map[T]struct{}` simula perfeitamente:

```go
// struct{} tem tamanho zero — sem desperdício de memória
visitados := make(map[string]struct{})

// Adicionar ao set
visitados["página1"] = struct{}{}
visitados["página2"] = struct{}{}

// Verificar pertencimento
_, existe := visitados["página1"]
fmt.Println(existe)  // true

// Remover
delete(visitados, "página1")
```

### Cache com Maps

```go
var cache = make(map[int]int)

func fibonacci(n int) int {
    if n <= 1 {
        return n
    }
    if v, ok := cache[n]; ok {
        return v  // cache hit
    }
    resultado := fibonacci(n-1) + fibonacci(n-2)
    cache[n] = resultado  // armazenar no cache
    return resultado
}
```

### `maps` package (Go 1.21+)

```go
import "maps"

m1 := map[string]int{"a": 1, "b": 2}
m2 := maps.Clone(m1)           // cópia rasa
maps.Copy(m1, m2)               // copiar m2 em m1
maps.DeleteFunc(m1, func(k string, v int) bool {
    return v > 1  // remove entradas onde v > 1
})
for k, v := range maps.All(m1) {}  // iterator
```

---

## Perguntas & Respostas Frequentes

**P: Maps em Go são thread-safe?**
R: **Não.** Acessar o mesmo map concorrentemente de múltiplas goroutines (sem sincronização) causa race condition e pode causar panic. Use `sync.RWMutex` para proteger ou `sync.Map` para casos de uso específicos:
```go
import "sync"
var mu sync.RWMutex
var m = make(map[string]int)

// Escrever
mu.Lock()
m["chave"] = 1
mu.Unlock()

// Ler
mu.RLock()
v := m["chave"]
mu.RUnlock()
```

**P: O que acontece se eu usar o mesmo Map em dois pacotes?**
R: Maps são passados por referência (o descritor contém um ponteiro). Modificações em qualquer pacote afetam o mesmo map subjacente. Diferente de Slices, não há aliasing sutil — maps sempre se comportam como referências.

**P: Por que não posso usar um Slice como chave de Map?**
R: Slices não são comparáveis — não têm operador `==` definido. Isso ocorre porque Slices são descritores (ponteiro + len + cap); dois Slices com o mesmo conteúdo teriam descritores diferentes. Use um Array (tamanho fixo) ou converta para string como chave.

**P: `delete` em um Map nil causa panic?**
R: Não. Assim como ler de um map nil retorna zero value sem panic, `delete` em map nil é silencioso. O que causa panic é **escrever** em um map nil.

**P: Como fazer um Map ordenado em Go?**
R: Go não tem Map ordenado nativo. Opções: (1) coletar chaves, ordenar, iterar; (2) usar uma estrutura de árvore de terceiros como `github.com/emirpasic/gods`; (3) manter uma slice de chaves paralela ordenada.

---

## Comparações com Outras Linguagens

### Dicionário/Map em diferentes linguagens

```python
# Python: dict — mantém ordem de inserção desde 3.7
d = {"nome": "Jorge", "idade": 30}
d["cidade"] = "SP"
del d["nome"]
if "idade" in d:  # verificação de existência
    print(d["idade"])
# KeyError se acessar chave ausente!
```

```java
// Java: HashMap (não ordenado), TreeMap (ordenado), LinkedHashMap (inserção)
Map<String, Integer> m = new HashMap<>();
m.put("chave", 1);
m.getOrDefault("ausente", 0);  // equivalente ao zero value
m.containsKey("chave");
```

```javascript
// JavaScript: Object (string keys) ou Map (qualquer tipo de chave)
const m = new Map();
m.set("chave", 1);
m.has("chave");   // true
m.get("ausente"); // undefined (similar a nil)
m.delete("chave");
for (const [k, v] of m) {} // iteração em ordem de inserção
```

```go
// Go: map — sem ordem garantida, tipo-seguro
m := make(map[string]int)
m["chave"] = 1
delete(m, "chave")
v, ok := m["ausente"]  // comma ok para distinguir ausente de zero value
```

---

## Armadilhas Comuns

1. **Escrever em Map nil**: `var m map[string]int; m["k"] = 1` → panic. Sempre use `make`.

2. **Confiar em ordem de iteração**: o Go embaralha deliberadamente a ordem para desencorajar dependência.

3. **Map não é thread-safe**: acesso concorrente sem mutex causa race condition ou panic.

4. **Deletar chave durante iteração**: o comportamento é definido pelo Go — é seguro deletar chaves durante `range`, mas as iterações seguintes podem ou não ver a chave deletada.

5. **Map aninhado não inicializado**:
   ```go
   m := make(map[string]map[string]int)
   m["a"]["b"] = 1  // PANIC: m["a"] é nil
   // Solução: inicializar antes de usar
   m["a"] = make(map[string]int)
   m["a"]["b"] = 1  // OK
   ```

---

## Quiz de Fixação

1. O que é uma Hash Table e por que garante O(1) em média?
2. Por que `var m map[string]int` causa panic ao escrever?
3. Como o comma ok idiom resolve a ambiguidade entre "chave ausente" e "zero value"?
4. O que acontece se você iterar sobre o mesmo Map duas vezes? A ordem é a mesma?
5. Por que `[]int` não pode ser chave de Map mas `[3]int` pode?
6. Como implementar um Set (conjunto sem repetição) em Go usando Map?

---

## 🔬 Dissecando a Sintaxe

### Declaração e criação de Map — as formas corretas e incorretas

```go
var m1 map[string]int           // [1] ERRADO para escrita
m2 := make(map[string]int)      // [2] CORRETO
m3 := map[string]int{"a": 1}   // [3] CORRETO com valores iniciais
```

```
[1]  var m1 map[string]int
              ────────────
                    └─ tipo do map: "map" + "[tipo da chave]" + "tipo do valor"
                        map[string]int = chaves são string, valores são int
              m1 = nil (zero value de map) — leitura OK, escrita = PANIC

[2]  make(map[string]int)
     ────  ────────────
       │         └─ tipo a criar
       └─ função builtin que inicializa a estrutura interna do map
           retorna um map pronto para leitura E escrita

[3]  map[string]int{"a": 1}
     ────────────   ───────
          │           └─ par chave:valor (literal de map)
          └─ tipo do map
          "{" "}" = abre/fecha o literal (como um struct literal)
          "a": 1 → chave "a" com valor 1
```

### O padrão de contagem — dissecar `contador[palavra]++`

```go
contador := make(map[string]int)
palavras := []string{"go", "go", "gopher"}

for _, palavra := range palavras {
    contador[palavra]++          // [1]
}
```

```
[1]  contador[palavra]++
      ────────  ──────   ──
          │       │       └─ "++" = incremento: soma 1 e armazena de volta
          │       └─ a chave: o valor da variável "palavra" nesta iteração
          └─ o map: acessar com uma chave

EXPANSÃO DO QUE ACONTECE:
  contador[palavra]++
  é exatamente igual a:
  contador[palavra] = contador[palavra] + 1

PRIMEIRA OCORRÊNCIA ("go", quando ainda não existe no map):
  contador["go"] = contador["go"] + 1
               ↑
               leitura de chave INEXISTENTE → retorna ZERO VALUE = 0
  contador["go"] = 0 + 1 = 1  ← guarda 1 no map

SEGUNDA OCORRÊNCIA ("go"):
  contador["go"] = contador["go"] + 1
                            ↑
                            chave EXISTE → retorna 1
  contador["go"] = 1 + 1 = 2  ← atualiza para 2
```

### Comma Ok Idiom — a distinção que salva vidas

```go
notas := map[string]int{"Ana": 10, "Bob": 0}

valor, ok := notas["Carlos"]    // [1]
```

```
valor  ,  ok  :=  notas  [  "Carlos"  ]
─────     ──   ─  ─────     ─────────
  │        │   │    │            └─ a chave buscada
  │        │   │    └─ o map
  │        │   └─ ":=" declara DUAS variáveis novas
  │        └─ segundo valor de retorno:
  │             true  = a chave "Carlos" EXISTE no map
  │             false = a chave "Carlos" NÃO existe
  └─ primeiro valor de retorno:
      se ok=true  → o valor associado à chave
      se ok=false → zero value do tipo valor (int=0, string="", bool=false...)

SEM comma ok: impossível distinguir "Carlos não existe" de "Carlos tirou 0"
COM comma ok: ok=false → definitivamente não existe
```

---

## 🎙️ Prompts para o NotebookLM

### 🎧 Audio Overview
> "Gere um Audio Overview explicando Maps em Go para alguém que conhece dicionários do Python. Cubra: por que o Map do Go não tem ordem (ao contrário do dict do Python 3.7+), por que escrever em map nil causa panic mas ler não, o padrão contador[palavra]++ com zero value, e o comma ok idiom com um exemplo de bug que ele previne."

### 📋 Briefing Doc
> "Crie um briefing sobre 'Maps em Go: as 5 coisas que todo iniciante precisa saber', cobrindo: criação com make vs literal, zero value como superpoder, comma ok para distinção, delete seguro, e a ausência de order garantida."

### 📚 Study Guide
> "Crie um guia de estudos do Capítulo 08 com: (1) tabela das operações de Map (criar, ler, escrever, deletar, verificar existência) com exemplos de código, (2) exercício de trace: dado um slice de palavras, trace o estado do map contador após cada iteração, (3) quiz de 8 perguntas sobre maps."

### 🖼️ Roteiro de Slides
> "Crie um roteiro de 6 slides sobre 'Maps em Go'. Slide 1: O que é uma Hash Table — analogia do guarda-volumes. Slide 2: make vs nil map — o crash que ninguém espera. Slide 3: Zero Value — o padrão contador++. Slide 4: Comma Ok — distinguindo ausente de zero. Slide 5: Iteração sem ordem — por que é intencional. Slide 6: Map vs Set — como simular Set em Go."

### 💬 Perguntas Profundas para o Chat
- "O que é uma Hash Table internamente? Como a função hash converte uma chave string em um índice de array?"
- "Por que escrever em um map nil causa panic, mas ler de um map nil retorna zero value sem panic? Qual a lógica por trás dessa assimetria?"
- "Maps em Go não são thread-safe. Me explique o que pode acontecer se duas goroutines acessam o mesmo map simultaneamente e como resolver."
- "Qual a diferença entre `map[string]struct{}` e `map[string]bool` para implementar um Set? Por que `struct{}` é preferido?"
