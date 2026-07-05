# 🏋️ Pilha de Exercícios — Capítulo 12: Interfaces

Exercícios extras em dificuldade crescente. Sem solução aqui — o esforço é o ponto.

**Conceitos disponíveis até este capítulo:** tudo até Interfaces (cap. 12): variáveis, tipos, controle de fluxo, for/range, funções, slices, maps, ponteiros, structs, métodos, interfaces, type assertion, type switch, any.

---

## 🟢 Fácil

### 1. O Contador de Caracteres Universal
Crie uma interface `TextoInfo` com um método `ContarCaracteres() int`. Implemente para os tipos `Frase` (campo `Texto string`) e `Paragrafo` (campo `Linhas []string`, onde `ContarCaracteres` retorna a soma dos caracteres de todas as linhas). Crie um slice `[]TextoInfo` com alguns itens e imprima a contagem de cada um.

### 2. O Formatador de Saída
Crie uma interface `Formatavel` com método `Formatar() string`. Implemente para: `Moeda` (campo `Valor float64`, formata como `"R$ X.XX"`), `Percentual` (campo `Valor float64`, formata como `"X.XX%"`), e `Temperatura` (campo `Celsius float64`, formata como `"X.X°C"`). Crie uma função `Exibir(f Formatavel)` que imprime o resultado formatado.

### 3. A Fila de Animais
Crie uma interface `Animal` com métodos `Nome() string` e `Som() string`. Implemente para `Gato`, `Cachorro` e `Passaro`. Crie um slice `[]Animal`, popule com vários animais de tipos diferentes e itere imprimindo `"[Nome] faz: [Som]"` para cada um.

### 4. O Validador
Crie uma interface `Validavel` com método `Validar() error`. Implemente para: `Email` (campo `Endereco string`, valida se contém `@`), `CPF` (campo `Numero string`, valida se tem exatamente 11 dígitos), e `Senha` (campo `Texto string`, valida se tem pelo menos 8 caracteres). Uma função `ValidarTodos(itens []Validavel)` imprime "OK" ou o erro de cada item.

---

## 🟡 Médio

### 5. O Banco de Dados em Memória
Crie uma interface `Repositorio` com métodos `Salvar(id string, valor any)`, `Buscar(id string) (any, bool)` e `Deletar(id string)`. Implemente um `RepositorioMemoria` usando `map[string]any` por baixo. Teste salvando structs diferentes (Produto, Usuario) no mesmo repositório.

### 6. O Renderizador de Relatório
Crie uma interface `Renderizavel` com método `Renderizar() string`. Implemente para: `RelatorioTexto` (retorna texto simples), `RelatorioHTML` (retorna HTML com tags `<h1>`, `<p>`), e `RelatorioMarkdown` (retorna markdown com `#` e `-`). Uma função `GerarRelatorio(r Renderizavel, titulo, conteudo string)` chama `r.Renderizar()` e imprime o resultado. O tipo concreto decide o formato.

### 7. O Ordenador Genérico
Crie uma interface `Comparavel` com método `MenorQue(outro Comparavel) bool`. Implemente para `NumeroInteiro` (campo `Valor int`) e `Texto` (campo `Valor string`, compara lexicograficamente). Implemente um Bubble Sort genérico que aceite `[]Comparavel` e ordene usando apenas `MenorQue`. Teste com slices de números e de textos.

### 8. O Pipeline de Transformação
Crie uma interface `Transformador` com método `Transformar(entrada string) string`. Implemente: `MaiusculoTransformador`, `TrimTransformador` (remove espaços), `SubstituirTransformador` (substitui `De` por `Para`, campos da struct). Crie uma função `AplicarPipeline(entrada string, transformadores []Transformador) string` que aplica cada transformação em sequência, passando o resultado de um para o próximo.

---

## 🔴 Desafio

### 9. O Event Bus (Publicador/Assinante)
Crie uma interface `Assinante` com método `Receber(evento string, dados any)`. Crie um struct `EventBus` com `map[string][]Assinante`. Implemente `Publicar(evento string, dados any)` (notifica todos os assinantes daquele evento) e `Assinar(evento string, a Assinante)`. Crie dois assinantes concretos: `LogAssinante` (imprime tudo no console) e `EmailAssinante` (simula envio de email). Teste publicando eventos com ambos assinando.

### 10. O Analisador de Tipos
Crie uma função `Analisar(valores []any) map[string]int` que conta quantos elementos de cada tipo há no slice. Use type switch. Retorne um map onde a chave é o nome do tipo (`"int"`, `"string"`, `"bool"`, `"float64"`, `"outro"`) e o valor é a contagem. Teste com um slice misto.

### 11. O Codificador/Decodificador
Crie uma interface `Codec` com métodos `Codificar(dados any) (string, error)` e `Decodificar(texto string) (any, error)`. Implemente `JSONCodec` usando `encoding/json`. Implemente `Base64Codec` que converte strings para base64 e de volta (use `encoding/base64`). Crie uma função `TransmitirDados(c Codec, dado any)` que codifica, exibe o resultado e decodifica de volta.

### 12. O Mini ORM
Crie uma interface `Modelo` com métodos `NomeTabela() string` e `ID() string`. Implemente para `Produto` (campos `Codigo, Nome string`) e `Cliente` (campos `CPF, Nome string`). Crie um `BancoDeDados` com `map[string]map[string]Modelo` (chave externa = nome da tabela, interna = ID). Implemente métodos `Inserir(m Modelo)`, `Encontrar(tabela, id string) (Modelo, bool)` e `ListarTodos(tabela string) []Modelo`. Teste com os dois tipos juntos.
