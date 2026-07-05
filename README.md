# 🦫 Go na Prática: Jornada de Aprendizado

> **"Menos é exponencialmente mais."** — Filosofia de design do Go.

Bem-vindo ao repositório de estudos baseado no livro/guia **Go na Prática: 30 capítulos para dominar a linguagem**. Este repositório foi estruturado para servir de diário de bordo prático, organizando anotações teóricas, códigos de exemplo e resoluções de desafios propostos.

---

## 🛠️ Como o Repositório Está Organizado

Cada capítulo possui sua própria pasta dentro do diretório `estudos-go/` e segue uma estrutura padrão e modular para evitar conflitos de compilação:

*   **`fonte.txt`**: transcrição bruta e fiel do conteúdo do livro para o capítulo — a fonte primária de tudo que vem depois (notas, exemplos e conteúdo de divulgação).
*   **`README.md`**: Anotações teóricas resumidas e explicação dos exercícios práticos, escritas a partir do `fonte.txt`.
*   **`exemplos/`**: Códigos demonstrados ao longo do capítulo (ex: `exemplos/ola-mundo/main.go`).
*   **`exercicios/`**: Subpastas isoladas para cada exercício, permitindo rodar e compilar as diferentes resoluções sem conflitos de pacotes.
*   **`teste-de-mesa.md`**: rastreamento manual (linha a linha, sem rodar o código) da execução de todo `main.go` do capítulo — prevê o estado das variáveis e a saída no terminal antes de confirmar com `go run`. Útil para expor bugs sutis (aliasing de slices, overflow silencioso, uso incorreto de `fmt`).

Use `estudos-go/_template/` como ponto de partida ao iniciar um novo capítulo, e consulte `estudos-go/GLOSSARIO.md` para revisar os termos-chave já estudados.

---

## ✍️ Produção de Conteúdo (LinkedIn & Medium)

A pasta `conteudo/` centraliza tudo relacionado à divulgação dos estudos nas redes:

*   **`conteudo/medium/`**: artigos completos, um por capítulo (`capituloXX.md`).
*   **`conteudo/linkedin/`**: posts curtos de divulgação, um por capítulo (`capituloXX.md`).
*   **`_template.md`**: modelo padrão em cada subpasta, com checklist de publicação.
*   **`conteudo/PAINEL.md`**: painel de status (rascunho/revisado/publicado) e links das publicações.

> **Regra de ouro:** todo artigo de Medium e post de LinkedIn deve ser baseado no `fonte.txt` do capítulo correspondente, não em memória ou suposição sobre o conteúdo do livro.

---

## 🏋️ Pilha de Exercícios (Prática Deliberada)

A pasta `pilha-exercicios/` traz, para cada capítulo concluído, no mínimo 10 exercícios extras (inspirados no formato de [exercicios.dunossauro.com](https://exercicios.dunossauro.com/)), em dificuldade crescente e **sem solução no arquivo** — o objetivo é praticar de verdade, não copiar resposta. Veja `pilha-exercicios/README.md` para as regras completas.

---

## 🚀 Como Executar os Códigos

Certifique-se de ter o Go SDK instalado (versão 1.20+ recomendada) e siga as diretrizes abaixo no terminal.

### 1. Rodar diretamente (Modo Desenvolvimento)
Para testar o código rapidamente sem gerar um executável permanente:
```bash
go run estudos-go/capitulo01/exemplos/ola-mundo/main.go
```

### 2. Compilar para Produção (Gerar Binário Estático)
Para gerar o executável na pasta atual:
```bash
# Navegar até o diretório ou rodar a compilação apontando o arquivo
go build -o bin/ola-mundo estudos-go/capitulo01/exemplos/ola-mundo/main.go

# Executar o binário gerado
./bin/ola-mundo
```

### 3. Formatar o Código
Lembre-se sempre de rodar o formatador automático antes de comitar:
```bash
go fmt ./...
```

---

## 📊 Painel de Progresso (30 Capítulos)

**Progresso Atual:** `[████████████░░░░░░░░░░]` **40%** (12 de 30 capítulos concluídos)

| Capítulo | Título | Status | Recursos Disponíveis |
| :---: | :--- | :---: | :---: |
| **01** | **Introdução & Origens do Go** | 🟢 *Concluído* | [Notas de Estudo](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo01/README.md) \| [Exemplos](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo01/exemplos) \| [Exercícios](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo01/exercicios) |
| **02** | **Variáveis, Constantes e Tipos de Dados** | 🟢 *Concluído* | [Notas de Estudo](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo02/README.md) \| [Exemplos](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo02/exemplos) \| [Exercícios](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo02/exercicios) |
| **03** | **Tipos Básicos — A Química dos Materiais** | 🟢 *Concluído* | [Notas de Estudo](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo03/README.md) \| [Exemplos](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo03/exemplos) \| [Exercícios](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo03/exercicios) |
| **04** | **Controle de Fluxo — O Guarda de Trânsito** | 🟢 *Concluído* | [Notas de Estudo](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo04/README.md) \| [Exemplos](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo04/exemplos) \| [Exercícios](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo04/exercicios) |
| **05** | **Laços de Repetição (For) — A Pista de Corrida** | 🟢 *Concluído* | [Notas de Estudo](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo05/README.md) \| [Exemplos](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo05/exemplos) \| [Exercícios](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo05/exercicios) |
| **06** | **Funções — As Mini-Fábricas Especializadas** | 🟢 *Concluído* | [Notas de Estudo](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo06/README.md) \| [Exemplos](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo06/exemplos) \| [Exercícios](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo06/exercicios) |
| **07** | **Arrays e Slices — A Lista Fixa e a Lista Mágica** | 🟢 *Concluído* | [Notas de Estudo](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo07/README.md) \| [Exemplos](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo07/exemplos) \| [Exercícios](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo07/exercicios) |
| **08** | **Maps — O Guarda-Volumes Inteligente** | 🟢 *Concluído* | [Notas de Estudo](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo08/README.md) \| [Exemplos](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo08/exemplos) \| [Exercícios](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo08/exercicios) |
| **09** | **Ponteiros — O Endereço da Casa** | 🟢 *Concluído* | [Notas de Estudo](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo09/README.md) \| [Exemplos](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo09/exemplos) \| [Exercícios](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo09/exercicios) |
| **10** | **Structs — A Ficha do Paciente** | 🟢 *Concluído* | [Notas de Estudo](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo10/README.md) \| [Exemplos](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo10/exemplos) \| [Exercícios](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo10/exercicios) |
| **11** | **Métodos — Dando Vida aos Dados** | 🟢 *Concluído* | [Notas de Estudo](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo11/README.md) \| [Exemplos](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo11/exemplos) \| [Exercícios](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo11/exercicios) |
| **12** | **Interfaces — O Plugue Universal** | 🟢 *Concluído* | [Notas de Estudo](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo12/README.md) \| [Exemplos](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo12/exemplos) \| [Exercícios](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo12/exercicios) |
| **13** | Capítulo 13 | 🟡 *Em Andamento* | [Notas de Estudo](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo13/README.md) |
| **14** | Capítulo 14 | 🔴 *Não Iniciado* | *A ser criado* |
| **15** | Capítulo 15 | 🔴 *Não Iniciado* | *A ser criado* |
| **16** | Capítulo 16 | 🔴 *Não Iniciado* | *A ser criado* |
| **17** | Capítulo 17 | 🔴 *Não Iniciado* | *A ser criado* |
| **18** | Capítulo 18 | 🔴 *Não Iniciado* | *A ser criado* |
| **19** | Capítulo 19 | 🔴 *Não Iniciado* | *A ser criado* |
| **20** | Capítulo 20 | 🔴 *Não Iniciado* | *A ser criado* |
| **21** | Capítulo 21 | 🔴 *Não Iniciado* | *A ser criado* |
| **22** | Capítulo 22 | 🔴 *Não Iniciado* | *A ser criado* |
| **23** | Capítulo 23 | 🔴 *Não Iniciado* | *A ser criado* |
| **24** | Capítulo 24 | 🔴 *Não Iniciado* | *A ser criado* |
| **25** | Capítulo 25 | 🔴 *Não Iniciado* | *A ser criado* |
| **26** | Capítulo 26 | 🔴 *Não Iniciado* | *A ser criado* |
| **27** | Capítulo 27 | 🔴 *Não Iniciado* | *A ser criado* |
| **28** | Capítulo 28 | 🔴 *Não Iniciado* | *A ser criado* |
| **29** | Capítulo 29 | 🔴 *Não Iniciado* | *A ser criado* |
| **30** | Capítulo 30 | 🔴 *Não Iniciado* | *A ser criado* |

---

*Painel atualizado automaticamente conforme progresso nos estudos. Bons estudos!*
