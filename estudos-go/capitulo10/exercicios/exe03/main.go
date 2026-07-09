//3. O sistema de Playlist (Desafio)
// Crie uma struct Musica (Titulo, Artista, DuracaoEmSegundos). Crie um struct Playlist(Nome, Musicas []Musica). Instancie uma playlist "Rock" com 3 músicas.
// Crie um loop que percorra a playlist e imprima:
// - Nome da música
// - Nome do artista
// - No final, a duração total da playlist em minutos (TotalSegundos / 60)

package main

import "fmt"

type Musica struct {
	Titulo            string
	Artista           string
	DuracaoEmSegundos int
}

type Playlist struct {
	Nome    string
	Musicas []Musica
}

func main() {
	musica1 := Musica{
		Titulo:            "The Pretender",
		Artista:           "Foo Fighters",
		DuracaoEmSegundos: 269,
	}
	musica2 := Musica{
		Titulo:            "Bohemian Rhapsody",
		Artista:           "Queen",
		DuracaoEmSegundos: 355,
	}
	musica3 := Musica{
		Titulo:            "Stairway to Heaven",
		Artista:           "Led Zeppelin",
		DuracaoEmSegundos: 482,
	}

	playlistRock := Playlist{
		Nome:    "Rock",
		Musicas: []Musica{musica1, musica2, musica3},
	}

	fmt.Println("Playlist:", playlistRock.Nome)
	for _, musica := range playlistRock.Musicas {
		fmt.Println("Título:", musica.Titulo)
		fmt.Println("Artista:", musica.Artista)
		fmt.Println("Duração em Segundos:", musica.DuracaoEmSegundos)
		fmt.Println()
	}

	totalSegundos := 0
	for _, musica := range playlistRock.Musicas {
		totalSegundos += musica.DuracaoEmSegundos
	}
	fmt.Println("Duração total da playlist em minutos:", totalSegundos/60)
}
