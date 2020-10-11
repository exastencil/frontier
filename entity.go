package main

import "github.com/hajimehoshi/ebiten/v2"

// Entity ...
type Entity struct {
	X, Y, Z               int
	Length, Width, Height int
	Image                 *ebiten.Image
}
