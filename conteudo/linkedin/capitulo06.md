🚀 Capítulo 6 da jornada Go na Prática concluído: Funções! 🔧

O que mais me chamou atenção neste capítulo não foi a sintaxe de função em si, mas como o Go trata erro de forma diferente:

- Função em Go pode devolver mais de um valor ao mesmo tempo — e é assim que erro deixa de ser exceção e passa a ser só mais um valor de retorno (`valor, err := funcao()`).
- Retorno nomeado permite o "naked return" — útil em função curta, perigoso em função longa.
- Parâmetro variádico (`...int`) deixa uma função aceitar de zero a N argumentos, igual o `fmt.Println` faz.
- E função em Go é valor: dá pra guardar uma função anônima numa variável e usá-la como callback.

Desafios resolvidos pra fixar: a calculadora com soma/subtração/multiplicação, o analisador de preço (retornando dois valores diferentes de uma vez) e o conversor de segundos para horas/minutos/segundos.

Confira todos os códigos e notas de estudo:
👉 Repositório no GitHub: https://github.com/jorgegabrielti/go-in-practice-book
👉 Artigo detalhado no Medium: [Link do artigo no Medium]

#Golang #Programacao #CleanCode #Backend #Estudos
