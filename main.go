// main.go - Loop principal do jogo
package main

import (
	"jogo/Client"
	"os"
)

func main() {
	// Inicializa a interface (termbox)
	Client.InterfaceIniciar()
	defer Client.InterfaceFinalizar()

	// Usa "mapa.txt" como arquivo padrão ou lê o primeiro argumento
	mapaFile := "mapa.txt"
	if len(os.Args) > 1 {
		mapaFile = os.Args[1]
	}

	// Inicializa o jogo
	jogo := Client.JogoNovo()
	if err := Client.JogoCarregarMapa(mapaFile, &jogo); err != nil {
		panic(err)
	}

	// Desenha o estado inicial do jogo
	Client.InterfaceDesenharJogo(&jogo)

	// Loop principal de entrada
	for {
		evento := Client.InterfaceLerEventoTeclado()
		if continuar := Client.PersonagemExecutarAcao(evento, &jogo); !continuar {
			break
		}
		Client.InterfaceDesenharJogo(&jogo)
	}
}
