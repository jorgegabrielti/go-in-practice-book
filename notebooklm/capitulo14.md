# Capítulo 14 — Pacotes e Módulos: Sua Caixa de Ferramentas

## Tema Central

Pacotes são a unidade de organização de código em Go — uma pasta com arquivos `.go` de propósito comum. Módulos são projetos completos com identidade própria, gerenciados pelo `go.mod`. A visibilidade é controlada por uma única convenção: a letra inicial do identificador.

---

## O que o Livro Cobre

- Analogia da oficina bagunçada — todo código em `main.go` vs. gavetas organizadas
- Pacote = pasta com arquivos `.go`; Módulo = o projeto inteiro
- `go mod init` — cria o `go.mod`, dá nome e identidade ao módulo
- Regra de visibilidade: Maiúscula = exportado, Minúscula = privado
- Criando um pacote próprio: pasta `matematica/` com `package matematica`
- Import pelo caminho completo: `"nome-do-modulo/nome-da-pasta"`
- `go get` — baixar dependências externas, atualiza `go.mod` e cria `go.sum`
- `go.sum` — checksum criptográfico de segurança das dependências
- `go mod tidy` — sincroniza `go.mod` com o código real
- Dica: evitar `package utils` — preferir nomes semânticos
- Exemplos: `conversor/medidas.go` com `KmParaMilhas` (público) e `milhasParaKm` (privado)
- 3 exercícios: A Organização da Geometria, O Segredo, A Biblioteca Externa (uuid)

---

## Aprofundando os Conceitos

### Como o Go resolve imports internamente

Quando você escreve `import "github.com/usuario/projeto/conversor"`, o Go faz:

1. Lê o `go.mod` para identificar o nome do módulo raiz
2. Mapeia o caminho de import para o sistema de arquivos (subtrai o nome do módulo, adiciona o diretório raiz)
3. Compila os arquivos `.go` na pasta correspondente
4. Verifica que o `package` declarado no topo dos arquivos bate com o nome da pasta

```
Import: "github.com/joao/calc/matematica"
Módulo: "github.com/joao/calc" (do go.mod)
Pasta:  ./matematica/
```

### O que o `go.mod` contém

```
module github.com/seu-usuario/calculadora

go 1.21

require (
    github.com/fatih/color v1.15.0
    github.com/google/uuid v1.4.0
)
```

- `module`: nome canônico do módulo (usado nos imports)
- `go`: versão mínima de Go requerida
- `require`: dependências externas com versão exata (semver)

### O que o `go.sum` faz

```
github.com/fatih/color v1.15.0 h1:kOqh6YHvy4Ph1WmfklE/EKg0mCVEoVECFMlefW6TBHE=
github.com/fatih/color v1.15.0/go.mod h1:0h5ZqXfHYED7Bhv2ZJamyIOUej8KcI2cs/...
```

Cada linha é um hash criptográfico (SHA-256) do conteúdo da dependência. Quando você faz `go build`, o Go verifica os hashes — se alguém modificou o código da biblioteca (supply chain attack), o build falha. É a garantia de reprodutibilidade.

### Pacotes internos: `internal/`

Go tem uma convenção especial: se você cria uma pasta chamada `internal/`, seu conteúdo só pode ser importado por código dentro do mesmo módulo (ou da subárvore acima). É uma forma de criar APIs privadas dentro de um módulo grande:

```
meu-modulo/
├── go.mod
├── main.go
├── api/          → qualquer módulo pode importar
└── internal/
    └── db/       → só código dentro de meu-modulo/ pode importar
```

```go
// Dentro de meu-modulo/main.go:
import "meu-modulo/internal/db"  // ✅ funciona

// De outro módulo externo:
import "meu-modulo/internal/db"  // ❌ erro: use of internal package
```

### A função `init()`

Todo pacote pode ter uma função `init()` — sem parâmetros, sem retorno. O Go a chama automaticamente ao importar o pacote, antes do `main()`:

```go
package config

var ambiente string

func init() {
    ambiente = os.Getenv("APP_ENV")
    if ambiente == "" {
        ambiente = "desenvolvimento"
    }
}

func Ambiente() string { return ambiente }
```

Regras do `init`:
- Um arquivo pode ter múltiplos `init()`
- Executam na ordem de declaração
- Não pode ser chamado explicitamente
- Executa após todas as variáveis do pacote serem inicializadas

### Versões semânticas (semver) no Go Modules

```bash
go get github.com/spf13/cobra@v1.8.0   # versão específica
go get github.com/spf13/cobra@latest    # versão mais recente
go get github.com/spf13/cobra@main      # branch específico
```

Go adota semver: `MAJOR.MINOR.PATCH`
- `PATCH`: correção de bug — compatível
- `MINOR`: nova funcionalidade — compatível
- `MAJOR`: quebra de API — import path muda! (`/v2`, `/v3`)

```go
// v1
import "github.com/spf13/cobra"
// v2 (quebra de API)
import "github.com/spf13/cobra/v2"
```

### A diretiva `replace` no go.mod

Útil para apontar para uma versão local de um módulo, sem precisar publicá-lo:

```
replace github.com/seu-usuario/biblioteca => ../biblioteca
```

Cenários comuns: desenvolvimento simultâneo de módulos relacionados, debugging de dependências, forks locais.

---

## Referências Oficiais

- **Go Modules Reference**: https://go.dev/ref/mod
- **Using Go Modules (tutorial)**: https://go.dev/blog/using-go-modules
- **Module-aware commands**: https://go.dev/ref/mod#go-commands
- **Especificação — Packages**: https://go.dev/ref/spec#Packages
- **Especificação — Exported identifiers**: https://go.dev/ref/spec#Exported_identifiers
- **Effective Go — Package names**: https://go.dev/doc/effective_go#package-names
- **Tour de Go — Packages**: https://go.dev/tour/basics/1
- **pkg.go.dev** — repositório público de módulos Go: https://pkg.go.dev

---

## Exemplos de Código Adicionais

### Pacote com múltiplos arquivos

Um pacote pode ter vários arquivos `.go` — todos com o mesmo `package`:

```
banco/
├── conta.go      (package banco)
├── transacao.go  (package banco)
└── relatorio.go  (package banco)
```

Todos os identificadores de todos os arquivos compartilham o mesmo espaço de nomes dentro do pacote. Isso permite dividir um pacote grande em arquivos menores sem criar subpacotes.

### Import com alias

```go
import (
    "fmt"
    meulog "github.com/usuario/projeto/logger"   // alias
    _ "github.com/lib/pq"                         // import só pelo efeito do init()
    . "github.com/usuario/projeto/shortcuts"      // importa no namespace local (evite)
)
```

- **Alias**: troca o nome local do pacote
- **`_`**: importa apenas para executar o `init()` (efeito colateral)
- **`.`**: importa todos os identificadores no namespace atual — evite, dificulta leitura

### Estrutura de projeto típica em Go

```
meu-sistema/
├── go.mod
├── go.sum
├── cmd/
│   └── server/
│       └── main.go     ← ponto de entrada
├── internal/
│   ├── banco/
│   │   └── conta.go
│   └── config/
│       └── config.go
└── pkg/
    └── conversor/      ← API pública reutilizável
        └── medidas.go
```

Convenção: `cmd/` para pontos de entrada, `internal/` para código privado ao módulo, `pkg/` para bibliotecas reutilizáveis publicamente.

---

## Perguntas & Respostas Frequentes

**P: O nome da pasta tem que ser igual ao `package` declarado no arquivo?**
R: Por convenção, sim — e o Go linter avisa quando diferem. A exceção mais comum é `package main` (que pode estar em qualquer pasta) e testes (`package foo_test`). Evite divergências — confunde quem lê o código.

**P: Posso ter dois arquivos com `package main` em pastas diferentes?**
R: Sim! Cada pasta com `package main` é um executável separado. É exatamente por isso que projetos com múltiplos binários usam a estrutura `cmd/bin1/main.go`, `cmd/bin2/main.go`.

**P: `go get` vs `go install` — qual a diferença?**
R: `go get` adiciona/atualiza uma dependência no `go.mod` do módulo atual. `go install` compila e instala um binário no `$GOPATH/bin` — útil para ferramentas CLI (como `golint`, `gopls`).

**P: Como funciona o cache de módulos do Go?**
R: Go mantém um cache global em `$GOPATH/pkg/mod/`. Uma vez baixado, o módulo fica disponível para todos os projetos — sem re-download. O cache é somente-leitura para garantir integridade.

**P: O `go.sum` deve ser commitado?**
R: Sim, sempre. Ele garante que todos na equipe e no CI usam exatamente as mesmas versões das dependências, prevenindo ataques de supply chain e garantindo builds reproduzíveis.

---

## Comparações com Outras Linguagens

### Sistema de módulos: Go vs Node vs Python vs Java

```javascript
// Node.js — package.json + node_modules
{
  "name": "meu-projeto",
  "dependencies": { "express": "^4.18.0" }
}
// npm install  →  node_modules/ (pode ter gigabytes)
// Go equivalent: go.mod + cache global compartilhado (sem node_modules)
```

```python
# Python — requirements.txt ou pyproject.toml
# pip install -r requirements.txt
# Ambiente virtual necessário para isolamento
# Go: sem ambientes virtuais — go.mod basta
```

```java
// Java — pom.xml (Maven) ou build.gradle
// Sistemas complexos, geração de código, plugins
// Go: go.mod é um arquivo de texto simples
```

```go
// Go — go.mod
module github.com/usuario/projeto
go 1.21
require github.com/fatih/color v1.15.0
// Simples, legível, sem XML ou JSON
```

### Visibilidade: Go vs outros

| Linguagem | Mecanismo | Verificação |
|---|---|---|
| Go | Maiúscula/Minúscula | Compilação |
| Java | `public`/`private`/`protected` | Compilação |
| Python | `_prefixo` (convenção) | Runtime (não forçado) |
| JavaScript | Closures / `#campo` privado | Runtime / Compilação (TS) |
| C# | `public`/`private`/`internal` | Compilação |

---

## Armadilhas Comuns

### 1. `package` diferente da pasta

```go
// Arquivo: matematica/somar.go
package calc  // ❌ convenção quebrada — deveria ser 'matematica'
```

### 2. Import circular

```go
// pacote A importa B, B importa A → erro de compilação
// Solução: extrair a dependência compartilhada para um terceiro pacote C
```

### 3. Esquecer de rodar `go mod tidy`

```go
// Você usa um pacote mas não fez go get
// go build → erro: no required module provides package ...
// Solução: go mod tidy
```

### 4. `package utils` — gaveta da bagunça

```go
// ❌ Tudo em utils
package utils
func ValidarCPF() {}
func FormatarData() {}
func HashSenha() {}

// ✅ Pacotes semânticos
package documentos  // ValidarCPF()
package tempo       // FormatarData()
package seguranca   // HashSenha()
```

---

## Quiz de Fixação

1. Qual a diferença entre pacote e módulo em Go?
2. O que o `go.mod` registra? E o `go.sum`?
3. Como Go determina se um identificador é exportado?
4. Como você importa um pacote que criou na pasta `conversor/` do seu módulo?
5. O que `go mod tidy` faz exatamente?
6. Para que serve a pasta `internal/` em Go?
7. Qual a diferença entre `go get` e `go install`?
8. Por que o `go.sum` deve ser commitado junto com o `go.mod`?

---

## 🔬 Dissecando a Sintaxe

### `go.mod` — anatomia linha a linha

```
module  github.com/seu-usuario/calculadora
──────  ─────────────────────────────────
  │              └─ caminho canônico do módulo
  │                 usado nos imports: "github.com/seu-usuario/calculadora/pacote"
  │                 não precisa existir no GitHub — é só um identificador único
  └─ palavra-chave: declara o nome do módulo

go  1.21
──  ────
 │    └─ versão mínima de Go requerida para compilar este módulo
 └─ palavra-chave

require (
───────
  └─ bloco de dependências externas

    github.com/fatih/color  v1.15.0
    ──────────────────────  ───────
             │                 └─ versão exata (semver)
             └─ caminho do módulo externo (mesma sintaxe do module acima)
)
```

### A regra de visibilidade — exemplos visuais

```go
// ✅ EXPORTADOS — acessíveis de fora do pacote
type  Conta    struct { ... }    // [1] tipo
func  Depositar(v float64)       // [2] função
var   LimiteMaximo = 10000.0     // [3] variável
const TaxaJuros   = 0.05         // [4] constante

// ❌ NÃO-EXPORTADOS — invisíveis fora do pacote
type  historico   struct { ... }  // [1] tipo interno
func  calcularFee(v float64)      // [2] detalhe de implementação
var   contador    = 0             // [3] estado interno
```

```
[1] Conta  vs  historico
     │              │
     └─ 'C' maiúsculo    └─ 'h' minúsculo
        → exportado          → privado
        banco.Conta          erro de compilação
```

### Import — dissecar o caminho completo

```go
import "github.com/seu-usuario/calculadora/matematica"
        ─────────────────────────────────  ──────────
                      │                        └─ nome da pasta (= nome do package)
                      └─ nome do módulo (do go.mod)
                         O Go subtrai o módulo do import para achar a pasta:
                         "calculadora/matematica" → ./matematica/

// Usar:
matematica.Somar(10, 20)
──────────  ─────
     │           └─ função exportada
     └─ último segmento do import path = nome local do pacote
```

---

## 🎙️ Prompts para o NotebookLM

### 🎧 Audio Overview
> "Gere um Audio Overview sobre Pacotes e Módulos em Go para um desenvolvedor vindo de Node.js. Um apresentador fica surpreso que não há `node_modules`. O outro explica o cache global, como o `go.mod` é mais simples que `package.json`, e por que a regra de maiúscula/minúscula substitui `public`/`private`. Inclua a armadilha do `package utils`."

### 📋 Briefing Doc
> "Crie um briefing sobre 'O ciclo de vida de uma dependência em Go': do `go get` ao `go.sum`, passando pelo cache global, verificação de integridade, e como `go mod tidy` mantém tudo sincronizado."

### 📚 Study Guide
> "Crie um guia de estudos do Capítulo 14 com: (1) tabela comparando sistemas de módulos: Go vs npm vs pip vs Maven, (2) exercício prático: criar um módulo do zero com dois pacotes internos e uma dependência externa, (3) quiz de 8 perguntas sobre visibilidade e módulos."

### 🖼️ Roteiro de Slides
> "Crie um roteiro de 6 slides sobre 'Pacotes e Módulos em Go'. Slide 1: A gaveta certa — analogia da oficina. Slide 2: `go.mod` — o RG do projeto. Slide 3: Maiúscula vs Minúscula — a regra de visibilidade. Slide 4: Criando seu próprio pacote. Slide 5: `go get` e o ecossistema externo. Slide 6: `go mod tidy` e boas práticas de nomenclatura."

### 💬 Perguntas Profundas para o Chat
- "Como o Go resolve o caminho de um import para uma pasta no sistema de arquivos? Me explique o processo completo do `go.mod` até encontrar o arquivo `.go`."
- "O que é o `go.sum` e por que ele deve ser commitado? O que acontece se alguém modifica uma dependência externa depois que você fez o download?"
- "Qual a diferença entre `internal/` e visibilidade por minúscula em Go? Quando usar cada um?"
- "Por que `package utils` é considerado um anti-pattern em Go? Me dê um exemplo de refatoração de um `utils` para pacotes semânticos."
