> **Antes de publicar:** conferir checklist no fim do arquivo.

Dia X/100 do desafio #100DaysOfCodeWithAngie Angela Hempkmeier Ceccato concluído com sucesso!

Dando continuidade aos estudos de Go na Prática, hoje estudei sobre Maps.

- Map é uma estrutura chave→valor onde a busca acontece em O(1) — independente de quantas entradas existam. Por baixo, o Go implementa isso como uma Hash Table: uma função matemática transforma a chave num índice de array e vai direto ao dado.
- `make(map[string]int)` é a forma correta de criar um map. Declarar sem `make` cria um map `nil` — escrever nele causa panic em runtime, um dos erros mais silenciosos para quem está começando.
- Acessar uma chave que não existe não gera erro: o Go devolve o zero value do tipo (`0` para `int`). Isso permite o padrão `contador[palavra]++` sem nenhum `if` de verificação — na primeira ocorrência pega `0`, soma `1`, guarda `1`. Elegante demais.
- O comma ok idiom (`traducao, ok := mapa[chave]`) é a forma de distinguir "chave ausente" de "chave com zero value". Aprendi aqui e já vi o mesmo padrão em type assertions e channels — é uma das ideias que o Go usa de forma consistente em toda a linguagem.
- Maps não preservam ordem de inserção. Para exibir ordenado, o jeito é extrair as chaves para um slice e usar `sort.Strings()` antes de iterar.

Resolvi dois desafios pra fixar: o Contador de Palavras (zero value + `contador[palavra]++`) e o Dicionário de Cores (comma ok idiom para busca segura).

👉 Artigo detalhado no Medium: [Link do artigo no Medium]
#Golang #Programacao #CleanCode #Backend #Estudos

---

**Checklist antes de publicar:**
- [x] Texto é original — nenhuma frase copiada ou parafraseada de perto
- [x] Gancho na primeira linha (dia + desafio + tema)
- [x] Máximo 5 bullets, frases curtas
- [ ] Substituir `X` pelo número do dia real
- [ ] Link do Medium preenchido (preencher após publicar o artigo)
- [x] Hashtags revisadas (5 relevantes)
- [ ] Status atualizado em `conteudo/PAINEL.md`
