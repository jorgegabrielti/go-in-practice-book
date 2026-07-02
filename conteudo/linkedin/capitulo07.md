> **Antes de publicar:** conferir checklist no fim do arquivo.

Dia X/100 do desafio #100DaysOfCodeWithAngie Angela Hempkmeier Ceccato concluído com sucesso!

Dando continuidade aos estudos de Go na Prática, hoje estudei sobre Arrays e Slices.

- Array tem tamanho fixo definido na declaração — e o tamanho faz parte do tipo: `[3]int` e `[4]int` são tipos diferentes e incompatíveis.
- Slice é a estrutura do dia a dia: dinâmico, cresce com `append` e por baixo dos panos vive sobre um Array real.
- O `append` sempre precisa ser reatribuído: `slice = append(slice, valor)` — sem o `=`, o novo valor é simplesmente descartado.
- Fatiar um Slice (`slice[min:max]`) **não copia os dados** — os dois Slices compartilham o mesmo Array. Mudar um muda o outro. Esse foi o bug mais silencioso que encontrei até agora nessa jornada.
- A forma segura de combinar partes de um Slice é criar um destino novo com `make()` e usar `append` — assim o original fica intacto.

Resolvi três desafios pra fixar: a Lista de Convidados (slice + loop + append), o Removedor de Itens (fatiamento seguro com make) e o Matador de Duplicatas (loops aninhados com variável bandeirinha).

👉 Artigo detalhado no Medium: [Link do artigo no Medium]
#Golang #Programacao #CleanCode #Backend #Estudos

---

**Checklist antes de publicar:**
- [x] Conteúdo confere com `estudos-go/capitulo07/fonte.txt` (sem invenções ou lacunas)
- [x] Texto é original — nenhuma frase copiada ou parafraseada de perto do `fonte.txt`
- [x] Gancho na primeira linha (emoji + número do capítulo + tema)
- [x] Máximo 5 bullets, frases curtas
- [ ] Link do GitHub e do Medium preenchidos (sem placeholder — preencher após publicar o artigo)
- [x] Hashtags revisadas (3 a 5, relevantes)
- [ ] Status atualizado em `conteudo/PAINEL.md`
