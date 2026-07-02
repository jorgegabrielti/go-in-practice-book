---

Controle de fluxo em Go: Tomando decisões de forma inteligente e eficiente

Seja muito bem-vindo ao 4º artigo da série Go na Prática. Até aqui, aprendemos sobre variáveis, constantes e tipos primitivos de dados. Nossos programas já sabem armazenar valores, mas eles ainda sofrem de uma limitação grave: eles só sabem executar instruções de cima para baixo, sem poder desviar ou contornar obstáculos.
Na vida real, o software precisa tomar decisões. É aí que entra o Controle de Fluxo, que funciona como um verdadeiro guarda de trânsito em nosso código, orientando para qual direção a execução deve seguir. Em Go, a palavra-chave é simplicidade: as estruturas condicionais foram reduzidas ao essencial para oferecer o máximo de performance e clareza de leitura.
Neste artigo, vamos explorar como Go simplifica a tomada de decisões com condicionais enxutas, switches inteligentes e boas práticas de estruturação de código.

---

1. 🚦 Condicionais Clássicos: `if` e `else` Sem Burocracia
Em linguagens como Java, C# ou JavaScript, somos obrigados a envolver a condição de um teste lógico entre parênteses. No Go, os criadores da linguagem eliminaram esse caractere redundante. Se a condição está explícita, o compilador consegue identificá-la sem precisar de parênteses.
Vejamos um exemplo simples de validação de acesso:
package main

import "fmt"

func main(){
  idade := 17
   
  if idade >= 18 {
    fmt.Println("Entrada permitida!")
  } else {
    fmt.Println("Entrada proibida!") 
  }
}
⚠️ Uma regra de formatação que não abre margem para Discussão
Em Go, a chave de abertura `{` deve obrigatoriamente estar na mesma linha do termo `if` ou `else`. Se você tentar quebrá-la para a linha de baixo, o compilador recusará a execução do código. A intenção por trás disso é louvável: unificar o estilo de escrita da comunidade e evitar discussões intermináveis (e inúteis) sobre estilo de identação em reuniões de time.
Para cenários com múltiplas alternativas, usamos a estrutura encadeada `else if`:
...
temperatura := 25

if temperatura > 30 {
  fmt.Println("Resfriar ambiente.")
} else if temperatura > 20 {
  fmt.Println("Temperatura sob controle.")
} else {
  fmt.Println("Aquecer ambiente.")
}
...

---

2. 💎 A Joia do escopo: `if` com inicialização curta (Short Statement)
Uma das técnicas mais elegantes do Go é a inicialização curta de variáveis dentro do próprio bloco do `if`. Muitas vezes, criamos uma variável auxiliar apenas para fazer um teste imediato. Se ela for criada no escopo principal, ela continuará poluindo a memória até o término da função.
Com o Short Statement, nós declaramos a variável temporária e a avaliamos na mesma linha:
// Sintaxe: if declaracao; condicao { … }
if resultado := obterStatus(); resultado == "erro" {
  fmt.Println("Houve uma falha interna.")
}
Esse recurso é muito usado em Go para tratamento de erros e chamadas rápidas. O compilador garante que a variável morra assim que a chave de fechamento `}` do condicional for executada, ajudando o coletor de lixo a manter o consumo de memória sob controle.

---

3. 🧠 Operadores Lógicos e a Inteligência do Curto-Circuito
Para combinar decisões lógicas mais complexas, contamos com três operadores padrão:
&& (AND): Todas as condições conjuntas precisam ser verdadeiras.
|| (OR): Pelo menos uma das condições precisa ser verdadeira.
! (NOT): Inverte o valor lógico de uma expressão.

...
temCinto := true
estaAlcoolizado := false

if temCinto && !estaAlcoolizado {
  fmt.Println("Direção segura autorizada.")
}
...
⚡Curto-Circuito: Performance por Padrão
O compilador de Go avalia as expressões lógicas usando a estratégia de curto-circuito. Em uma operação `A && B`, se o termo `A` for avaliado como falso, o Go nem sequer perde tempo processando o termo `B`, pois sabe que a expressão completa obrigatoriamente resultará em falso. Da mesma forma, em `A || B`, se `A` for verdadeiro, `B` é ignorado.

---

4. 🎛️ O `switch` Moderno: Código sem poluição visual
Quando precisamos tratar muitos caminhos possíveis, a escada de `else if` se torna um pesadelo de leitura. É para isso que serve o `switch`. Em Go, o `switch` foi repensado para corrigir as maiores dores das linguagens clássicas.
Principais Vantagens do Switch em Go:
1. Fim do `break` manual: Esqueça o pesadelo de deixar passar o fluxo e disparar casos indesejados. O Go insere um `break` implícito ao final de cada `case`.
2. Múltiplos valores: É possível associar várias correspondências no mesmo bloco separando-as por vírgula.
3. Qualquer tipo comparável: Aceita strings, inteiros, decimais ou structs.

dia := "sabado"

switch dia {

case "segunda", "terça", "quarta", "quinta", "sexta":
  fmt.Println("Dia de trabalho regular.")

case "sabado", "domingo":
  fmt.Println("Fim de semana! Tempo de lazer.")

default:
  fmt.Println("Dia inválido.")
}
O Poder do Switch Sem Expressão (Tagless)
Um padrão de design extremamente idiomático em Go é escrever o `switch` sem definir uma variável ao lado da palavra-chave. Isso é equivalente a um `switch true`, onde o Go testará a validade de cada `case` individual e executará o primeiro que for avaliado como verdadeiro. É o substituto ideal para longos encadeamentos condicionais:
switch {

case nota >= 9.0:
  fmt.Println("Conceito Excelente!")

case nota >= 7.0:
  fmt.Println("Conceito Aprovado.")

default:
  fmt.Println("Conceito Reprovado.")
}

---

5. 💡 Dica do Gopher: Evite aninhamentos com Early Return
Códigos com dezenas de `if`s aninhados - o famoso padrão "código pirâmide" ou "caminho de flecha" - são difíceis de depurar e manter. O design idiomático do Go prioriza o chamado caminho feliz (happy path).
Em verdade, a boa prática indica tratar erros ou retornos rápidos no início da função usando cláusulas de guarda (*Guard Clauses*) e interromper o fluxo com um `return` rápido. Com isso, evitamos identações profundas e tornamos o código limpo, linear e muito fácil de ler.

---

🔬 Desafios
Como parte dos meus estudos deste capítulo, resolvi três desafios práticos que consolidaram estes conceitos:

1. O Validador de Acesso: Uma lógica clássica de segurança cruzando login e senha.
package main

import "fmt"

var login string = "admin"
var senha string = "12345"

func main() {
  if login == "admin" && senha == "12345" {
    fmt.Println("Acesso liberado!")
  } else {
    fmt.Println("Login ou senha inválidos.")
  }
}

2. A Tabela Periódica: Um conversor de siglas de elementos químicos para nomes completos usando switch com controle padrão (default).
package main

import "fmt"

var sigla string = "He"

func main() {
  switch sigla {
  case "He":
    fmt.Println("He => Hélio")
  case "O":
    fmt.Println("O => Oxigênio")
  case "Li":
    fmt.Println("Li => Lítio")
  default:
    fmt.Println("Elemento desconhecido")
  }
}

3. O Detetive de Ano Bissexto: Uma lógica de divisibilidade matemática implementada de forma limpa para determinar se anos informados são bissextos.
package main

import "fmt"

func main() {
  ano := 2026

  switch {
  case ano%4 == 0:
    fmt.Println("Ano bissexto")
  default:
    fmt.Println("Ano não bissexto")
  }
}

---

🎯 Conclusão
Go nos mostra, a cada etapa, que o design de uma linguagem de programação pode nos induzir a escrever códigos melhores. Ao remover parênteses redundantes, simplificar o `switch` e desencorajar ninhos profundos de condicionais, a linguagem nos ajuda a focar no que realmente importa: resolver o problema com o máximo de clareza.
Gostou do conteúdo? O código fonte de todos os meus exemplos e desafios deste estudo está publicado no GitHub:
👉 [Repositório Go na Prática](https://github.com/jorgegabrielti/go-in-practice-book)
No próximo capítulo, continuaremos a nossa jornada explorando o único e versátil laço de repetição do Go: O Loop `for`. Até a próxima! 🦫✨