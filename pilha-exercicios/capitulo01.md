# Pilha de Exercícios — Capítulo 01: Introdução ao Go

> Conceitos disponíveis até aqui: `package main`, `import`, `func main`, `fmt.Println`/`Print`/`Printf` (com valores literais), regra de exportação (maiúscula/minúscula), `go fmt`. Ainda **não** use variáveis, `if` ou `for` — isso vem nos próximos capítulos.

#### Exercício 01 🟢

Faça um programa que imprima, em três linhas separadas (três chamadas de `fmt.Println`), seu nome, sua cidade e o livro que você está estudando.

#### Exercício 02 🟢

Usando várias chamadas de `fmt.Println`, desenhe uma linha horizontal feita de `=` (por exemplo, 20 caracteres) acima e abaixo da frase "Bem-vindo ao Go!".

#### Exercício 03 🟢

Usando apenas `fmt.Println`, desenhe um quadrado 5x5 feito de asteriscos (`*`), uma linha por `Println`.

#### Exercício 04 🟡

Use `fmt.Print` (sem quebra de linha) três vezes seguidas para montar uma única frase na tela, e finalize com `fmt.Println` para garantir a quebra de linha no final.

#### Exercício 05 🟡

Usando `fmt.Printf` com os verbos `%s` e `%d`, imprima uma frase contendo um nome e uma idade, ambos escritos diretamente como literais (sem variáveis ainda).

#### Exercício 06 🟡

Escreva um pequeno poema de 4 linhas usando uma *raw string* (entre crases) e imprima-o com uma única chamada de `fmt.Println`.

#### Exercício 07 🟡

Adicione um comentário de uma linha (`//`) explicando, em suas próprias palavras, o que cada uma das quatro partes do "Olá Mundo" faz: `package main`, `import "fmt"`, `func main()` e `fmt.Println(...)`.

#### Exercício 08 🔴

Provoque (e depois corrija) três erros de compilação diferentes dos vistos no capítulo: troque `func main` por `func Main`, esqueça de fechar um parêntese em `fmt.Println(`, e remova as aspas de uma string. Anote em comentários qual foi a mensagem de erro do compilador em cada caso.

#### Exercício 09 🔴

Escreva um programa de propósito mal formatado (chaves na linha errada, espaços em vez de tabs, linhas desalinhadas). Rode `go fmt` nele e compare o antes/depois. Documente em um comentário o que mudou.

#### Exercício 10 🔴

Pesquise (sem precisar usar no código) a diferença entre o `Println` do pacote `fmt` e o `println` *builtin* da linguagem (sem necessidade de `import`). Documente a diferença em um comentário no topo do arquivo.
