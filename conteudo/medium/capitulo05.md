---

Repetição sem complicações em Go: Dominando o Loop for

Seja muito bem-vindo ao 5º artigo da série Go na Prática. Desta vez, vamos explorar o poder da automação.
Computadores são incrivelmente rápidos e precisos por uma razão muito simples: eles adoram repetir tarefas. Ao contrário de nós, humanos, eles não ficam entediados, não sentem fadiga e podem executar a mesma instrução bilhões de vezes por segundo sem errar.
Na maioria das linguagens de programação populares, quando você estuda estruturas de repetição, depara-se com um cardápio confuso: while, do-while, for, for-each, repeat-until. O Go, fiel à sua filosofia de simplicidade e clareza, resolveu unificar tudo.
Em Go, só existe uma palavra-chave para repetições: for.
No entanto, essa única palavra-chave é extremamente flexível e se comporta de diferentes formas conforme a sintaxe que você adota. Vamos entender como pilotar os loops em Go!

---

1. 🏁 O for Clássico: O Piloto de Fórmula 1
Se você já programou em C, C++, Java ou JavaScript, o for clássico é um velho conhecido. Ele se comporta como uma corrida estruturada com três partes essenciais:
1. Inicialização: A definição de onde começamos (ex: i := 0).
2. Condição: A regra que diz se continuamos correndo. O loop continua ativo enquanto a condição for verdadeira.
3. Pós-execução: O ajuste que fazemos ao final de cada volta (ex: i++).

package main

import "fmt"

func main() {
  for i := 0; i < 5; i++ {
    fmt.Println("Volta número:", i)
  }
}

💡 Escopo Limpo
A variável i declarada na inicialização do for só existe dentro das chaves do loop. O compilador garante que ela seja limpa da memória assim que o loop termina.

---

2. 🏃‍♂️ O for como while: O Maratonista Condicional
Em muitas linguagens, usamos o while quando não sabemos exatamente quantas voltas daremos, apenas sabemos que devemos continuar repetindo enquanto uma condição for atendida.
Em Go, você cria esse comportamento omitindo a inicialização e o passo de pós-execução, deixando apenas a condição lógica:
...
energia := 100

for energia > 0 {
  fmt.Println("Mantendo o ritmo... Energia:", energia)
  energia -= 20
}
...
Sem a necessidade de uma nova palavra-chave while, a leitura flui naturalmente como: "enquanto a energia for maior que zero, execute".

---

3. ♾️ O Loop Infinito: O Servidor Eterno
Existem momentos na arquitetura de sistemas onde seu programa precisa rodar indefinidamente. É o caso de servidores HTTP aguardando novas conexões de rede ou robôs monitorando sinais vitais de hardware.
No Go, um loop infinito é declarado de forma minimalista, bastando omitir qualquer termo:
...
for {
  fmt.Println("Aguardando novas tarefas...")
  
  if tarefaConcluida() {
    break
  }
}
...

---

4. 🛑 Controlando o Fluxo: break e continue
Mesmo dentro de um ciclo de repetição ativo, podemos alterar seu comportamento com duas palavras-chave:
break: Sai do loop na mesma hora, ignorando qualquer condição ou volta restante. Ele interrompe o bloco e passa a execução para a próxima instrução fora do loop.
continue: Aborta a execução do código desta volta específica e salta direto para o topo do loop, prosseguindo com a próxima iteração.
...
for i := 1; i <= 5; i++ {
  if i == 3 {
    fmt.Println("Pulando o processamento do número 3.")
    continue
  }
  fmt.Println("Item processado:", i)
}
...

---

5. 🗺️ Percorrendo Coleções com range
Quando lidamos com coleções de dados (como arrays, slices, strings ou maps), o Go nos fornece uma ferramenta fantástica chamada range. Ela percorre a estrutura automaticamente, nos fornecendo o índice e o valor de cada item a cada iteração:
package main

import "fmt"

func main() {
  nomes := []string{"Ana", "Beto", "Carla"}

  for indice, nome := range nomes {
    fmt.Printf("Posição %d: %s\n", indice, nome)
  }
}

🎯 O Gopher e o Blank Identifier _
O Go é muito exigir quanto ao uso de variáveis declaradas. Se você capturar o índice e o valor com o range, mas usar apenas o valor, seu programa não compilará.
Para resolver isso, usamos o Blank Identifier (_) para descartar explicitamente o que não precisamos:
// Descartando o índice na leitura da coleção
for _, nome := range nomes {
  fmt.Println("Nome:", nome)
}

---

🔬 Desafios
Como parte dos meus estudos deste capítulo, resolvi três desafios práticos que consolidaram estes conceitos:

1. A Tabuada: Um loop clássico que imprime a tabuada de um número de forma organizada de 1 a 10.
package main

import "fmt"

func main() {
  numero := 7
  for i := 1; i <= 10; i++ {
    resultado := numero * i
    fmt.Printf("%d x %d = %d\n", numero, i, resultado)
  }
}

2. O Desafio FizzBuzz: O tradicional teste lógico de lógica matemática que avalia divisores de 3 ("Fizz"), 5 ("Buzz") e ambos ("FizzBuzz") em um intervalo de 1 a 50.
package main

import "fmt"

func main() {
  for i := 1; i <= 50; i++ {
    if i%3 == 0 && i%5 == 0 {
      fmt.Println("FizzBuzz")
    } else if i%3 == 0 {
      fmt.Println("Fizz")
    } else if i%5 == 0 {
      fmt.Println("Buzz")
    } else {
      fmt.Println(i)
    }
  }
}

3. Soma dos Dígitos: Um excelente algoritmo lógico usando divisões sucessivas (/ 10) e resto da divisão (% 10) dentro de um loop condicional para somar os algarismos de um número inteiro grande.
package main

import "fmt"

func main() {
  numero := 12345
  soma := 0
  for numero > 0 {
    soma += numero % 10
    numero /= 10
  }

  fmt.Println("Soma dos digitos: ", soma)
}

---

🎯 Conclusão
Ao simplificar todas as estruturas de repetição em torno da palavra-chave for, o Go reduz drasticamente a carga cognitiva de quem escreve ou lê o código. Menos sintaxes para decorar significam menos espaço para erros e uma produtividade muito maior.
Gostou do conteúdo? O código fonte de todos os meus exemplos e desafios deste estudo está publicado no GitHub:
👉 [Repositório Go na Prática](https://github.com/jorgegabrielti/go-in-practice-book)
No próximo capítulo, entraremos de cabeça no universo das coleções estruturadas, desvendando Arrays e Slices no Go. Vejo você lá! 🦫✨
