# Pilha de Exercícios — Capítulo 06: Funções

> Conceitos disponíveis até aqui: tudo dos capítulos 01-05, mais funções com parâmetros e retorno único, múltiplos retornos, retornos nomeados/naked return, funções variádicas e funções anônimas/closures.

#### Exercício 01 🟢

Crie uma função `ehPar(n int) bool` que retorna `true` se o número for par. Na `main`, teste com três valores fixos diferentes e imprima o resultado de cada chamada.

#### Exercício 02 🟢

Crie uma função `saudacao(nome string, idade int) string` que retorna uma frase formatada combinando os dois parâmetros. Chame-a com valores fixos e imprima o retorno.

#### Exercício 03 🟢

Crie uma função `maiorEntre(a, b int) int` que retorna o maior dos dois valores recebidos, usando os parâmetros simplificados (mesmo tipo, sem repetir `int`).

#### Exercício 04 🟡

Crie uma função `dividirComResto(a, b int) (int, int)` que retorna o quociente e o resto da divisão inteira em um único `return`. Teste com 17 e 5.

#### Exercício 05 🟡

Crie uma função `validarSenha(senha string) (bool, string)` que retorna se a senha é válida (mínimo 8 caracteres, usando `len`) e uma mensagem explicando o motivo da reprovação (ou "Senha válida").

#### Exercício 06 🟡

Usando retornos nomeados e "naked return", crie uma função `estatisticas(notas ...float64) (soma, media float64)` que recebe uma quantidade variável de notas e calcula soma e média.

#### Exercício 07 🟡

Crie uma função variádica `concatenar(separador string, palavras ...string) string` que junta as palavras recebidas usando o separador entre elas, sem usar `strings.Join` (faça a concatenação manualmente em um loop com `range`).

#### Exercício 08 🔴

Crie uma função `aplicarDesconto(preco float64) (precoFinal float64, erro error)` que retorna erro se o preço for negativo, e aplica 10% de desconto caso contrário. No `main`, trate o erro com o padrão `resultado, err := funcao()` para os dois casos (preço válido e inválido).

#### Exercício 09 🔴

Crie uma função anônima dentro da `main`, guardada em uma variável `validarIdade`, que recebe um `int` e retorna um `bool` indicando se a pessoa é maior de idade (≥18). Use essa função anônima dentro de um `for` para classificar um slice fixo de idades.

#### Exercício 10 🔴

Crie uma função "gerente" `processarPedido(valor float64) (float64, error)` que internamente chama duas funções menores e privadas — uma para validar o valor (não pode ser zero ou negativo) e outra para calcular um imposto fixo de 5% — seguindo a Dica do Gopher de funções pequenas com responsabilidade única. Documente em um comentário por que separar essa lógica em três funções facilita testar cada parte isoladamente.
