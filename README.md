# 🦫 Go na Prática: Jornada de Aprendizado

> **"Menos é exponencialmente mais."** — Filosofia de design do Go.

Bem-vindo ao repositório de estudos baseado no livro/guia **Go na Prática: 30 capítulos para dominar a linguagem**. Este repositório foi estruturado para servir de diário de bordo prático, organizando anotações teóricas, códigos de exemplo e resoluções de desafios propostos.

---

## 🛠️ Como o Repositório Está Organizado

Cada capítulo possui sua própria pasta dentro do diretório `estudos-go/` e segue uma estrutura padrão e modular para evitar conflitos de compilação:

*   **`README.md`**: Anotações teóricas resumidas e explicação dos exercícios práticos.
*   **`exemplos/`**: Códigos demonstrados ao longo do capítulo (ex: `exemplos/ola-mundo/main.go`).
*   **`exercicios/`**: Subpastas isoladas para cada exercício, permitindo rodar e compilar as diferentes resoluções sem conflitos de pacotes.

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

**Progresso Atual:** `[███░░░░░░░░░░░░░░░░░░░]` **13.3%** (4 de 30 capítulos concluídos)

| Capítulo | Título | Status | Recursos Disponíveis |
| :---: | :--- | :---: | :---: |
| **01** | **Introdução & Origens do Go** | 🟢 *Concluído* | [Notas de Estudo](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo01/README.md) \| [Exemplos](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo01/exemplos) \| [Exercícios](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo01/exercicios) |
| **02** | **Variáveis, Constantes e Tipos de Dados** | 🟢 *Concluído* | [Notas de Estudo](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo02/README.md) \| [Exemplos](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo02/exemplos) \| [Exercícios](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo02/exercicios) |
| **03** | **Tipos Básicos — A Química dos Materiais** | 🟢 *Concluído* | [Notas de Estudo](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo03/README.md) \| [Exemplos](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo03/exemplos) \| [Exercícios](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo03/exercicios) |
| **04** | **Controle de Fluxo — O Guarda de Trânsito** | 🟢 *Concluído* | [Notas de Estudo](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo04/README.md) \| [Exemplos](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo04/exemplos) \| [Exercícios](file:///c:/Users/jorge/OneDrive/Documentos/PROJECTS/go-in-practice-book/estudos-go/capitulo04/exercicios) |
| **05** | Capítulo 05 | 🔴 *Não Iniciado* | *A ser criado* |
| **06** | Capítulo 06 | 🔴 *Não Iniciado* | *A ser criado* |
| **07** | Capítulo 07 | 🔴 *Não Iniciado* | *A ser criado* |
| **08** | Capítulo 08 | 🔴 *Não Iniciado* | *A ser criado* |
| **09** | Capítulo 09 | 🔴 *Não Iniciado* | *A ser criado* |
| **10** | Capítulo 10 | 🔴 *Não Iniciado* | *A ser criado* |
| **11** | Capítulo 11 | 🔴 *Não Iniciado* | *A ser criado* |
| **12** | Capítulo 12 | 🔴 *Não Iniciado* | *A ser criado* |
| **13** | Capítulo 13 | 🔴 *Não Iniciado* | *A ser criado* |
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
