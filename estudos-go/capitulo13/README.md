# 📖 Capítulo 13: Tratamento de Erros — Sem Pânico!

> **Livro Go na Prática: 30 capítulos para dominar a linguagem**

---

> **Fluxo de trabalho deste capítulo:**
> 1. Transcrever o conteúdo do livro para `fonte.txt` (texto bruto, fiel ao material original).
> 2. Escrever este `README.md` como síntese própria a partir do `fonte.txt`.
> 3. Resolver exemplos/exercícios em `exemplos/` e `exercicios/`.
> 4. Produzir o artigo de Medium e o post de LinkedIn em `conteudo/`, **sempre baseados no `fonte.txt`** — não em memória ou suposição do que o livro disse, e **escritos com palavras próprias** (sem copiar/parafrasear de perto o `fonte.txt`).
> 5. Rastrear a execução de todo o código do capítulo em `teste-de-mesa.md`, prevendo a saída linha a linha antes de rodar `go run`.

> 🧮 [Teste de mesa de todo o código deste capítulo](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo13/teste-de-mesa.md)

---

## 🎯 Tema Central

Go trata erros como **valores comuns** — não como explosões. Enquanto Java lança uma exceção que interrompe tudo, Go retorna um `error` junto com o resultado normal. Você decide o que fazer com ele. Isso é "Errors as Values".

---

## 📊 Resumo dos Conceitos

### 1. O tipo `error`

`error` é uma interface builtin com um único método:

```go
type error interface {
    Error() string
}
```

Qualquer tipo que implemente `Error() string` é um erro. O padrão idiomático é retornar dois valores — resultado e erro:

```go
func AbrirArquivo(nome string) (Arquivo, error)

arquivo, err := AbrirArquivo("dados.txt")
if err != nil {
    fmt.Println("Problema:", err)
    return
}
// caminho feliz continua aqui
```

### 2. Criando erros

**`errors.New`** — para mensagens fixas:
```go
var ErrSemGasolina = errors.New("tanque vazio")
```

**`fmt.Errorf`** — para mensagens com contexto dinâmico. O verbo `%w` "embrulha" (wrap) um erro existente, permitindo inspeção posterior com `errors.Is` e `errors.As`:
```go
return fmt.Errorf("falha ao logar %s: %w", usuario, ErrSenhaInvalida)
```

### 3. Sentinel Errors

Variáveis de erro pré-declaradas que podem ser comparadas com `errors.Is`:

```go
var ErrDivisaoPorZero = errors.New("matemática diz não: divisão por zero")

func dividir(a, b float64) (float64, error) {
    if b == 0 {
        return 0, ErrDivisaoPorZero
    }
    return a / b, nil
}

if errors.Is(err, ErrDivisaoPorZero) {
    // tratamento específico
}
```

### 4. `panic` — A Opção Nuclear

Para situações **genuinamente irrecuperáveis** (bug lógico grave, estado corrompido):

```go
panic("EXPLOSÃO NO REATOR 4!")
```

**Regra de ouro:** nunca use `panic` para controle de fluxo normal (arquivo não encontrado, usuário inválido). Para esses casos, retorne `error`.

### 5. `recover` — O Bombeiro

Captura um `panic` antes que ele derrube o programa. **Só funciona dentro de um `defer`**:

```go
defer func() {
    if r := recover(); r != nil {
        fmt.Println("Panic capturado:", r)
    }
}()
```

Uso típico: servidores web que não podem derrubar o processo inteiro por causa de uma requisição ruim.

### 6. `defer` — A Limpeza Garantida

Agenda a execução de uma função para o **último momento antes da função atual retornar**, não importa como ela retorne (return normal, return com erro, panic):

```go
func LerArquivo() {
    arquivo := abrir("dados.txt")
    defer arquivo.Fechar()  // executará sempre, não importa o que aconteça abaixo

    processar(arquivo)
}
```

Múltiplos `defer` funcionam como uma **pilha LIFO**: o último registrado é o primeiro a executar.

---

## 💡 Dica do Gopher — Line of Sight

Evite o `else` no tratamento de erros. O padrão idiomático do Go é o "Happy Path na esquerda":

```go
// ❌ Iniciante — else aninhado
if err != nil {
    tratarErro()
} else {
    continuar()
}

// ✅ Idiomático — Line of Sight
if err != nil {
    return err
}
continuar()
```

Tratar o erro e retornar cedo mantém o código principal sem aninhamento.

---

## 🔬 Exemplos Práticos no Repositório

### 1. [Exemplo 01: Tratamento Completo](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo13/exemplos/exe01/main.go)
*   **Conceito:** sentinel error, `errors.Is`, panic + recover com defer, defer LIFO.

---

## 🔬 Exercícios Práticos Resolvidos

### 1. [Exercício 01: A Compra Segura](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo13/exercicios/exe01/main.go)
*   **Objetivo:** retornar erro quando saldo é insuficiente.
*   **Conceito:** `errors.New`, múltiplos retornos, padrão `if err != nil`.

### 2. [Exercício 02: O Validador de Senha](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo13/exercicios/exe02/main.go)
*   **Objetivo:** validar senha com múltiplas regras, retornando erro descritivo.
*   **Conceito:** `fmt.Errorf`, `errors.New`, funções auxiliares.

### 3. [Exercício 03: A Pilha de Pratos](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo13/exercicios/exe03/main.go)
*   **Objetivo:** observar e explicar a ordem LIFO dos defers em loop.
*   **Conceito:** `defer` em loop, pilha LIFO, ordem de execução.

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
- [x] Artigo de Medium criado em `conteudo/medium/capitulo13.md`
- [x] Post de LinkedIn criado em `conteudo/linkedin/capitulo13.md`
- [x] Pilha de exercícios extras criada em `pilha-exercicios/capitulo13.md`
- [x] Teste de mesa criado em `teste-de-mesa.md`
- [x] Fonte NotebookLM criada em `notebooklm/capitulo13.md`
