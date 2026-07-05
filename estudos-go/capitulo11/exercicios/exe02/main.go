//2. O Relógio Digital (Médio)
// Crie uma struct Relogia com Horas e Minutos. Crie um método
// AdicionarMinutos(m int). O método deve somar
// os minutos. Se passar de 60, deve incrementar a hora e zerar os minutos (ou manter o resto). Se passar de 24h, deve voltar para 00h. Exemplo:
// 23.50 + 20min -> 00:10.
// Dica: Use operador módulo %.

package main

import "fmt"

type Relogia struct {
	Horas   int
	Minutos int
}

func (r *Relogia) AdicionarMinutos(m int) {
	r.Minutos += m
	if r.Minutos >= 60 {
		r.Horas += r.Minutos / 60
		r.Minutos %= 60
	}
	if r.Horas >= 24 {
		r.Horas %= 24
	}
}

func main() {
	relogio := Relogia{Horas: 23, Minutos: 50}
	relogio.AdicionarMinutos(60)
	fmt.Printf("Horario final: %d:%d\n", relogio.Horas, relogio.Minutos)
}
