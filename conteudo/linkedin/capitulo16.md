> **Antes de publicar:** conferir checklist no fim do arquivo.

---

🔗 **Por que o Go não usa mutex por padrão para concorrência?**

Porque existe algo melhor: channels.

O mantra oficial da linguagem é: *"Não compartilhe memória para se comunicar. Comunique-se para compartilhar memória."*

Na prática, isso significa:

📦 **Channel = tubo tipado entre goroutines.** `chan int` só aceita inteiros. `chan string` só aceita strings. O compilador rejeita o tipo errado — erro em compilação, não em produção.

🔒 **Bloqueio automático como sincronização.** Canais unbuffered (padrão) só transferem dados quando as duas goroutines estão prontas ao mesmo tempo. Sem `time.Sleep`, sem mutex, sem race condition. A troca é a sincronização.

💀 **Deadlock detectado em runtime.** Se todas as goroutines ficam esperando e nenhuma avança, o Go Runtime percebe e encerra com `fatal error: all goroutines are asleep - deadlock!` — em vez de travar silenciosamente para sempre.

🔄 **`close` + `for range` = pipeline elegante.** O produtor fecha o canal quando termina. O consumidor usa `for range` — o loop sai sozinho. Sem contador, sem sinal manual de "acabou".

🎯 **Canais direcionais como contrato.** `chan<- int` (send-only) e `<-chan int` (receive-only) na assinatura da função. O compilador garante que ninguém faz o que não deveria — segurança de tipos em tempo de compilação.

No Capítulo 16 do meu livro de estudos Go, implementei Ping Pong com sincronização perfeita, pipeline com `close`/`for range`, e temporizador manual usando apenas channels — sem nenhuma primitiva de sincronização extra.

---
→ Repositório nos comentários.

#Go #Golang #Channels #Concorrência #Backend #SoftwareEngineering #Programming

---

**Checklist antes de publicar:**
- [x] Conteúdo original — sem copiar do fonte.txt
- [x] 5 bullets de valor técnico
- [ ] Revisar ortografia
- [ ] Adicionar link do repositório nos comentários
- [ ] Publicar no perfil correto
