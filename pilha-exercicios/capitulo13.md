# 🏋️ Pilha de Exercícios — Capítulo 13: Tratamento de Erros

> Exercícios extras sem solução. Resolva por conta própria em `estudos-go/capitulo13/exercicios/` ou em um scratch separado.

---

## 🟢 Fácil

### 1. A Divisão Segura

Crie uma função `Dividir(a, b float64) (float64, error)`. Se `b` for zero, retorne um sentinel error `ErrDivisaoPorZero`. Caso contrário, retorne o resultado e `nil`. Na `main`, teste com `Dividir(10, 2)` e `Dividir(5, 0)` e imprima o resultado ou o erro.

### 2. O Conversor Seguro

Crie uma função `ConverterParaInteiro(s string) (int, error)` que tenta converter uma string para inteiro usando `strconv.Atoi`. Se falhar, retorne `0` e um erro formatado com `fmt.Errorf` incluindo a string original na mensagem. Teste com `"42"`, `"abc"` e `""`.

### 3. O Leitor de Idade

Crie uma função `ValidarIdade(idade int) error`. Regras:
- Se `idade < 0`: retorne `errors.New("idade não pode ser negativa")`
- Se `idade > 150`: retorne `errors.New("idade improvável")`
- Caso contrário: retorne `nil`

Teste com `-1`, `25`, `200`.

### 4. O Defer Simples

Crie uma função `AbrirConexao()` que imprime "Conectando..." e usa `defer` para imprimir "Desconectando." ao final. Chame a função na `main`. Adicione também um `fmt.Println("Usando conexão...")` antes do return para observar a ordem.

---

## 🟡 Médio

### 5. O Banco de Dados Simulado

Crie dois sentinel errors: `ErrUsuarioNaoEncontrado` e `ErrSenhaIncorreta`. Crie uma função `Login(usuario, senha string) error` que:
- Retorna `ErrUsuarioNaoEncontrado` se o usuário não for `"admin"`
- Retorna `ErrSenhaIncorreta` se a senha não for `"1234"`
- Retorna `nil` se ambos estiverem corretos

Na `main`, use `errors.Is` para distinguir os dois tipos de erro e imprima mensagens diferentes para cada caso.

### 6. A Pilha de Cleanup

Crie uma função que simule a abertura de três recursos (arquivo, conexão de banco, mutex) com `fmt.Println("Abrindo X...")`. Use `defer` para fechar cada um na ordem correta (o último a abrir deve ser o primeiro a fechar). Observe que os defers executam na ordem LIFO e verifique se a ordem de fechamento faz sentido.

### 7. O Wrapper de Erro

Crie um sentinel error `ErrPermissaoNegada`. Crie uma função `Acessar(usuario, recurso string) error` que retorna `ErrPermissaoNegada` quando `usuario != "admin"`. Na `main`, chame `Acessar`, e se houver erro, use `fmt.Errorf` com `%w` para "embrulhar" o erro original com contexto adicional (ex: `"ao acessar /etc/passwd: %w"`). Verifique com `errors.Is` que o erro embrulhado ainda é detectável.

### 8. O Recuperador de Panics

Crie uma função `ExecutarComSeguranca(fn func()) (err error)` que executa `fn` dentro de um `defer` + `recover`. Se `fn` causar um `panic`, a função deve capturá-lo e retorná-lo como um `error` (use `fmt.Errorf("panic capturado: %v", r)`). Teste passando funções que causam panic (acesso a índice inválido, divisão explícita, etc.) e funções normais.

---

## 🔴 Desafio

### 9. O Tipo de Erro Customizado

Crie um tipo `ErrValidacao` com campos `Campo string` e `Motivo string`. Implemente o método `Error() string` retornando `"campo 'X': Y"`. Crie uma função `ValidarUsuario(nome, email string) error` que retorna `*ErrValidacao` para nome vazio ou email sem `@`. Na `main`, use `errors.As` (não `errors.Is`) para extrair o `*ErrValidacao` e acessar os campos `Campo` e `Motivo`.

### 10. O Pipeline de Validações

Crie uma função `ValidarCadastro(nome, email, senha string) []error` que retorna uma **lista** de erros — um para cada campo inválido — em vez de parar no primeiro. Regras: nome não pode ser vazio, email precisa de `@`, senha precisa ter 8+ caracteres. Na `main`, itere sobre os erros e imprima todos de uma vez, como um formulário de cadastro real faria.

### 11. O Middleware de Recover

Simule um servidor HTTP simples: crie uma função `ServirRequisicao(rota string)` que pode causar `panic("rota inválida")` ou `panic("banco de dados indisponível")`. Crie um wrapper `HandlerSeguro(rota string)` que usa `defer + recover` para capturar qualquer panic, logar o erro (com `rota` no contexto) e continuar sem derrubar o programa. Chame 3 rotas diferentes — uma normal, uma que causa panic de rota, uma de banco.

### 12. O Rastreador de Erros (Cadeia de Wrapping)

Crie 3 camadas de funções: `repositorio()` → `servico()` → `handler()`. A camada `repositorio` retorna um sentinel error `ErrConexaoRecusada`. Cada camada acima usa `fmt.Errorf("contexto: %w", err)` para adicionar contexto ao propagar. Na `main`, use `errors.Is` para verificar que o erro original ainda é detectável no topo da cadeia, e `errors.Unwrap` para inspecionar a cadeia camada por camada.
