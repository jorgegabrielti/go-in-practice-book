# Pilha de Exercícios — Capítulo 04: Controle de Fluxo

> Conceitos disponíveis até aqui: tudo dos capítulos 01-03, mais `if`/`else`, `if` com inicialização curta, operadores lógicos (`&&`, `||`, `!`), `switch` clássico, `switch` tagless e `fallthrough`.

#### Exercício 01 🟢

Dado um valor fixo de IMC, use `if`/`else if` para classificá-lo em "Abaixo do peso", "Normal", "Sobrepeso" ou "Obesidade".

#### Exercício 02 🟢

Usando `if` com inicialização curta, verifique se um número fixo é múltiplo de 7 e imprima uma mensagem informando o resultado.

#### Exercício 03 🟢

Usando o operador `&&`, verifique (com valores fixos) se uma pessoa pode votar (maior de 16 anos) e se é obrigada a votar (entre 18 e 70 anos).

#### Exercício 04 🟡

Converta uma nota numérica fixa (0 a 10) em conceito (A, B, C, D ou E) usando um `switch` tagless com faixas de valores.

#### Exercício 05 🟡

Dado três lados fixos de um triângulo, use um `switch` tagless para classificá-lo em equilátero, isósceles ou escaleno.

#### Exercício 06 🟡

Monte um `switch` de "nível de acesso" (1 a 3) em que o nível 1 usa `fallthrough` para também imprimir a mensagem do nível 2. Documente em um comentário por que esse comportamento pode causar bugs em código real.

#### Exercício 07 🟡

Escreva primeiro uma versão "ruim" com três `if`s aninhados (deixe-a comentada) e depois refatore para a versão limpa usando *guard clauses* (early return).

#### Exercício 08 🔴

Expanda o desafio do Ano Bissexto do capítulo: além de dizer se o ano é bissexto, informe também quantos dias tem o mês de fevereiro naquele ano específico.

#### Exercício 09 🔴

Escreva um `switch` tagless sem cláusula `default` e teste com um valor que não se encaixa em nenhum `case`. Observe e documente em um comentário o que acontece quando nenhum caso bate.

#### Exercício 10 🔴

Usando operadores lógicos, valide (com valores booleanos fixos simulando os critérios) se uma senha "tem pelo menos 8 caracteres" E "contém ao menos um dígito". Use variáveis `bool` para representar cada critério já avaliado.
