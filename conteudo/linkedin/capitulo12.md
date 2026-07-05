> **Antes de publicar:** conferir checklist no fim do arquivo.

Dia X/100 do desafio #100DaysOfCodeWithAngie Angela Hempkmeier Ceccato concluído com sucesso!

Dando continuidade aos estudos de Go na Prática, hoje aprendi sobre Interfaces.

- Interfaces em Go são contratos implícitos: nenhum `implements`, nenhuma assinatura no cartório. Se o seu tipo tem o método que a interface exige, ele já é compatível — o compilador verifica e pronto. É o que chamam de Duck Typing: "se anda como pato e grasna como pato, é um pato."
- O poder real das interfaces é o Polimorfismo. Uma função `func IncomodarVizinho(b Barulhento)` aceita um cachorro e um alarme sem saber a diferença entre eles — só sabe chamar `.FazerBarulho()`. Adicionar um terceiro tipo amanhã não muda uma linha dessa função.
- A interface vazia `any` aceita absolutamente tudo (é assim que `fmt.Println` funciona). Mas usar `any` em excesso é transformar Go em JavaScript — você perde o compilador como guardião. Use com propósito.
- Type Assertion (`v.(Tipo)`) e Type Switch (`switch v.(type)`) são as ferramentas para "abrir a caixa" e trabalhar com o tipo concreto dentro de uma interface. A forma segura é sempre com comma ok: `dog, ok := v.(Cachorro)`.
- A dica mais valiosa do capítulo: defina interfaces onde elas são **usadas**, não onde são criadas. Interfaces pequenas (1 ou 2 métodos) são as mais poderosas — `io.Reader` e `fmt.Stringer` são a prova.

Resolvi três exercícios: A Impressora (polimorfismo com slice de interfaces), O Processador de Pagamentos (Boleto e Cartão tratados pela mesma função) e O Type Switch (classificar qualquer valor sem panic).

👉 Artigo detalhado no Medium: [Link do artigo no Medium]
#Golang #Programacao #CleanCode #Backend #Estudos

---

**Checklist antes de publicar:**
- [x] Conteúdo confere com `estudos-go/capitulo12/fonte.txt`
- [x] Texto é original — nenhuma frase copiada do fonte.txt
- [x] Gancho na primeira linha (dia + desafio + tema)
- [x] 5 bullets com frases curtas
- [ ] Substituir `X` pelo número do dia real
- [ ] Link do Medium preenchido (preencher após publicar o artigo)
- [x] Hashtags revisadas (5 relevantes)
- [ ] Status atualizado em `conteudo/PAINEL.md`
