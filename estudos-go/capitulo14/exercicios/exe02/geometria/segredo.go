package geometria

// Variável privada (letra minúscula).
// Não é acessível fora do pacote geometria.
var numeroMagico = 42

// PegarSegredo é uma função pública (letra maiúscula).
// Permite que pacotes externos acessem indiretamente a variável privada numeroMagico.
func PegarSegredo() int {
	return numeroMagico
}
