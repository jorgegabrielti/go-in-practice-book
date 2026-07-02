> **Antes de publicar:** conferir checklist no fim do arquivo.

Dia X/100 do desafio #100DaysOfCodeWithAngie Angela Hempkmeier Ceccato concluído com sucesso!

Dando continuidade aos estudos de Go na Prática, hoje entrei na Parte 2 do livro com Ponteiros.

- Ponteiro não guarda um valor — guarda *onde* o valor vive na memória. A analogia que grudou: você não entrega a casa para o amigo, você entrega o endereço escrito num papel. Se ele pintar uma parede, a sua casa real muda.
- `&variavel` pega o endereço; `*ponteiro` viaja até ele. São operações opostas e você as usa sempre juntas.
- Por padrão, Go passa tudo por valor (cópia). Se você quer que uma função modifique o original, precisa passar o endereço com `&` e receber com `*Tipo`.
- O zero value de um ponteiro é `nil`. Desreferenciar `nil` causa panic em runtime — a verificação `if p != nil` antes de usar é hábito que evita muita dor de cabeça.
- Regra de Ouro do capítulo: use ponteiro para modificar o original ou quando o dado for grande. Para `int`, `bool` e valores pequenos, cópia é mais rápida que ponteiro — o compilador do Go é inteligente e a CPU agradece.

Resolvi três exercícios: O Trocador (múltipla atribuição `*a, *b = *b, *a`), O Incrementador (contador compartilhado via ponteiro em loop) e O Detetive de Endereços (prova empírica de que passar por valor cria uma cópia em outro endereço).

👉 Artigo detalhado no Medium: [Link do artigo no Medium]
#Golang #Programacao #CleanCode #Backend #Estudos

---

**Checklist antes de publicar:**
- [x] Conteúdo confere com `estudos-go/capitulo09/fonte.txt`
- [x] Texto é original — nenhuma frase copiada do fonte.txt
- [x] Gancho na primeira linha (dia + desafio + tema)
- [x] 5 bullets com frases curtas
- [ ] Substituir `X` pelo número do dia real
- [ ] Link do Medium preenchido (preencher após publicar o artigo)
- [x] Hashtags revisadas (5 relevantes)
- [ ] Status atualizado em `conteudo/PAINEL.md`
