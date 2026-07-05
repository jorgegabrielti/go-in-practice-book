# Capítulo 01 — Introdução ao Go: A Filosofia da Linguagem

## Tema Central

Go (também chamada de Golang) foi criada em 2007 por engenheiros do Google — Robert Griesemer, Rob Pike e Ken Thompson — e tornou-se open-source em 2009. A motivação foi resolver problemas reais de escala: builds lentos, dependências obscuras, inconsistência de estilo entre equipes, e código difícil de ler. O resultado é uma linguagem que preza por **clareza acima de esperteza**.

---

## O que o Livro Cobre

- Go é compilada para binário nativo (sem interpretador em produção)
- `package main` e `func main()` como ponto de entrada
- Exportação por maiúscula (letra maiúscula = público)
- `go fmt` como formatador oficial
- Filosofia "Less is exponentially more"
- Apenas 25 palavras-chave

---

## Aprofundando os Conceitos

### Compilado vs. Interpretado vs. JVM

| Característica | Go | Python | Java |
|---|---|---|---|
| Modelo | Compilado para binário nativo | Interpretado (bytecode + CPython) | Compilado para bytecode (JVM) |
| Precisa de runtime? | Não (embute o runtime no binário) | Sim (Python instalado) | Sim (JVM instalada) |
| Velocidade de startup | Muito rápida | Lenta (inicializar interpretador) | Lenta (inicializar JVM) |
| Deploy | Copiar o binário | Copiar código + dependências | Copiar .jar + JVM |
| Cross-compile | `GOOS=linux GOARCH=amd64 go build` | N/A | Via JVM (escreva uma vez, rode em qualquer lugar) |

O binário Go é **autocontido**: inclui o garbage collector, o scheduler de goroutines e todas as dependências. Você copia um único arquivo para o servidor e roda. Sem `apt install python3`, sem `java -jar`.

### As 25 Palavras-Chave de Go

```
break        default      func         interface    select
case         defer        go           map          struct
chan         else         goto         package      switch
const        fallthrough  if           range        type
continue     for          import       return       var
```

Compare com Java (50+), C++ (86+) ou Python (35). Menos palavras-chave significa menos coisas para aprender e menos ambiguidade ao ler código de outra pessoa.

### Por que o Go não tem `while`?

Uma decisão deliberada de design. O `for` do Go cobre todos os casos:

```go
// for clássico (C-style)
for i := 0; i < 10; i++ { }

// for como while
for condicao { }

// loop infinito
for { }

// range (iterator)
for i, v := range slice { }
```

Uma única palavra para quatro padrões. Menos surpresas, mais consistência.

### Por que a chave `{` na mesma linha?

Não é gosto do designer — é uma consequência técnica. O Go insere ponto-e-vírgula automaticamente no fim de certas linhas durante a análise léxica (Automatic Semicolon Insertion). Se você colocar `{` na linha de baixo:

```go
// ERRO DE COMPILAÇÃO
func main()
{
    fmt.Println("Olá")
}
```

O lexer insere `;` depois de `main()`, quebrando a sintaxe. A chave obrigatoriamente na mesma linha não é opinião — é gramática da linguagem.

### `go fmt`: por que isso importa?

Em projetos com muitas pessoas, debates sobre estilo consomem tempo e energia. Go elimina a discussão: há um único estilo oficial, aplicado automaticamente. Código Go de qualquer empresa, qualquer projeto, qualquer país tem a mesma aparência. Isso reduz a carga cognitiva ao ler código alheio.

Ferramentas como `goimports` (extensão de `go fmt` que também organiza imports) são integradas a todos os editores populares.

### O Toolchain Go

```bash
go build    # compila o pacote atual
go run      # compila + executa (atalho para desenvolvimento)
go test     # roda testes (*_test.go)
go fmt      # formata o código
go vet      # analisa padrões problemáticos (sem compilar)
go mod      # gerencia módulos e dependências
go install  # compila + instala binário em $GOPATH/bin
go doc      # mostra documentação de um pacote/função
```

---

## Referências Oficiais

- **Página inicial do Go**: https://go.dev/
- **Por que Go foi criada (blog oficial)**: https://go.dev/blog/go-brand
- **FAQ oficial da linguagem**: https://go.dev/doc/faq
- **Histórico de versões**: https://go.dev/doc/devel/release
- **Especificação da linguagem (o documento definitivo)**: https://go.dev/ref/spec
- **Effective Go (guia de boas práticas)**: https://go.dev/doc/effective_go
- **Tour interativo de Go**: https://go.dev/tour/
- **Go Playground (rodar código no browser)**: https://go.dev/play/

---

## Exemplos de Código Adicionais

### Hello, World — o clássico

```go
package main

import "fmt"

func main() {
    fmt.Println("Olá, Gopher!")
}
```

### Cross-compilation — compilar para Linux no Windows

```bash
GOOS=linux GOARCH=amd64 go build -o app-linux ./...
```

Isso gera um binário `app-linux` executável em um servidor Linux, compilado na sua máquina Windows ou Mac. Sem Docker para cross-compile.

### Múltiplos arquivos no mesmo pacote

```go
// arquivo: saudacao.go
package main

func saudacao(nome string) string {
    return "Olá, " + nome + "!"
}
```

```go
// arquivo: main.go
package main

import "fmt"

func main() {
    fmt.Println(saudacao("Jorge"))
}
```

Ambos são `package main`. O `go build` os compila juntos automaticamente.

---

## Perguntas & Respostas Frequentes

**P: Por que o Go não tem classes?**
R: Go usa Structs e Interfaces em vez de classes. A filosofia é preferir composição sobre herança. Isso evita as hierarquias profundas de classes que tornam o código Java e C++ difíceis de manter. Veja capítulo 10 (Structs) e capítulo 11 (Métodos).

**P: Go é orientada a objetos?**
R: Depende da definição. Go tem tipos, métodos (funções associadas a tipos) e interfaces (polimorfismo), mas não tem herança de classe. O próprio FAQ do Go responde: "Yes and no." Veja: https://go.dev/doc/faq#Is_Go_an_object-oriented_language

**P: Por que importar um pacote e não usar gera erro de compilação?**
R: Importar sem usar é sintoma de código sujo. O compilador Go se recusa a compilar código com imports não usados para forçar higiene desde o início. Mesmo comportamento com variáveis declaradas e não usadas (exceto com `_`).

**P: O que é `GOPATH` e `GOROOT`?**
R: `GOROOT` é onde o Go está instalado (binários do compilador). `GOPATH` era onde ficavam seus projetos e dependências antes dos Módulos Go (go.mod). Desde Go 1.11+, o sistema de módulos tornou o GOPATH opcional para a maioria dos projetos.

**P: O que é um módulo Go?**
R: Um módulo é um conjunto de pacotes Go com um arquivo `go.mod` na raiz, definindo o nome do módulo e as dependências. É o sistema de gerenciamento de dependências moderno do Go.

---

## Comparações com Outras Linguagens

### Exportação: Go vs Python vs Java vs JavaScript

```python
# Python: convenção (underscore = privado, mas não enforçado)
def _privado():   # "privado" por convenção
    pass
def publico():    # público por convenção
    pass
```

```java
// Java: keywords explícitas
private void privado() { }
public void publico() { }
```

```javascript
// JavaScript: sem controle de acesso real em objetos simples
// ES2022+: # prefix para campos privados em classes
class Exemplo {
  #privado = 1;
  publico = 2;
}
```

```go
// Go: letra maiúscula = exportado (público), minúscula = não exportado (privado ao pacote)
func privado() { }   // só acessível dentro do pacote
func Publico() { }   // acessível de outros pacotes
```

A convenção Go é **enforçada pelo compilador** — não é apenas estilo, é sintaxe.

---

## Armadilhas Comuns

1. **`var` vs `:=` fora de funções**: `:=` só funciona dentro de funções. Em nível de pacote, use `var`.

2. **Import circular**: dois pacotes que se importam mutuamente causam erro de compilação. A solução é extrair a dependência comum para um terceiro pacote.

3. **`go run` vs `go build`**: `go run main.go` não deixa binário em disco. Use `go build` quando quiser distribuir.

4. **Package name ≠ folder name**: o nome do pacote (declarado no topo do arquivo) e o nome da pasta podem ser diferentes — mas por convenção devem coincidir (exceto `package main`).

---

## Quiz de Fixação

1. Quantas palavras-chave o Go tem? Por que isso importa?
2. O que acontece se você deixar uma variável declarada sem usar? E um import?
3. Por que a chave `{` deve estar na mesma linha em Go?
4. Qual é a diferença entre `go run` e `go build`?
5. Como o Go decide se um identificador é público ou privado?
6. O que é o `go fmt` e por que ele é importante em projetos de equipe?

---

## 🔬 Dissecando a Sintaxe

Esta seção desmonta cada pedaço do código para você entender o papel de cada elemento.

### Anatomia do programa mais simples possível

```go
package main         // [1]

import "fmt"         // [2]

func main() {        // [3] [4] [5] [6]
    fmt.Println("Olá, Gopher!")  // [7] [8] [9]
}                    // [10]
```

```
[1]  package main
      ───────  ────
         │       └─ nome do pacote: "main" é especial — diz ao compilador
         │           que este arquivo gera um binário executável (não uma biblioteca)
         └─ palavra-chave obrigatória no topo de todo arquivo .go

[2]  import "fmt"
      ──────  ───
         │      └─ nome do pacote importado (entre aspas duplas)
         │          "fmt" = Format — pacote da biblioteca padrão para I/O
         └─ palavra-chave que torna outro pacote disponível neste arquivo

[3]  func
      ────
         └─ palavra-chave que declara uma função

[4]  main
      ────
         └─ nome da função: "main" é especial — é o ponto de entrada do programa
             o sistema operacional chama exatamente esta função ao executar o binário

[5]  ()
      ──
         └─ lista de parâmetros (vazia aqui): main não recebe argumentos
             argumentos do sistema vêm via os.Args, não via parâmetros de main

[6]  {
      ─
         └─ abre o bloco de código da função
             DEVE estar na mesma linha que "func main()" — regra gramatical do Go

[7]  fmt
      ───
         └─ nome do pacote — o mesmo que importamos em [2]
             é como um "namespace": agrupa funções relacionadas

[8]  .Println
      ────────
         └─ "." acessa algo dentro do pacote fmt
             "Println" = Print + line: imprime e adiciona quebra de linha no final
             "P" maiúscula = função exportada (pública), acessível de fora do pacote

[9]  ("Olá, Gopher!")
      ─────────────────
         └─ argumento da função: um valor do tipo string (texto entre aspas duplas)
             Println é variádica: aceita zero ou mais argumentos de qualquer tipo

[10] }
      ─
         └─ fecha o bloco da função main
```

### Exportação por maiúscula vs minúscula

```go
//     [1]         [2]
fmt.Println("visível de fora")   // maiúscula = exportada
fmt.println("x")                 // ERRO: 'println' não existe em fmt (minúscula = inexistente)

//  [3]
var minhaVar = 10   // minúscula = privada ao pacote atual
var MinhaVar = 10   // maiúscula = exportada (pública para outros pacotes)
```

```
[1] fmt.Println  → "P" maiúsculo: função EXPORTADA, acessível de qualquer pacote
[2] fmt.println  → "p" minúsculo: seria privada ao pacote fmt (não existe na API pública)
[3] var minhaVar → variável com "m" minúsculo: privada ao pacote onde foi declarada
```

---

## 🎙️ Prompts para o NotebookLM

### 🎧 Audio Overview (Podcast)
Clique em **"Generate"** no painel Audio Overview, ou use este prompt de contexto antes:

> "Gere um Audio Overview com dois apresentadores: um veterano em Go e um iniciante vindo do Python. O iniciante faz perguntas sobre por que o Go funciona diferente do Python — sem classes, sem null, sem while — e o veterano explica a filosofia por trás de cada escolha de design, usando exemplos simples do cotidiano."

### 📋 Briefing Doc
> "Crie um briefing de meia página sobre a filosofia de design do Go, cobrindo: (1) por que foi criado, (2) as 3 principais diferenças em relação a Python e Java, (3) o que 'Less is exponentially more' significa na prática."

### 📚 Study Guide
> "Crie um guia de estudos do Capítulo 01 com: definição das 5 palavras-chave mais importantes (package, import, func, var, const), 8 perguntas de verdadeiro/falso com resposta e explicação, e 3 perguntas dissertativas para reflexão."

### 🖼️ Roteiro de Slides
> "Crie um roteiro de 6 slides para apresentar 'Introdução ao Go' para um desenvolvedor Python. Slide 1: O que é Go e por que existe. Slide 2: Diferenças filosóficas vs Python. Slide 3: Estrutura de um programa Go (package/import/func). Slide 4: Toolchain (go run, go build, go fmt). Slide 5: Exportação por maiúscula. Slide 6: Próximos passos."

### 💬 Perguntas Profundas para o Chat
Cole estas perguntas diretamente no chat do NotebookLM:

- "Por que o Go foi criado em 2007 dentro do Google? Quais problemas específicos ele resolveu que C++ e Java não resolviam?"
- "Me explique a diferença entre `go run` e `go build` como se eu fosse um iniciante que só conhece Python com `python script.py`."
- "O que acontece internamente quando o Go compila um programa? Por que o binário resultante não precisa de nenhuma dependência instalada no servidor?"
- "Por que importar um pacote e não usar causa erro de compilação? Em Python isso não é erro — qual é a filosofia por trás dessa decisão do Go?"
