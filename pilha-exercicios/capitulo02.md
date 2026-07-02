# Pilha de Exercícios — Capítulo 02: Variáveis, Constantes e Tipos de Dados

> Conceitos disponíveis até aqui: tudo do capítulo 01, mais `var`, `:=`, Zero Value, `const`, declaração agrupada e *shadowing*.

#### Exercício 01 🟢

Declare variáveis para nome, idade e cidade usando `:=` e imprima tudo com `fmt.Printf`, usando os verbos corretos (`%s`, `%d`) para cada tipo.

#### Exercício 02 🟢

Usando `var` (sem inicializar), declare uma variável de cada tipo — `int`, `float64`, `string`, `bool` — e imprima o Zero Value de cada uma.

#### Exercício 03 🟢

Declare uma `const` com o fator de conversão de quilômetros para milhas (`0.621371`) e use-a para converter uma distância fixa (ex: 100 km) para milhas.

#### Exercício 04 🟡

Declare duas variáveis `a` e `b` com valores diferentes e troque os valores entre elas usando uma terceira variável temporária. Imprima antes e depois da troca.

#### Exercício 05 🟡

Tente reatribuir um novo valor a uma `const` já declarada (deixe a linha comentada) e documente, em um comentário, a mensagem de erro que o compilador apresentaria.

#### Exercício 06 🟡

Reescreva o Exercício 01 usando exclusivamente `var` em vez de `:=`. Compare nos comentários qual versão você considera mais legível e por quê.

#### Exercício 07 🟡

Declare uma variável `total` no escopo do pacote (fora de `main`) e outra `total` dentro de `main` (*shadowing*). Imprima o valor de dentro de `main` e explique em um comentário por que a variável externa fica inacessível ali.

#### Exercício 08 🔴

Declare uma variável `idade` com um valor fixo e imprima o dobro, o triplo e a metade dela. Preste atenção ao calcular a metade: o resultado deveria ser decimal?

#### Exercício 09 🔴

Declare um bloco `const ( ... )` representando os dias da semana como números de 0 (Domingo) a 6 (Sábado) e imprima todos eles.

#### Exercício 10 🔴

Escreva um programa que assuma (de forma proposital e errada) que uma variável `bool` declarada sem valor seria `true` por padrão. Execute, observe o resultado real e corrija o comentário explicando qual é o Zero Value correto.
