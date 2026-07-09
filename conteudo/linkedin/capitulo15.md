> **Antes de publicar:** conferir checklist no fim do arquivo.

---

🚀 **Por que o Discord consegue servir milhões de usuários com Go?**

Antes de usar Go, eles tinham um problema clássico: cada conexão aberta era uma thread. Thread custa memória. Muita memória.

Go resolveu isso com goroutines — e a diferença é absurda:

🧵 **Thread tradicional:** ~1MB de memória só para existir. 10.000 conexões = 10GB apenas em threads paradas.

⚡ **Goroutine Go:** ~2KB. O mesmo servidor, 500x mais leve. E a pilha cresce automaticamente quando precisa.

🧠 **O segredo:** o Go Runtime Scheduler. Ele distribui milhares de goroutines entre poucos núcleos de CPU — invisível para você. Quando uma goroutine fica esperando I/O, o scheduler imediatamente coloca outra no lugar. Nenhum núcleo desperdiçado.

🎯 **A sintaxe:** uma palavra. `go minhaFuncao()`. Sem pool de threads para configurar, sem executor service, sem promises. Só `go`.

⚠️ **O gotcha clássico:** goroutine anônima dentro de loop captura variável por referência. Quando a goroutine finalmente roda, o loop já terminou. Resultado: todos os valores são iguais ao último. Correção: passe como parâmetro.

No Capítulo 15 do meu livro de estudos Go, demonstrei 10 goroutines completando 20 segundos de trabalho em 2 segundos reais. E o código para fazer isso tem 3 linhas.

Concorrência que escala, sem a complexidade que te faz querer largar tudo.

---
→ Repositório completo nos comentários.

#Go #Golang #Concorrência #Backend #SoftwareEngineering #Goroutines #Programming

---

**Checklist antes de publicar:**
- [x] Conteúdo original — sem copiar do fonte.txt
- [x] 5 bullets de valor técnico
- [ ] Revisar ortografia
- [ ] Adicionar link do repositório nos comentários
- [ ] Publicar no perfil correto
