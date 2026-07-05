> **Antes de publicar:** conferir checklist no fim do arquivo.

Dia X/100 do desafio #100DaysOfCodeWithAngie Angela Hempkmeier Ceccato concluído com sucesso!

Dando continuidade aos estudos de Go na Prática, hoje estudei Structs.

- Struct é a "Pasta Amarela" do Go: em vez de `nome1`, `idade1`, `peso1` voando soltos, você grampeia tudo numa entidade coerente — `type Pessoa struct { Nome string; Idade int; Peso float64 }` — e passa a pasta inteira para quem precisar.
- Sempre use a forma literal com nomes dos campos (`Pessoa{Nome: "João", Idade: 30}`). A forma posicional (`Pessoa{"João", 30, 80.5}`) economiza caracteres mas quebra silenciosamente se a ordem dos campos mudar. Aprendi da forma certa antes de aprender da forma errada.
- Ponteiro para Struct é o padrão mais comum em Go: `p2 := &p1` e depois `p2.Preco = 200` — o compilador desreferencia automaticamente, sem precisar de `(*p2).Preco`. Modificar via ponteiro altera o original.
- Structs se aninham naturalmente: `u.Endereco.Cidade` acessa em profundidade. Um campo pode ser de qualquer tipo, inclusive outro tipo struct.
- Go não tem herança. No lugar de `extends`, usa Embedding — você coloca uma struct dentro da outra sem nome de campo e os campos são "promovidos": `cachorro.Vida` funciona direto, como se fosse do próprio Cachorro. Composição no lugar de hierarquia de classes.

Resolvi três exercícios: Cadastro de Gamer (instância simples), Comparação de Retângulos (função que recebe struct por valor) e Sistema de Playlist (struct aninhada + loop + cálculo de duração total).

👉 Artigo detalhado no Medium: [Link do artigo no Medium]
#Golang #Programacao #CleanCode #Backend #Estudos

---

**Checklist antes de publicar:**
- [x] Conteúdo confere com `estudos-go/capitulo10/fonte.txt`
- [x] Texto é original — nenhuma frase copiada do fonte.txt
- [x] Gancho na primeira linha (dia + desafio + tema)
- [x] 5 bullets com frases curtas
- [ ] Substituir `X` pelo número do dia real
- [ ] Link do Medium preenchido (preencher após publicar o artigo)
- [x] Hashtags revisadas (5 relevantes)
- [ ] Status atualizado em `conteudo/PAINEL.md`
