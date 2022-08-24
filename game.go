package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	world  *World
	pixels []byte
	input  *Input
}

func (g *Game) Update() error {
	g.world.Update()

	if dir, ok := g.input.Dir(); ok {
		if Dir.String(dir) == "Left" && TPS != 5 {
			TPS = TPS - 5
		}
		if Dir.String(dir) == "Right" {
			TPS = TPS + 5
		}

		ebiten.SetMaxTPS(TPS)
	}

	fmt.Println(ebiten.CurrentTPS(), TPS)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.pixels == nil {
		g.pixels = make([]byte, screenWidth*screenHeight*4)
	}
	g.world.Draw(g.pixels)
	screen.ReplacePixels(g.pixels)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f\nFPS: %0.2f", ebiten.CurrentTPS(), ebiten.CurrentFPS()))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
