# 📖 Capítulo 16: Channels - O Mensageiro Confiável

> **Livro Go na Prática: 30 capítulos para dominar a linguagem**

---

> **Fluxo de trabalho deste capítulo:**
> 1. Transcrever o conteúdo do livro para `fonte.txt` (texto bruto, fiel ao material original).
> 2. Escrever este `README.md` como síntese própria a partir do `fonte.txt`.
> 3. Resolver exemplos/exercícios em `exemplos/` e `exercicios/`.
> 4. Produzir o artigo de Medium e o post de LinkedIn em `conteudo/`, **sempre baseados no `fonte.txt`** — não em memória ou suposição do que o livro disse, e **escritos com palavras próprias** (sem copiar/parafrasear de perto o `fonte.txt`).
> 5. Rastrear a execução de todo o código do capítulo em `teste-de-mesa.md`, prevendo a saída linha a linha antes de rodar `go run`.

> 🧮 [Teste de mesa de todo o código deste capítulo](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo16/teste-de-mesa.md)

---

## 🎯 Tema Central

Canais (Channels) em Go como mecanismo primário de comunicação e sincronização de dados entre Goroutines, permitindo compartilhar dados de forma segura sem a necessidade de travas manuais (como mutexes) ou compartilhamento direto de memória.

---

## 📊 Resumo dos Conceitos

- **O Mantra do Go**: *"Não compartilhe memória para se comunicar. Comunique-se para compartilhar memória."*
- **Tubo Tipado**: Canais são fortemente tipados (`chan T`). Apenas dados do tipo definido podem trafegar neles.
- **Sintaxe da Seta (`<-`)**:
  - Enviar dados: `c <- valor` (seta aponta para o canal)
  - Receber dados: `valor := <-c` (seta aponta para fora do canal)
- **Bloqueio (Blocking)**: Por padrão, canais em Go não possuem buffer (*unbuffered*). Isso significa que operações de envio (`c <- x`) ou recebimento (`<-c`) bloqueiam a goroutine atual até que a outra parte esteja pronta para realizar a operação oposta.
- **Deadlock**: Ocorre se tentarmos ler/escrever em canais numa única goroutine (como na `main` sozinha) ou se todas as goroutines estiverem dormindo esperando por operações em canais que nunca ocorrerão.
- **Fechamento de Canal (`close`)**: O remetente pode fechar um canal (`close(c)`). Leitores podem ler continuamente usando `for range`, que encerra o loop automaticamente assim que o canal for fechado.
- **Canais Direcionais**: Restrição de uso de canais para segurança de tipos:
  - `chan<- T`: Apenas envio (Send-only)
  - `<-chan T`: Apenas recebimento (Receive-only)

---

## 💡 Dica do Gopher

Sempre delegue a responsabilidade de fechar o canal ao **remetente** (quem envia). Enviar dados em um canal fechado causa um *panic*, enquanto ler de um canal fechado retorna imediatamente o valor zero do tipo do canal.

---

## 🔬 Exemplos Práticos no Repositório

- [Exemplo 1: Jogo de Ping Pong](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo16/exemplos/exe01/main.go) - Demonstração de sincronização perfeita entre duas goroutines simulando a troca de uma bola através de um canal compartilhado.
- [Exemplo 2: Produtor e Consumidor](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo16/exemplos/exe02/main.go) - Uso de canais direcionais, fechamento com `close` e iteração segura com `for range`.

---

## 🔬 Exercícios Práticos Resolvidos

- [Exercício 1: O Correio Elegante](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo16/exercicios/exe01/main.go) - Envio de múltiplas strings em goroutine e leitura manual em `main`, com explicação sobre o comportamento de Deadlock.
- [Exercício 2: A Soma Distribuída](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo16/exercicios/exe02/main.go) - Divisão de um slice ao meio para cálculo de soma paralela em duas goroutines, combinando o resultado final na `main`.
- [Exercício 3: O Temporizador Manual](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo16/exercicios/exe03/main.go) - Implementação de um temporizador de 2 segundos que desbloqueia a execução da `main` através de sinalização por canal.

---

## ✅ Checklist antes de marcar como concluído

- [x] `fonte.txt` com a transcrição completa do capítulo
- [x] Teoria revisada e resumida neste README (com base no `fonte.txt`)
- [x] Todos os exemplos do capítulo têm pasta própria em `exemplos/`
- [x] Todos os exercícios resolvidos têm pasta própria em `exercicios/`
- [x] `go build ./...` e `go vet ./...` passam sem erros
- [x] `go fmt ./...` executado
- [ ] Termos novos adicionados ao `estudos-go/GLOSSARIO.md`
- [ ] Painel de progresso no `README.md` raiz atualizado
- [ ] Artigo de Medium criado em `conteudo/medium/capitulo16.md`
- [ ] Post de LinkedIn criado em `conteudo/linkedin/capitulo16.md`
- [ ] Pilha de exercícios extras criada em `pilha-exercicios/capitulo16.md`
- [ ] Teste de mesa criado em `teste-de-mesa.md`
- [ ] Fonte NotebookLM criada em `notebooklm/capitulo16.md`
