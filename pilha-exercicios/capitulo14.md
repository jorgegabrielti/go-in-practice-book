# 🏋️ Pilha de Exercícios — Capítulo 14: Pacotes e Módulos

> Exercícios extras sem solução. Resolva por conta própria criando módulos reais no terminal.

---

## 🟢 Fácil

### 1. A Calculadora Modular

Crie um módulo `go mod init github.com/seu-usuario/calculadora`. Dentro dele, crie um pacote `operacoes/` com funções exportadas `Somar`, `Subtrair`, `Multiplicar` e `Dividir` (retorna float64, error). Na `main.go`, importe o pacote e teste as quatro operações.

### 2. A Saudação Personalizada

Crie um pacote `saudacao/` com uma função exportada `Ola(nome string) string` que retorna `"Olá, {nome}! Bem-vindo ao Go."`. Crie também uma função privada `formatarNome(s string) string` que coloca a primeira letra em maiúscula (use `strings.Title` ou implemente manualmente). A função `Ola` usa `formatarNome` internamente. Na main, verifique que só `Ola` é acessível.

### 3. O Inventário de Tipos

Crie um pacote `formas/` com: tipo exportado `Retangulo struct { Largura, Altura float64 }`, método exportado `Area() float64`, e campo privado interno (use uma função exportada como getter). Importe na main e use.

### 4. O Módulo com Alias

Crie um módulo que importa dois pacotes próprios com nomes parecidos, como `matematica/basica` e `matematica/avancada`. Na `main.go`, use import com alias para diferenciar: `import basica "meu-modulo/matematica/basica"` e `import avancada "meu-modulo/matematica/avancada"`. Chame uma função de cada.

---

## 🟡 Médio

### 5. O Pacote de Validação

Crie um pacote `validacao/` com as funções exportadas `CPF(s string) bool`, `Email(s string) bool` e `Senha(s string) error`. Use `strings`, `regexp` ou lógica manual. Crie um sentinel error `ErrSenhaFraca`. Na main, valide alguns valores de exemplo.

### 6. O Logger Semântico

Em vez de criar `utils/`, crie um pacote `logger/` com funções `Info(msg string)`, `Aviso(msg string)` e `Erro(msg string)` — cada uma imprime com um prefixo diferente (`[INFO]`, `[AVISO]`, `[ERRO]`). Use uma variável privada no pacote para controlar o nível mínimo de log (`nivel`). Exponha uma função `DefinirNivel(n int)` para configurá-la.

### 7. O Módulo com go get

Instale a biblioteca `github.com/google/uuid` via `go get`. Crie um pacote próprio `identificador/` com uma função exportada `Novo() string` que usa `uuid.New().String()` internamente. Na main, importe `identificador` (não `uuid` diretamente) e gere 3 IDs. Verifique que o `go.mod` registrou a dependência do uuid.

### 8. A Faxina do go mod tidy

Crie um módulo e importe intencionalmente dois pacotes externos (um que você vai usar de verdade e um que você vai apagar do código depois). Rode `go mod tidy` e observe o que ele remove do `go.mod` e `go.sum`. Documente o antes e depois dos arquivos.

---

## 🔴 Desafio

### 9. A Biblioteca Interna Multi-Pacote

Crie um módulo com a seguinte estrutura:
```
meu-sistema/
├── go.mod
├── main.go
├── banco/
│   ├── conta.go       (tipo Conta, funções Depositar, Sacar)
│   └── transacao.go   (tipo Transacao, histórico)
└── relatorio/
    └── gerador.go     (função Gerar(contas []banco.Conta) string)
```
O pacote `relatorio` importa `banco`. A `main` importa ambos. Implemente com visibilidade correta — campos internos de `Conta` (saldo) devem ser privados, acessíveis só por métodos exportados.

### 10. O Pacote com `init()`

Crie um pacote `config/` com uma função especial `init()` (sem parâmetros, sem retorno — Go chama automaticamente ao importar o pacote). Dentro do `init`, leia uma variável de ambiente `APP_ENV` com `os.Getenv` e defina uma variável privada `ambiente string`. Exponha uma função `Ambiente() string` que retorna o valor. Na main, importe o pacote e imprima `config.Ambiente()` — sem chamar nenhum setup manual.

### 11. O Versionamento de Módulo

Explore o sistema de versões do Go Modules: crie um módulo, faça `go get github.com/spf13/cobra@v1.8.0` (versão específica), e depois `go get github.com/spf13/cobra@latest`. Inspecione o `go.mod` antes e depois de cada comando. Crie um programa simples usando `cobra` para ter um CLI com dois subcomandos.

### 12. O Módulo Replace

Crie dois módulos: `biblioteca/` e `app/`. A `app` importa `biblioteca`. Em vez de publicar a `biblioteca` no GitHub, use a diretiva `replace` no `go.mod` da `app` para apontar para o caminho local:
```
replace github.com/seu-usuario/biblioteca => ../biblioteca
```
Implemente uma função na biblioteca, use na app, e confirme que o `replace` funciona sem precisar publicar nada.
