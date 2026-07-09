package conversor

// KmParaMilhas converte quilômetros para milhas.
// É pública (Maiúscula) — acessível de qualquer pacote.
func KmParaMilhas(km float64) float64 {
	return km * 0.621371
}

// milhasParaKm converte o inverso.
// É privada (Minúscula) — só o pacote 'conversor' pode usar.
// Tente chamar isso da main e veja o erro de compilação!
func milhasParaKm(milhas float64) float64 {
	return milhas / 0.621371
}
