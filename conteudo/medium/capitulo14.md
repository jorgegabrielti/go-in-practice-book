> **Antes de publicar:** conferir checklist no fim do arquivo.

# 📦 Como o Go organiza o caos: pacotes, módulos e a regra da letra maiúscula

Todo projeto começa arrumado. Algumas semanas depois, você tem um `main.go` com 800 linhas, funções de banco de dados misturadas com lógica de apresentação, e uma vontade crescente de fingir que o arquivo não existe.

Go tem uma resposta para isso: pacotes e módulos. E a solução é mais simples do que parece.

---

## A gaveta certa para cada ferramenta

Pense no seu projeto como uma oficina. Ferramentas de elétrica ficam numa gaveta. Ferramentas de mecânica, em outra. Você não mistura. Quando precisa de um alicate, sabe exatamente onde procurar.

Em Go, cada gaveta é um **pacote** — uma pasta com arquivos `.go` que compartilham um propósito. E o conjunto de todas as gavetas é o **módulo** — o projeto inteiro, com identidade própria.

O módulo nasce de um único comando no terminal:

```bash
go mod init github.com/seu-usuario/nome-do-projeto
```

Isso cria o `go.mod`, o "registro civil" do seu projeto. Ele diz quem você é e quais dependências externas você usa. Sem ele, o Go não sabe o que está compilando.

---

## A regra que substitui public e private

Java tem `public`, `private`, `protected`. Python tem convenção de underscore. Go tem algo mais elegante: a letra inicial do nome.

**Começa com maiúscula → exportado (público)**. Qualquer pacote que importar o seu pode usar.
**Começa com minúscula → não-exportado (privado)**. Só o próprio pacote enxerga.

```go
package financeira

func CalcularImposto() float64 { ... }  // ✅ acessível de fora: financeira.CalcularImposto()
func ajustarBase() float64      { ... }  // ❌ invisível para outros pacotes
```

Essa regra se aplica a tudo: funções, tipos, variáveis, constantes, campos de struct. Uma letra maiúscula é a declaração de intenção — "isso é parte da API pública deste pacote".

O compilador reforça a regra em tempo de compilação. Tentar acessar um identificador minúsculo de fora do pacote resulta em erro imediato — sem surpresas em produção.

---

## Estruturando um projeto real

Suponha que você está construindo uma calculadora de conversão de unidades. Em vez de tudo em `main.go`, você separa:

```
meu-projeto/
├── go.mod
├── main.go
└── conversor/
    └── medidas.go
```

O arquivo `conversor/medidas.go` declara `package conversor` e expõe apenas o que deve ser público:

```go
package conversor

// KmParaMilhas é pública — começa com K maiúsculo
func KmParaMilhas(km float64) float64 {
    return km * 0.621371
}

// milhasParaKm é privada — detalhe de implementação
func milhasParaKm(milhas float64) float64 {
    return milhas / 0.621371
}
```

Na `main.go`, o import usa o caminho completo: nome do módulo + nome da pasta:

```go
import "github.com/seu-usuario/meu-projeto/conversor"

func main() {
    fmt.Println(conversor.KmParaMilhas(100)) // 62.14
    // conversor.milhasParaKm(62) → erro de compilação
}
```

A fronteira entre público e privado é clara, verificada pelo compilador, e não depende de nenhuma convenção informal.

---

## O ecossistema externo: go get

Go tem um repositório imenso de pacotes públicos. Você não precisa implementar tudo do zero — basta saber onde buscar.

```bash
go get github.com/fatih/color
```

Esse único comando faz três coisas: baixa o código, registra a dependência no `go.mod` e cria um `go.sum` com o checksum criptográfico da versão baixada. O `go.sum` é a garantia de que o código que você tem é exatamente o que o autor publicou — sem alterações.

Depois é só importar e usar:

```go
import "github.com/fatih/color"

color.Red("algo deu errado")
color.Green("tudo certo")
```

---

## go mod tidy: a faxina automática

Com o tempo, o `go.mod` acumula entulho — dependências que você importou e depois removeu do código, ou pacotes que você usa mas esqueceu de adicionar. O comando:

```bash
go mod tidy
```

Lê todo o código, descobre o que está sendo usado de fato, baixa o que falta e remove o que sobra. É o equivalente a reorganizar a oficina antes de fechar o expediente. Boa prática: rodar sempre antes de commitar.

---

## 💡 Nomeie pelo o que é, não pelo o que faz

A tentação do iniciante é criar um pacote `utils/` e jogar tudo lá. Formatação de data, validação de CPF, helpers de string — vira uma gaveta da bagunça digital.

Go incentiva o oposto: nomes que dizem o domínio, não a utilidade.

Em vez de `utils.ValidarCPF()`, prefira `documentos.ValidarCPF()`.
Em vez de `helpers.FormatarData()`, prefira o formato dentro do próprio pacote de domínio.

Pacote bom responde à pergunta "o que ele representa?", não "pra que ele serve?".

---

## 🦫 O que ficou

Pacotes em Go são simplesmente pastas. Módulos são projetos com identidade. A visibilidade é controlada por uma letra. O ecossistema externo entra com um comando. E `go mod tidy` mantém tudo limpo.

É simples porque foi projetado para ser. Num projeto Go bem organizado, você sabe exatamente onde cada coisa mora — e o compilador te ajuda a manter as fronteiras.

Repositório completo: [Link do repositório]

---

**Checklist antes de publicar:**
- [x] Conteúdo confere com `estudos-go/capitulo14/fonte.txt`
- [x] Texto é original — sem frases copiadas do fonte.txt
- [x] Todos os blocos de código testados e funcionais
- [ ] Link para o repositório no GitHub incluído
- [ ] Revisão ortográfica feita
- [ ] Capa/imagem de destaque escolhida no Medium
- [ ] Tags do Medium definidas (Golang, Programming, Modules, Backend, CleanCode)
- [ ] Status atualizado em `conteudo/PAINEL.md`
