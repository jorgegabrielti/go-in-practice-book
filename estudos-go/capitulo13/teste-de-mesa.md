# 🧮 Teste de Mesa — Capítulo 13: Tratamento de Erros

Rastreamento manual da execução do `exemplos/exe01/main.go` linha a linha, **antes** de rodar `go run`.

---

## Código rastreado

`exemplos/exe01/main.go` — funções `dividir`, `operacaoArriscada` e `main`.

---

## Estado inicial

| Variável | Tipo | Valor |
|---|---|---|
| `ErrDivisaoPorZero` | `error` | `&errors.errorString{"matemática diz não: divisão por zero"}` |

---

## Bloco 1 — Tratamento Padrão (`dividir(10, 0)`)

```
resultado, err := dividir(10, 0)
```

**Dentro de `dividir(a=10, b=0)`:**

| Passo | Expressão | Resultado |
|---|---|---|
| 1 | `b == 0` | `true` |
| 2 | `return 0, ErrDivisaoPorZero` | sai da função |

**De volta em `main`:**

| Variável | Valor |
|---|---|
| `resultado` | `0.0` |
| `err` | `ErrDivisaoPorZero` (não é `nil`) |

```
if err != nil          → true (err tem valor)
  if errors.Is(err, ErrDivisaoPorZero)  → true (mesmo ponteiro)
    fmt.Println("Erro Específico: Você tentou dividir por zero.")
```

**Saída:**
```
Erro Específico: Você tentou dividir por zero.
```

---

## Bloco 2 — Teste de `panic` e `recover`

```
fmt.Println("\n--- Teste de Pânico ---")
```
**Saída:**
```

--- Teste de Pânico ---
```

**Chamada `operacaoArriscada()`:**

| Passo | Ação | Observação |
|---|---|---|
| 1 | `defer func() { ... }()` | registra a função anônima na pilha de defers — **ainda não executa** |
| 2 | `fmt.Println("Iniciando operação nuclear...")` | imprime normalmente |
| 3 | `panic("EXPLOSÃO NO REATOR 4!")` | inicia o unwinding — execução normal para |
| 4 | `fmt.Println("Isso nunca será impresso.")` | **NUNCA EXECUTADO** — o panic já interrompeu |
| 5 | defer ativa: `r := recover()` | captura `"EXPLOSÃO NO REATOR 4!"` |
| 6 | `r != nil` | `true` |
| 7 | `fmt.Println("🚒 OPA! Recuperei de um pânico:", r)` | imprime |
| 8 | função retorna normalmente | o panic foi "engolido" pelo recover |

**Saída de `operacaoArriscada`:**
```
Iniciando operação nuclear...
🚒 OPA! Recuperei de um pânico: EXPLOSÃO NO REATOR 4!
```

**De volta em `main`** (o recover salvou o programa):
```
fmt.Println("O programa continuou rodando! O recover salvou o dia.")
```
**Saída:**
```
O programa continuou rodando! O recover salvou o dia.
```

---

## Bloco 3 — Defer na Prática (função anônima)

```
fmt.Println("\n--- Teste de Defer ---")
```
**Saída:**
```

--- Teste de Defer ---
```

**Dentro da função anônima `func() { ... }()`:**

| Passo | Linha | Ação | Saída imediata |
|---|---|---|---|
| 1 | `fmt.Println("Abrindo porta...")` | executa | `Abrindo porta...` |
| 2 | `defer fmt.Println("Fechando porta (Defer).")` | registra na pilha — **não imprime agora** | — |
| 3 | `fmt.Println("Entrando na casa...")` | executa | `Entrando na casa...` |
| 4 | `fmt.Println("Fazendo bagunça...")` | executa | `Fazendo bagunça...` |
| 5 | função termina → defer executa | `fmt.Println("Fechando porta (Defer).")` | `Fechando porta (Defer).` |

---

## Saída Total Esperada

```
Erro Específico: Você tentou dividir por zero.

--- Teste de Pânico ---
Iniciando operação nuclear...
🚒 OPA! Recuperei de um pânico: EXPLOSÃO NO REATOR 4!
O programa continuou rodando! O recover salvou o dia.

--- Teste de Defer ---
Abrindo porta...
Entrando na casa...
Fazendo bagunça...
Fechando porta (Defer).
```

---

## Exercícios — Trace Rápido

### exe01 — A Compra Segura

```
comprar(100, 50)
  preco (50) > saldo (100) → false
  return 50.0, nil
→ novoSaldo=50.0, sem erro

comprar(50, 100)
  preco (100) > saldo (50) → true
  return 50.0, errors.New("saldo insuficiente")
→ novoSaldo=50.0, err="saldo insuficiente"
```

### exe02 — Validador de Senha

```
validarSenha("12345")
  len("12345") = 5 < 8 → true
  return fmt.Errorf("senha muito curta: tem apenas 5 caracteres")

validarSenha("senhasegura")
  len = 11 >= 8 → ok
  temNumero("senhasegura") → false (nenhum unicode.IsDigit)
  return errors.New("senha precisa de número")

validarSenha("senha12345")
  len = 10 >= 8 → ok
  temNumero → true (encontrou '1')
  return nil
```

**Saída esperada:**
```
Senha: 12345        | Erro: senha muito curta: tem apenas 5 caracteres
Senha: senhasegura  | Erro: senha precisa de número
Senha: senha12345   | OK (Válida)
```

### exe03 — A Pilha de Pratos (LIFO)

Loop `for i := 0; i < 5; i++` registra 5 defers:

| Iteração | defer registrado |
|---|---|
| i=0 | `defer fmt.Println(0)` — entra na pilha |
| i=1 | `defer fmt.Println(1)` — entra na pilha |
| i=2 | `defer fmt.Println(2)` — entra na pilha |
| i=3 | `defer fmt.Println(3)` — entra na pilha |
| i=4 | `defer fmt.Println(4)` — entra na pilha |

`fmt.Println("Fim do Loop")` → imprime antes dos defers.

Ao sair de `main`, pilha desempilha (LIFO — Last In, First Out):

**Saída esperada:**
```
Fim do Loop
4
3
2
1
0
```
