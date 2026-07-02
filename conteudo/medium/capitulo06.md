---

Funções em Go: por que sua linguagem te obriga a tratar erro na hora

Seja muito bem-vindo ao 6º artigo da série Go na Prática. Nos capítulos anteriores aprendemos a tomar decisões com `if`/`switch` e a repetir tarefas com `for`. Só que repetir tarefa dentro de um loop é uma coisa — repetir o mesmo bloco de lógica em pontos diferentes do programa é outra, bem mais perigosa. E é exatamente esse problema que as funções resolvem.

Pensa numa situação boba: você tem um sistema de pedidos e, em três lugares diferentes do código, calcula o valor do frete da mesma forma. Funciona até o dia em que a regra do frete muda — e você precisa lembrar de atualizar os três lugares. Esquece um, e já tem um bug de produção esperando pra acontecer. Esse tipo de duplicação tem um nome conhecido no mercado, o princípio DRY (não se repita), e funções são a ferramenta mais básica pra cumpri-lo: você escreve a lógica uma vez, dá um nome a ela, e chama esse nome sempre que precisar.

O que me chamou atenção neste capítulo foi notar como o Go trata funções de um jeito bem mais "adulto" que linguagens como Java e C, principalmente em um ponto: a forma de lidar com erro.

---

1. 🔧 Uma função é só uma caixa com entrada e saída
No fim das contas, declarar uma função em Go é simples:
package main

func somar(a int, b int) int {
  resultado := a + b
  return resultado
}
`func` diz que uma nova caixa está sendo criada, `somar` é o rótulo dela, `(a int, b int)` define o que ela aceita receber — e o Go é exigente aqui: se a chamada não bater com os tipos esperados, nem compila — e o `int` solto depois dos parênteses é o que ela promete entregar de volta via `return`.
Um detalhe pequeno que economiza digitação: quando os parâmetros são do mesmo tipo, não precisa repetir — `func somar(a, b int)` já basta.

---

2. 🎯 O que realmente separa Go das outras linguagens: retornar mais de uma coisa
Em Java ou C, uma função só devolve um único valor. Se você precisa devolver "o resultado" e também avisar "se algo deu errado", normalmente vira gambiarra: lançar exceção, devolver um objeto inchado, usar ponteiro de saída. Go simplesmente permite que uma função devolva quantos valores quiser, lado a lado:
func dividir(a, b int) (int, error) {
  if b == 0 {
    return 0, fmt.Errorf("não é possível dividir por zero")
  }
  return a / b, nil
}
Repare como o erro vira só mais um valor de retorno, não uma exceção que precisa ser capturada em algum lugar distante. Quem chama a função recebe os dois valores de uma vez:
resultado, erro := dividir(10, 0)

if erro != nil {
  fmt.Println("Ops, deu ruim:", erro)
} else {
  fmt.Println("O resultado é:", resultado)
}
Essa construção `valor, err := funcao()` aparece o tempo inteiro em código Go, e não é estilo — é a própria filosofia da linguagem te empurrando pra tratar erro no exato lugar onde ele pode acontecer, em vez de deixar pra depois (ou nunca).

---

3. 🏷️ Quando o próprio retorno já vem com nome
O Go também permite batizar as variáveis de saída direto na assinatura da função:
func coordenadas() (latitude float64, longitude float64) {
  latitude = -23.5505
  longitude = -46.6333

  // como as variáveis já existem e já têm valor, um "return"
  // sozinho já basta para devolvê-las
  return
}
Esse `return` sem nada na frente costuma ser chamado de naked return. Funciona bem em funções curtas, mas em funções grandes ele vira um problema de leitura: se a função tem 40, 50 linhas, chegar no final e ver só um `return` te obriga a rolar a tela pra trás pra lembrar o que cada variável guarda. Na prática, eu uso isso só quando a função inteira cabe na tela sem rolar.

---

4. 🧺 Quando você não sabe quantos argumentos vai receber
Tem um caso bem comum: uma função que precisa aceitar uma quantidade variável de valores. O exemplo mais óbvio do dia a dia é o próprio `fmt.Println`, que aceita de um a vários argumentos sem reclamar. Pra escrever algo parecido, o Go usa `...` antes do tipo:
func somarTudo(numeros ...int) int {
  total := 0
  for _, n := range numeros {
    total += n
  }
  return total
}

somarTudo(10, 20)     // 30
somarTudo(1, 1, 1, 1) // 4
Dentro da função, esse parâmetro variádico se comporta exatamente como um slice — por isso o `range` funciona normalmente em cima dele.

---

5. 🎭 Função guardada numa variável (sim, isso existe)
Outro detalhe que estranhei no começo: em Go, função é tratada como qualquer outro valor — número, string, struct. Isso significa que dá pra guardar uma função dentro de uma variável:
func main() {
  dobrar := func(x int) int {
    return x * 2
  }
  fmt.Println(dobrar(5)) // 10
}
Na prática isso é usado quando você precisa de uma lógica que só vai existir naquele ponto específico do código, ou quando quer passar um comportamento como argumento pra outra função — o que abre porta pra callbacks, algo que volta com força mais pra frente no livro.

---

💡 Dica do Gopher
Função deve fazer uma coisa só, e fazer bem. Se o nome dela já é uma frase inteira — tipo `CalcularSalarioEEnviarEmailEAtualizarBanco()` — é sinal de que ela está fazendo trabalho demais. Separar em três funções menores (`CalcularSalario`, `EnviarEmail`, `AtualizarBanco`) e criar uma quarta que só orquestra as outras três deixa tudo mais fácil de testar isoladamente, mais fácil de ler e mais fácil de reaproveitar quando a regra mudar só em uma das partes.

---

🔬 Desafios
Pra fixar o conteúdo, resolvi os três exercícios propostos no capítulo:

1. Três funções (`soma`, `subtracao`, `multiplicacao`) recebendo dois inteiros cada e devolvendo um inteiro, chamadas com os valores 10 e 5.
package main

import "fmt"

func soma(a, b int) int        { return a + b }
func subtracao(a, b int) int   { return a - b }
func multiplicacao(a, b int) int { return a * b }

func main() {
  fmt.Println(soma(10, 5), subtracao(10, 5), multiplicacao(10, 5))
}

2. Uma função que classifica um preço em "caro" ou "barato" devolvendo dois valores diferentes de uma vez — um `bool` e uma `string`.
func analisarPreco(preco float64) (bool, string) {
  if preco >= 100.0 {
    return true, "Caro"
  }
  return false, "Barato"
}

3. Conversão de segundos totais em horas, minutos e segundos, usando retorno nomeado — 3661 segundos precisa virar `1, 1, 1`.
func converterTempo(tempoSegundos int) (horas, minutos, segundos int) {
  horas = tempoSegundos / 3600
  minutos = tempoSegundos % 3600 / 60
  segundos = tempoSegundos % 60
  return
}

---

🎯 Conclusão
Saindo desse capítulo, o que fica é menos sobre sintaxe e mais sobre hábito: o Go não te dá exceções pra capturar em qualquer lugar, te dá um segundo valor de retorno pra checar logo ali. Parece pouco, mas muda completamente como você escreve (e lê) código de verdade. Na próxima parte da série a jornada continua. Gostou do conteúdo? O código fonte de todos os meus exemplos e desafios deste estudo está publicado no GitHub:
👉 [Repositório Go na Prática](https://github.com/jorgegabrielti/go-in-practice-book)
