# 🏋️ Pilha de Exercícios — Capítulo 10: Structs

Exercícios extras em dificuldade crescente para praticar Structs em Go. **Sem solução no arquivo** — o objetivo é resolver de verdade.

---

## 🟢 Fáceis

### Ex01 — Ficha de Livro
Crie uma struct `Livro` com os campos `Titulo` (string), `Autor` (string) e `Paginas` (int). Instancie 3 livros diferentes e imprima cada um usando `%+v`.

### Ex02 — Atualizar Campo
Crie uma struct `Conta` com `Titular` (string) e `Saldo` (float64). Instancie uma conta com saldo `1000.00`. Crie uma função `depositar(c *Conta, valor float64)` que soma o valor ao saldo. Chame-a duas vezes e imprima o saldo final.

### Ex03 — Zero Value Consciente
Declare `var p Produto` (usando a struct do livro, sem inicializar). Imprima cada campo e confirme os zero values: `""`, `0`, `false`. Depois preencha os campos um a um usando o ponto e imprima novamente.

### Ex04 — Struct como Retorno
Crie uma função `novoProduto(nome string, preco float64) Produto` que recebe os dados e retorna uma struct `Produto` já preenchida. Use-a na `main` para criar 2 produtos.

---

## 🟡 Intermediários

### Ex05 — Catálogo de Produtos
Crie um slice `[]Produto` com pelo menos 5 produtos. Percorra o slice com `range` e imprima apenas os produtos onde `Ativo == true` e `EmEstoque > 0`.

### Ex06 — Struct Aninhada: Endereço de Entrega
Crie as structs `Endereco` (Rua, Cidade, CEP) e `Cliente` (Nome, Email, Endereco). Instancie um cliente completo e imprima o nome e a cidade de entrega no formato: `"Entregar para Ana em São Paulo"`.

### Ex07 — Maior e Menor
Dado um slice de structs `Temperatura` com campos `Cidade` (string) e `Graus` (float64), percorra e encontre a cidade mais quente e a mais fria. Imprima ambas.

### Ex08 — Atualizador via Ponteiro
Crie uma função `aplicarDesconto(p *Produto, percentual float64)` que reduz o preço do produto original pelo percentual informado. Teste com `desconto de 20%` num produto de `R$ 100.00` e confirme que `p.Preco` virou `80.00`.

---

## 🔴 Difíceis

### Ex09 — Inventário com Map
Crie um `map[string]Produto` onde a chave é o código do produto (ex: `"SKU-001"`). Adicione 3 produtos. Crie uma função `buscarProduto(inv map[string]Produto, codigo string)` que usa o comma ok idiom para retornar o produto e um bool indicando se foi encontrado.

### Ex10 — Carrinho de Compras
Usando `Produto` e uma struct `Carrinho` com `Itens []Produto`, implemente:
- `adicionarItem(c *Carrinho, p Produto)` — adiciona ao slice
- `removerItem(c *Carrinho, nome string)` — remove o primeiro item com aquele nome
- `calcularTotal(c Carrinho) float64` — soma todos os preços

Teste com pelo menos 3 itens, remova um e imprima o total.

### Ex11 — Embedding na Prática
Crie as structs `Veiculo` (Marca string, Velocidade int) e `Carro` (com Embedding de `Veiculo` e campo extra `Portas int`). Crie uma função `acelerar(v *Veiculo, delta int)` que aumenta a velocidade. Chame-a passando `&carro.Veiculo` e verifique que `carro.Velocidade` (promoted field) foi atualizado.

### Ex12 — Ranking de Jogadores
Crie uma struct `Jogador` com `Nome` (string) e `Pontuacao` (int). Crie um slice com 5 jogadores. Implemente um bubble sort manual (sem `sort.Slice`) que ordene o slice por `Pontuacao` decrescente. Imprima o ranking numerado do 1º ao 5º lugar.
