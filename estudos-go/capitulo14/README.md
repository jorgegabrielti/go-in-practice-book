# 📖 Capítulo 14: Pacotes e Módulos — Sua Caixa de Ferramentas

> **Livro Go na Prática: 30 capítulos para dominar a linguagem**

---

> **Fluxo de trabalho deste capítulo:**
> 1. Transcrever o conteúdo do livro para `fonte.txt` (texto bruto, fiel ao material original).
> 2. Escrever este `README.md` como síntese própria a partir do `fonte.txt`.
> 3. Resolver exemplos/exercícios em `exemplos/` e `exercicios/`.
> 4. Produzir o artigo de Medium e o post de LinkedIn em `conteudo/`, **sempre baseados no `fonte.txt`** — não em memória ou suposição do que o livro disse, e **escritos com palavras próprias** (sem copiar/parafrasear de perto o `fonte.txt`).
> 5. Rastrear a execução de todo o código do capítulo em `teste-de-mesa.md`, prevendo a saída linha a linha antes de rodar `go run`.

> 🧮 [Teste de mesa de todo o código deste capítulo](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo14/teste-de-mesa.md)

---

## 🎯 Tema Central

Pacotes organizam o código em gavetas com propósito único. Módulos são o projeto inteiro, com identidade própria. A visibilidade é controlada por uma regra visual: **letra maiúscula = público, letra minúscula = privado** — verificada pelo compilador.

---

## 📊 Resumo dos Conceitos

### 1. Pacote vs. Módulo

**Pacote** = pasta com arquivos `.go` que compartilham o mesmo `package`. É a unidade de organização.

**Módulo** = o projeto inteiro, identificado pelo `go.mod`. Pode conter vários pacotes.

```bash
go mod init github.com/seu-usuario/calculadora
# Cria go.mod — o "RG" do projeto
```

### 2. A Regra de Visibilidade

Sem `public`, `private` ou `protected`. Só a letra inicial:

```go
package financeira

func CalcularImposto() float64 { ... }  // Maiúscula → exportado ✅
func ajustarBase() float64      { ... }  // Minúscula → privado ❌ (fora do pacote)
```

Vale para funções, tipos, variáveis, constantes e campos de struct.

### 3. Criando seu Próprio Pacote

```
meu-projeto/
├── go.mod                   (module github.com/usuario/meu-projeto)
├── main.go
└── conversor/
    └── medidas.go           (package conversor)
```

```go
// conversor/medidas.go
package conversor

func KmParaMilhas(km float64) float64 { return km * 0.621371 }   // pública
func milhasParaKm(m float64) float64  { return m / 0.621371 }    // privada
```

```go
// main.go
import "github.com/usuario/meu-projeto/conversor"  // nome módulo + pasta

conversor.KmParaMilhas(100)  // ✅
conversor.milhasParaKm(62)   // ❌ erro de compilação
```

### 4. Importando Código de Terceiros (`go get`)

```bash
go get github.com/fatih/color
```

O Go baixa o código, registra no `go.mod` e cria o `go.sum` (checksum de segurança).

```go
import "github.com/fatih/color"

color.Red("Erro!")
color.Green("Sucesso!")
```

### 5. `go mod tidy` — A Faxina

```bash
go mod tidy
```

Lê o código, baixa o que falta e remove dependências que não estão mais em uso. Rodar antes de todo commit é boa prática.

---

## 💡 Dica do Gopher — Nomeie pelo domínio, não pela utilidade

```
❌ package utils   → "gaveta da bagunça"
✅ package documentos, package seguranca, package relatorio
```

Pacote bom responde "o que ele representa?", não "pra que ele serve?".

---

## 🔬 Exemplos Práticos no Repositório

### 1. [Exemplo 01: Conversor de Unidades](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo14/exemplos/exe01/main.go)
*   **Conceito:** pacote `conversor` com função pública e privada; import pelo caminho do módulo.

---

## 🔬 Exercícios Práticos Resolvidos

### 1. [Exercício 01: A Organização da Geometria](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo14/exercicios/exe01/main.go)
*   **Objetivo:** criar pacote `geometria` e importá-lo na main.
*   **Conceito:** estrutura de módulo, criação de pacote, import path.

### 2. [Exercício 02: O Segredo](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo14/exercicios/exe02/main.go)
*   **Objetivo:** demonstrar visibilidade pública vs privada na prática.
*   **Conceito:** regra da letra maiúscula/minúscula, getter para estado privado.

### 3. [Exercício 03: A Biblioteca Externa](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo14/exercicios/exe03/main.go)
*   **Objetivo:** instalar e usar `github.com/google/uuid` via `go get`.
*   **Conceito:** `go get`, `go.mod`, `go.sum`, ecossistema de pacotes Go.

---

## ✅ Checklist antes de marcar como concluído

- [x] `fonte.txt` com a transcrição completa do capítulo
- [x] Teoria revisada e resumida neste README (com base no `fonte.txt`)
- [x] Todos os exemplos do capítulo têm pasta própria em `exemplos/exe01`
- [x] Todos os exercícios resolvidos têm pasta própria em `exercicios/exe01`, `exe02`, `exe03`
- [x] `go build ./...` e `go vet ./...` passam sem erros
- [x] `go fmt ./...` executado
- [x] Termos novos adicionados ao `estudos-go/GLOSSARIO.md`
- [x] Painel de progresso no `README.md` raiz atualizado
- [x] Artigo de Medium criado em `conteudo/medium/capitulo14.md`
- [x] Post de LinkedIn criado em `conteudo/linkedin/capitulo14.md`
- [x] Pilha de exercícios extras criada em `pilha-exercicios/capitulo14.md`
- [x] Teste de mesa criado em `teste-de-mesa.md`
- [x] Fonte NotebookLM criada em `notebooklm/capitulo14.md`
