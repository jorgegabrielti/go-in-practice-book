# 🧮 Teste de Mesa — Capítulo 14: Pacotes e Módulos

> **Nota:** Este capítulo envolve múltiplas pastas e arquivos. O "teste de mesa" aqui foca em rastrear o fluxo de chamadas entre pacotes e a visibilidade de identificadores, não apenas a execução sequencial de um único `main.go`.

---

## Código rastreado

`exemplos/exe01/` — pacote `conversor` + `main.go`

---

## Estrutura de Arquivos

```
exemplos/exe01/
├── go.mod           (module github.com/seu-usuario/capitulo14)
├── main.go          (package main)
└── conversor/
    └── medidas.go   (package conversor)
```

---

## Rastreamento: `conversor/medidas.go`

| Identificador | Visibilidade | Acessível de fora? |
|---|---|---|
| `KmParaMilhas` | Maiúscula → **Exportada** | ✅ Sim — `conversor.KmParaMilhas(...)` |
| `milhasParaKm` | Minúscula → **Não-exportada** | ❌ Não — erro de compilação |

---

## Rastreamento: `main.go`

```
import "github.com/seu-usuario/capitulo14/conversor"
→ Go localiza a pasta conversor/ dentro do módulo
→ carrega as definições exportadas do package conversor
```

```
km := 100.0
→ km = float64(100.0)

milhas := conversor.KmParaMilhas(km)
→ chama KmParaMilhas(100.0)
   dentro da função: return 100.0 * 0.621371
   = 62.1371
→ milhas = float64(62.1371)

fmt.Printf("%.2f Km é igual a %.2f Milhas\n", km, milhas)
→ formata km com 2 casas: "100.00"
→ formata milhas com 2 casas: "62.14"
```

**Saída esperada:**
```
100.00 Km é igual a 62.14 Milhas
```

---

## Rastreamento: Linha comentada (visibilidade)

```go
// conversor.milhasParaKm(50) // ERRO: cannot refer to unexported name
```

| O que o compilador faz | Resultado |
|---|---|
| Verifica se `milhasParaKm` está exportada | Não — começa com minúscula |
| Tenta referenciar do pacote externo | Erro de compilação: `cannot refer to unexported name conversor.milhasParaKm` |
| O programa **não compila** | Nunca chega a rodar |

---

## Exercícios — Trace de Visibilidade

### exe02 — O Segredo

```
// No pacote geometria:
var numeroMagico = 42       // minúscula → privado
func PegarSegredo() int { return numeroMagico }  // maiúscula → público

// Na main:
geometria.numeroMagico      // ❌ ERRO DE COMPILAÇÃO
geometria.PegarSegredo()    // ✅ retorna 42
```

### exe03 — UUID externo

```
go get github.com/google/uuid
→ baixa o módulo, atualiza go.mod e cria go.sum

uuid.New()         // gera um UUID v4 aleatório (tipo uuid.UUID)
uuid.New().String() // converte para string: "f47ac10b-58cc-4372-a567-0e02b2c3d479"
fmt.Println(...)   // imprime o UUID — valor diferente a cada execução
```

**Saída esperada (exemplo — valor é aleatório):**
```
f47ac10b-58cc-4372-a567-0e02b2c3d479
```

---

## Resumo: o que o compilador verifica em imports

| Situação | Resultado |
|---|---|
| Import de pacote que existe na pasta | ✅ Compila |
| Import de pacote não presente e não no go.mod | ❌ Erro: `cannot find module` |
| Usar identificador exportado de outro pacote | ✅ Compila |
| Usar identificador não-exportado de outro pacote | ❌ Erro: `cannot refer to unexported name` |
| `go.mod` ausente na raiz | ❌ Erro: `go.mod file not found` |
