package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
)

type Game struct{}

func (g *Game) Draw(screen *ebiten.Image) {
	//create world
	world = NewWorld(xWidth, yHeight)
	world.GenerateGrid(alivePercentage)
}

func (g *Game) Update(screen *ebiten.Image) error {
	_ = screen.Fill(color.RGBA{0, 0, 0, 0xff})
	background, _ := ebiten.NewImage(xWidth, yHeight, ebiten.FilterDefault)

	world.Print(background)
	world.Next()

	_ = screen.DrawImage(background, &ebiten.DrawImageOptions{})
	return nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return xWidth, yHeight
}
