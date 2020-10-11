package main

import (
	"fmt"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2/inpututil"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	viewWidth  = 640
	viewHeight = 480
)

// Game ...
type Game struct {
	world World
	debug bool
}

// Update ...
func (g *Game) Update() error {
	// Toggle Debug Mode
	if inpututil.IsKeyJustPressed(ebiten.KeyD) {
		g.debug = !g.debug
	}
	return nil
}

// Draw ...
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Clear()
	g.world.Draw(screen)
	if g.debug {
		debug := fmt.Sprintf(
			"TPS: %0.2f\nFPS: %0.2f\nCamera: %v, %v",
			ebiten.CurrentTPS(),
			ebiten.CurrentFPS(),
			g.world.camera.X,
			g.world.camera.Y,
		)
		ebitenutil.DebugPrint(screen, debug)
	}
}

// Layout ...
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return viewWidth, viewHeight
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Frontier")
	game := Game{}
	image, _, _ := ebitenutil.NewImageFromFile("res/sprites/placeholder.png")
	entity := Entity{X: 0, Y: 0, Length: 16, Width: 16, Height: 10, Image: image}
	chunk := &Chunk{Entities: []Entity{entity}}
	game.world = World{
		chunks: []*Chunk{chunk},
	}

	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
