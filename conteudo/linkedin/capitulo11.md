> **Antes de publicar:** conferir checklist no fim do arquivo.

Dia X/100 do desafio #100DaysOfCodeWithAngie Angela Hempkmeier Ceccato concluído com sucesso!

Dando continuidade aos estudos de Go na Prática, hoje aprendi sobre Métodos.

- No capítulo anterior, Structs eram como cachorros de pelúcia: tinham nome, cor, ocupavam espaço — mas não faziam nada. Métodos são o que fazem o cachorro latir sozinho. Em vez de `latir(rex)`, você escreve `rex.Latir()`. É a diferença entre carregar o objeto até a função e deixar o objeto agir por conta própria.
- Todo método tem um **Receiver** — o parâmetro especial antes do nome que amarra a função ao tipo. `func (c Cachorro) Latir()` diz: "Latir pertence exclusivamente ao tipo Cachorro". Sem herança, sem `this` implícito, tudo explícito e legível.
- A decisão central do capítulo: Value Receiver `(c Tipo)` faz uma cópia — seguro, somente leitura. Pointer Receiver `(c *Tipo)` passa o endereço — eficiente e permite modificar o original. Regra: na dúvida, use Pointer Receiver.
- Métodos não são exclusividade de Structs. Qualquer tipo criado com `type` aceita métodos — `type Dinheiro float64` pode ter um método `String()` que sempre formata como `"R$ X.XX"`. Abstração poderosa em cima de tipos primitivos.
- Sem Getters e Setters desnecessários. Campo com letra maiúscula já é público em Go — `p.Nome = "Ana"` é o caminho certo. Só crie um Setter quando houver validação real. Go preza pela simplicidade.

Resolvi três exercícios: O Carro Acelerado (Pointer Receiver modificando velocidade), O Relógio Digital (lógica de wrap de 60min/24h com módulo) e A Calculadora Orientada a Métodos (acúmulo de resultado via Pointer Receiver).

👉 Artigo detalhado no Medium: [Link do artigo no Medium]
#Golang #Programacao #CleanCode #Backend #Estudos

---

**Checklist antes de publicar:**
- [x] Conteúdo confere com `estudos-go/capitulo11/fonte.txt`
- [x] Texto é original — nenhuma frase copiada do fonte.txt
- [x] Gancho na primeira linha (dia + desafio + tema)
- [x] 5 bullets com frases curtas
- [ ] Substituir `X` pelo número do dia real
- [ ] Link do Medium preenchido (preencher após publicar o artigo)
- [x] Hashtags revisadas (5 relevantes)
- [ ] Status atualizado em `conteudo/PAINEL.md`
