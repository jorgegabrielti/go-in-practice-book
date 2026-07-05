# 🏋️ Pilha de Exercícios — Capítulo 11: Métodos

Exercícios extras em dificuldade crescente para praticar Métodos em Go. **Sem solução no arquivo.**

---

## 🟢 Fáceis

### Ex01 — Termômetro
Crie uma struct `Termometro` com campo `Celsius float64`. Adicione dois métodos: `ParaFahrenheit() float64` (retorna `C*9/5 + 32`) e `ParaKelvin() float64` (retorna `C + 273.15`). Ambos Value Receivers. Teste com `36.6°C`.

### Ex02 — Contador com Reset
Crie uma struct `Contador` com campo `Valor int`. Métodos com Pointer Receiver: `Incrementar()`, `Decrementar()` (mínimo 0) e `Resetar()`. Faça uma sequência de chamadas e imprima o valor após cada uma.

### Ex03 — Ponto no Plano
Crie uma struct `Ponto` com `X` e `Y` (float64). Método `DistanciaOrigem() float64` (Value Receiver) que retorna `√(X²+Y²)` usando `math.Sqrt`. Método `Mover(dx, dy float64)` (Pointer Receiver) que desloca o ponto. Imprima a distância antes e depois de mover.

### Ex04 — Tipo Temperatura
Crie `type Celsius float64`. Adicione um método `String() string` que retorna `"X.XX°C"`. Faça o mesmo para `type Fahrenheit float64`. Instancie um de cada e imprima com `fmt.Println` (que chama `String()` automaticamente se definido).

---

## 🟡 Intermediários

### Ex05 — Pilha (Stack) com Métodos
Crie uma struct `Pilha` com campo `itens []int`. Métodos com Pointer Receiver: `Empurrar(v int)`, `Retirar() (int, bool)` (retorna o topo e `false` se vazia) e `Tamanho() int`. Teste empurrando 5 valores e retirando 3.

### Ex06 — Conta com Histórico
Expanda a `ContaBancaria` do livro adicionando um campo `Historico []string`. Cada operação (`Depositar`, `Sacar`) deve adicionar uma entrada descritiva ao histórico. Crie um método `ImprimirHistorico()` que percorre e imprime cada entrada.

### Ex07 — Fila (Queue) com Métodos
Crie uma struct `Fila` com campo `itens []string`. Métodos: `Enfileirar(s string)` (adiciona no fim), `Desenfileirar() (string, bool)` (remove do início) e `Vazia() bool`. Simule uma fila de atendimento com pelo menos 4 clientes.

### Ex08 — Tipo Dinheiro Completo
Expanda o tipo `Dinheiro float64` do livro. Adicione os métodos: `Somar(outro Dinheiro) Dinheiro`, `Subtrair(outro Dinheiro) Dinheiro`, `Aplicar(percentual float64) Dinheiro` (ex: 10% de juros). Encadeie as operações e imprima o resultado final.

---

## 🔴 Difíceis

### Ex09 — Matriz 2x2
Crie uma struct `Matriz` com campo `dados [2][2]float64`. Métodos: `Set(linha, col int, val float64)`, `Get(linha, col int) float64`, `Transposta() Matriz` (retorna nova matriz com linhas e colunas trocadas) e `String() string` que imprime a matriz formatada. Teste criando uma matriz, alterando um valor, transpondo e imprimindo.

### Ex10 — Sistema de RPG
Crie uma struct `Personagem` com `Nome`, `Vida`, `Ataque` e `Defesa` (todos int). Métodos:
- `Atacar(alvo *Personagem)` — causa dano de `c.Ataque - alvo.Defesa` (mínimo 1) em `alvo.Vida`
- `EstaVivo() bool` — retorna `Vida > 0`
- `String() string` — retorna `"Nome (Vida HP)"`

Simule um combate entre dois personagens em loop até um deles morrer.

### Ex11 — Pipeline de Transformação
Crie um tipo `Pipeline` que encapsula um `[]float64`. Métodos com Pointer Receiver que retornam `*Pipeline` (para encadeamento): `Filtrar(fn func(float64) bool)`, `Mapear(fn func(float64) float64)` e `Resultado() []float64`. Use assim: `p.Filtrar(positivos).Mapear(dobrar).Resultado()`.

### Ex12 — Máquina de Estados
Crie um tipo `Estado string` e uma struct `Semaforo` com campo `estado Estado`. Estados possíveis: `"verde"`, `"amarelo"`, `"vermelho"`. Método `Avancar()` que transita entre estados na sequência correta (verde→amarelo→vermelho→verde). Imprima o estado após cada avanço, simulando 7 ciclos.
