package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 480
	screenHeight = 320
)

var pixelColor = []byte{255, 255, 255, 255}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	g := &Game{
		world: NewWorld(screenWidth, screenHeight, int((screenWidth*screenHeight)/4)),
	}

	fmt.Println(pixelColor)

	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Game of Life")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}

}
