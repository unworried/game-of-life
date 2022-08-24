package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 480
	screenHeight = 320
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	/*flag.IntVar(&xWidth, "width", 480, "-width=640")
	flag.IntVar(&yHeight, "height", 320, "-height=480")
	flag.IntVar(&alivePercentage, "percent", 20, "-percent=30")
	flag.Parse()*/
	g := &Game{
		world: NewWorld(screenWidth, screenHeight, int((screenWidth*screenHeight)/10)),
	}

	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Game of Life")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}

}
