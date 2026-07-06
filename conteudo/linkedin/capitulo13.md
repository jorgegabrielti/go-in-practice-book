> **Antes de publicar:** conferir checklist no fim do arquivo.

Dia X/100 do desafio #100DaysOfCodeWithAngie Angela Hempkmeier Ceccato concluído com sucesso!

Dando continuidade aos estudos de Go na Prática, hoje aprendi sobre Tratamento de Erros.

- Em Go, erro não é uma exceção — é um valor de retorno como qualquer outro. Uma função retorna `(Resultado, error)`, você verifica `if err != nil` e decide o que fazer. Sem `try/catch`, sem fluxo mágico. Explícito por design.
- Sentinel errors são variáveis de erro nomeadas que você pode comparar com `errors.Is`. Com `fmt.Errorf` e o verbo `%w`, você embrulha erros sem perder a causa raiz — o rastreamento funciona mesmo através de várias camadas.
- `panic` é o botão de emergência — para situações que **não deveriam existir**: estado corrompido, bug lógico grave. Para cenários normais como "arquivo não encontrado" ou "senha inválida", use `error`. Regra simples: se é previsível, é `error`; se é impossível, é `panic`.
- `recover` captura um `panic` antes que ele derrube o processo — mas só funciona dentro de um `defer`. Uso típico: servidores web que precisam isolar falhas de uma requisição sem afetar as demais.
- `defer` agenda uma função para rodar no último segundo antes do retorno, não importa como a função termina. É assim que você garante limpeza de recursos sem `try/finally`. Múltiplos defers empilham em LIFO — o último registrado é o primeiro a executar.

O padrão "Line of Sight": trate o erro, retorne cedo, deixe o caminho feliz na esquerda sem `else`. Código Go idiomático lê de cima para baixo sem aninhamento.

👉 Artigo detalhado no Medium: [Link do artigo no Medium]
#Golang #Programacao #ErrorHandling #Backend #Estudos

---

**Checklist antes de publicar:**
- [x] Conteúdo confere com `estudos-go/capitulo13/fonte.txt`
- [x] Texto é original — nenhuma frase copiada do fonte.txt
- [x] Gancho na primeira linha (dia + desafio + tema)
- [x] 5 bullets com frases curtas
- [ ] Substituir `X` pelo número do dia real
- [ ] Link do Medium preenchido (preencher após publicar o artigo)
- [x] Hashtags revisadas (5 relevantes)
- [ ] Status atualizado em `conteudo/PAINEL.md`
