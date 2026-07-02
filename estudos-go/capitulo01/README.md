# 📖 Capítulo 01: Introdução ao Go

> **Livro Go na Prática: 30 capítulos para dominar a linguagem**

> 🧮 [Teste de mesa de todo o código deste capítulo](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo01/teste-de-mesa.md)

---

## 🍽️ A Cozinha Caótica e a Máquina Perfeita
Imagine que você foi contratado para gerenciar a cozinha de um restaurante gigantesco que serve milhões de pratos por segundo. 
*   **O Caos (Outras Linguagens):** 
    *   Alguns chefs escrevem receitas que só eles entendem (como em **C++**). Quando um chef sai de férias, ninguém consegue dar manutenção.
    *   Outros precisam ler a receita em voz alta, linha por linha, enquanto cozinham (como em **Python** ou **Ruby**). É fácil de iniciar, mas no horário de pico, ler e cozinhar ao mesmo tempo gera gargalos de velocidade.
    *   Ingredientes mudam de lugar, fogões explodem se você usar a panela errada e gasta-se mais tempo limpando do que cozinhando.
*   **A Solução (Go):** O Go é como demolir essa cozinha antiga e instalar uma linha de montagem industrial futurista. As receitas são universais (todos escrevem do mesmo jeito) e um chef novo se adapta no primeiro dia.

---

## 🤖 Elementos do Ecossistema Go

### 🛡️ O Compilador
Antes de ligar o fogo, um robô de segurança (o compilador) verifica cada passo da receita. Se você colocar sal no açúcar, ele impede a execução e te avisa. Nada explode em produção. Depois de aprovado, o código roda de forma extremamente rápida em máquinas de precisão.

### 🏭 O Nascimento no Google e a Espera de 45 Minutos
Em 2007, os sistemas internos do Google escritos em C++ demoravam até **45 minutos** para compilar. Diante desse problema de produtividade, três engenheiros lendários projetaram o Go:
1.  **Ken Thompson:** Criador do Unix e da linguagem B (antecessora do C).
2.  **Rob Pike:** Co-criador do UTF-8 e gigante do Unix.
3.  **Robert Griesemer:** Um dos criadores do motor V8 de JavaScript do Google Chrome.

### 🧩 A Filosofia: *Less is Exponentially More*
Muitas linguagens evoluem adicionando recursos até se tornarem pesadas e complexas. Go tomou o caminho oposto:
*   **Apenas 25 palavras-chave (keywords):** Em comparação com as mais de 100 do C#.
*   **Curva de aprendizado rápida:** É possível aprender a especificação da linguagem em poucas horas.
*   **Código óbvio:** Não há "mágica" por trás do código. O que você lê é o que ele faz.

---

## 📀 Compilado vs. Interpretado

> [!NOTE]
> **A Analogia Definitiva:**
> *   **Interpretador (Python, Ruby):** É um músico lendo a partitura e tocando ao vivo. É flexível (dá para mudar no meio), mas tem limite de velocidade. Se houver um erro no final da música, você só descobrirá quando chegar lá no meio do show (erro em tempo de execução).
> *   **Compilador (Go, C, Rust):** É o estúdio de gravação. O código-fonte é a partitura. O compilador grava, mixa, masteriza e gera um CD/MP3 (o executável binário). Para tocar a música (rodar o programa), você só precisa do arquivo executável final, sem necessidade de carregar o compilador ou o código-fonte junto.

---

## 🐹 O Gopher
Desenhado por Renée French (esposa de Rob Pike), o Gopher é uma marmota azul que representa um operário pragmático. Ele constrói coisas e foca no simples. A comunidade do Go se intitula **Gophers**, valorizando o código limpo, óbvio e que funciona.

---

## 🛠️ Configuração do Ambiente

1.  **Instalação do SDK:** Baixe no site oficial [go.dev/dl](https://go.dev/dl).
    *   **Windows:** Execute o `.msi` e siga o instalador padrão.
    *   **Linux/Mac:** Extraia e adicione a pasta `bin` ao seu `PATH`.
    *   **Verificação:** Digite `go version` no terminal.
2.  **VS Code:** Instale a extensão oficial do **Go (Go Team at Google)**. Ao abrir um arquivo `.go` pela primeira vez, clique em **Install All** na notificação de *"Analysis tools missing"* para instalar utilitários cruciais como o `gopls`.

> [!TIP]
> **Obsessão por Formatação (`go fmt`):**
> Esqueça discussões sobre espaços vs. tabs ou onde colocar as chaves. O comando `go fmt` (ou o auto-save do VS Code) formata seu código automaticamente no padrão oficial de Go (uso de tabs e chaves na mesma linha). Isso zera a fadiga cognitiva ao ler códigos de terceiros.

---

## 🔬 Anatomia do "Olá Mundo"

Análise do arquivo [exemplos/ola-mundo/main.go](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo01/exemplos/ola-mundo/main.go):

```go
package main

import "fmt"

func main() {
    fmt.Println("Olá mundo, eu sou gopher")
}
```

*   **`package main`:** Declara que este arquivo pertence ao pacote principal (`main`). Indica que este código gerará um programa executável e não apenas uma biblioteca.
*   **`import "fmt"`:** Importa o pacote da biblioteca padrão responsável pela formatação de texto e escrita/leitura no console.
*   **`func main()`:** É o ponto de entrada (a chave de ignição) do programa.
*   **`fmt.Println(...)`:** Executa a função `Println` contida no pacote `fmt`.
    > [!IMPORTANT]
    > **A Regra de Visibilidade:**
    > O "P" maiúsculo em `Println` significa que a função é **Pública** (exportada para uso externo). Se fosse minúsculo (`println`), seria privada e inacessível de fora do pacote `fmt`.

---

## 🚀 Como Executar

Abra o terminal na pasta do exemplo e escolha um dos métodos:

### Método 1: Execução Rápida
Útil para testes rápidos durante o desenvolvimento:
```bash
go run exemplos/ola-mundo/main.go
```

### Método 2: Compilação Estática
Gera um executável independente que pode ser distribuído sem precisar do Go instalado:
```bash
# Compilar
go build -o ola-mundo exemplos/ola-mundo/main.go

# Rodar o binário gerado
./ola-mundo      # Linux/macOS
.\ola-mundo.exe  # Windows
```

---

## 📝 Exercícios de Fixação

Abaixo estão os exercícios propostos no final do capítulo. Resolva-os dentro da pasta `exercicios/` para praticar a memória muscular.

- [x] **Exercício 1: O identificador pessoal (Fácil)**
  Modifique o programa para exibir três linhas separadas: seu nome, sua cidade e sua comida favorita.
  *Resolução em:* [exercicios/ex01_identificador/main.go](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo01/exercicios/ex01_identificador/main.go)
  
- [x] **Exercício 2: O artista ASCII médio**
  Use caracteres como asteriscos (`*`) para desenhar um quadrado ou triângulo na tela.
  *Resolução em:* [exercicios/ex02_artista/main.go](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo01/exercicios/ex02_artista/main.go)
  
- [ ] **Exercício 3: O quebrador de código (Debug)**
  Provoque erros intencionais no código para aprender a ler os erros do compilador:
  1. Altere `package main` para `package batata`.
  2. Remova a linha `import "fmt"`.
  3. Altere `fmt.Println` para `fmt.println` (letra minúscula).
  *Resolução em:* [exercicios/ex03_quebrador/main.go](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo01/exercicios/ex03_quebrador/main.go)
