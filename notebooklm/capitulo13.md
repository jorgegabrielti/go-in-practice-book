# Capítulo 13 — Tratamento de Erros: Sem Pânico!

## Tema Central

Go adota a filosofia **"Errors as Values"**: erros são variáveis comuns retornadas por funções, não mecanismos de controle de fluxo especiais. Essa escolha é deliberada — torna erros explícitos, rastreáveis e parte formal do contrato de uma função.

---

## O que o Livro Cobre

- Analogia da impressora + papelzinho — Go entrega o problema, você decide o que fazer
- `error` como interface: `Error() string`
- Padrão `val, err := f()` + `if err != nil { return }`
- `errors.New` — erro com mensagem fixa
- `fmt.Errorf` com `%w` — erro com contexto + wrapping
- Sentinel errors + `errors.Is` para comparação específica
- `panic` — opção nuclear para erros irrecuperáveis
- `recover` — captura panic dentro de defer (O Bombeiro)
- `defer` — agendamento de limpeza garantida, LIFO
- Dica "Line of Sight" — Happy Path na esquerda, sem else
- `main.go`: `ErrDivisaoPorZero`, `dividir`, `operacaoArriscada` (defer+recover), demo de defer anônimo
- 3 exercícios: A Compra Segura, O Validador de Senha, A Pilha de Pratos

---

## Aprofundando os Conceitos

### Por que Go não tem exceções

A decisão foi tomada pelos criadores com base em problemas observados em linguagens com exceções:

1. **Fluxo invisível**: com `throw/catch`, o leitor não consegue saber, só lendo a assinatura da função, se ela pode falhar e de que forma.
2. **Tratamento postergado**: o `catch` incentiva tratar o erro longe da origem — o que dificulta o diagnóstico.
3. **Performance**: a criação de stack traces para exceções é cara. Erros como valores têm custo zero quando não ocorrem.

Go acredita que se algo pode falhar, isso deve aparecer explicitamente na assinatura — `(T, error)` não deixa dúvida.

### A interface `error` por dentro

```go
// Definição builtin do Go
type error interface {
    Error() string
}
```

Qualquer tipo com `Error() string` é um `error`. Você pode criar tipos de erro customizados com campos extras:

```go
type ErrValidacao struct {
    Campo  string
    Motivo string
}

func (e *ErrValidacao) Error() string {
    return fmt.Sprintf("campo '%s': %s", e.Campo, e.Motivo)
}
```

### Cadeia de Wrapping (`%w`) e `errors.Unwrap`

O verbo `%w` em `fmt.Errorf` cria uma cadeia rastreável:

```go
var ErrBase = errors.New("conexão recusada")

// Camada 1 — repositório
func buscarNoBanco() error {
    return ErrBase
}

// Camada 2 — serviço
func buscarUsuario(id int) error {
    err := buscarNoBanco()
    return fmt.Errorf("usuário %d não encontrado: %w", id, err)
}

// Camada 3 — handler
func handleRequisicao() error {
    err := buscarUsuario(42)
    return fmt.Errorf("falha na requisição: %w", err)
}

// Na main
err := handleRequisicao()
errors.Is(err, ErrBase)  // true — atravessa toda a cadeia
```

`errors.Unwrap(err)` retorna o erro embrulhado um nível abaixo. `errors.Is` e `errors.As` percorrem a cadeia inteira automaticamente.

### `errors.Is` vs `errors.As`

| Função | Uso | Quando usar |
|---|---|---|
| `errors.Is(err, target)` | Compara por **identidade** (mesmo ponteiro) | Sentinel errors: `ErrNotFound`, `ErrPermissao` |
| `errors.As(err, &target)` | Extrai um tipo específico de erro | Tipos customizados com campos extras |

```go
// errors.Is — verifica se é aquele sentinel
if errors.Is(err, ErrDivisaoPorZero) { ... }

// errors.As — extrai o tipo para acessar campos
var valErr *ErrValidacao
if errors.As(err, &valErr) {
    fmt.Println("Campo com problema:", valErr.Campo)
}
```

### `defer`: ordem de execução e captura de variáveis

```go
func exemplo() {
    // Atenção: defer captura o valor de 'i' no momento do defer
    i := 0
    defer fmt.Println("defer capturou i =", i)  // imprime 0, não 10
    i = 10
    fmt.Println("i agora é", i)
}
// Saída:
// i agora é 10
// defer capturou i = 0
```

Para capturar o valor **final** de uma variável, use um ponteiro ou uma função anônima:

```go
defer func() { fmt.Println("i final =", i) }()  // imprime 10
```

### `panic` e `recover` — contrato de uso

| Situação | Usar |
|---|---|
| Arquivo não encontrado | `error` |
| Senha incorreta | `error` |
| Índice fora do slice | `panic` (bug do código) |
| Configuração inválida na inicialização | `panic` (estado corrompido desde o início) |
| Requisição HTTP malformada | `error` |
| Assert interno que nunca deveria falhar | `panic` |

`recover` **não é catch**. Ele apenas impede que o processo morra — não repara o estado corrompido. Após um recover, o estado interno pode estar inconsistente. Use apenas para isolar falhas (servidores, workers), nunca para lógica de negócio.

### `defer` para mais do que fechar arquivos

```go
// Medir tempo de execução de qualquer função
func crono(nome string) func() {
    inicio := time.Now()
    return func() {
        fmt.Printf("%s levou %v\n", nome, time.Since(inicio))
    }
}

func minhaFuncao() {
    defer crono("minhaFuncao")()  // inicia o timer e agenda o log
    // ... código da função
}
```

```go
// Destravar mutex sempre
mu.Lock()
defer mu.Unlock()
```

```go
// Fechar body de resposta HTTP sempre
resp, err := http.Get(url)
if err != nil { return err }
defer resp.Body.Close()
```

---

## Referências Oficiais

- **Especificação — Errors**: https://go.dev/ref/spec#Errors
- **Effective Go — Errors**: https://go.dev/doc/effective_go#errors
- **Blog: Error handling and Go**: https://go.dev/blog/error-handling-and-go
- **Blog: Working with Errors in Go 1.13**: https://go.dev/blog/go1.13-errors
- **Blog: Defer, Panic, and Recover**: https://go.dev/blog/defer-panic-and-recover
- **Pacote `errors`**: https://pkg.go.dev/errors
- **`fmt.Errorf` com %w**: https://pkg.go.dev/fmt#Errorf
- **Tour de Go — Error**: https://go.dev/tour/methods/19
- **Tour de Go — Defer**: https://go.dev/tour/flowcontrol/12

---

## Exemplos de Código Adicionais

### Tipo de erro customizado completo

```go
type ErrHTTP struct {
    Codigo  int
    Metodo  string
    URL     string
    Motivo  string
}

func (e *ErrHTTP) Error() string {
    return fmt.Sprintf("HTTP %d %s %s: %s", e.Codigo, e.Metodo, e.URL, e.Motivo)
}

func fazerRequisicao(url string) error {
    // Simula uma resposta 404
    return &ErrHTTP{
        Codigo: 404,
        Metodo: "GET",
        URL:    url,
        Motivo: "recurso não encontrado",
    }
}

err := fazerRequisicao("/api/usuario/999")
var httpErr *ErrHTTP
if errors.As(err, &httpErr) {
    if httpErr.Codigo == 404 {
        fmt.Println("Recurso não existe, criando novo...")
    }
}
```

### Middleware de recover para servidor

```go
func recuperar(handler func()) {
    defer func() {
        if r := recover(); r != nil {
            // log do panic sem derrubar o servidor
            log.Printf("PANIC recuperado: %v\n%s", r, debug.Stack())
        }
    }()
    handler()
}

// Uso em servidor HTTP
http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    recuperar(func() {
        processarRequisicao(w, r)
    })
})
```

---

## Perguntas & Respostas Frequentes

**P: Por que `if err != nil` repete tanto em Go? Não é verboso demais?**
R: É intencional. A verbosidade é o custo da explicitidade. Cada `if err != nil` é uma decisão consciente: "aqui pode falhar, e eu estou tratando". A comunidade Go prefere clareza à concisão nesses casos. Propostas de sintaxe para reduzir a repetição foram discutidas na spec e rejeitadas exatamente por isso.

**P: `errors.New` vs `fmt.Errorf` — quando usar cada um?**
R: `errors.New` quando a mensagem é fixa e você vai comparar com `errors.Is` (sentinel). `fmt.Errorf` quando precisa incluir dados dinâmicos (nome de usuário, valor, etc.) ou embrulhar um erro existente com `%w`.

**P: `defer` executa se o programa receber um `os.Exit`?**
R: Não. `os.Exit` encerra o processo imediatamente, sem executar defers. É diferente de `panic`, que executa os defers antes de encerrar.

**P: Posso usar `recover` fora de `defer`?**
R: Pode chamar, mas retorna `nil`. `recover` só tem efeito quando chamado **diretamente** dentro de uma função `defer`, durante o unwinding de um `panic`.

**P: Quando usar `log.Fatal` vs retornar `error` vs `panic`?**
R: `log.Fatal` = problema na inicialização que impede o programa de funcionar (chama `os.Exit`). `panic` = bug interno que nunca deveria acontecer. `error` = todo o resto — qualquer situação esperada que pode dar errado.

---

## Comparações com Outras Linguagens

### Tratamento de erros: Go vs Java vs Python vs TypeScript

```java
// Java — exceção interrompe o fluxo
try {
    Arquivo a = AbrirArquivo("dados.txt");
    processar(a);
} catch (IOException e) {
    System.out.println("Erro: " + e.getMessage());
} finally {
    fechar();  // equivalente ao defer
}
```

```python
# Python — exceção + context manager
try:
    with open("dados.txt") as f:  # with = defer automático
        processar(f)
except FileNotFoundError as e:
    print(f"Erro: {e}")
```

```typescript
// TypeScript — exceções + Result type (funcional)
// Alternativa moderna: Result<T, E>
type Result<T, E> = { ok: true; value: T } | { ok: false; error: E };
```

```go
// Go — erro é valor, defer é explícito
f, err := os.Open("dados.txt")
if err != nil {
    return err
}
defer f.Close()  // garantia de fechamento
processar(f)
```

**Diferença chave:** em Go, olhando a assinatura `(T, error)`, você já sabe que a operação pode falhar — sem ler documentação extra. Em Java e Python, você precisa saber quais exceções uma função pode lançar (que podem não estar documentadas).

---

## Armadilhas Comuns

### 1. Ignorar o erro com `_`

```go
resultado, _ := dividir(10, 0)  // ❌ erro silenciado
// resultado será zero value (0.0) sem aviso
```

Só ignore com `_` quando você tem certeza absoluta de que o erro não pode ocorrer naquele contexto.

### 2. `recover` fora de `defer`

```go
func errado() {
    r := recover()  // ❌ sempre retorna nil aqui
    // ...
    panic("ops")
}
```

### 3. Wrapping sem `%w` (perde rastreabilidade)

```go
// ❌ embrulha mas perde o erro original
return fmt.Errorf("falha: %v", err)  // errors.Is não vai encontrar o original

// ✅ preserva o erro original na cadeia
return fmt.Errorf("falha: %w", err)
```

### 4. `defer` em loop com valor capturado errado

```go
for i := 0; i < 5; i++ {
    defer fmt.Println(i)  // captura o valor atual de i — correto neste caso
    // mas se fosse defer fmt.Println(&i), todos imprimiriam 5
}
```

---

## Quiz de Fixação

1. Qual a diferença entre `errors.New` e `fmt.Errorf`?
2. O que `%w` faz em `fmt.Errorf` que `%v` não faz?
3. Em que ordem executam múltiplos `defer` na mesma função?
4. Por que `recover` só funciona dentro de `defer`?
5. Qual a diferença entre `errors.Is` e `errors.As`?
6. O que acontece com os `defer` se o programa chamar `os.Exit`?
7. Por que Go não tem `try/catch`?
8. Quando é correto usar `panic` vs retornar `error`?

---

## 🔬 Dissecando a Sintaxe

### O padrão `val, err := f()` — anatomia completa

```go
arquivo,   err   :=  AbrirArquivo( "dados.txt" )
────────   ───   ──  ────────────  ────────────
   │         │    │       │              └─ argumento da função
   │         │    │       └─ função que retorna (T, error)
   │         │    └─ ":=" declara AMBAS as variáveis (arquivo e err) como novas
   │         └─ segundo retorno: sempre do tipo error (interface)
   │              nil  = nenhum problema, pode usar 'arquivo'
   │              !nil = algo deu errado, 'arquivo' pode ser zero value
   └─ primeiro retorno: o valor útil (se err != nil, pode ser inválido)

if err != nil {
── ───  ──  ───
 │  │    │   └─ nil = zero value da interface error (tipo=nil, valor=nil)
 │  │    └─ operador "diferente de"
 │  └─ a variável err que veio da linha acima
 └─ guarda: se entrar aqui, tem problema — tratar e retornar

    return  // [1]
    ──────
      └─ encerra a função atual imediatamente
         deixa o "caminho feliz" fora do if (sem else)
}
// [1] Se chegou aqui: err É nil — 'arquivo' está disponível com segurança
```

### `defer` — quando exatamente executa

```go
func LerArquivo() {
    arquivo := abrir("dados.txt")  // [1] abre agora
    defer arquivo.Fechar()         // [2] AGENDA para depois — não executa agora
    ─────
      └─ "adiar" — a chamada é registrada na pilha de defers da função

    processar(arquivo)             // [3] executa agora
    return                         // [4] return dispara os defers (LIFO)
                                   //     arquivo.Fechar() executa AQUI
}
```

```
Linha do tempo:
[1] abrir()       → arquivo disponível
[2] defer         → "Fechar() vai rodar no fim" (registrado, não executado)
[3] processar()   → usa arquivo normalmente
[4] return        → agora sim: defer executa → arquivo.Fechar()
```

### `panic` + `recover` — o fluxo exato

```go
func operacaoArriscada() {
    defer func() {                    // [1] registra o defer
        if r := recover(); r != nil { // [4] recover() captura o valor do panic
            fmt.Println("capturei:", r)
        }
    }()                               // () — executa imediatamente como defer

    fmt.Println("antes")             // [2] executa normalmente
    panic("BOOM")                    // [3] inicia unwinding
    fmt.Println("nunca")             // NUNCA EXECUTADO
}

// Fluxo:
// [1] defer func registrado na pilha
// [2] "antes" impresso
// [3] panic → para execução normal → inicia defers
// [4] defer func executa → recover() pega "BOOM" → imprime
// função retorna normalmente (panic foi interceptado)
```

---

## 🎙️ Prompts para o NotebookLM

### 🎧 Audio Overview
> "Gere um Audio Overview sobre Tratamento de Erros em Go para um desenvolvedor que vem de Java. Um apresentador fica surpreso que não há `try/catch`. O outro explica por que erros como valores são mais previsíveis, como `defer` substitui `finally`, e quando usar `panic` vs `error`. Inclua a dica de Line of Sight."

### 📋 Briefing Doc
> "Crie um briefing sobre as '4 ferramentas de erros do Go': `error` (o valor), `errors.Is/As` (a inspeção), `panic/recover` (o emergencial), `defer` (a limpeza). Para cada uma: quando usar, quando NÃO usar, e armadilha comum."

### 📚 Study Guide
> "Crie um guia de estudos do Capítulo 13 com: (1) tabela comparando `errors.New` vs `fmt.Errorf` vs tipo customizado, (2) fluxograma de decisão: erro → retornar error ou panic?, (3) exercício: dado um código com `recover` fora de defer, identifique o bug, (4) quiz de 8 perguntas sobre erros em Go."

### 🖼️ Roteiro de Slides
> "Crie um roteiro de 6 slides sobre 'Tratamento de Erros em Go'. Slide 1: O bilhete vs o alarme de incêndio — Errors as Values. Slide 2: `(T, error)` — o contrato explícito. Slide 3: Sentinel errors e `errors.Is`. Slide 4: `panic` — quando o impossível acontece. Slide 5: `defer` — a limpeza garantida e LIFO. Slide 6: Line of Sight — o Happy Path na esquerda."

### 💬 Perguntas Profundas para o Chat
- "Qual a diferença fundamental entre `errors.Is` e `errors.As` em Go? Me dê um exemplo de quando cada um é a ferramenta certa."
- "Como o wrapping de erros com `%w` funciona internamente? Como `errors.Is` atravessa a cadeia de erros embrulhados?"
- "Por que `recover` só funciona dentro de `defer`? O que acontece no runtime do Go durante o unwinding de um panic?"
- "Compare o padrão Line of Sight do Go com o uso de `else` em Java/Python. Por que Go prefere o retorno antecipado?"
