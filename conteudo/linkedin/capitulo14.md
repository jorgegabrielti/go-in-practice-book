> **Antes de publicar:** conferir checklist no fim do arquivo.

Dia X/100 do desafio #100DaysOfCodeWithAngie Angela Hempkmeier Ceccato concluído com sucesso!

Dando continuidade aos estudos de Go na Prática, hoje aprendi sobre Pacotes e Módulos.

- Um **módulo** começa com `go mod init` — ele dá identidade ao projeto e gerencia dependências no `go.mod`. Um **pacote** é uma pasta com arquivos `.go` de propósito comum. Módulo é a oficina inteira; pacote é cada gaveta organizada.
- A regra de visibilidade do Go é genial na sua simplicidade: **letra maiúscula = exportado (público), letra minúscula = privado**. Sem `public`, `private` ou `protected` — o compilador reforça a fronteira automaticamente.
- Criar seu próprio pacote é só criar uma pasta: `conversor/medidas.go` com `package conversor`. O import usa o caminho completo: `"github.com/usuario/projeto/conversor"`. Nada de configuração extra.
- `go get github.com/...` instala qualquer biblioteca externa em três passos: baixa o código, registra no `go.mod` e cria o `go.sum` (checksum de segurança). O ecossistema Go é enorme — UUID, cores no terminal, HTTP, banco de dados, tudo disponível.
- `go mod tidy` é a faxina: lê o código, baixa o que falta e remove o que sobra do `go.mod`. Rodar antes de todo commit é boa prática.

Dica que ficou: evite `package utils`. Prefira nomes semânticos — `documentos`, `seguranca`, `relatorio`. Pacote bom responde "o que ele representa?", não "pra que ele serve?".

👉 Artigo detalhado no Medium: [Link do artigo no Medium]
#Golang #Programacao #GoModules #Backend #Estudos

---

**Checklist antes de publicar:**
- [x] Conteúdo confere com `estudos-go/capitulo14/fonte.txt`
- [x] Texto é original — nenhuma frase copiada do fonte.txt
- [x] Gancho na primeira linha (dia + desafio + tema)
- [x] 5 bullets com frases curtas
- [ ] Substituir `X` pelo número do dia real
- [ ] Link do Medium preenchido (preencher após publicar o artigo)
- [x] Hashtags revisadas (5 relevantes)
- [ ] Status atualizado em `conteudo/PAINEL.md`
