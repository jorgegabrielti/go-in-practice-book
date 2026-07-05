# Capítulo 04 — Controle de Fluxo: if, switch e Guard Clauses

## Tema Central

Go tem controle de fluxo deliberadamente simples. O `if` e o `switch` seguem regras claras que eliminam classes inteiras de bugs comuns em outras linguagens — como esquecer o `break` em um `switch` de C.

---

## O que o Livro Cobre

- `if`/`else` sem parênteses na condição; chave `{` na mesma linha (obrigatório)
- `if` com inicialização curta: `if x := f(); cond { }`
- Operadores lógicos: `&&`, `||`, `!`, com curto-circuito
- `switch`: `break` implícito por padrão; múltiplos valores por `case`
- `switch` tagless (equivalente a `if/else if` encadeado)
- `fallthrough`: força execução do próximo `case`
- Guard clauses / early return

---

## Aprofundando os Conceitos

### Por que não tem parênteses no `if`?

Em C, C++, Java e JavaScript, os parênteses em `if (cond)` são sintaxe obrigatória. Em Go, eles são opcionais no parser — mas `go fmt` os remove se você os adicionar. A justificativa: parênteses são ruído visual que não acrescentam clareza. A chave `{` já delimita o bloco.

```go
// CORRETO (Go-idiomático)
if x > 0 {
    fmt.Println("positivo")
}

// COMPILA mas go fmt remove os parênteses
if (x > 0) {
    fmt.Println("positivo")
}
```

### `if` com inicialização — escopo limitado

```go
// x só existe dentro do bloco if/else
if x := calcular(); x > 100 {
    fmt.Println("grande:", x)
} else {
    fmt.Println("pequeno:", x)
}
// x não existe aqui — erro de compilação se tentar usar
```

Isso é poderoso para erros:

```go
if err := arquivo.Fechar(); err != nil {
    log.Printf("erro ao fechar: %v", err)
}
// err não "vaza" para o escopo externo
```

### Curto-circuito em operadores lógicos

```go
// && para quando o PRIMEIRO é false
if arquivo != nil && arquivo.Existe() {
    // Se arquivo for nil, arquivo.Existe() NUNCA É CHAMADO
    // Sem curto-circuito, isso causaria panic (nil pointer dereference)
}

// || para quando o PRIMEIRO é true
if cache.Hit() || banco.Buscar() {
    // Se o cache acertou, banco.Buscar() não é executado (evita I/O desnecessário)
}
```

Curto-circuito não é apenas otimização — frequentemente é a diferença entre o código funcionar ou causar panic.

### `switch` em Go vs C/Java

A diferença mais importante: em C e Java, **você precisa do `break`** para sair de cada `case`. Esquecer o `break` é um bug clássico (fall-through acidental). Em Go, o `break` é **implícito** — você precisa de `fallthrough` explícito quando QUER o comportamento de fall-through.

```go
// Go: break implícito — cada case é independente
switch status {
case "ativo":
    fmt.Println("usuário ativo")
    // sem break necessário
case "inativo":
    fmt.Println("usuário inativo")
}

// Múltiplos valores no mesmo case
switch diaDaSemana {
case "sábado", "domingo":
    fmt.Println("fim de semana")
default:
    fmt.Println("dia útil")
}
```

```java
// Java: PRECISA do break — esquecer é bug silencioso
switch (status) {
    case "ativo":
        System.out.println("ativo");
        break;  // sem isso, cai no próximo case!
    case "inativo":
        System.out.println("inativo");
        break;
}
```

### `switch` tagless — o `if/else if` legível

```go
// Com if/else if — difícil de ler com muitas condições
if nota >= 90 {
    fmt.Println("A")
} else if nota >= 80 {
    fmt.Println("B")
} else if nota >= 70 {
    fmt.Println("C")
} else {
    fmt.Println("Reprovado")
}

// Com switch tagless — mais limpo
switch {
case nota >= 90:
    fmt.Println("A")
case nota >= 80:
    fmt.Println("B")
case nota >= 70:
    fmt.Println("C")
default:
    fmt.Println("Reprovado")
}
```

### `fallthrough` — quando você realmente quer o fall-through

```go
switch versao {
case 3:
    fmt.Println("executando migração v2→v3")
    fallthrough
case 2:
    fmt.Println("executando migração v1→v2")
    fallthrough
case 1:
    fmt.Println("executando migração v0→v1")
}
// Se versao = 3: imprime as três mensagens (cumulativo)
// Se versao = 2: imprime as duas últimas
// Se versao = 1: imprime só a última
```

`fallthrough` é raro em código Go idiomático — um sinal de que talvez uma abordagem diferente seja melhor.

### Guard Clauses: a técnica mais importante do capítulo

Guard clause é retornar cedo quando uma condição não é satisfeita, evitando aninhamento profundo:

```go
// Sem guard clause — "arrow code" (pirâmide de doom)
func processar(usuario *Usuario) error {
    if usuario != nil {
        if usuario.Ativo {
            if usuario.Saldo > 0 {
                // lógica principal aqui, enterrada no nível 3
                return nil
            } else {
                return errors.New("saldo insuficiente")
            }
        } else {
            return errors.New("usuário inativo")
        }
    } else {
        return errors.New("usuário nil")
    }
}

// Com guard clauses — linear, legível
func processar(usuario *Usuario) error {
    if usuario == nil {
        return errors.New("usuário nil")
    }
    if !usuario.Ativo {
        return errors.New("usuário inativo")
    }
    if usuario.Saldo <= 0 {
        return errors.New("saldo insuficiente")
    }
    
    // lógica principal aqui, limpa e no nível 0
    return nil
}
```

Guard clauses estão intimamente ligadas ao padrão `if err != nil { return err }` do Go — a forma idiomática de tratar erros.

---

## Referências Oficiais

- **Especificação — Statements (If, Switch)**: https://go.dev/ref/spec#Statements
- **Effective Go — Control structures**: https://go.dev/doc/effective_go#control-structures
- **Effective Go — If**: https://go.dev/doc/effective_go#if
- **Effective Go — Switch**: https://go.dev/doc/effective_go#switch
- **Tour de Go — Flow control**: https://go.dev/tour/flowcontrol/1

---

## Exemplos de Código Adicionais

### Type switch — switch sobre o tipo de uma interface

```go
func descrever(i interface{}) {
    switch v := i.(type) {
    case int:
        fmt.Printf("Inteiro: %d\n", v)
    case string:
        fmt.Printf("String com %d chars: %s\n", len(v), v)
    case bool:
        fmt.Printf("Booleano: %v\n", v)
    default:
        fmt.Printf("Tipo desconhecido: %T\n", v)
    }
}

descrever(42)       // Inteiro: 42
descrever("olá")    // String com 3 chars: olá
descrever(true)     // Booleano: true
descrever(3.14)     // Tipo desconhecido: float64
```

O type switch é essencial ao trabalhar com interfaces e com `any` (alias de `interface{}`).

### `switch` com inicialização

```go
switch err := operacao(); {
case err == nil:
    fmt.Println("sucesso")
case errors.Is(err, ErrNaoEncontrado):
    fmt.Println("não encontrado")
default:
    fmt.Printf("erro inesperado: %v\n", err)
}
```

---

## Perguntas & Respostas Frequentes

**P: O `if` do Go pode ter `else if`?**
R: Sim: `if ... {} else if ... {} else {}`. Mas para muitas condições, o `switch` tagless é mais legível.

**P: O `switch` do Go avalia os `case` em ordem?**
R: Sim. O primeiro `case` que corresponder é executado. Isso importa quando os casos se sobrepõem (embora casos que se sobrepõem sejam sinal de código confuso).

**P: Posso usar `switch` com strings?**
R: Sim. O Go aceita qualquer tipo comparável com `==` em um `switch`, incluindo `string`, `int`, `bool`, ponteiros, etc.

**P: O `fallthrough` testa a condição do próximo `case`?**
R: Não. `fallthrough` executa o corpo do próximo `case` **incondicionalmente**, sem testar sua condição. É diferente do fall-through de C.

**P: Qual a diferença entre `break` dentro de `switch` e dentro de `for`?**
R: Dentro de `switch`, `break` encerra o `switch` (mas como o break já é implícito, raramente é necessário explícito). Dentro de `for`, `break` encerra o loop. Se um `switch` está dentro de um `for` e você quer sair do `for`, use um **label**:
```go
loop:
    for i := 0; i < 10; i++ {
        switch i {
        case 5:
            break loop  // sai do for, não do switch
        }
    }
```

---

## Comparações com Outras Linguagens

### `switch` fall-through

```c
// C: fall-through é o padrão (bug histórico)
switch (x) {
    case 1: printf("um");  // CUIDADO: sem break, cai no próximo!
    case 2: printf("dois");
    break;
}
// Se x=1: imprime "umdois" — bug silencioso clássico
```

```python
# Python: não tem switch (antes do 3.10)
# Python 3.10+: match/case (sem fall-through)
match status:
    case "ativo":
        print("ativo")
    case "inativo":
        print("inativo")
```

```go
// Go: break implícito — fall-through é explícito e raro
switch x {
case 1:
    fmt.Print("um")  // para aqui automaticamente
case 2:
    fmt.Print("dois")
}
```

### Guard clauses

```python
# Python equivalente (sem guard clause)
def processar(usuario):
    if usuario is not None:
        if usuario.ativo:
            # lógica enterrada
            pass

# Com guard clause
def processar(usuario):
    if usuario is None:
        return
    if not usuario.ativo:
        return
    # lógica no nível 0
```

---

## Armadilhas Comuns

1. **`fallthrough` no último `case`**: causa erro de compilação — não há próximo case para executar.

2. **`switch` vs `if` para comparação de tipos**: para type assertion, use type switch. Para comparação de valores, use switch normal.

3. **Esquecer o `default` em switch**: nem sempre necessário, mas bom para documentar que outros casos foram considerados.

4. **Condição no `if` deve ser `bool`**: ao contrário de Python (que aceita qualquer valor como truthy/falsy), Go exige expressão booleana explícita:
   ```go
   if x { }       // ERRO se x for int
   if x != 0 { }  // CORRETO
   ```

---

## Quiz de Fixação

1. O que acontece se você esquecer o `break` em um `switch` do Go?
2. Qual a diferença entre `switch x { case 1: }` e `switch { case x == 1: }`?
3. Por que `if x { }` não compila se `x` for um inteiro?
4. O que é `fallthrough` e por que é raro no Go idiomático?
5. O que é curto-circuito e por que importa na expressão `arquivo != nil && arquivo.Existe()`?
6. Reescreva este código com guard clauses:
   ```go
   if usuario != nil {
       if usuario.Ativo {
           processar(usuario)
       }
   }
   ```

---

## 🔬 Dissecando a Sintaxe

### `if` com inicialização curta — o bloco mais compacto do Go

```go
if x := calcular(); x > 100 {
    fmt.Println("grande:", x)
} else {
    fmt.Println("pequeno:", x)
}
```

```
if  x := calcular()  ;  x > 100  {
───  ──────────────   ─  ────────  ─
 │         │          │      │     └─ abre o bloco do if
 │         │          │      └─ CONDIÇÃO: expressão booleana avaliada após o ";"
 │         │          └─ separador: divide a inicialização da condição
 │         └─ INICIALIZAÇÃO: declara "x" com o retorno de calcular()
 │             ":=" = declara + atribui (x só existe dentro deste bloco if/else)
 └─ palavra-chave if

DETALHE IMPORTANTE: "x" existe DENTRO do if e do else, mas NÃO fora deles
→ se tentar usar "x" após o bloco, erro de compilação: "undefined: x"
→ isso é proposital: limita o escopo da variável ao bloco onde faz sentido
```

### `switch` tagless — lendo a estrutura

```go
switch {                    // [1]
case nota >= 90:            // [2]
    fmt.Println("A")
case nota >= 80:            // [3]
    fmt.Println("B")
default:                    // [4]
    fmt.Println("Reprovado")
}
```

```
[1]  switch {
      ──────
         └─ switch SEM expressão após a palavra-chave
             equivale a "switch true { }"
             cada case fornece sua própria condição booleana

[2]  case nota >= 90:
      ────  ──────────
        │        └─ condição completa: qualquer expressão que resulte em bool
        └─ palavra-chave: "se esta condição for verdadeira, execute o bloco abaixo"
            ":": dois-pontos finaliza a condição do case (sem parênteses, sem chaves)
            BREAK IMPLÍCITO: ao executar este case, o switch termina (sem fallthrough)

[3]  case nota >= 80:
         └─ só executado se [2] for falso (Go avalia em ordem, para no primeiro true)

[4]  default:
      ───────
         └─ executado se NENHUM case anterior for verdadeiro
             equivale ao "else" do if/else if
```

### Guard Clause — lendo o padrão

```go
func processar(u *Usuario) error {
    if u == nil {                          // [1]
        return errors.New("usuário nil")   // [2]
    }
    if !u.Ativo {                          // [3]
        return errors.New("inativo")
    }
    // [4] lógica principal aqui
    return nil
}
```

```
[1]  if u == nil {
          ─  ───
          │    └─ nil = zero value de ponteiro (equivale a "nenhum valor")
          └─ u = o parâmetro do tipo *Usuario (ponteiro)
              "u == nil" pergunta: "u está apontando para algo válido?"

[2]  return errors.New("usuário nil")
      ──────  ─────────────────────
         │          └─ cria um novo valor de erro com a mensagem fornecida
         └─ retorna IMEDIATAMENTE — a função para aqui
             este é o "guard": bloqueia a entrada de dados inválidos cedo

[3]  !u.Ativo
      ─ ──────
      │    └─ "u.Ativo" = acessa o campo Ativo da struct que u aponta
      └─ "!" = NOT: inverte o booleano (true→false, false→true)
          "!u.Ativo" = "se o usuário NÃO estiver ativo"

[4]  Chegou aqui = todos os guards passaram = dados válidos garantidos
     A lógica principal fica no "nível 0" — sem aninhamento
```

---

## 🎙️ Prompts para o NotebookLM

### 🎧 Audio Overview
> "Gere um Audio Overview onde dois desenvolvedores debatem: 'Guard Clauses vs if/else aninhado — qual é melhor e por quê?' Um deles defende o estilo tradicional com aninhamento, o outro explica por que Go favorece guard clauses e early return. Inclua exemplos de código e o impacto na legibilidade."

### 📋 Briefing Doc
> "Crie um briefing sobre as diferenças entre o `switch` do Go e o `switch` de C/Java/JavaScript, destacando: (1) break implícito vs explícito, (2) switch tagless como alternativa ao if/else if, (3) o caso especial de fallthrough."

### 📚 Study Guide
> "Crie um guia de estudos do Capítulo 04 com: (1) tabela comparando if/switch/guard clause — quando usar cada um, (2) 5 exercícios de refatoração: transforme código com if aninhado em guard clauses, (3) quiz de 8 perguntas sobre controle de fluxo."

### 🖼️ Roteiro de Slides
> "Crie um roteiro de 5 slides sobre 'Controle de Fluxo em Go'. Slide 1: if sem parênteses — por quê. Slide 2: if com inicialização curta — o superpoder. Slide 3: switch com break implícito vs C/Java. Slide 4: Guard Clauses — antes e depois. Slide 5: O custo cognitivo do aninhamento."

### 💬 Perguntas Profundas para o Chat
- "O que é 'Pyramid of Doom' ou 'Arrow Code' e como guard clauses resolvem esse problema?"
- "Qual a diferença entre `switch x { case 1: }` e `switch { case x == 1: }`? Me dê um exemplo onde um é melhor que o outro."
- "O `fallthrough` do Go é diferente do fall-through de C — como exatamente?"
- "Por que Go exige que a condição do `if` seja booleana? Em Python `if 0:` é falso — no Go isso compila?"
