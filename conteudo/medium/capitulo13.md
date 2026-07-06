> **Antes de publicar:** conferir checklist no fim do arquivo.

# 🧾 Go não grita com você quando algo dá errado

Imagine que você contratou um funcionário muito eficiente. Sempre que ele encontra um problema no trabalho, ele não liga o alarme de incêndio, não corre para o corredor gritando, não para a empresa inteira. Ele simplesmente coloca um bilhetinho na sua mesa: "Não consegui entregar o relatório. O arquivo estava vazio."

Você lê o bilhete, decide o que fazer, e o dia continua.

Esse é o modelo de erros do Go.

---

## O bilhete, não o alarme

Em linguagens que usam exceções — Java, Python, C# — quando algo falha, o programa interrompe o fluxo normal e dispara um mecanismo de controle especial: `throw`, `catch`, `finally`. É poderoso, mas cria um fluxo invisível que pode surpreender.

Go faz diferente: erro é só mais um valor de retorno. Nada de especial, nada de mágico.

```go
func buscarUsuario(id int) (Usuario, error)
```

Essa função pode te dar um usuário ou um erro. Você trata os dois do mesmo jeito — como dados.

```go
usuario, err := buscarUsuario(42)
if err != nil {
    fmt.Println("Problema ao buscar usuário:", err)
    return
}
// usuário disponível aqui com segurança
```

Isso é o que a comunidade chama de **"Errors as Values"**. O erro não é uma catástrofe — é informação.

---

## Fabricando seus próprios bilhetes

Existem duas ferramentas principais para criar erros nas suas funções.

A primeira é `errors.New`, para quando a mensagem é fixa:

```go
var ErrTanqueVazio = errors.New("tanque vazio")
```

A segunda é `fmt.Errorf`, para quando você precisa incluir detalhes do contexto:

```go
return fmt.Errorf("falha ao autenticar usuário %s: senha inválida", nome)
```

O verbo `%w` é especial: ele "embrulha" um erro existente dentro do novo, preservando a identidade do erro original para inspeção posterior. Isso é o que permite criar cadeias de erros rastreáveis — cada camada adiciona contexto sem perder a causa raiz.

---

## Sentinel errors: bilhetes com nome

Uma prática muito comum em Go é criar variáveis de erro nomeadas no topo do arquivo — os chamados **sentinel errors**:

```go
var ErrDivisaoPorZero = errors.New("matemática diz não: divisão por zero")

func dividir(a, b float64) (float64, error) {
    if b == 0 {
        return 0, ErrDivisaoPorZero
    }
    return a / b, nil
}
```

Por que isso é útil? Porque quem chama a função pode verificar **exatamente qual erro** ocorreu:

```go
if errors.Is(err, ErrDivisaoPorZero) {
    fmt.Println("Tente usar um denominador diferente de zero.")
}
```

`errors.Is` é esperto: ele atravessa cadeias de wrapping. Então mesmo que alguém tenha embrulhado o erro com contexto extra, ele ainda encontra o sentinel original lá dentro.

---

## O botão de emergência: `panic`

Até agora falamos de erros recuperáveis. Mas e quando a situação é realmente sem saída? Quando o banco de dados sumiu, quando o índice acessado não existe, quando o estado do programa é tão corrompido que não tem mais como continuar?

Para isso, o Go tem o `panic`:

```go
panic("estado impossível: lista deveria ter pelo menos um elemento")
```

`panic` para tudo imediatamente e começa a desfazer a pilha de chamadas. A regra é clara: não use `panic` para situações normais de negócio — "arquivo não encontrado" é um `error`, não um `panic`. Reserve para bugs graves que nunca deveriam acontecer.

---

## O bombeiro: `recover`

É possível interceptar um `panic` antes que ele encerre o processo. Mas há um detalhe importante: `recover` só funciona dentro de um `defer`.

```go
func executarComSeguranca() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Capturei um panic:", r)
        }
    }()

    // código que pode panic aqui
}
```

Isso é muito útil em servidores web: uma requisição mal-formada não pode derrubar o servidor inteiro. O `recover` age como um para-quedas de emergência — raramente usado, mas vital quando necessário.

---

## O faxineiro garantido: `defer`

`defer` agenda uma função para rodar no **último momento antes do retorno**, independente de como a função termina — return normal, return com erro, panic capturado.

```go
func processarArquivo(nome string) error {
    f, err := os.Open(nome)
    if err != nil {
        return err
    }
    defer f.Close()  // vai fechar, sempre — não importa o que aconteça abaixo

    return processarDados(f)
}
```

Sem `defer`, você precisaria lembrar de chamar `f.Close()` em cada caminho de retorno. Com `defer`, você declara a intenção uma vez e ela é garantida.

Quando você usa vários `defer` na mesma função, eles executam em ordem **LIFO** (Last In, First Out) — o último registrado é o primeiro a rodar. Isso é intencional: se você abriu A, B e C nessa ordem, faz sentido fechar C, B e A.

---

## A regra do Happy Path

Uma convenção que separa o código Go idiomático do novato: evite o `else` no tratamento de erros.

```go
// Estilo que parece razoável, mas não é idiomático
if err != nil {
    return err
} else {
    processar(dado)
}

// Estilo Go: trate o problema e retorne cedo
if err != nil {
    return err
}
processar(dado)  // caminho feliz, sem aninhamento
```

Esse padrão se chama **"Line of Sight"**: o fluxo principal fica sempre na coluna esquerda, sem indentação extra. Erros são tratados e descartados rapidamente. O código fica mais fácil de ler de cima para baixo.

---

## 🦫 O que ficou

Go escolheu deliberadamente não ter exceções. Erros são valores porque valores são previsíveis, rastreáveis e parte explícita do contrato de uma função. Quando você vê `(Arquivo, error)` na assinatura, já sabe que essa operação pode falhar — não precisa ler a documentação para descobrir.

`defer` garante limpeza. `panic` sinaliza o impossível. `recover` protege sistemas críticos. E o `if err != nil` repetido que parece verboso? É clareza explícita — e em Go, explícito é sempre preferido ao implícito.

Repositório completo: [Link do repositório]

---

**Checklist antes de publicar:**
- [x] Conteúdo confere com `estudos-go/capitulo13/fonte.txt`
- [x] Texto é original — sem frases copiadas do fonte.txt
- [x] Todos os blocos de código testados e funcionais
- [ ] Link para o repositório no GitHub incluído
- [ ] Revisão ortográfica feita
- [ ] Capa/imagem de destaque escolhida no Medium
- [ ] Tags do Medium definidas (Golang, Programming, ErrorHandling, Backend, CleanCode)
- [ ] Status atualizado em `conteudo/PAINEL.md`
