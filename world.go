package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

// World ...
type World struct {
	camera Entity
	chunks []*Chunk
	op     ebiten.DrawImageOptions
}

// Draw ...
func (w *World) Draw(dst *ebiten.Image) {
	for _, chunk := range w.chunks {
		for _, e := range chunk.Entities {
			w.op.GeoM.Reset()
			// Offset relative to entity center
			w.op.GeoM.Translate(
				float64(-(e.Length + e.Width)),
				float64(-((e.Length+e.Width)/2 + e.Height)),
			)

			// Offset relative to chunk center
			// op.GeoM.Translate(float64(e.X), float64(e.Y))

			// Offset chunk to world center

			// Offset world center to camera
			w.op.GeoM.Translate(
				float64(viewWidth/2),
				float64(viewHeight/2),
			)

			dst.DrawImage(e.Image, &w.op)
		}
	}
	// Center point for debugging
	dst.Set(viewWidth/2, viewHeight/2, color.White)
}
