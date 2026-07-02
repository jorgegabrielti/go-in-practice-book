# 📖 Capítulo 06: Funções — As Mini-Fábricas Especializadas

> **Livro Go na Prática: 30 capítulos para dominar a linguagem**

---

## 🎯 Tema Central

Toda vez que um bloco de código é copiado e colado mais de duas vezes, nasce uma dívida técnica — é o oposto do princípio **DRY (Don't Repeat Yourself)**. A solução são as **Funções**: pense em uma função como uma **mini-fábrica** (um liquidificador). Você joga ingredientes dentro (parâmetros), o motor processa (corpo da função) e sai um produto pronto (retorno) — sem precisar saber como o motor funciona por dentro. Isso é **abstração**.

Em Go, funções são cidadãos de primeira classe e têm um superpoder raro entre linguagens como Java ou C: **múltiplos retornos**.

---

## 📊 Resumo dos Conceitos

### 1. Anatomia de uma Função

```go
func somar(a int, b int) int {
    resultado := a + b
    return resultado
}
```

* `func`: declara uma nova função.
* `somar`: o nome da função.
* `(a int, b int)`: os parâmetros — a "boca do funil". Tipos errados na chamada não compilam.
* `int` (depois dos parênteses): o tipo do valor de retorno.
* `return`: devolve o valor e termina a execução da função.

**Simplificando parâmetros:** se os parâmetros são do mesmo tipo, `func somar(a, b int)` é equivalente a `func somar(a int, b int)`.

### 2. Múltiplos Retornos

Em C ou Java, uma função só retorna um valor — para sinalizar erro junto com um resultado, seria preciso usar objetos complexos, exceções ou ponteiros. Em Go, uma função pode retornar quantos valores quiser:

```go
func dividir(a, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("não é possível dividir por zero")
    }
    return a / b, nil
}
```

No lado de quem chama:

```go
resultado, erro := dividir(10, 0)
if erro != nil {
    fmt.Println("Ops, deu ruim:", erro)
} else {
    fmt.Println("O resultado é:", resultado)
}
```

O padrão `resultado, err := funcao()` é a assinatura visual do Go: ele força o tratamento do erro logo após ele acontecer, em vez de deixar o programa explodir mais tarde.

### 3. Retornos Nomeados (Named Returns)

É possível nomear as variáveis de retorno já na assinatura da função, como se as caixas de saída já viessem etiquetadas:

```go
func coordenadas() (latitude float64, longitude float64) {
    latitude = -23.5505
    longitude = -46.6333

    // "Naked Return": como as variáveis já têm nome e já foram
    // preenchidas, basta dizer "return".
    return
}
```

**Cuidado:** o "naked return" pode confundir em funções longas — se a função tem 50 linhas, ao chegar no `return` sozinho é preciso subir tudo para lembrar o que está sendo devolvido. Usar com moderação, só em funções curtas.

### 4. Funções Variádicas (o "Buffet Livre")

Quando não se sabe quantos parâmetros serão recebidos, usa-se `...` antes do tipo. Dentro da função, o parâmetro variádico se comporta como um slice:

```go
func somarTudo(numeros ...int) int {
    total := 0
    for _, n := range numeros {
        total += n
    }
    return total
}

somarTudo(10, 20)        // 30
somarTudo(1, 1, 1, 1)    // 4
```

É assim que `fmt.Println` aceita qualquer quantidade de argumentos.

### 5. Funções Anônimas e Closures

Em Go, funções são valores — podem ser guardadas em variáveis como qualquer outro dado:

```go
func main() {
    dobrar := func(x int) int {
        return x * 2
    }
    fmt.Println(dobrar(5)) // 10
}
```

Útil para lógica descartável usada só naquele ponto, ou para passar comportamento como parâmetro de outra função (callbacks).

---

## 💡 Dica do Gopher

> **Funções pequenas são funções felizes.** Uma função deve fazer apenas **uma coisa** e fazê-la bem. Se o nome da função é `CalcularSalarioEEnviarEmailEAtualizarBancoDeDados()`, há um problema — quebre em `CalcularSalario()`, `EnviarEmail()` e `AtualizarBancoDeDados()`, e crie uma quarta função "gerente" que chama as três. Vantagens: fica mais fácil testar (calcular o salário sem mandar e-mail de verdade), mais fácil ler, e mais fácil reutilizar isoladamente no futuro.

---

## 🔬 Exemplos Práticos no Repositório

### 1. [Exemplo 01: Funções — visão geral](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo06/exemplos/ex01/main.go)
*   **Conceito:** reúne os quatro casos do capítulo em um único programa — função simples sem retorno (`boasVindas`), múltiplos retornos (`geometria`, área e perímetro de um retângulo), tratamento de erro com `errors.New` (`dividir`) e função variádica (`calcularMedia`).

---

## 🔬 Exercícios Práticos Resolvidos

### 1. [Exercício 01: A Calculadora Básica](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo06/exercicios/ex01/main.go)
*   **Objetivo:** criar `soma`, `subtracao` e `multiplicacao`, cada uma recebendo dois inteiros e retornando um inteiro, e chamá-las com 10 e 5.
*   **Conceito:** anatomia básica de função com retorno único.

### 2. [Exercício 02: O Analisador de Preços](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo06/exercicios/ex02/main.go)
*   **Objetivo:** `analisarPreco(preco float64)` retornando um `bool` (se é "caro", acima de 100.0) e uma `string` ("Caro"/"Barato"), testada com 150.0 e 50.0.
*   **Conceito:** múltiplos retornos com tipos diferentes (`bool` + `string`).

### 3. [Exercício 03: O Conversor Seguro](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo06/exercicios/ex03/main.go)
*   **Objetivo:** `converterTempo(segundos int)` retornando `(horas, minutos, segundos)` — 3661 segundos deve virar `1, 1, 1`.
*   **Conceito:** retornos nomeados e "naked return".

---

## ✅ Checklist antes de marcar como concluído

- [x] `fonte.txt` com a transcrição completa do capítulo
- [x] Teoria revisada e resumida neste README (com base no `fonte.txt`)
- [x] Todos os exemplos do capítulo têm pasta própria em `exemplos/ex01`, `ex02`... (dois dígitos)
- [x] Todos os exercícios resolvidos têm pasta própria em `exercicios/ex01`, `ex02`... (dois dígitos)
- [ ] `go build ./...` e `go vet ./...` passam sem erros
- [ ] `go fmt ./...` executado
- [x] Termos novos adicionados ao `estudos-go/GLOSSARIO.md`
- [x] Painel de progresso no `README.md` raiz atualizado
- [x] Artigo de Medium criado em `conteudo/medium/capitulo06.md` (a partir do `_template.md`, baseado no `fonte.txt`)
- [x] Post de LinkedIn criado em `conteudo/linkedin/capitulo06.md` (a partir do `_template.md`, baseado no `fonte.txt`)
- [x] Status atualizado em `conteudo/PAINEL.md`
- [x] Pilha de exercícios extras (mín. 10, dificuldade crescente) criada em `pilha-exercicios/capitulo06.md`
- [x] Teste de mesa de todo o código do capítulo criado em `teste-de-mesa.md`
